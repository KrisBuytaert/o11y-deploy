// Copyright 2023 The O11y Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prometheus

// GetDashboards returns pointers to grafana.com dashboards
func (m *Module) GetDashboards() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"dashboard_id": 3662,
			"revision_id":  2,
			"datasource":   "prometheus",
		},
	}
}

// GetDashboardFiles returns included dashboard files
func (m *Module) GetDashboardFiles() map[string][]byte {
	return make(map[string][]byte)
}
