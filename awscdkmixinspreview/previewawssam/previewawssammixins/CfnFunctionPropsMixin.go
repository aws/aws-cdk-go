package previewawssammixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssam/previewawssammixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var assumeRolePolicyDocument interface{}
//
//   cfnFunctionPropsMixin := awscdkmixinspreview.Mixins.NewCfnFunctionPropsMixin(&CfnFunctionMixinProps{
//   	Architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	AssumeRolePolicyDocument: assumeRolePolicyDocument,
//   	AutoPublishAlias: jsii.String("autoPublishAlias"),
//   	AutoPublishCodeSha256: jsii.String("autoPublishCodeSha256"),
//   	CodeSigningConfigArn: jsii.String("codeSigningConfigArn"),
//   	CodeUri: jsii.String("codeUri"),
//   	DeadLetterQueue: &DeadLetterQueueProperty{
//   		TargetArn: jsii.String("targetArn"),
//   		Type: jsii.String("type"),
//   	},
//   	DeploymentPreference: &DeploymentPreferenceProperty{
//   		Alarms: []*string{
//   			jsii.String("alarms"),
//   		},
//   		Enabled: jsii.Boolean(false),
//   		Hooks: &HooksProperty{
//   			PostTraffic: jsii.String("postTraffic"),
//   			PreTraffic: jsii.String("preTraffic"),
//   		},
//   		Role: jsii.String("role"),
//   		Type: jsii.String("type"),
//   	},
//   	Description: jsii.String("description"),
//   	Environment: &FunctionEnvironmentProperty{
//   		Variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	EphemeralStorage: &EphemeralStorageProperty{
//   		Size: jsii.Number(123),
//   	},
//   	EventInvokeConfig: &EventInvokeConfigProperty{
//   		DestinationConfig: &EventInvokeDestinationConfigProperty{
//   			OnFailure: &DestinationProperty{
//   				Destination: jsii.String("destination"),
//   				Type: jsii.String("type"),
//   			},
//   			OnSuccess: &DestinationProperty{
//   				Destination: jsii.String("destination"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		MaximumEventAgeInSeconds: jsii.Number(123),
//   		MaximumRetryAttempts: jsii.Number(123),
//   	},
//   	Events: map[string]interface{}{
//   		"eventsKey": &EventSourceProperty{
//   			"properties": &AlexaSkillEventProperty{
//   				"skillId": jsii.String("skillId"),
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	FileSystemConfigs: []interface{}{
//   		&FileSystemConfigProperty{
//   			Arn: jsii.String("arn"),
//   			LocalMountPath: jsii.String("localMountPath"),
//   		},
//   	},
//   	FunctionName: jsii.String("functionName"),
//   	FunctionUrlConfig: &FunctionUrlConfigProperty{
//   		AuthType: jsii.String("authType"),
//   		Cors: jsii.String("cors"),
//   		InvokeMode: jsii.String("invokeMode"),
//   	},
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
//   	ImageUri: jsii.String("imageUri"),
//   	InlineCode: jsii.String("inlineCode"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Layers: []*string{
//   		jsii.String("layers"),
//   	},
//   	MemorySize: jsii.Number(123),
//   	PackageType: jsii.String("packageType"),
//   	PermissionsBoundary: jsii.String("permissionsBoundary"),
//   	Policies: jsii.String("policies"),
//   	ProvisionedConcurrencyConfig: &ProvisionedConcurrencyConfigProperty{
//   		ProvisionedConcurrentExecutions: jsii.String("provisionedConcurrentExecutions"),
//   	},
//   	ReservedConcurrentExecutions: jsii.Number(123),
//   	Role: jsii.String("role"),
//   	Runtime: jsii.String("runtime"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Timeout: jsii.Number(123),
//   	Tracing: jsii.String("tracing"),
//   	VersionDescription: jsii.String("versionDescription"),
//   	VpcConfig: &VpcConfigProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html
//
type CfnFunctionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFunctionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
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


// Create a mixin to apply properties to `AWS::Serverless::Function`.
func NewCfnFunctionPropsMixin(props *CfnFunctionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFunctionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFunctionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFunctionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Serverless::Function`.
func NewCfnFunctionPropsMixin_Override(c CfnFunctionPropsMixin, props *CfnFunctionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_sam.mixins.CfnFunctionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFunctionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

