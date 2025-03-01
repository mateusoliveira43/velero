//go:build windows
// +build windows

/*
Copyright The Velero Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kopia

import (
	"context"
	"fmt"

	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/snapshot/restore"
)

type BlockOutput struct {
	*restore.FilesystemOutput

	targetFileName string
}

func (o *BlockOutput) WriteFile(ctx context.Context, relativePath string, remoteFile fs.File, progressCb restore.FileWriteProgress) error {
	return fmt.Errorf("block mode is not supported for Windows")
}

func (o *BlockOutput) BeginDirectory(ctx context.Context, relativePath string, e fs.Directory) error {
	return fmt.Errorf("block mode is not supported for Windows")
}
