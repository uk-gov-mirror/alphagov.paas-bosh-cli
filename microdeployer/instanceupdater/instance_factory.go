package instanceupdater

import (
	boshlog "github.com/cloudfoundry/bosh-agent/logger"
	boshsys "github.com/cloudfoundry/bosh-agent/system"
	bmagentclient "github.com/cloudfoundry/bosh-micro-cli/microdeployer/agentclient"
	bmas "github.com/cloudfoundry/bosh-micro-cli/microdeployer/applyspec"
)

type InstanceFactory interface {
	Create(string) Instance
}

type instanceFactory struct {
	agentClientFactory     bmagentclient.Factory
	templatesSpecGenerator TemplatesSpecGenerator
	applySpecFactory       bmas.Factory
	fs                     boshsys.FileSystem
	logger                 boshlog.Logger
}

func NewInstanceFactory(
	agentClientFactory bmagentclient.Factory,
	templatesSpecGenerator TemplatesSpecGenerator,
	applySpecFactory bmas.Factory,
	fs boshsys.FileSystem,
	logger boshlog.Logger,
) InstanceFactory {
	return &instanceFactory{
		agentClientFactory:     agentClientFactory,
		templatesSpecGenerator: templatesSpecGenerator,
		applySpecFactory:       applySpecFactory,
		fs:                     fs,
		logger:                 logger,
	}
}

func (f *instanceFactory) Create(mbusURL string) Instance {
	agentClient := f.agentClientFactory.Create(mbusURL)

	return NewInstance(
		agentClient,
		f.templatesSpecGenerator,
		f.applySpecFactory,
		mbusURL,
		f.fs,
		f.logger,
	)
}
