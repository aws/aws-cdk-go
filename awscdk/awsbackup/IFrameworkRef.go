package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Framework.
// Experimental.
type IFrameworkRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Framework resource.
	// Experimental.
	FrameworkRef() *FrameworkReference
}

// The jsii proxy for IFrameworkRef
type jsiiProxy_IFrameworkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IFrameworkRef) FrameworkRef() *FrameworkReference {
	var returns *FrameworkReference
	_jsii_.Get(
		j,
		"frameworkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFrameworkRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFrameworkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

