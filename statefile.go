package appconfig

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"time"
)

// StateFile is a collection of Data along with a timestamp.
type StateFile struct {
	Dttm       float64 `json:"dttm"`
	Collection `json:"data"`
}

// Time returns time.Time from the StateFile's Dttm.
func (s *StateFile) Time() time.Time {
	return timeFromFloat64(s.Dttm)
}

// KafkaMSG is how the statefile arrives in Kafka.
type KafkaMSG struct {
	ENV     string `json:"env"`
	ASI     string `json:"asi"`
	EASI    string `json:"easi"`
	Node    string `json:"node"`
	Message string `json:"message"`
}

// StateFile returns the StateFile from a KafkaMSG.
func (k *KafkaMSG) StateFile() (stateFile StateFile, err error) {
	err = json.Unmarshal([]byte(k.Message), &stateFile)
	stateFile.AssignADs(stateFile.findDefaultADs())
	return
}

// SavedFile returns a SavedFile from a KafkaMSG.
func (k *KafkaMSG) SavedFile() (savedFile SavedFile, err error) {
	var stateFile StateFile
	err = json.Unmarshal([]byte(k.Message), &stateFile)
	if err != nil {
		return
	}
	stateFile.AssignADs(stateFile.findDefaultADs())
	savedFile = SavedFile{
		ENV:       k.ENV,
		ASI:       k.ASI,
		EASI:      k.EASI,
		EASIN:     k.EASIN(),
		Node:      k.Node,
		StateFile: stateFile,
	}
	return
}

// EASIN returns the EASIN string from a KafkaMSG.
func (k *KafkaMSG) EASIN() string {
	return k.EASI + `:` + k.Node
}

// SHA returns the Sha1 string using the EASIN.
func (k *KafkaMSG) SHA() string {
	b := []byte(k.EASI + `:` + k.Node)
	return fmt.Sprintf("%x", sha1.Sum(b))
}

func timeFromFloat64(ts float64) time.Time {
	secs := int64(ts)
	nsecs := int64((ts - float64(secs)) * 1e9)
	return time.Unix(secs, nsecs)
}

func (s *StateFile) findDefaultADs() string {
	var appDomain string
	ads := filterUnique(s.FromType(TypeSimple).Get(`appdomain`))
	switch {
	case len(ads) > 1:
		appDomain = `MULTIPLE`
	case len(ads) < 1:
		appDomain = `NA`
	case len(ads) == 1:
		switch a := ads[0]; a {
		case "":
			appDomain = `NA`
		default:
			appDomain = a
		}
	}
	return appDomain
}
