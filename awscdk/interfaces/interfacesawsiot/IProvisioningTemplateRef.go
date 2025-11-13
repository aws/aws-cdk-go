package interfacesawsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProvisioningTemplate.
// Experimental.
type IProvisioningTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ProvisioningTemplate resource.
	// Experimental.
	ProvisioningTemplateRef() *ProvisioningTemplateReference
}

// The jsii proxy for IProvisioningTemplateRef
type jsiiProxy_IProvisioningTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IProvisioningTemplateRef) ProvisioningTemplateRef() *ProvisioningTemplateReference {
	var returns *ProvisioningTemplateReference
	_jsii_.Get(
		j,
		"provisioningTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProvisioningTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProvisioningTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

