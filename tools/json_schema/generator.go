package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// --------------------------------------------------------------------------------
//
// DEFINITION OF INPUT STRUCTURES
//
// --------------------------------------------------------------------------------

// Source is the head node of the document
type Source struct {
	SwaggerVersion string                 `json:"swaggerVersion"`
	APIVersion     string                 `json:"apiVersion"`
	BasePath       string                 `json:"basePath"`
	ResourcePath   string                 `json:"resourcePath"`
	Info           SourceInfo             `json:"info"`
	Apis           []interface{}          `json:"apis"`
	Models         map[string]SourceModel `json:"models"`
}

// SourceInfo gives some details about the document
type SourceInfo struct {
	Title       string `json:"id"`
	Description string `json:"description"`
}

// SourceModel defines an object
type SourceModel struct {
	ID          string                    `json:"id"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Required    []string                  `json:"required"`
	Properties  map[string]SourceProperty `json:"properties"`
}

// SourceProperty defines a property
type SourceProperty struct {
	Type        string     `json:"type,omitempty"`
	Description string     `json:"description,omitempty"`
	Format      string     `json:"format,omitempty"`
	Reference   string     `json:"$ref,omitempty"`
	Items       SourceItem `json:"items,omitempty"`
}

// SourceItem defines the content of an array
type SourceItem struct {
	Type        string `json:"type,omitempty"`
	Reference   string `json:"$ref,omitempty"`
	Description string `json:"description,omitempty"`
}

// --------------------------------------------------------------------------------
//
// DEFINITION OF OUTPUT STRUCTURES
//
// --------------------------------------------------------------------------------

// Node is an interface for all nodes
type Node interface {
	getNodeType() string
}

// HeadNode is the head node of the schema. It is defined apart as it holds specific attributes
type HeadNode struct {
	ID          string          `json:"$id"`
	Type        string          `json:"type,omitempty"`
	Schema      string          `json:"$schema"`
	Description string          `json:"description,omitempty"`
	Required    []string        `json:"required,omitempty"`
	Properties  map[string]Node `json:"properties,omitempty"`
}

func (o HeadNode) getNodeType() string {
	return "HeadNode"
}

// ObjectNode is a node defining an object
type ObjectNode struct {
	ID                   string            `json:"-"`
	Type                 string            `json:"type,omitempty"`
	Description          string            `json:"description,omitempty"`
	Required             []string          `json:"required,omitempty"`
	Properties           map[string]Node   `json:"properties,omitempty"`
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty"`
}

func (o ObjectNode) getNodeType() string {
	return "ObjectNode"
}

// AttributeNode is a node defining an attribute inside an object
type AttributeNode struct {
	ID          string `json:"-"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

func (o AttributeNode) getNodeType() string {
	return "AttributeNode"
}

// ArrayNode is a node defining an array of object or attribute
type ArrayNode struct {
	ID          string `json:"-"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Items       Node   `json:"items,omitempty"`
}

func (o ArrayNode) getNodeType() string {
	return "ArrayNode"
}

// --------------------------------------------------------------------------------
//
// BEGINNING OF CODE
//
// --------------------------------------------------------------------------------

func main() {

	coreModels := loadSpec(filepath.Join(homeDir(), "go", "src", "github.com", "kubernetes", "api", "swagger-spec", "v1.json"))
	appsModels := loadSpec(filepath.Join(homeDir(), "go", "src", "github.com", "kubernetes", "api", "swagger-spec", "apps_v1.json"))
	batchModels := loadSpec(filepath.Join(homeDir(), "go", "src", "github.com", "kubernetes", "api", "swagger-spec", "batch_v1.json"))
	batchModelsBeta1 := loadSpec(filepath.Join(homeDir(), "go", "src", "github.com", "kubernetes", "api", "swagger-spec", "batch_v1beta1.json"))
	rbacModels := loadSpec(filepath.Join(homeDir(), "go", "src", "github.com", "kubernetes", "api", "swagger-spec", "rbac.authorization.k8s.io_v1.json"))
	networkingModels := loadSpec(filepath.Join(homeDir(), "go", "src", "github.com", "kubernetes", "api", "swagger-spec", "networking.k8s.io_v1.json"))
	storageModels := loadSpec(filepath.Join(homeDir(), "go", "src", "github.com", "kubernetes", "api", "swagger-spec", "storage.k8s.io_v1.json"))

	generateSchema(coreModels, "v1.Namespace", "namespace_schema.json")
	generateSchema(coreModels, "v1.Node", "node_schema.json")
	generateSchema(coreModels, "v1.PersistentVolume", "persistent_volume_schema.json")
	generateSchema(rbacModels, "v1.ClusterRole", "cluster_role_schema.json")
	generateSchema(rbacModels, "v1.ClusterRoleBinding", "cluster_role_binding_schema.json")
	generateSchema(storageModels, "v1.StorageClass", "storage_class_schema.json")

	generateSchema(coreModels, "v1.Service", "service_schema.json")
	generateSchema(coreModels, "v1.Pod", "pod_schema.json")
	generateSchema(coreModels, "v1.PersistentVolumeClaim", "persistent_volume_claim_schema.json")
	generateSchema(coreModels, "v1.ConfigMap", "config_map_schema.json")
	generateSchema(coreModels, "v1.ReplicationController", "replication_controller_schema.json")
	generateSchema(coreModels, "v1.Secret", "secret_schema.json")
	generateSchema(coreModels, "v1.ServiceAccount", "service_account_schema.json")
	generateSchema(appsModels, "v1.Deployment", "deployment_schema.json")
	generateSchema(appsModels, "v1.StatefulSet", "stateful_set_schema.json")
	generateSchema(appsModels, "v1.DaemonSet", "daemon_set_schema.json")
	generateSchema(appsModels, "v1.ReplicaSet", "replica_set_schema.json")
	generateSchema(batchModels, "v1.Job", "job_schema.json")
	generateSchema(batchModelsBeta1, "v1beta1.CronJob", "cron_job_schema.json")
	generateSchema(networkingModels, "v1.NetworkPolicy", "network_policy_schema.json")
	generateSchema(rbacModels, "v1.Role", "role_schema.json")
	generateSchema(rbacModels, "v1.RoleBinding", "role_binding_schema.json")
}

// loadSpec loads the models from a specification file
func loadSpec(fileName string) map[string]SourceModel {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		err = jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var source Source

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	if err = json.Unmarshal(bytes, &source); err != nil {
		panic(err)
	}

	fmt.Printf("Found %v models in file %v\n", len(source.Models), fileName)

	return source.Models
}

// generateSchema generates a schema file for the given object
func generateSchema(models map[string]SourceModel, objectName string, fileName string) {
	node := formatHead(buildObject(models, objectName))

	schema, err := json.MarshalIndent(node, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	output, err := os.Create(filepath.Join(homeDir(), "go", "src", "github.com", "twuillemin", "kuboxy", "docs", "json_schemas", fileName))
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		err = output.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = output.Write(schema)
	if err != nil {
		fmt.Println(err)
	}
}

// buildObject convert the requested object to a json schema structure. Note that, as this is a recursive function, the
// head node of the structure is return just as an ObjectNode instead of an HeadNode
func buildObject(models map[string]SourceModel, name string) Node {

	model, ok := models[name]
	if !ok {
		fmt.Printf("Missing definition for %v\n", name)
	}

	properties := make(map[string]Node)

	for propertyName, property := range model.Properties {
		if len(property.Reference) > 0 {
			node := buildObject(models, property.Reference)
			properties[propertyName] = node
		} else if len(property.Type) > 0 {
			switch property.Type {

			case "string", "integer", "boolean":

				properties[propertyName] = AttributeNode{
					model.ID + "/" + propertyName,
					property.Type,
					property.Description,
				}

			case "array":

				var subNode Node
				// If array with simple property
				if len(property.Items.Type) > 0 {
					subNode = AttributeNode{
						model.ID + "/" + propertyName + "/items",
						property.Type,
						property.Description,
					}
				} else if len(property.Items.Reference) > 0 {
					subNode = buildObject(models, property.Items.Reference)
				} else {
					fmt.Printf("Unsuported Array type %v: %v\n", propertyName, property.Type)
				}

				properties[propertyName] = ArrayNode{
					model.ID + "/" + propertyName,
					"array",
					property.Description,
					subNode,
				}

			case "object":

				additionalProperties := make(map[string]string)
				additionalProperties["type"] = "string"

				properties[propertyName] = ObjectNode{
					model.ID + "/" + propertyName,
					property.Type,
					property.Description,
					nil,
					nil,
					additionalProperties,
				}

			default:

				fmt.Printf("Unsuported Property type %v: %v\n", propertyName, property.Type)

			}
		} else {
			fmt.Printf("Unable to process property Property %v\n", propertyName)
		}
	}

	return ObjectNode{
		model.ID,
		"object",
		model.Description,
		model.Required,
		properties,
		nil,
	}
}

// formatHead converts the head node of a json schema from a standard ObjectNode to a HeadNode
func formatHead(node Node) HeadNode {

	objectNode, ok := node.(ObjectNode)
	if !ok {
		panic(fmt.Errorf("the node given for head is not an ObjectNode"))
	}

	return HeadNode{
		objectNode.ID,
		objectNode.Type,
		"http://json-schema.org/draft-07/schema#",
		objectNode.Description,
		objectNode.Required,
		objectNode.Properties,
	}
}

// homeDir returns the home directory of the user depending on the OS
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
