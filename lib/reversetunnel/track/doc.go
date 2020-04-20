/*
Copyright 2020 Gravitational, Inc.

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

// Proxy Tracking
//
// The Tracker type provided by this package provides a simple interface for
// tracking known proxies by endpoint/name and correctly handling expiration
// and exclusivity.  The Tracker also wraps a workpool.Pool, updating per-key
// counts as new proxies are discovered and/or old proxies are expired.
//
package track
