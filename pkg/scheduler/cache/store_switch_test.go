/*
Copyright 2023 The Godel Scheduler Authors.

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

package cache

import (
	"reflect"
	"testing"

	"github.com/kubewharf/godel-scheduler/pkg/scheduler/cache/commonstores"
	nodestore "github.com/kubewharf/godel-scheduler/pkg/scheduler/cache/commonstores/node_store"
	pdbstore "github.com/kubewharf/godel-scheduler/pkg/scheduler/cache/commonstores/pdb_store"
	podstore "github.com/kubewharf/godel-scheduler/pkg/scheduler/cache/commonstores/pod_store"
	podgroupstore "github.com/kubewharf/godel-scheduler/pkg/scheduler/cache/commonstores/podgroup_store"
	preemptionstore "github.com/kubewharf/godel-scheduler/pkg/scheduler/cache/commonstores/preemption_store"
	unitstatusstore "github.com/kubewharf/godel-scheduler/pkg/scheduler/cache/commonstores/unit_status_store"
	"github.com/kubewharf/godel-scheduler/pkg/scheduler/cache/handler"
)

func Test_makeStoreSwitch(t *testing.T) {
	type args struct {
		handler   handler.CacheHandler
		storeType commonstores.StoreType
	}

	tests := []struct {
		name    string
		prepare func()
		cleanup func()
		args    args
		want    []commonstores.StoreName
	}{
		{
			name: "normal: no preemption, no queuechecker",
			args: args{
				handler: handler.MakeCacheHandlerWrapper().Obj(),
			},
			want: []commonstores.StoreName{
				podgroupstore.Name,
				unitstatusstore.Name,
				nodestore.Name,
				podstore.Name,
			},
		},
		{
			name: "normal: has preemption, no queuechecker",
			args: args{
				handler: handler.MakeCacheHandlerWrapper().
					EnableStore(string(preemptionstore.Name)).
					Obj(),
			},
			want: []commonstores.StoreName{
				pdbstore.Name,
				podgroupstore.Name,
				preemptionstore.Name,
				unitstatusstore.Name,
				nodestore.Name,
				podstore.Name,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare and Cleanup
			{
				if tt.prepare != nil {
					tt.prepare()
				}
				if tt.cleanup != nil {
					defer tt.cleanup()
				}
			}

			store := makeStoreSwitch(tt.args.handler, tt.args.storeType)

			gotStoreNames := []commonstores.StoreName{}
			store.Range(func(cs commonstores.CommonStore) error {
				gotStoreNames = append(gotStoreNames, cs.Name())
				return nil
			})

			if !reflect.DeepEqual(gotStoreNames, tt.want) {
				t.Errorf("makeStoreSwitch() = %v, want %v", gotStoreNames, tt.want)
			}
		})
	}
}
