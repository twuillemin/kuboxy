package configuration

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ApplicationConfiguration is the configuration of the application, ie its global parameters
type ApplicationConfiguration struct {
	Address                      string `json:"address,omitempty" yaml:"address,omitempty"`
	RestPort                     int    `json:"restPort,omitempty" yaml:"restPort,omitempty"`
	WebSocketPort                int    `json:"webSocketPort,omitempty" yaml:"webSocketPort,omitempty"`
	CertificateFileName          string `json:"certificateFileName,omitempty" yaml:"certificateFileName,omitempty"`
	PrivateKeyFileName           string `json:"privateKeyFileName,omitempty" yaml:"privateKeyFileName,omitempty"`
	KubeContextConfigurationFile string `json:"kubeContextConfigurationFile,omitempty" yaml:"kubeContextConfigurationFile,omitempty"`
}

var currentConfiguration *ApplicationConfiguration

// GetConfiguration returns the the current configuration of the application
func GetConfiguration() (ApplicationConfiguration, error) {

	if currentConfiguration == nil {
		config, err := readConfiguration()
		if err != nil {
			return ApplicationConfiguration{}, err
		}
		currentConfiguration = &config
	}

	return *currentConfiguration, nil
}

// readConfiguration read the configuration from all the possible sources
func readConfiguration() (ApplicationConfiguration, error) {

	// The default configuration
	config := ApplicationConfiguration{
		Address:                      "localhost",
		RestPort:                     8080,
		WebSocketPort:                8081,
		CertificateFileName:          "",
		PrivateKeyFileName:           "",
		KubeContextConfigurationFile: filepath.Join(homeDir(), ".kuboxy", "kube.config"),
	}

	// Read values from flag on the command line
	configurationFilePtr := flag.String("configurationFile", "", "The file having the configuration for the server Properties of this file can be overwritten by passing directly parameters to the application")

	// Get command line configuration
	commandLineConfiguration := ApplicationConfiguration{
		Address:                      *flag.String("address", "", "The IP address for listening incoming connections"),
		RestPort:                     *flag.Int("restPort", -1, "The port for the REST services"),
		WebSocketPort:                *flag.Int("webSocketPort", -1, "The port for the WebSocket for events"),
		CertificateFileName:          *flag.String("certificateFileName", "", "The  name of the public certificate file"),
		PrivateKeyFileName:           *flag.String("privateKeyFileName", "", "The  name of the private key file"),
		KubeContextConfigurationFile: *flag.String("kubeContextConfigurationFile", "", "The  name of the file keeping the configuration of the context/cluster to connect to"),
	}

	// Parse the flags
	flag.Parse()

	// Get home configuration
	homeConfiguration, err := getHomeConfigurationFile()
	if err != nil {
		return config, err
	}
	updateConfiguration(&config, homeConfiguration)

	// Get specific configuration for overwriting home if asked
	if len(*configurationFilePtr) > 0 {
		specificConfiguration, e := getSpecificConfigurationFile(*configurationFilePtr)
		if e != nil {
			return config, e
		}
		updateConfiguration(&config, specificConfiguration)
	}

	updateConfiguration(&config, commandLineConfiguration)

	return config, nil
}

// Print prints the configuration in the standard output
func (conf ApplicationConfiguration) Print() {
	fmt.Printf("Configuration:\n")
	fmt.Printf("\taddress:                       %v\n", conf.Address)
	fmt.Printf("\trestPort:                      %v\n", conf.RestPort)
	fmt.Printf("\twebSocketPort:                 %v\n", conf.WebSocketPort)
	fmt.Printf("\tcertificateFileName:           %v\n", conf.CertificateFileName)
	fmt.Printf("\tprivateKeyFileName:            %v\n", conf.PrivateKeyFileName)
	fmt.Printf("\tkubeContextConfigurationFile:  %v\n", conf.KubeContextConfigurationFile)
}

// getHomeConfigurationFile read the configuration file from the current user directory. If the file is missing, no
// error is raised
func getHomeConfigurationFile() (ApplicationConfiguration, error) {

	configFileName := filepath.Join(homeDir(), ".kuboxy", "application.config")

	applicationConfiguration := ApplicationConfiguration{}

	// If the file does not exist or is not readable, just return empty structure without error
	if _, err := os.Stat(configFileName); err != nil {
		return applicationConfiguration, nil
	}

	// Read the file
	yamlFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return applicationConfiguration, fmt.Errorf("unable to read configuration file due to %v", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &applicationConfiguration)
	if err != nil {
		return applicationConfiguration, fmt.Errorf("unable to unmarshall configuration file due to %v", err.Error())
	}

	return applicationConfiguration, nil
}

// getSpecificConfigurationFile read the configuration file in a specific path. If the file is missing an error is
// raised
func getSpecificConfigurationFile(configurationFileName string) (ApplicationConfiguration, error) {

	applicationConfiguration := ApplicationConfiguration{}

	// If the file does not exist or is not readable, return err
	if _, err := os.Stat(configurationFileName); err != nil {
		return applicationConfiguration, err
	}

	// Read the file
	yamlFile, err := ioutil.ReadFile(configurationFileName)
	if err != nil {
		return applicationConfiguration, fmt.Errorf("unable to read configuration file due to %v", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &applicationConfiguration)
	if err != nil {
		return applicationConfiguration, fmt.Errorf("unable to unmarshall configuration file due to %v", err.Error())
	}

	return applicationConfiguration, nil
}

// homeDir returns the home directory of the user depending on the OS
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

// updateConfiguration updates a configuration with all valid parameters from another one. The result is always
// valid
func updateConfiguration(toUpdate *ApplicationConfiguration, source ApplicationConfiguration) {
	if len(source.Address) > 0 {
		toUpdate.Address = source.Address
	}
	if source.RestPort > 0 && source.RestPort < 65535 {
		toUpdate.RestPort = source.RestPort
	}
	if source.WebSocketPort > 0 && source.WebSocketPort < 65535 {
		toUpdate.WebSocketPort = source.WebSocketPort
	}
	if len(source.CertificateFileName) != 0 && len(source.PrivateKeyFileName) != 0 {
		toUpdate.CertificateFileName = source.CertificateFileName
		toUpdate.PrivateKeyFileName = source.PrivateKeyFileName
	}
	if len(source.KubeContextConfigurationFile) > 0 {
		toUpdate.KubeContextConfigurationFile = source.KubeContextConfigurationFile
	}
}
