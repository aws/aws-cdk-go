package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions/internal"
)

// Represents a Step Functions Activity https://docs.aws.amazon.com/step-functions/latest/dg/concepts-activities.html.
type IActivity interface {
	awscdk.IResource
	// The ARN of the activity.
	ActivityArn() *string
	// The name of the activity.
	ActivityName() *string
}

// The jsii proxy for IActivity
type jsiiProxy_IActivity struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IActivity) ActivityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActivity) ActivityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activityName",
		&returns,
	)
	return returns
}

