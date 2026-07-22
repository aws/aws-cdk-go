package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::QuickSight::AssetBundleExportJob.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAssetBundleExportJobPropsMixin := awscdkcfnpropertymixins.Aws_quicksight.NewCfnAssetBundleExportJobPropsMixin(&CfnAssetBundleExportJobMixinProps{
//   	AssetBundleExportJobId: jsii.String("assetBundleExportJobId"),
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	ExportFormat: jsii.String("exportFormat"),
//   	IncludeAllDependencies: jsii.Boolean(false),
//   	IncludeFolderMembers: jsii.String("includeFolderMembers"),
//   	IncludeFolderMemberships: jsii.Boolean(false),
//   	IncludePermissions: jsii.Boolean(false),
//   	IncludeTags: jsii.Boolean(false),
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html
//
type CfnAssetBundleExportJobPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAssetBundleExportJobMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAssetBundleExportJobPropsMixin
type jsiiProxy_CfnAssetBundleExportJobPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAssetBundleExportJobPropsMixin) Props() *CfnAssetBundleExportJobMixinProps {
	var returns *CfnAssetBundleExportJobMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetBundleExportJobPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::AssetBundleExportJob`.
func NewCfnAssetBundleExportJobPropsMixin(props *CfnAssetBundleExportJobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAssetBundleExportJobPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAssetBundleExportJobPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAssetBundleExportJobPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnAssetBundleExportJobPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::AssetBundleExportJob`.
func NewCfnAssetBundleExportJobPropsMixin_Override(c CfnAssetBundleExportJobPropsMixin, props *CfnAssetBundleExportJobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnAssetBundleExportJobPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAssetBundleExportJobPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssetBundleExportJobPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnAssetBundleExportJobPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssetBundleExportJobPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnAssetBundleExportJobPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssetBundleExportJobPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAssetBundleExportJobPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

