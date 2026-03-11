package awsappflow

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsappflow/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new connector profile associated with your AWS account .
//
// There is a soft quota of 100 connector profiles per AWS account . If you need more connector profiles than this quota allows, you can submit a request to the Amazon AppFlow team through the Amazon AppFlow support channel. In each connector profile that you create, you can provide the credentials and properties for only one connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnConnectorPropsMixin := awscdkcfnpropertymixins.Aws_appflow.NewCfnConnectorPropsMixin(&CfnConnectorMixinProps{
//   	ConnectorLabel: jsii.String("connectorLabel"),
//   	ConnectorProvisioningConfig: &ConnectorProvisioningConfigProperty{
//   		Lambda: &LambdaConnectorProvisioningConfigProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   		},
//   	},
//   	ConnectorProvisioningType: jsii.String("connectorProvisioningType"),
//   	Description: jsii.String("description"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-connector.html
//
type CfnConnectorPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnConnectorMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConnectorPropsMixin
type jsiiProxy_CfnConnectorPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnConnectorPropsMixin) Props() *CfnConnectorMixinProps {
	var returns *CfnConnectorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppFlow::Connector`.
func NewCfnConnectorPropsMixin(props *CfnConnectorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnConnectorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConnectorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnectorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appflow.CfnConnectorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppFlow::Connector`.
func NewCfnConnectorPropsMixin_Override(c CfnConnectorPropsMixin, props *CfnConnectorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appflow.CfnConnectorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnConnectorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_appflow.CfnConnectorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_appflow.CfnConnectorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectorPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnConnectorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

