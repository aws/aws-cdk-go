package awsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppBlock.
// Experimental.
type IAppBlockRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AppBlock resource.
	// Experimental.
	AppBlockRef() *AppBlockReference
}

// The jsii proxy for IAppBlockRef
type jsiiProxy_IAppBlockRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAppBlockRef) AppBlockRef() *AppBlockReference {
	var returns *AppBlockReference
	_jsii_.Get(
		j,
		"appBlockRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppBlockRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppBlockRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

