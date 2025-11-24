package mixinsawssynthetics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awssynthetics/mixinsawssynthetics/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates a canary.
//
// Canaries are scripts that monitor your endpoints and APIs from the outside-in. Canaries help you check the availability and latency of your web services and troubleshoot anomalies by investigating load time data, screenshots of the UI, logs, and metrics. You can set up a canary to run continuously or just once.
//
// Canaries are automated scripts that run at specified intervals against an endpoint. They include Python or Node.js code to create a Lambda function. This code needs to be packaged in a certain way, depending on the language. For more information, see [Writing a canary script](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html) .
//
// To create canaries, you must have the `CloudWatchSyntheticsFullAccess` policy. If you are creating a new IAM role for the canary, you also need the the `iam:CreateRole` , `iam:CreatePolicy` and `iam:AttachRolePolicy` permissions. For more information, see [Necessary Roles and Permissions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Roles) .
//
// Do not include secrets or proprietary information in your canary names. The canary name makes up part of the Amazon Resource Name (ARN) for the canary, and the ARN is included in outbound calls over the internet. For more information, see [Security Considerations for Synthetics Canaries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCanaryPropsMixin := awscdkmixinspreview.Mixins.NewCfnCanaryPropsMixin(&CfnCanaryMixinProps{
//   	ArtifactConfig: &ArtifactConfigProperty{
//   		S3Encryption: &S3EncryptionProperty{
//   			EncryptionMode: jsii.String("encryptionMode"),
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   	ArtifactS3Location: jsii.String("artifactS3Location"),
//   	BrowserConfigs: []interface{}{
//   		&BrowserConfigProperty{
//   			BrowserType: jsii.String("browserType"),
//   		},
//   	},
//   	Code: &CodeProperty{
//   		BlueprintTypes: []*string{
//   			jsii.String("blueprintTypes"),
//   		},
//   		Dependencies: []interface{}{
//   			&DependencyProperty{
//   				Reference: jsii.String("reference"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Handler: jsii.String("handler"),
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
//   		S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   		Script: jsii.String("script"),
//   		SourceLocationArn: jsii.String("sourceLocationArn"),
//   	},
//   	DeleteLambdaResourcesOnCanaryDeletion: jsii.Boolean(false),
//   	DryRunAndUpdate: jsii.Boolean(false),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	FailureRetentionPeriod: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	ProvisionedResourceCleanup: jsii.String("provisionedResourceCleanup"),
//   	ResourcesToReplicateTags: []*string{
//   		jsii.String("resourcesToReplicateTags"),
//   	},
//   	RunConfig: &RunConfigProperty{
//   		ActiveTracing: jsii.Boolean(false),
//   		EnvironmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		EphemeralStorage: jsii.Number(123),
//   		MemoryInMb: jsii.Number(123),
//   		TimeoutInSeconds: jsii.Number(123),
//   	},
//   	RuntimeVersion: jsii.String("runtimeVersion"),
//   	Schedule: &ScheduleProperty{
//   		DurationInSeconds: jsii.String("durationInSeconds"),
//   		Expression: jsii.String("expression"),
//   		RetryConfig: &RetryConfigProperty{
//   			MaxRetries: jsii.Number(123),
//   		},
//   	},
//   	StartCanaryAfterCreation: jsii.Boolean(false),
//   	SuccessRetentionPeriod: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VisualReference: &VisualReferenceProperty{
//   		BaseCanaryRunId: jsii.String("baseCanaryRunId"),
//   		BaseScreenshots: []interface{}{
//   			&BaseScreenshotProperty{
//   				IgnoreCoordinates: []*string{
//   					jsii.String("ignoreCoordinates"),
//   				},
//   				ScreenshotName: jsii.String("screenshotName"),
//   			},
//   		},
//   		BrowserType: jsii.String("browserType"),
//   	},
//   	VisualReferences: []interface{}{
//   		&VisualReferenceProperty{
//   			BaseCanaryRunId: jsii.String("baseCanaryRunId"),
//   			BaseScreenshots: []interface{}{
//   				&BaseScreenshotProperty{
//   					IgnoreCoordinates: []*string{
//   						jsii.String("ignoreCoordinates"),
//   					},
//   					ScreenshotName: jsii.String("screenshotName"),
//   				},
//   			},
//   			BrowserType: jsii.String("browserType"),
//   		},
//   	},
//   	VpcConfig: &VPCConfigProperty{
//   		Ipv6AllowedForDualStack: jsii.Boolean(false),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html
//
type CfnCanaryPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCanaryMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCanaryPropsMixin
type jsiiProxy_CfnCanaryPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCanaryPropsMixin) Props() *CfnCanaryMixinProps {
	var returns *CfnCanaryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanaryPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Synthetics::Canary`.
func NewCfnCanaryPropsMixin(props *CfnCanaryMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCanaryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCanaryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCanaryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Synthetics::Canary`.
func NewCfnCanaryPropsMixin_Override(c CfnCanaryPropsMixin, props *CfnCanaryMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCanaryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCanaryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCanaryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_synthetics.mixins.CfnCanaryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCanaryPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCanaryPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

