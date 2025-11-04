package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamingDistribution.
// Experimental.
type IStreamingDistributionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a StreamingDistribution resource.
	// Experimental.
	StreamingDistributionRef() *StreamingDistributionReference
}

// The jsii proxy for IStreamingDistributionRef
type jsiiProxy_IStreamingDistributionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IStreamingDistributionRef) StreamingDistributionRef() *StreamingDistributionReference {
	var returns *StreamingDistributionReference
	_jsii_.Get(
		j,
		"streamingDistributionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamingDistributionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamingDistributionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

