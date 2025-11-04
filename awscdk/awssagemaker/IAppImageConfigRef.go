package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppImageConfig.
// Experimental.
type IAppImageConfigRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AppImageConfig resource.
	// Experimental.
	AppImageConfigRef() *AppImageConfigReference
}

// The jsii proxy for IAppImageConfigRef
type jsiiProxy_IAppImageConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAppImageConfigRef) AppImageConfigRef() *AppImageConfigReference {
	var returns *AppImageConfigReference
	_jsii_.Get(
		j,
		"appImageConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppImageConfigRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppImageConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

