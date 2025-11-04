package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PushTemplate.
// Experimental.
type IPushTemplateRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PushTemplate resource.
	// Experimental.
	PushTemplateRef() *PushTemplateReference
}

// The jsii proxy for IPushTemplateRef
type jsiiProxy_IPushTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPushTemplateRef) PushTemplateRef() *PushTemplateReference {
	var returns *PushTemplateReference
	_jsii_.Get(
		j,
		"pushTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPushTemplateRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPushTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

