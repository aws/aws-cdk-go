package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents source code for an AppSync Function or Resolver.
//
// Example:
//   var api eventApi
//   var lambdaDataSource appSyncLambdaDataSource
//
//
//   // Lambda data source for publish handler
//   api.AddChannelNamespace(jsii.String("lambda-ns"), &ChannelNamespaceOptions{
//   	Code: appsync.Code_FromInline(jsii.String("/* event handler code here.*/")),
//   	PublishHandlerConfig: &HandlerConfig{
//   		DataSource: lambdaDataSource,
//   	},
//   })
//
//   // Direct Lambda data source for publish handler
//   api.AddChannelNamespace(jsii.String("lambda-direct-ns"), &ChannelNamespaceOptions{
//   	PublishHandlerConfig: &HandlerConfig{
//   		DataSource: lambdaDataSource,
//   		Direct: jsii.Boolean(true),
//   	},
//   })
//
//   api.AddChannelNamespace(jsii.String("lambda-direct-async-ns"), &ChannelNamespaceOptions{
//   	PublishHandlerConfig: &HandlerConfig{
//   		DataSource: lambdaDataSource,
//   		Direct: jsii.Boolean(true),
//   		LambdaInvokeType: appsync.LambdaInvokeType_EVENT,
//   	},
//   })
//
type Code interface {
	// Bind source code to an AppSync Function or resolver.
	Bind(scope constructs.Construct) *CodeConfig
}

// The jsii proxy struct for Code
type jsiiProxy_Code struct {
	_ byte // padding
}

func NewCode_Override(c Code) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.Code",
		nil, // no parameters
		c,
	)
}

// Loads the function code from a local disk path.
func Code_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.Code",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Inline code for AppSync function.
//
// Returns: `InlineCode` with inline code.
func Code_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.Code",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Code) Bind(scope constructs.Construct) *CodeConfig {
	if err := c.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

