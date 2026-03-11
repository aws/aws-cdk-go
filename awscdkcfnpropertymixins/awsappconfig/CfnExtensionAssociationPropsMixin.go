package awsappconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// When you create an extension or configure an AWS authored extension, you associate the extension with an AWS AppConfig application, environment, or configuration profile.
//
// For example, you can choose to run the `AWS AppConfig deployment events to Amazon SNS` AWS authored extension and receive notifications on an Amazon SNS topic anytime a configuration deployment is started for a specific application. Defining which extension to associate with an AWS AppConfig resource is called an *extension association* . An extension association is a specified relationship between an extension and an AWS AppConfig resource, such as an application or a configuration profile. For more information about extensions and associations, see [Extending workflows](https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html) in the *AWS AppConfig User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnExtensionAssociationPropsMixin := awscdkcfnpropertymixins.Aws_appconfig.NewCfnExtensionAssociationPropsMixin(&CfnExtensionAssociationMixinProps{
//   	ExtensionIdentifier: jsii.String("extensionIdentifier"),
//   	ExtensionVersionNumber: jsii.Number(123),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extensionassociation.html
//
type CfnExtensionAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnExtensionAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnExtensionAssociationPropsMixin
type jsiiProxy_CfnExtensionAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnExtensionAssociationPropsMixin) Props() *CfnExtensionAssociationMixinProps {
	var returns *CfnExtensionAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExtensionAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppConfig::ExtensionAssociation`.
func NewCfnExtensionAssociationPropsMixin(props *CfnExtensionAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnExtensionAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnExtensionAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnExtensionAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appconfig.CfnExtensionAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppConfig::ExtensionAssociation`.
func NewCfnExtensionAssociationPropsMixin_Override(c CfnExtensionAssociationPropsMixin, props *CfnExtensionAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appconfig.CfnExtensionAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnExtensionAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnExtensionAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_appconfig.CfnExtensionAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnExtensionAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_appconfig.CfnExtensionAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnExtensionAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnExtensionAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

