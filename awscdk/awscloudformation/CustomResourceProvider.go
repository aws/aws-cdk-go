package awscloudformation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Represents a provider for an AWS CloudFormation custom resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//
//   customResourceProvider := awscdk.Aws_cloudformation.CustomResourceProvider_FromLambda(function_)
//
// Deprecated: use core.CustomResource instead
type CustomResourceProvider interface {
	ICustomResourceProvider
	// the ServiceToken which contains the ARN for this provider.
	// Deprecated: use core.CustomResource instead
	ServiceToken() *string
	// Called when this provider is used by a `CustomResource`.
	// Deprecated: use core.CustomResource instead
	Bind(_arg awscdk.Construct) *CustomResourceProviderConfig
}

// The jsii proxy struct for CustomResourceProvider
type jsiiProxy_CustomResourceProvider struct {
	jsiiProxy_ICustomResourceProvider
}

func (j *jsiiProxy_CustomResourceProvider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}


// The Lambda provider that implements this custom resource.
//
// We recommend using a lambda.SingletonFunction for this.
// Deprecated: use core.CustomResource instead
func CustomResourceProvider_FromLambda(handler awslambda.IFunction) CustomResourceProvider {
	_init_.Initialize()

	if err := validateCustomResourceProvider_FromLambdaParameters(handler); err != nil {
		panic(err)
	}
	var returns CustomResourceProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudformation.CustomResourceProvider",
		"fromLambda",
		[]interface{}{handler},
		&returns,
	)

	return returns
}

// The SNS Topic for the provider that implements this custom resource.
// Deprecated: use core.CustomResource instead
func CustomResourceProvider_FromTopic(topic awssns.ITopic) CustomResourceProvider {
	_init_.Initialize()

	if err := validateCustomResourceProvider_FromTopicParameters(topic); err != nil {
		panic(err)
	}
	var returns CustomResourceProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudformation.CustomResourceProvider",
		"fromTopic",
		[]interface{}{topic},
		&returns,
	)

	return returns
}

// Use AWS Lambda as a provider.
// Deprecated: use `fromLambda`.
func CustomResourceProvider_Lambda(handler awslambda.IFunction) CustomResourceProvider {
	_init_.Initialize()

	if err := validateCustomResourceProvider_LambdaParameters(handler); err != nil {
		panic(err)
	}
	var returns CustomResourceProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudformation.CustomResourceProvider",
		"lambda",
		[]interface{}{handler},
		&returns,
	)

	return returns
}

// Use an SNS topic as the provider.
// Deprecated: use `fromTopic`.
func CustomResourceProvider_Topic(topic awssns.ITopic) CustomResourceProvider {
	_init_.Initialize()

	if err := validateCustomResourceProvider_TopicParameters(topic); err != nil {
		panic(err)
	}
	var returns CustomResourceProvider

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudformation.CustomResourceProvider",
		"topic",
		[]interface{}{topic},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResourceProvider) Bind(_arg awscdk.Construct) *CustomResourceProviderConfig {
	if err := c.validateBindParameters(_arg); err != nil {
		panic(err)
	}
	var returns *CustomResourceProviderConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

