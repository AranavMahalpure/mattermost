// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package jobs

import (
	"github.com/mattermost/mattermost-server/server/public/model"
)

type IndexerJobInterface interface {
	MakeWorker() model.Worker
}
