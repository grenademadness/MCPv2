/*
 * Copyright Daniel Hawton
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package cache

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
)

var (
	ErrEntryNotFound = bigcache.ErrEntryNotFound
	cache            *bigcache.BigCache
)

func Setup() error {
	var err error
	cache, err = bigcache.New(context.Background(), bigcache.DefaultConfig(3*time.Minute))
	if err != nil {
		panic(err)
	}

	return nil
}

func Get(key string) ([]byte, error) {
	return cache.Get(key)
}

func Set(key string, value []byte) error {
	return cache.Set(key, value)
}
