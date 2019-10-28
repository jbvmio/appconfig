package appconfig

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

// SavedState constains saved StateFiles.
type SavedState []SavedFile

// Collection returns the aggregated Collection for all data contained in the SavedState.
func (s SavedState) Collection() Collection {
	var data []Data
	for _, SS := range s {
		data = append(data, SS.Collection()...)
	}
	return data
}

// GetAll returns everything.
func (s SavedState) GetAll() (envs, asis, easis, easins, nodes []string) {
	var wg sync.WaitGroup
	wg.Add(5)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		envs = s.ENVs()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		asis = s.ASIs()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		easis = s.EASIs()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		easins = s.EASINs()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		nodes = s.Nodes()
	}(&wg)
	wg.Wait()
	return
}

// ENVs returns all the ENVs found in the SavedState.
func (s SavedState) ENVs() (envs []string) {
	dupe := make(map[string]bool)
	for _, d := range s {
		if !dupe[d.ENV] {
			dupe[d.ENV] = true
			envs = append(envs, d.ENV)
		}
	}
	return
}

// ASIs returns all the ASIs found in the SavedState.
func (s SavedState) ASIs() (asis []string) {
	dupe := make(map[string]bool)
	for _, d := range s {
		if !dupe[d.ASI] {
			dupe[d.ASI] = true
			asis = append(asis, d.ASI)
		}
	}
	return
}

// EASIs returns all the EASIs found in the SavedState.
func (s SavedState) EASIs() (easis []string) {
	dupe := make(map[string]bool)
	for _, d := range s {
		if !dupe[d.EASI] {
			dupe[d.EASI] = true
			easis = append(easis, d.EASI)
		}
	}
	return
}

// EASINs returns all the EASINs found in the SavedState.
func (s SavedState) EASINs() (easins []string) {
	dupe := make(map[string]bool)
	for _, d := range s {
		if !dupe[d.EASIN] {
			dupe[d.EASIN] = true
			easins = append(easins, d.EASIN)
		}
	}
	return
}

// Nodes returns all the Nodes found in the SavedState.
func (s SavedState) Nodes() (nodes []string) {
	dupe := make(map[string]bool)
	for _, d := range s {
		if !dupe[d.Node] {
			dupe[d.Node] = true
			nodes = append(nodes, d.Node)
		}
	}
	return
}

// SavedFile is a StateFile in a saved state prepared for retrieval.
type SavedFile struct {
	ENV       string    `json:"env"`
	ASI       string    `json:"asi"`
	EASI      string    `json:"easi"`
	Node      string    `json:"node"`
	EASIN     string    `json:"easin"`
	StateFile StateFile `json:"statefile"`
}

// Collection returns the underlying Collection from the SavedFile.
func (s *SavedFile) Collection() Collection {
	return s.StateFile.Collection
}

// HasENV returns true if the entered string matches the env, false otherwise.
func (s *SavedFile) HasENV(env string) bool {
	return s.ENV == env
}

// HasASI returns true if the entered string matches the asi, false otherwise.
func (s *SavedFile) HasASI(asi string) bool {
	return s.ASI == asi
}

// HasEASI returns true if the entered string matches the easi, false otherwise.
func (s *SavedFile) HasEASI(easi string) bool {
	return s.EASI == easi
}

// HasNode returns true if the entered string matches the node, false otherwise.
func (s *SavedFile) HasNode(node string) bool {
	return s.Node == node
}

// HasEASIN returns true if the entered string matches the easin, false otherwise.
func (s *SavedFile) HasEASIN(easin string) bool {
	return s.EASIN == easin
}

// SHA returns the Sha1 string using the EASIN.
func (s *SavedFile) SHA() string {
	b := []byte(s.EASI + `:` + s.Node)
	return fmt.Sprintf("%x", sha1.Sum(b))
}
