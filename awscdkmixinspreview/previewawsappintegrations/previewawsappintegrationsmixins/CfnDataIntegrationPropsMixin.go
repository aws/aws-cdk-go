package previewawsappintegrationsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappintegrations/previewawsappintegrationsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates and persists a DataIntegration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var filters interface{}
//   var objectConfiguration interface{}
//
//   cfnDataIntegrationPropsMixin := awscdkmixinspreview.Mixins.NewCfnDataIntegrationPropsMixin(&CfnDataIntegrationMixinProps{
//   	Description: jsii.String("description"),
//   	FileConfiguration: &FileConfigurationProperty{
//   		Filters: filters,
//   		Folders: []*string{
//   			jsii.String("folders"),
//   		},
//   	},
//   	KmsKey: jsii.String("kmsKey"),
//   	Name: jsii.String("name"),
//   	ObjectConfiguration: objectConfiguration,
//   	ScheduleConfig: &ScheduleConfigProperty{
//   		FirstExecutionFrom: jsii.String("firstExecutionFrom"),
//   		Object: jsii.String("object"),
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	SourceUri: jsii.String("sourceUri"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html
//
type CfnDataIntegrationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataIntegrationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataIntegrationPropsMixin
type jsiiProxy_CfnDataIntegrationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataIntegrationPropsMixin) Props() *CfnDataIntegrationMixinProps {
	var returns *CfnDataIntegrationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataIntegrationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppIntegrations::DataIntegration`.
func NewCfnDataIntegrationPropsMixin(props *CfnDataIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataIntegrationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataIntegrationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataIntegrationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnDataIntegrationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppIntegrations::DataIntegration`.
func NewCfnDataIntegrationPropsMixin_Override(c CfnDataIntegrationPropsMixin, props *CfnDataIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnDataIntegrationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataIntegrationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataIntegrationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnDataIntegrationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataIntegrationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appintegrations.mixins.CfnDataIntegrationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataIntegrationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDataIntegrationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

