package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Lambda::MicrovmImage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMicrovmImagePropsMixin := awscdkcfnpropertymixins.Aws_lambda.NewCfnMicrovmImagePropsMixin(&CfnMicrovmImageMixinProps{
//   	AdditionalOsCapabilities: []*string{
//   		jsii.String("additionalOsCapabilities"),
//   	},
//   	BaseImageArn: jsii.String("baseImageArn"),
//   	BaseImageVersion: jsii.String("baseImageVersion"),
//   	BuildRoleArn: jsii.String("buildRoleArn"),
//   	CodeArtifact: &CodeArtifactProperty{
//   		Uri: jsii.String("uri"),
//   	},
//   	CpuConfigurations: []interface{}{
//   		&CpuConfigurationProperty{
//   			Architecture: jsii.String("architecture"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	EgressNetworkConnectors: []*string{
//   		jsii.String("egressNetworkConnectors"),
//   	},
//   	EnvironmentVariables: []interface{}{
//   		&EnvironmentVariableProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Hooks: &HooksProperty{
//   		MicrovmHooks: &MicrovmHooksProperty{
//   			Resume: jsii.String("resume"),
//   			ResumeTimeoutInSeconds: jsii.Number(123),
//   			Run: jsii.String("run"),
//   			RunTimeoutInSeconds: jsii.Number(123),
//   			Suspend: jsii.String("suspend"),
//   			SuspendTimeoutInSeconds: jsii.Number(123),
//   			Terminate: jsii.String("terminate"),
//   			TerminateTimeoutInSeconds: jsii.Number(123),
//   		},
//   		MicrovmImageHooks: &MicrovmImageHooksProperty{
//   			Ready: jsii.String("ready"),
//   			ReadyTimeoutInSeconds: jsii.Number(123),
//   			Validate: jsii.String("validate"),
//   			ValidateTimeoutInSeconds: jsii.Number(123),
//   		},
//   		Port: jsii.Number(123),
//   	},
//   	Logging: &LoggingProperty{
//   		CloudWatch: &CloudWatchLoggingProperty{
//   			LogGroup: jsii.String("logGroup"),
//   			LogStream: jsii.String("logStream"),
//   		},
//   		Disabled: jsii.Boolean(false),
//   	},
//   	Name: jsii.String("name"),
//   	Resources: []interface{}{
//   		&ResourcesProperty{
//   			MinimumMemoryInMiB: jsii.Number(123),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html
//
type CfnMicrovmImagePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMicrovmImageMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMicrovmImagePropsMixin
type jsiiProxy_CfnMicrovmImagePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMicrovmImagePropsMixin) Props() *CfnMicrovmImageMixinProps {
	var returns *CfnMicrovmImageMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrovmImagePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lambda::MicrovmImage`.
func NewCfnMicrovmImagePropsMixin(props *CfnMicrovmImageMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMicrovmImagePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMicrovmImagePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMicrovmImagePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnMicrovmImagePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lambda::MicrovmImage`.
func NewCfnMicrovmImagePropsMixin_Override(c CfnMicrovmImagePropsMixin, props *CfnMicrovmImageMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnMicrovmImagePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMicrovmImagePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMicrovmImagePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnMicrovmImagePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMicrovmImagePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnMicrovmImagePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMicrovmImagePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMicrovmImagePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

