// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package opendata

import "testing"

//TestTaipeiPetsData :Test if Taipei Pet data still exist
func TestTaipeiPetsData(t *testing.T) {
	t := new(TaipeiPets)
	url := t.GetURL()

	if url == "" {
		t.Error("URL error")
	}

	Pets := t.GetPets()
	if len(Pets) == 0 {
		t.Error("Empty in Taipei Pets")
	}
}