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

package pnm

import . "github.com/kkdai/pnm/opendata"

//Pets :All pet related API
type Pets struct {
	allPets    []Pet
	queryIndex int
}

//NewPets :
func NewPets() *Pets {
	p := new(Pets)
	p.getPets()
	return p
}

//GetNextPet :
func (p *Pets) GetNextPet() *Pet {
	if len(p.allPets) == 0 {
		p.getPets()
	}

	retPet := &p.allPets[p.getNextIndex()]
	return retPet
}

//GetPreviousPet :
func (p *Pets) GetPreviousPet() *Pet {
	if len(p.allPets) == 0 {
		p.getPets()
	}

	retPet := &p.allPets[p.getPreviousIndex()]
	return retPet
}

//GetNextDog :
func (p *Pets) GetNextDog() *Pet {
	if len(p.allPets) == 0 {
		p.getPets()
	}

	var retPet *Pet
	for {
		retPet = &p.allPets[p.getNextIndex()]
		if retPet.PetType() == Dog {
			break
		}
	}

	return retPet
}

//GetNextCat :
func (p *Pets) GetNextCat() *Pet {
	if len(p.allPets) == 0 {
		p.getPets()
	}

	var retPet *Pet
	for {
		retPet = &p.allPets[p.getNextIndex()]
		if retPet.PetType() == Cat {
			break
		}
	}
	return retPet
}

//GetPetsCount :
func (p *Pets) GetPetsCount() int {
	return len(p.allPets)
}

func (p *Pets) getPets() {
	db := NewPetDB(AllRegion)
	p.allPets = db.GetPets()
}

func (p *Pets) getPreviousIndex() int {
	if p.queryIndex < 0 {
		p.queryIndex = 0
	}

	retInt := p.queryIndex
	p.queryIndex--
	return retInt
}

func (p *Pets) getNextIndex() int {
	if p.queryIndex >= len(p.allPets) {
		p.queryIndex = 0
	}

	retInt := p.queryIndex
	p.queryIndex++
	return retInt
}
