package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ViewVersion.
// Experimental.
type IViewVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ViewVersion resource.
	// Experimental.
	ViewVersionRef() *ViewVersionReference
}

// The jsii proxy for IViewVersionRef
type jsiiProxy_IViewVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IViewVersionRef) ViewVersionRef() *ViewVersionReference {
	var returns *ViewVersionReference
	_jsii_.Get(
		j,
		"viewVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IViewVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IViewVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

