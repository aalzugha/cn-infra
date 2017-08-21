// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpcsync

import (
	"github.com/ligato/cn-infra/core"
	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/datasync/syncbase"
	"github.com/ligato/cn-infra/servicelabel"
)

// PluginID used in the Agent Core flavors
const PluginID core.PluginName = "grpc-sync"

// Plugin grpcsync implements Plugin interface therefore can be loaded with other plugins
type Plugin struct {
	Adapter      datasync.KeyValProtoWatcher
	ServiceLabel servicelabel.ReaderAPI
}

// Init uses provided connection to build new transport adapter
func (plugin *Plugin) Init() error {
	grpcAdapter := NewAdapter()
	plugin.Adapter = &syncbase.Adapter{Watcher: grpcAdapter}

	return nil
}

// Close resources
func (plugin *Plugin) Close() error {
	return nil
}