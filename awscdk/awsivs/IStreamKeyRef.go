package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamKey.
// Experimental.
type IStreamKeyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a StreamKey resource.
	// Experimental.
	StreamKeyRef() *StreamKeyReference
}

// The jsii proxy for IStreamKeyRef
type jsiiProxy_IStreamKeyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IStreamKeyRef) StreamKeyRef() *StreamKeyReference {
	var returns *StreamKeyReference
	_jsii_.Get(
		j,
		"streamKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamKeyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamKeyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

