package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Amazon Bedrock intelligent prompt routing provides a single serverless endpoint for efficiently routing requests between different foundational models within the same model family.
//
// It can help you optimize for response quality and cost.
//
// Intelligent prompt routing predicts the performance of each model for each request,
// and dynamically routes each request to the model that it predicts is most likely
// to give the desired response at the lowest cost.
//
// Example:
//   // Create a prompt router for intelligent model selection
//   promptRouter := bedrock.PromptRouter_FromDefaultId(bedrock.DefaultPromptRouterIdentifier_ANTHROPIC_CLAUDE_V1(), jsii.String("us-east-1"))
//
//   // Use the prompt router with a prompt variant
//   variant := bedrock.PromptVariant_Text(&TextPromptVariantProps{
//   	VariantName: jsii.String("variant1"),
//   	PromptText: jsii.String("What is the capital of France?"),
//   	Model: promptRouter,
//   })
//
//   bedrock.NewPrompt(this, jsii.String("Prompt"), &PromptProps{
//   	PromptName: jsii.String("prompt-router-test"),
//   	Variants: []IPromptVariant{
//   		variant,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-routing.html
//
// Experimental.
type PromptRouter interface {
	IBedrockInvokable
	IPromptRouter
	// The ARN used for invoking this prompt router.
	//
	// This equals to the promptRouterArn property, useful for implementing IBedrockInvokable interface.
	// Experimental.
	InvokableArn() *string
	// The ARN of the prompt router.
	// Experimental.
	PromptRouterArn() *string
	// The ID of the prompt router.
	// Experimental.
	PromptRouterId() *string
	// The inference endpoints (cross-region profiles) that this router will route to.
	//
	// These are created automatically based on the routing models and region.
	// Experimental.
	RoutingEndpoints() *[]IBedrockInvokable
	// Grants the necessary permissions to invoke this prompt router and all its routing endpoints.
	//
	// This method grants permissions to:
	// - Get prompt router details (bedrock:GetPromptRouter)
	// - Invoke models through the router (bedrock:InvokeModel)
	// - Use all underlying models and cross-region profiles
	// [disable-awslint:no-grants].
	//
	// Returns: An IAM Grant object representing the granted permissions.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for PromptRouter
type jsiiProxy_PromptRouter struct {
	jsiiProxy_IBedrockInvokable
	jsiiProxy_IPromptRouter
}

func (j *jsiiProxy_PromptRouter) InvokableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptRouter) PromptRouterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptRouterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptRouter) PromptRouterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptRouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptRouter) RoutingEndpoints() *[]IBedrockInvokable {
	var returns *[]IBedrockInvokable
	_jsii_.Get(
		j,
		"routingEndpoints",
		&returns,
	)
	return returns
}


// Experimental.
func NewPromptRouter(props *PromptRouterProps, region *string) PromptRouter {
	_init_.Initialize()

	if err := validateNewPromptRouterParameters(props, region); err != nil {
		panic(err)
	}
	j := jsiiProxy_PromptRouter{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.PromptRouter",
		[]interface{}{props, region},
		&j,
	)

	return &j
}

// Experimental.
func NewPromptRouter_Override(p PromptRouter, props *PromptRouterProps, region *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.PromptRouter",
		[]interface{}{props, region},
		p,
	)
}

// Creates a PromptRouter from a default router identifier.
//
// Returns: A new PromptRouter instance configured with the default settings.
// Experimental.
func PromptRouter_FromDefaultId(defaultRouter DefaultPromptRouterIdentifier, region *string) PromptRouter {
	_init_.Initialize()

	if err := validatePromptRouter_FromDefaultIdParameters(defaultRouter, region); err != nil {
		panic(err)
	}
	var returns PromptRouter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptRouter",
		"fromDefaultId",
		[]interface{}{defaultRouter, region},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PromptRouter) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := p.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		p,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

