# Introduction
Kuboxy (KUBernetes rOXY) is a software proxy that allows a user to connect to multiple Kubernetes cluster from a single 
server.

Kuboxy offers also several interesting features:

  * Classical REST endpoints for accessing all object (JSON format)
  * Swagger interface with all objects fully documented
  * Interface WebSocket for receiving events from the clusters
  * Kuboxy con also be itself hosted in a Kubernetes cluster
  * The JSON schemas of the Kubernetes objects are provided

# Usage

```bash
go run cmd/main.go [parameters]
```

or for building the application

```bash
go build -o kuboxy.exe cmd/main.go
./kuboxy.exe [parameters]
```

# Configuration

## Options
Options are the following:

| Name | Usage | Default Value | Example |
| --- | --- | --- | --- |
| configurationFilePtr | Indicates a YAML/JSON file defining all the following parameters | ~/.kuboxy/application.config | ```./kuboxy.exe -configurationFilePtr="~/.kuboxy/config.json"``` |
| address | The IP address used by the service | 127.0.0.1| ```./kuboxy.exe -address=192.168.1.1``` |   
| restPort | The IP port used by the service for REST endpoints | 8080 | ```./kuboxy.exe -restPort=8080``` |
| webSocketPort | The IP port used by the event service | 8081 | ```./kuboxy.exe -webSocketPort=8081``` |
| certificateFileName | The certificate used for providing HTTPS connections | _none_  | ```./kuboxy.exe -certificateFileName="~/.kuboxy/cert.pem"``` |
| privateKeyFileName | The private key of the certificate | _none_  | ```./kuboxy.exe -privateKeyFileName="~/.kuboxy/key.pem"``` |
| kubeContextConfigurationFile | The file storing the credentials of the clusters | ~/.kuboxy/kube.config | ```./kuboxy.exe -kubeContextConfigurationFile="~/.kuboxy/kube.config"``` |

The options, save for ```configurationFilePtr``` can be defined permanently in a YAML file (JSON file is also 
acceptable as it is a subset of YAML). The equivalent of the above example are:

```yaml
address: "192.168.1.1", 
restPort: 8080,
webSocketPort: 8081, 
certificateFileName: "~/.kuboxy/cert.pem",
privateKeyFileName: "~/.kuboxy/key.pem",
kubeContextConfigurationFile: "~/.kuboxy/kube.config"
```

or

```json
{
  "address": "192.168.1.1", 
  "restPort": 8080,
  "webSocketPort": 8081, 
  "certificateFileName": "~/.kuboxy/cert.pem",
  "privateKeyFileName": "~/.kuboxy/key.pem",
  "kubeContextConfigurationFile": "~/.kuboxy/kube.config"
}
```

By default the configuration is located in ```~/.kuboxy/application.config```.

## Order of evaluations
The options are evaluated in the following order:

  1) The default values
  2) The default configuration file (```~/.kuboxy/application.config```) is loaded and its values override the default values
  3) If a specific configuration file (option ```configurationFilePtr```) is given, it is loaded and its values override the existing ones
  4) The parameters given on the command line override the existing values

Although it may seems a bit convoluted, it is not necessary to use all possibilities. In production use, defining 
everything in ```~/.kuboxy/application.config```.

# REST API
All the endpoints are available: https://localhost:8080/swagger/index.html. The endpoints are grouped by families:

 * Configuration
 * Labels
 * Objects at the cluster level
 * Objects at the namespace level
 * Search and summary
 
## Configuration endpoints
The configuration endpoints allows to configure the application, more precisely the cluster referenced by the *Cuboxy*. 
As by Kubernetes standards, the following information can be configured:

 * The cluster: name, server, certificate
 * The user: name, certificate or password
 * The context: name, cluster and user with an optional namespace
 
## Labels
This single endpoint allows to retrieve easily all the labels and their possible values for context/namespace. 

## Objects at the cluster level
These endpoints allows CRUD operations on the following objects:

 * clusterRoleBindings
 * clusterRoles
 * namespaces
 * nodeMetricses
 * nodes
 * persistentVolumes
 * storageClasses
 
## Objects at the namespace level
These endpoints allows CRUD operations on the following objects:

 * configMaps
 * cronJobs
 * daemonSets
 * deployments
 * jobs
 * networkPolicies
 * persistentVolumeClaims
 * podMetricses
 * pods
 * replicaSets
 * replicationControllers
 * roleBindings
 * roles
 * secrets
 * serviceAccounts
 * services
 * statefulSets
 
## Search and summary
This two endpoints allows to easily search objects in a cluster and to generate a high level overview of state of the 
cluster

# WebSocket events
It is possible for a client to subscribe to a context events. The subscription is running over a WebSocket connection, 
so once the chanel is open, a client can't manage its subscription and receive events without further connection.

The process is the following:
 
 1) The client connect to the WebSocket (by default https://localhost:8081/api/v1/events/)
 2) The client send a request for subscribing/unsubscribing to event. Requests are defined here after
 3) As soon as change (new/update/delete) is done in the cluster about an object, the client will receive the object 
 over the WebSocket
 
The command object has the following structure:
 
```json
{
  "command": "The command",
  "objectType": "The type of object (optional)",
  "contextName": "The name of the context (optional)",
  "namespaceName": "The namespace (optional)"
}
```

The command sent by the client is one of the following:

 * "AddSource": For subscribing a new source of event.
 * "RemoveSource": For unsubscribing to an existing source.
 * "RemoveAllSources": For unsubscribing to ... all sources.
 
# Generating documentation

## Generating Swagger documentation 
Swagger/swaggo are using a definition located at `docs/docs.go`. This documentation is not generated automatically at
the compilation time. So in case of changes in the API, it is necessary to regenerate this file with the following
procedure:
  
  1) Install, if needed, Swaggo : `go get github.com/swaggo/swag/cmd/swag` that will generate a _swag_ executable
  2) Run `swag init` in the root of the project
  3) In the folder _docs_ of the project, run `update_docs.go`

## Generating the JSON schemas
The documentation also includes the JSON schemas for the various object managed by the application. This schemas are
located in the folder `docs/json_schemas`. Although not directly used by the application, the may be used by the client.
For generating the schemas, run the file `generator.go` located in `tools/json_schema`.


# License

Copyright 2019 Thomas Wuillemin  <thomas.wuillemin@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this project or its content except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.