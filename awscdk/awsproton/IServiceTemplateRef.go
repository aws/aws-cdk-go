package awsproton

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsproton/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceTemplate.
// Experimental.
type IServiceTemplateRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ServiceTemplate resource.
	// Experimental.
	ServiceTemplateRef() *ServiceTemplateReference
}

// The jsii proxy for IServiceTemplateRef
type jsiiProxy_IServiceTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IServiceTemplateRef) ServiceTemplateRef() *ServiceTemplateReference {
	var returns *ServiceTemplateReference
	_jsii_.Get(
		j,
		"serviceTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceTemplateRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

