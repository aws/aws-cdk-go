package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
)

// Task to train a machine learning model using Amazon SageMaker.
// Experimental.
type ISageMakerTask interface {
	awsiam.IGrantable
}

// The jsii proxy for ISageMakerTask
type jsiiProxy_ISageMakerTask struct {
	internal.Type__awsiamIGrantable
}

