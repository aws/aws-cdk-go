package awsmgn

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmgn/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Account level Launch Configuration Template for AWS Application Migration Service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLaunchConfigurationTemplatePropsMixin := awscdkcfnpropertymixins.Aws_mgn.NewCfnLaunchConfigurationTemplatePropsMixin(&CfnLaunchConfigurationTemplateMixinProps{
//   	AssociatePublicIpAddress: jsii.Boolean(false),
//   	BootMode: jsii.String("bootMode"),
//   	CopyPrivateIp: jsii.Boolean(false),
//   	CopyTags: jsii.Boolean(false),
//   	EnableMapAutoTagging: jsii.Boolean(false),
//   	EnableParametersEncryption: jsii.Boolean(false),
//   	LargeVolumeConf: &LaunchTemplateDiskConfProperty{
//   		Iops: jsii.Number(123),
//   		Throughput: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	LaunchDisposition: jsii.String("launchDisposition"),
//   	Licensing: &LicensingProperty{
//   		OsByol: jsii.Boolean(false),
//   	},
//   	MapAutoTaggingMpeId: jsii.String("mapAutoTaggingMpeId"),
//   	ParametersEncryptionKey: jsii.String("parametersEncryptionKey"),
//   	PostLaunchActions: &PostLaunchActionsProperty{
//   		CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   		Deployment: jsii.String("deployment"),
//   		S3LogBucket: jsii.String("s3LogBucket"),
//   		S3OutputKeyPrefix: jsii.String("s3OutputKeyPrefix"),
//   		SsmDocuments: []interface{}{
//   			&SsmDocumentProperty{
//   				ActionName: jsii.String("actionName"),
//   				ExternalParameters: map[string]interface{}{
//   					"externalParametersKey": &SsmExternalParameterProperty{
//   						"dynamicPath": jsii.String("dynamicPath"),
//   					},
//   				},
//   				MustSucceedForCutover: jsii.Boolean(false),
//   				Parameters: map[string]interface{}{
//   					"parametersKey": []interface{}{
//   						&SsmParameterStoreParameterProperty{
//   							"parameterName": jsii.String("parameterName"),
//   							"parameterType": jsii.String("parameterType"),
//   						},
//   					},
//   				},
//   				SsmDocumentName: jsii.String("ssmDocumentName"),
//   				TimeoutSeconds: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SmallVolumeConf: &LaunchTemplateDiskConfProperty{
//   		Iops: jsii.Number(123),
//   		Throughput: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	SmallVolumeMaxSize: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetInstanceTypeRightSizingMethod: jsii.String("targetInstanceTypeRightSizingMethod"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html
//
type CfnLaunchConfigurationTemplatePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLaunchConfigurationTemplateMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLaunchConfigurationTemplatePropsMixin
type jsiiProxy_CfnLaunchConfigurationTemplatePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLaunchConfigurationTemplatePropsMixin) Props() *CfnLaunchConfigurationTemplateMixinProps {
	var returns *CfnLaunchConfigurationTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfigurationTemplatePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MGN::LaunchConfigurationTemplate`.
func NewCfnLaunchConfigurationTemplatePropsMixin(props *CfnLaunchConfigurationTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLaunchConfigurationTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLaunchConfigurationTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLaunchConfigurationTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnLaunchConfigurationTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MGN::LaunchConfigurationTemplate`.
func NewCfnLaunchConfigurationTemplatePropsMixin_Override(c CfnLaunchConfigurationTemplatePropsMixin, props *CfnLaunchConfigurationTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnLaunchConfigurationTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLaunchConfigurationTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLaunchConfigurationTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnLaunchConfigurationTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLaunchConfigurationTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mgn.CfnLaunchConfigurationTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLaunchConfigurationTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLaunchConfigurationTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

