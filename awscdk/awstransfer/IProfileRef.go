package awstransfer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Profile.
// Experimental.
type IProfileRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Profile resource.
	// Experimental.
	ProfileRef() *ProfileReference
}

// The jsii proxy for IProfileRef
type jsiiProxy_IProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IProfileRef) ProfileRef() *ProfileReference {
	var returns *ProfileReference
	_jsii_.Get(
		j,
		"profileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfileRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

