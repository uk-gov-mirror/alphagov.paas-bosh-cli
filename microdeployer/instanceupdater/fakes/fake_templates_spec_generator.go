package fakes

import (
	bmdepl "github.com/cloudfoundry/bosh-micro-cli/deployment"
	bminsup "github.com/cloudfoundry/bosh-micro-cli/microdeployer/instanceupdater"
	bmstemcell "github.com/cloudfoundry/bosh-micro-cli/stemcell"
)

type FakeTemplatesSpecGenerator struct {
	CreateInputs        []CreateInput
	CreateTemplatesSpec bminsup.TemplatesSpec
	CreateErr           error
	CreateCalled        bool
}

type CreateInput struct {
	DeploymentJob  bmdepl.Job
	StemcellJob    bmstemcell.Job
	DeploymentName string
	Properties     map[string]interface{}
	MbusURL        string
}

func NewFakeTemplatesSpecGenerator() *FakeTemplatesSpecGenerator {
	return &FakeTemplatesSpecGenerator{
		CreateInputs: []CreateInput{},
	}
}

func (g *FakeTemplatesSpecGenerator) Create(
	deploymentJob bmdepl.Job,
	stemcellJob bmstemcell.Job,
	deploymentName string,
	properties map[string]interface{},
	mbusURL string,
) (bminsup.TemplatesSpec, error) {
	g.CreateInputs = append(g.CreateInputs, CreateInput{
		DeploymentJob:  deploymentJob,
		StemcellJob:    stemcellJob,
		DeploymentName: deploymentName,
		Properties:     properties,
		MbusURL:        mbusURL,
	})

	g.CreateCalled = true
	return g.CreateTemplatesSpec, g.CreateErr
}

func (g *FakeTemplatesSpecGenerator) SetCreateBehavior(createTemplatesSpec bminsup.TemplatesSpec, err error) {
	g.CreateTemplatesSpec = createTemplatesSpec
	g.CreateErr = err
}
