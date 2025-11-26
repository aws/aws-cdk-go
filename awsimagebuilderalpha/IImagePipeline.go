package awsimagebuilderalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/internal"
)

// An EC2 Image Builder Image Pipeline.
// Experimental.
type IImagePipeline interface {
	awscdk.IResource
	// Grant custom actions to the given grantee for the image pipeline.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grants the default permissions for building an image to the provided execution role.
	// Experimental.
	GrantDefaultExecutionRolePermissions(grantee awsiam.IGrantable) *[]awsiam.Grant
	// Grant read permissions to the given grantee for the image pipeline.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permissions to the given grantee to start an execution of the image pipeline.
	// Experimental.
	GrantStartExecution(grantee awsiam.IGrantable) awsiam.Grant
	// Creates an EventBridge rule for Image Builder CVE detected events.
	// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/integ-eventbridge.html
	//
	// Experimental.
	OnCVEDetected(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Creates an EventBridge rule for Image Builder events.
	// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/integ-eventbridge.html
	//
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Creates an EventBridge rule for Image Builder image build completion events.
	// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/integ-eventbridge.html
	//
	// Experimental.
	OnImageBuildCompleted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Creates an EventBridge rule for Image Builder image build failure events.
	// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/integ-eventbridge.html
	//
	// Experimental.
	OnImageBuildFailed(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Creates an EventBridge rule for Image Builder image state change events.
	// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/integ-eventbridge.html
	//
	// Experimental.
	OnImageBuildStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Creates an EventBridge rule for Image Builder image success events.
	// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/integ-eventbridge.html
	//
	// Experimental.
	OnImageBuildSucceeded(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Creates an EventBridge rule for Image Builder image pipeline automatically disabled events.
	// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/integ-eventbridge.html
	//
	// Experimental.
	OnImagePipelineAutoDisabled(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Creates an EventBridge rule for Image Builder wait for action events.
	// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/integ-eventbridge.html
	//
	// Experimental.
	OnWaitForAction(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The ARN of the image pipeline.
	// Experimental.
	ImagePipelineArn() *string
	// The name of the image pipeline.
	// Experimental.
	ImagePipelineName() *string
}

// The jsii proxy for IImagePipeline
type jsiiProxy_IImagePipeline struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IImagePipeline) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImagePipeline) GrantDefaultExecutionRolePermissions(grantee awsiam.IGrantable) *[]awsiam.Grant {
	if err := i.validateGrantDefaultExecutionRolePermissionsParameters(grantee); err != nil {
		panic(err)
	}
	var returns *[]awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDefaultExecutionRolePermissions",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImagePipeline) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImagePipeline) GrantStartExecution(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantStartExecutionParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantStartExecution",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImagePipeline) OnCVEDetected(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnCVEDetectedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCVEDetected",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImagePipeline) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImagePipeline) OnImageBuildCompleted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnImageBuildCompletedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onImageBuildCompleted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImagePipeline) OnImageBuildFailed(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnImageBuildFailedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onImageBuildFailed",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImagePipeline) OnImageBuildStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnImageBuildStateChangeParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onImageBuildStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImagePipeline) OnImageBuildSucceeded(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnImageBuildSucceededParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onImageBuildSucceeded",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImagePipeline) OnImagePipelineAutoDisabled(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnImagePipelineAutoDisabledParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onImagePipelineAutoDisabled",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IImagePipeline) OnWaitForAction(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnWaitForActionParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onWaitForAction",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IImagePipeline) ImagePipelineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePipelineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImagePipeline) ImagePipelineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePipelineName",
		&returns,
	)
	return returns
}

