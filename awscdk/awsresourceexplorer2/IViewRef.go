package awsresourceexplorer2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsresourceexplorer2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a View.
// Experimental.
type IViewRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a View resource.
	// Experimental.
	ViewRef() *ViewReference
}

// The jsii proxy for IViewRef
type jsiiProxy_IViewRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IViewRef) ViewRef() *ViewReference {
	var returns *ViewReference
	_jsii_.Get(
		j,
		"viewRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IViewRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IViewRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

