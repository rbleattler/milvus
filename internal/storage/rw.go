// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"context"
	"fmt"
	sio "io"

	"github.com/samber/lo"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/internal/allocator"
	"github.com/milvus-io/milvus/internal/storagecommon"
	"github.com/milvus-io/milvus/internal/storagev2/packed"
	"github.com/milvus-io/milvus/pkg/v2/proto/datapb"
	"github.com/milvus-io/milvus/pkg/v2/util/merr"
)

const (
	StorageV1 int64 = 0
	StorageV2 int64 = 2
)

type rwOptions struct {
	version             int64
	bufferSize          int64
	downloader          func(ctx context.Context, paths []string) ([][]byte, error)
	uploader            func(ctx context.Context, kvs map[string][]byte) error
	multiPartUploadSize int64
	columnGroups        []storagecommon.ColumnGroup
}

type RwOption func(*rwOptions)

func DefaultRwOptions() *rwOptions {
	return &rwOptions{
		bufferSize:          packed.DefaultWriteBufferSize,
		multiPartUploadSize: packed.DefaultMultiPartUploadSize,
	}
}

func WithVersion(version int64) RwOption {
	return func(options *rwOptions) {
		options.version = version
	}
}

func WithBufferSize(bufferSize int64) RwOption {
	return func(options *rwOptions) {
		options.bufferSize = bufferSize
	}
}

func WithMultiPartUploadSize(multiPartUploadSize int64) RwOption {
	return func(options *rwOptions) {
		options.multiPartUploadSize = multiPartUploadSize
	}
}

func WithDownloader(downloader func(ctx context.Context, paths []string) ([][]byte, error)) RwOption {
	return func(options *rwOptions) {
		options.downloader = downloader
	}
}

func WithUploader(uploader func(ctx context.Context, kvs map[string][]byte) error) RwOption {
	return func(options *rwOptions) {
		options.uploader = uploader
	}
}

func WithColumnGroups(columnGroups []storagecommon.ColumnGroup) RwOption {
	return func(options *rwOptions) {
		options.columnGroups = columnGroups
	}
}

func NewBinlogRecordReader(ctx context.Context, binlogs []*datapb.FieldBinlog, schema *schemapb.CollectionSchema, option ...RwOption) (RecordReader, error) {
	rwOptions := DefaultRwOptions()
	for _, opt := range option {
		opt(rwOptions)
	}
	switch rwOptions.version {
	case StorageV1:
		itr := 0
		return newCompositeBinlogRecordReader(schema, func() ([]*Blob, error) {
			if len(binlogs) <= 0 {
				return nil, sio.EOF
			}
			paths := make([]string, len(binlogs))
			for i, fieldBinlog := range binlogs {
				if itr >= len(fieldBinlog.GetBinlogs()) {
					return nil, sio.EOF
				}
				paths[i] = fieldBinlog.GetBinlogs()[itr].GetLogPath()
			}
			itr++
			values, err := rwOptions.downloader(ctx, paths)
			if err != nil {
				return nil, err
			}
			blobs := lo.Map(values, func(v []byte, i int) *Blob {
				return &Blob{Key: paths[i], Value: v}
			})
			return blobs, nil
		})
	case StorageV2:
		if len(binlogs) <= 0 {
			return nil, sio.EOF
		}
		binlogLists := lo.Map(binlogs, func(fieldBinlog *datapb.FieldBinlog, _ int) []*datapb.Binlog {
			return fieldBinlog.GetBinlogs()
		})
		paths := make([][]string, len(binlogLists[0]))
		for _, binlogs := range binlogLists {
			for j, binlog := range binlogs {
				paths[j] = append(paths[j], binlog.GetLogPath())
			}
		}
		return newPackedRecordReader(paths, schema, rwOptions.bufferSize)
	}
	return nil, merr.WrapErrServiceInternal(fmt.Sprintf("unsupported storage version %d", rwOptions.version))
}

func NewBinlogRecordWriter(ctx context.Context, collectionID, partitionID, segmentID UniqueID,
	schema *schemapb.CollectionSchema, allocator allocator.Interface, chunkSize uint64, rootPath string, maxRowNum int64,
	option ...RwOption,
) (BinlogRecordWriter, error) {
	rwOptions := DefaultRwOptions()
	for _, opt := range option {
		opt(rwOptions)
	}
	blobsWriter := func(blobs []*Blob) error {
		kvs := make(map[string][]byte, len(blobs))
		for _, blob := range blobs {
			kvs[blob.Key] = blob.Value
		}
		return rwOptions.uploader(ctx, kvs)
	}
	switch rwOptions.version {
	case StorageV1:
		return newCompositeBinlogRecordWriter(collectionID, partitionID, segmentID, schema,
			blobsWriter, allocator, chunkSize, rootPath, maxRowNum,
		)
	case StorageV2:
		return newPackedBinlogRecordWriter(collectionID, partitionID, segmentID, schema,
			blobsWriter, allocator, chunkSize, rootPath, maxRowNum,
			rwOptions.bufferSize, rwOptions.multiPartUploadSize, rwOptions.columnGroups,
		)
	}
	return nil, merr.WrapErrServiceInternal(fmt.Sprintf("unsupported storage version %d", rwOptions.version))
}
