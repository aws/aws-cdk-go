package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EIP.
// Experimental.
type IEIPRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EIP resource.
	// Experimental.
	EipRef() *EIPReference
}

// The jsii proxy for IEIPRef
type jsiiProxy_IEIPRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEIPRef) EipRef() *EIPReference {
	var returns *EIPReference
	_jsii_.Get(
		j,
		"eipRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEIPRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEIPRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

