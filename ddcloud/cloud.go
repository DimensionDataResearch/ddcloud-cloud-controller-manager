package ddcloud

import (
	"io/ioutil"
	"io"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	
	"github.com/DimensionDataResearch/go-dd-cloud-compute/compute"
	
	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/kubernetes/pkg/controller"
)

const providerName string = "ddcloud"

func init() {
	cloudprovider.RegisterCloudProvider(providerName, NewProvider)
}

// Config represents the configuration for the Dimension Data cloud controller provider.
type Config {
	Region string `yaml:"region"`
	UserName string `yaml:"user"`
	Password string `yaml:"password"`

	NetworkDomain `yaml:"network_domain"`
}

// The Dimension Data cloud controller provider.
type cloud struct {
	config Config
	client *compute.Client
}

// NewProvider creates a new instance of the Dimension Data cloud controller provider.
func NewProvider(configFile io.Reader) (cloudprovider.Interface, error) {
	configData, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read configuration")
	}

	config := Config{}
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse configuration")
	}

	cloud := &cloud{
		config: config,
		client: compute.NewClient(config.Region, config.UserName, config.Password),
	}

	return cloud, nil
}

func (cloud *cloud) Initialize(clientBuilder controller.ControllerClientBuilder) {
	
}

func (cloud *cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return nil, false
}

func (cloud *cloud) Instances() (cloudprovider.Instances, bool) {
	return nil, false
}

func (cloud *cloud) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

func (cloud *cloud) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

func (cloud *cloud) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

func (cloud *cloud) ProviderName() string {
	return providerName
}

func (cloud *cloud) ScrubDNS(nameservers, searches []string) (nsOut, srchOut []string) {
	return nil, nil
}

func (cloud *cloud) HasClusterID() bool {
	return false
}
