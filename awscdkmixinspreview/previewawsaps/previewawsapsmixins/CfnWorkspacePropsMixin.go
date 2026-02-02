package previewawsapsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsaps/previewawsapsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Amazon Managed Service for Prometheus workspace is a logical and isolated Prometheus server dedicated to ingesting, storing, and querying your Prometheus-compatible metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkspacePropsMixin := awscdkmixinspreview.Mixins.NewCfnWorkspacePropsMixin(&CfnWorkspaceMixinProps{
//   	AlertManagerDefinition: jsii.String("alertManagerDefinition"),
//   	Alias: jsii.String("alias"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		LogGroupArn: jsii.String("logGroupArn"),
//   	},
//   	QueryLoggingConfiguration: &QueryLoggingConfigurationProperty{
//   		Destinations: []interface{}{
//   			&LoggingDestinationProperty{
//   				CloudWatchLogs: &CloudWatchLogDestinationProperty{
//   					LogGroupArn: jsii.String("logGroupArn"),
//   				},
//   				Filters: &LoggingFilterProperty{
//   					QspThreshold: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkspaceConfiguration: &WorkspaceConfigurationProperty{
//   		LimitsPerLabelSets: []interface{}{
//   			&LimitsPerLabelSetProperty{
//   				LabelSet: []interface{}{
//   					&LabelProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Limits: &LimitsPerLabelSetEntryProperty{
//   					MaxSeries: jsii.Number(123),
//   				},
//   			},
//   		},
//   		RetentionPeriodInDays: jsii.Number(123),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html
//
type CfnWorkspacePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWorkspaceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkspacePropsMixin
type jsiiProxy_CfnWorkspacePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWorkspacePropsMixin) Props() *CfnWorkspaceMixinProps {
	var returns *CfnWorkspaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspacePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::APS::Workspace`.
func NewCfnWorkspacePropsMixin(props *CfnWorkspaceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWorkspacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkspacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkspacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::APS::Workspace`.
func NewCfnWorkspacePropsMixin_Override(c CfnWorkspacePropsMixin, props *CfnWorkspaceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWorkspacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkspacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkspacePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWorkspacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

