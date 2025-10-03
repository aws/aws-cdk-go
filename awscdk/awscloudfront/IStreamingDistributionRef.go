package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamingDistribution.
// Experimental.
type IStreamingDistributionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStreamingDistributionRef
type jsiiProxy_IStreamingDistributionRef struct {
	internal.Type__constructsIConstruct
}

