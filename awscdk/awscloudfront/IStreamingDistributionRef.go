package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamingDistribution.
// Experimental.
type IStreamingDistributionRef interface {
	constructs.IConstruct
	// A reference to a StreamingDistribution resource.
	// Experimental.
	StreamingDistributionRef() *StreamingDistributionReference
}

// The jsii proxy for IStreamingDistributionRef
type jsiiProxy_IStreamingDistributionRef struct {
	internal.Type__constructsIConstruct
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

