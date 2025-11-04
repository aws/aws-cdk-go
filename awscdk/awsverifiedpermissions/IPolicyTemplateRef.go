package awsverifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsverifiedpermissions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyTemplate.
// Experimental.
type IPolicyTemplateRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PolicyTemplate resource.
	// Experimental.
	PolicyTemplateRef() *PolicyTemplateReference
}

// The jsii proxy for IPolicyTemplateRef
type jsiiProxy_IPolicyTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPolicyTemplateRef) PolicyTemplateRef() *PolicyTemplateReference {
	var returns *PolicyTemplateReference
	_jsii_.Get(
		j,
		"policyTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyTemplateRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

