package awsshield

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsshield/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DRTAccess.
// Experimental.
type IDRTAccessRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DRTAccess resource.
	// Experimental.
	DrtAccessRef() *DRTAccessReference
}

// The jsii proxy for IDRTAccessRef
type jsiiProxy_IDRTAccessRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDRTAccessRef) DrtAccessRef() *DRTAccessReference {
	var returns *DRTAccessReference
	_jsii_.Get(
		j,
		"drtAccessRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDRTAccessRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDRTAccessRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

