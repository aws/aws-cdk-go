package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProcessingJob.
// Experimental.
type IProcessingJobRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProcessingJobRef
type jsiiProxy_IProcessingJobRef struct {
	internal.Type__constructsIConstruct
}

