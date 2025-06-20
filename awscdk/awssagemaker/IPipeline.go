package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
)

// The interface for a SageMaker Pipeline resource.
type IPipeline interface {
	awscdk.IResource
	// Permits an IAM principal to start this pipeline execution.
	GrantStartPipelineExecution(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the pipeline.
	PipelineArn() *string
	// The name of the pipeline.
	PipelineName() *string
}

// The jsii proxy for IPipeline
type jsiiProxy_IPipeline struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPipeline) GrantStartPipelineExecution(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantStartPipelineExecutionParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantStartPipelineExecution",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPipeline) PipelineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipeline) PipelineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineName",
		&returns,
	)
	return returns
}

