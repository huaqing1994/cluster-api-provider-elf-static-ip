/*
Copyright 2022.

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

package manager

import "time"

const (
	defaultPrefix = "cape-ip-"

	// DefaultSyncPeriod is the default value for the eponymous
	// manager option.
	DefaultSyncPeriod = time.Minute * 10

	// DefaultPodName is the default value for the eponymous manager option.
	DefaultPodName = defaultPrefix + "controller-manager"

	// DefaultLeaderElectionID is the default value for the eponymous manager option.
	DefaultLeaderElectionID = DefaultPodName + "-runtime"
)