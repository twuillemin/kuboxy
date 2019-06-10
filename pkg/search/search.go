package search

//go:generate go run gen/gen_search_cluster.go
//go:generate go run gen/gen_search_namespace.go

import (
	"github.com/twuillemin/kuboxy/pkg/types"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"regexp"
)

// Parameter groups all the possible parameters for searching objects
type Parameter struct {
	Name        string             `json:"name"`
	Namespace   string             `json:"namespace"`
	Label       string             `json:"label"`
	LabelValue  string             `json:"labelValue"`
	ObjectTypes []types.ObjectType `json:"objectTypes"`
}

type preparedParameter struct {
	Name        *regexp.Regexp
	Namespace   *regexp.Regexp
	Label       *regexp.Regexp
	LabelValue  *regexp.Regexp
	objectTypes map[types.ObjectType]bool
}

// Search searches all objects matching the given parameters
func Search(contextName string, parameter Parameter) (map[types.ObjectType][]interface{}, error) {

	searchParameter, err := prepareParameters(parameter)
	if err != nil {
		return nil, err
	}

	results := make(map[types.ObjectType][]interface{})

	// If a namespace is specified, do not search the object at the cluster level as they don't have namespace...
	if searchParameter.Namespace == nil {
		err = searchClusterObjects(contextName, searchParameter, results)
		if err != nil {
			return nil, err
		}
	}
	err = searchNamespaceObjects(contextName, searchParameter, results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func prepareParameters(searchParameter Parameter) (*preparedParameter, error) {

	prepared := preparedParameter{}

	if len(searchParameter.Name) > 0 {
		regex, err := regexp.Compile(searchParameter.Name)
		if err != nil {
			return nil, err
		}
		prepared.Name = regex
	}

	if len(searchParameter.Namespace) > 0 {
		regex, err := regexp.Compile(searchParameter.Namespace)
		if err != nil {
			return nil, err
		}
		prepared.Namespace = regex
	}

	if len(searchParameter.Label) > 0 {
		regex, err := regexp.Compile(searchParameter.Label)
		if err != nil {
			return nil, err
		}
		prepared.Label = regex
	}

	if len(searchParameter.LabelValue) > 0 {
		regex, err := regexp.Compile(searchParameter.LabelValue)
		if err != nil {
			return nil, err
		}
		prepared.LabelValue = regex
	}

	objectTypes := make(map[types.ObjectType]bool)
	if len(searchParameter.ObjectTypes) > 0 {
		for _, objectType := range searchParameter.ObjectTypes {
			objectTypes[objectType] = true
		}
	} else {
		for _, objectType := range types.ClusterObjectDefinitions {
			objectTypes[objectType.Type] = true
		}
		for _, objectType := range types.NamespaceObjectDefinitions {
			objectTypes[objectType.Type] = true
		}
	}

	prepared.objectTypes = objectTypes

	return &prepared, nil
}

func isValidNamespaceObject(meta meta.ObjectMeta, searchParameter *preparedParameter) bool {
	if searchParameter.Name != nil {
		if match := searchParameter.Name.Match([]byte(meta.Name)); !match {
			return false
		}
	}
	if searchParameter.Namespace != nil {
		if match := searchParameter.Namespace.Match([]byte(meta.Namespace)); !match {
			return false
		}
	}

	return hasLabel(meta, searchParameter.Label, searchParameter.LabelValue)
}

func isValidClusterObject(meta meta.ObjectMeta, searchParameter *preparedParameter) bool {
	if searchParameter.Name != nil {
		if match := searchParameter.Name.Match([]byte(meta.Name)); !match {
			return false
		}
	}
	return hasLabel(meta, searchParameter.Label, searchParameter.LabelValue)
}

// hasLabel checks if the given meta has the searched label and value. In case of searched label and value being both
// null, the result is always positive
func hasLabel(meta meta.ObjectMeta, searchedLabel *regexp.Regexp, searchedValue *regexp.Regexp) bool {

	// If nothing searched, always true
	if searchedLabel == nil && searchedValue == nil {
		return true
	}

	// If a label is specified
	if searchedLabel != nil {
		// Search for the label and optionally for the value
		for label, value := range meta.Labels {
			if searchedLabel.Match([]byte(label)) {
				if searchedValue == nil || searchedValue.Match([]byte(value)) {
					return true
				}
			}
		}
	} else {
		// Search just for the value (that can not be null at this point)
		for _, value := range meta.Labels {
			if searchedValue.Match([]byte(value)) {
				return true
			}
		}
	}

	return false
}
