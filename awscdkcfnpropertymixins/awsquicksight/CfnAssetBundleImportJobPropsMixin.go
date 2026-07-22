package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::QuickSight::AssetBundleImportJob.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAssetBundleImportJobPropsMixin := awscdkcfnpropertymixins.Aws_quicksight.NewCfnAssetBundleImportJobPropsMixin(&CfnAssetBundleImportJobMixinProps{
//   	AssetBundleImportJobId: jsii.String("assetBundleImportJobId"),
//   	AssetBundleImportSource: &AssetBundleImportSourceDescriptionProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	FailureAction: jsii.String("failureAction"),
//   	OverrideValidationStrategy: &AssetBundleImportJobOverrideValidationStrategyProperty{
//   		StrictModeForAllResources: jsii.Boolean(false),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleimportjob.html
//
type CfnAssetBundleImportJobPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAssetBundleImportJobMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAssetBundleImportJobPropsMixin
type jsiiProxy_CfnAssetBundleImportJobPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAssetBundleImportJobPropsMixin) Props() *CfnAssetBundleImportJobMixinProps {
	var returns *CfnAssetBundleImportJobMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetBundleImportJobPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::AssetBundleImportJob`.
func NewCfnAssetBundleImportJobPropsMixin(props *CfnAssetBundleImportJobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAssetBundleImportJobPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAssetBundleImportJobPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAssetBundleImportJobPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnAssetBundleImportJobPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::AssetBundleImportJob`.
func NewCfnAssetBundleImportJobPropsMixin_Override(c CfnAssetBundleImportJobPropsMixin, props *CfnAssetBundleImportJobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnAssetBundleImportJobPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAssetBundleImportJobPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssetBundleImportJobPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnAssetBundleImportJobPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssetBundleImportJobPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnAssetBundleImportJobPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssetBundleImportJobPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAssetBundleImportJobPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

