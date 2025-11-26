package previewawslambdamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslambda/previewawslambdamixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lambda::Function` resource creates a Lambda function.
//
// To create a function, you need a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) and an [execution role](https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html) . The deployment package is a .zip file archive or container image that contains your function code. The execution role grants the function permission to use AWS services, such as Amazon CloudWatch Logs for log streaming and AWS X-Ray for request tracing.
//
// You set the package type to `Image` if the deployment package is a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html) . For these functions, include the URI of the container image in the Amazon ECR registry in the [`ImageUri` property of the `Code` property](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-imageuri) . You do not need to specify the handler and runtime properties.
//
// You set the package type to `Zip` if the deployment package is a [.zip file archive](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html#gettingstarted-package-zip) . For these functions, specify the Amazon S3 location of your .zip file in the `Code` property. Alternatively, for Node.js and Python functions, you can define your function inline in the [`ZipFile` property of the `Code` property](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-zipfile) . In both cases, you must also specify the handler and runtime properties.
//
// You can use [code signing](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html) if your deployment package is a .zip file archive. To enable code signing for this function, specify the ARN of a code-signing configuration. When a user attempts to deploy a code package with `UpdateFunctionCode` , Lambda checks that the code package has a valid signature from a trusted publisher. The code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.
//
// When you update a `AWS::Lambda::Function` resource, CloudFormation calls the [UpdateFunctionConfiguration](https://docs.aws.amazon.com/lambda/latest/api/API_UpdateFunctionConfiguration.html) and [UpdateFunctionCode](https://docs.aws.amazon.com/lambda/latest/api/API_UpdateFunctionCode.html) Lambda APIs under the hood. Because these calls happen sequentially, and invocations can happen between these calls, your function may encounter errors in the time between the calls. For example, if you remove an environment variable, and the code that references that environment variable in the same CloudFormation update, you may see invocation errors related to a missing environment variable. To work around this, you can invoke your function against a version or alias by default, rather than the `$LATEST` version.
//
// Note that you configure [provisioned concurrency](https://docs.aws.amazon.com/lambda/latest/dg/provisioned-concurrency.html) on a `AWS::Lambda::Version` or a `AWS::Lambda::Alias` .
//
// For a complete introduction to Lambda functions, see [What is Lambda?](https://docs.aws.amazon.com/lambda/latest/dg/lambda-welcome.html) in the *Lambda developer guide.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFunctionPropsMixin := awscdkmixinspreview.Mixins.NewCfnFunctionPropsMixin(&CfnFunctionMixinProps{
//   	Architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	Code: &CodeProperty{
//   		ImageUri: jsii.String("imageUri"),
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
//   		S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   		SourceKmsKeyArn: jsii.String("sourceKmsKeyArn"),
//   		ZipFile: jsii.String("zipFile"),
//   	},
//   	CodeSigningConfigArn: jsii.String("codeSigningConfigArn"),
//   	DeadLetterConfig: &DeadLetterConfigProperty{
//   		TargetArn: jsii.String("targetArn"),
//   	},
//   	Description: jsii.String("description"),
//   	Environment: &EnvironmentProperty{
//   		Variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	EphemeralStorage: &EphemeralStorageProperty{
//   		Size: jsii.Number(123),
//   	},
//   	FileSystemConfigs: []interface{}{
//   		&FileSystemConfigProperty{
//   			Arn: jsii.String("arn"),
//   			LocalMountPath: jsii.String("localMountPath"),
//   		},
//   	},
//   	FunctionName: jsii.String("functionName"),
//   	Handler: jsii.String("handler"),
//   	ImageConfig: &ImageConfigProperty{
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		EntryPoint: []*string{
//   			jsii.String("entryPoint"),
//   		},
//   		WorkingDirectory: jsii.String("workingDirectory"),
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Layers: []*string{
//   		jsii.String("layers"),
//   	},
//   	LoggingConfig: &LoggingConfigProperty{
//   		ApplicationLogLevel: jsii.String("applicationLogLevel"),
//   		LogFormat: jsii.String("logFormat"),
//   		LogGroup: jsii.String("logGroup"),
//   		SystemLogLevel: jsii.String("systemLogLevel"),
//   	},
//   	MemorySize: jsii.Number(123),
//   	PackageType: jsii.String("packageType"),
//   	RecursiveLoop: jsii.String("recursiveLoop"),
//   	ReservedConcurrentExecutions: jsii.Number(123),
//   	Role: jsii.String("role"),
//   	Runtime: jsii.String("runtime"),
//   	RuntimeManagementConfig: &RuntimeManagementConfigProperty{
//   		RuntimeVersionArn: jsii.String("runtimeVersionArn"),
//   		UpdateRuntimeOn: jsii.String("updateRuntimeOn"),
//   	},
//   	SnapStart: &SnapStartProperty{
//   		ApplyOn: jsii.String("applyOn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TenancyConfig: &TenancyConfigProperty{
//   		TenantIsolationMode: jsii.String("tenantIsolationMode"),
//   	},
//   	Timeout: jsii.Number(123),
//   	TracingConfig: &TracingConfigProperty{
//   		Mode: jsii.String("mode"),
//   	},
//   	VpcConfig: &VpcConfigProperty{
//   		Ipv6AllowedForDualStack: jsii.Boolean(false),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html
//
type CfnFunctionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFunctionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFunctionPropsMixin
type jsiiProxy_CfnFunctionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFunctionPropsMixin) Props() *CfnFunctionMixinProps {
	var returns *CfnFunctionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lambda::Function`.
func NewCfnFunctionPropsMixin(props *CfnFunctionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFunctionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFunctionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFunctionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lambda::Function`.
func NewCfnFunctionPropsMixin_Override(c CfnFunctionPropsMixin, props *CfnFunctionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFunctionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFunctionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFunctionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnFunctionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFunctionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

