package awsbedrockalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Represents a Prompt, either created with CDK or imported.
// Experimental.
type IPrompt interface {
	awscdk.IResource
	// Grant the given identity permissions to get the prompt.
	// Experimental.
	GrantGet(grantee awsiam.IGrantable) awsiam.Grant
	// Optional KMS encryption key associated with this prompt.
	// Experimental.
	KmsKey() awskms.IKey
	// The ARN of the prompt.
	//
	// Example:
	//   "arn:aws:bedrock:us-east-1:123456789012:prompt/PROMPT12345"
	//
	// Experimental.
	PromptArn() *string
	// The ID of the prompt.
	//
	// Example:
	//   "PROMPT12345"
	//
	// Experimental.
	PromptId() *string
	// The version of the prompt.
	// Default: "DRAFT".
	//
	// Experimental.
	PromptVersion() *string
}

// The jsii proxy for IPrompt
type jsiiProxy_IPrompt struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPrompt) GrantGet(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantGetParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantGet",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPrompt) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrompt) PromptArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrompt) PromptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrompt) PromptVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptVersion",
		&returns,
	)
	return returns
}

