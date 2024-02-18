package resourceswriter

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/dragondrop-cloud/cloud-concierge/main/internal/hclcreate"
	"github.com/dragondrop-cloud/cloud-concierge/main/internal/implementations/markdowncreation"
	terraformValueObjects "github.com/dragondrop-cloud/cloud-concierge/main/internal/implementations/terraform_value_objects"
	"github.com/dragondrop-cloud/cloud-concierge/main/internal/interfaces"
)

// Factory is a struct that generates instances implementing of the ResourcesWriter interface.
type Factory struct{}

// Instantiate creates an instance that implements the ResourcesWriter interface, with the implementation
// depending on the current environment.
func (f *Factory) Instantiate(environment string, vcs interfaces.VCS, provider terraformValueObjects.Provider, hclConfig hclcreate.Config, jobName string) (interfaces.ResourcesWriter, error) {
	switch environment {
	case "isolated":
		return new(IsolatedResourcesWriter), nil
	default:
		return f.bootstrappedResourceWriter(vcs, provider, hclConfig, jobName)
	}
}

// bootstrappedResourceWriter creates a complete implementation of the ResourcesWriter interface with
// configuration specified via environment variables.
func (f *Factory) bootstrappedResourceWriter(vcs interfaces.VCS, provider terraformValueObjects.Provider, hclConfig hclcreate.Config, jobName string) (interfaces.ResourcesWriter, error) {
	hclCreate, err := hclcreate.NewHCLCreate(hclConfig, provider)
	if err != nil {
		log.Errorf("[cannot instantiate hclCreate config]%s", err.Error())
		return nil, fmt.Errorf("[cannot instantiate hclCreate config]%w", err)
	}

	return NewTerraformResourceWriter(hclCreate, vcs, markdowncreation.NewMarkdownCreator(), jobName), nil
}
