package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks/internal"
)

// Task to train a machine learning model using Amazon SageMaker.
type ISageMakerTask interface {
	awsiam.IGrantable
}

// The jsii proxy for ISageMakerTask
type jsiiProxy_ISageMakerTask struct {
	internal.Type__awsiamIGrantable
}

