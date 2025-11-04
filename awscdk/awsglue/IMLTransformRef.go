package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MLTransform.
// Experimental.
type IMLTransformRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MLTransform resource.
	// Experimental.
	MlTransformRef() *MLTransformReference
}

// The jsii proxy for IMLTransformRef
type jsiiProxy_IMLTransformRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMLTransformRef) MlTransformRef() *MLTransformReference {
	var returns *MLTransformReference
	_jsii_.Get(
		j,
		"mlTransformRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMLTransformRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMLTransformRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

