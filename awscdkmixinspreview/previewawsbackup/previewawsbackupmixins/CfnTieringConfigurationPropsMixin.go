package previewawsbackupmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbackup/previewawsbackupmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Backup::TieringConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTieringConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnTieringConfigurationPropsMixin(&CfnTieringConfigurationMixinProps{
//   	BackupVaultName: jsii.String("backupVaultName"),
//   	ResourceSelection: []interface{}{
//   		&ResourceSelectionProperty{
//   			Resources: []*string{
//   				jsii.String("resources"),
//   			},
//   			ResourceType: jsii.String("resourceType"),
//   			TieringDownSettingsInDays: jsii.Number(123),
//   		},
//   	},
//   	TieringConfigurationName: jsii.String("tieringConfigurationName"),
//   	TieringConfigurationTags: map[string]*string{
//   		"tieringConfigurationTagsKey": jsii.String("tieringConfigurationTags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-tieringconfiguration.html
//
type CfnTieringConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTieringConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTieringConfigurationPropsMixin
type jsiiProxy_CfnTieringConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTieringConfigurationPropsMixin) Props() *CfnTieringConfigurationMixinProps {
	var returns *CfnTieringConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTieringConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Backup::TieringConfiguration`.
func NewCfnTieringConfigurationPropsMixin(props *CfnTieringConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTieringConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTieringConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTieringConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnTieringConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Backup::TieringConfiguration`.
func NewCfnTieringConfigurationPropsMixin_Override(c CfnTieringConfigurationPropsMixin, props *CfnTieringConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnTieringConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTieringConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTieringConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnTieringConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTieringConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnTieringConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTieringConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTieringConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

