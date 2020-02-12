package simplemath

import "fmt"

// to modify any kind of state its absolute necessary to use pointers

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion {
		major: major,
		minor: minor, 
		patch: patch,
	}
}

// first method for semantic version, over the method declaration String() is tied to the semantic method
func (sv SemanticVersion) String() string{
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}


// at using * pointer into the semantiv version object we can actualy change the memory alogation and by that the STATE of the data
func (sv *SemanticVersion) IncrementMayor() {
	sv.major += 1
}

func (sv *SemanticVersion) IncrementMinor() {	
	sv.minor += 1
}

func (sv *SemanticVersion) IncrementPatch() {
	sv.patch += 1
}	