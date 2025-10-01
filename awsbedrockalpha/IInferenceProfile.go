package awsbedrockalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents an Inference Profile, either created with CDK or imported.
// Experimental.
type IInferenceProfile interface {
	// Grants appropriate permissions to use the inference profile.
	//
	// Each profile type requires different permissions based on its usage pattern.
	//
	// Returns: An IAM Grant object representing the granted permissions.
	// Experimental.
	GrantProfileUsage(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the inference profile.
	// Experimental.
	InferenceProfileArn() *string
	// The unique identifier of the inference profile.
	// Experimental.
	InferenceProfileId() *string
	// The type of inference profile.
	// Experimental.
	Type() InferenceProfileType
}

// The jsii proxy for IInferenceProfile
type jsiiProxy_IInferenceProfile struct {
	_ byte // padding
}

func (i *jsiiProxy_IInferenceProfile) GrantProfileUsage(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantProfileUsageParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantProfileUsage",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IInferenceProfile) InferenceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInferenceProfile) InferenceProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInferenceProfile) Type() InferenceProfileType {
	var returns InferenceProfileType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

