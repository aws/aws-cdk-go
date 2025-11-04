package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Owner.
// Experimental.
type IOwnerRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Owner resource.
	// Experimental.
	OwnerRef() *OwnerReference
}

// The jsii proxy for IOwnerRef
type jsiiProxy_IOwnerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IOwnerRef) OwnerRef() *OwnerReference {
	var returns *OwnerReference
	_jsii_.Get(
		j,
		"ownerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOwnerRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOwnerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

