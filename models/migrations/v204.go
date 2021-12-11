// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func addPinnedColumn(x *xorm.Engine) error {
	// ProjectIssue saves relation from issue to a project
	type Repository struct {
		IsPinned bool `xorm:"NOT NULL DEFAULT FALSE"`
	}

	return x.Sync2(new(Repository))
}
