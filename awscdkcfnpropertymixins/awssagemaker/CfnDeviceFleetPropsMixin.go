package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SageMaker::DeviceFleet` resource is an Amazon SageMaker resource type that allows you to create a DeviceFleet that manages your SageMaker Edge Manager Devices.
//
// You must register your devices against the `DeviceFleet` separately.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDeviceFleetPropsMixin := awscdkcfnpropertymixins.Aws_sagemaker.NewCfnDeviceFleetPropsMixin(&CfnDeviceFleetMixinProps{
//   	Description: jsii.String("description"),
//   	DeviceFleetName: jsii.String("deviceFleetName"),
//   	OutputConfig: &EdgeOutputConfigProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		S3OutputLocation: jsii.String("s3OutputLocation"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-devicefleet.html
//
type CfnDeviceFleetPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDeviceFleetMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDeviceFleetPropsMixin
type jsiiProxy_CfnDeviceFleetPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDeviceFleetPropsMixin) Props() *CfnDeviceFleetMixinProps {
	var returns *CfnDeviceFleetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceFleetPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::DeviceFleet`.
func NewCfnDeviceFleetPropsMixin(props *CfnDeviceFleetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDeviceFleetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDeviceFleetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDeviceFleetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDeviceFleetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::DeviceFleet`.
func NewCfnDeviceFleetPropsMixin_Override(c CfnDeviceFleetPropsMixin, props *CfnDeviceFleetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDeviceFleetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDeviceFleetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeviceFleetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDeviceFleetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeviceFleetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnDeviceFleetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeviceFleetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDeviceFleetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

