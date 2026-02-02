package previewawsimagebuildermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsimagebuilder/previewawsimagebuildermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A distribution configuration allows you to specify the name and description of your output AMI, authorize other AWS account s to launch the AMI, and replicate the AMI to other AWS Regions .
//
// It also allows you to export the AMI to Amazon S3 .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var amiDistributionConfiguration interface{}
//   var containerDistributionConfiguration interface{}
//
//   cfnDistributionConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnDistributionConfigurationPropsMixin(&CfnDistributionConfigurationMixinProps{
//   	Description: jsii.String("description"),
//   	Distributions: []interface{}{
//   		&DistributionProperty{
//   			AmiDistributionConfiguration: amiDistributionConfiguration,
//   			ContainerDistributionConfiguration: containerDistributionConfiguration,
//   			FastLaunchConfigurations: []interface{}{
//   				&FastLaunchConfigurationProperty{
//   					AccountId: jsii.String("accountId"),
//   					Enabled: jsii.Boolean(false),
//   					LaunchTemplate: &FastLaunchLaunchTemplateSpecificationProperty{
//   						LaunchTemplateId: jsii.String("launchTemplateId"),
//   						LaunchTemplateName: jsii.String("launchTemplateName"),
//   						LaunchTemplateVersion: jsii.String("launchTemplateVersion"),
//   					},
//   					MaxParallelLaunches: jsii.Number(123),
//   					SnapshotConfiguration: &FastLaunchSnapshotConfigurationProperty{
//   						TargetResourceCount: jsii.Number(123),
//   					},
//   				},
//   			},
//   			LaunchTemplateConfigurations: []interface{}{
//   				&LaunchTemplateConfigurationProperty{
//   					AccountId: jsii.String("accountId"),
//   					LaunchTemplateId: jsii.String("launchTemplateId"),
//   					SetDefaultVersion: jsii.Boolean(false),
//   				},
//   			},
//   			LicenseConfigurationArns: []*string{
//   				jsii.String("licenseConfigurationArns"),
//   			},
//   			Region: jsii.String("region"),
//   			SsmParameterConfigurations: []interface{}{
//   				&SsmParameterConfigurationProperty{
//   					AmiAccountId: jsii.String("amiAccountId"),
//   					DataType: jsii.String("dataType"),
//   					ParameterName: jsii.String("parameterName"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-distributionconfiguration.html
//
type CfnDistributionConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDistributionConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDistributionConfigurationPropsMixin
type jsiiProxy_CfnDistributionConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDistributionConfigurationPropsMixin) Props() *CfnDistributionConfigurationMixinProps {
	var returns *CfnDistributionConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ImageBuilder::DistributionConfiguration`.
func NewCfnDistributionConfigurationPropsMixin(props *CfnDistributionConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDistributionConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDistributionConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDistributionConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnDistributionConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ImageBuilder::DistributionConfiguration`.
func NewCfnDistributionConfigurationPropsMixin_Override(c CfnDistributionConfigurationPropsMixin, props *CfnDistributionConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnDistributionConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDistributionConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDistributionConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnDistributionConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDistributionConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnDistributionConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDistributionConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDistributionConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

