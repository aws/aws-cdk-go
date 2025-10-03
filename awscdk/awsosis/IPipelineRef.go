package awsosis

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsosis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Pipeline.
// Experimental.
type IPipelineRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPipelineRef
type jsiiProxy_IPipelineRef struct {
	internal.Type__constructsIConstruct
}

