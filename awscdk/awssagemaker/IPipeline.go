package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker"
	"github.com/aws/constructs-go/constructs/v10"
)

// The interface for a SageMaker Pipeline resource.
type IPipeline interface {
	interfacesawssagemaker.IPipelineRef
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
	internal.Type__interfacesawssagemakerIPipelineRef
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

func (i *jsiiProxy_IPipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IPipeline) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
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

func (j *jsiiProxy_IPipeline) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipeline) PipelineRef() *interfacesawssagemaker.PipelineReference {
	var returns *interfacesawssagemaker.PipelineReference
	_jsii_.Get(
		j,
		"pipelineRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipeline) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

