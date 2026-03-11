package awsservicecatalog

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A self-service action association consisting of the Action ID, the Product ID, and the Provisioning Artifact ID.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnServiceActionAssociationPropsMixin := awscdkcfnpropertymixins.Aws_servicecatalog.NewCfnServiceActionAssociationPropsMixin(&CfnServiceActionAssociationMixinProps{
//   	ProductId: jsii.String("productId"),
//   	ProvisioningArtifactId: jsii.String("provisioningArtifactId"),
//   	ServiceActionId: jsii.String("serviceActionId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceactionassociation.html
//
type CfnServiceActionAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnServiceActionAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServiceActionAssociationPropsMixin
type jsiiProxy_CfnServiceActionAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnServiceActionAssociationPropsMixin) Props() *CfnServiceActionAssociationMixinProps {
	var returns *CfnServiceActionAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceActionAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ServiceCatalog::ServiceActionAssociation`.
func NewCfnServiceActionAssociationPropsMixin(props *CfnServiceActionAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnServiceActionAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServiceActionAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServiceActionAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_servicecatalog.CfnServiceActionAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ServiceCatalog::ServiceActionAssociation`.
func NewCfnServiceActionAssociationPropsMixin_Override(c CfnServiceActionAssociationPropsMixin, props *CfnServiceActionAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_servicecatalog.CfnServiceActionAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnServiceActionAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServiceActionAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_servicecatalog.CfnServiceActionAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServiceActionAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_servicecatalog.CfnServiceActionAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServiceActionAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnServiceActionAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

