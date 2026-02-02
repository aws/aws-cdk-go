package previewawsservicecatalogmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsservicecatalog/previewawsservicecatalogmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Provisions the specified product.
//
// A provisioned product is a resourced instance of a product. For example, provisioning a product based on a AWS CloudFormation template launches a AWS CloudFormation stack and its underlying resources. You can check the status of this request using [DescribeRecord](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeRecord.html) .
//
// If the request contains a tag key with an empty list of values, there is a tag conflict for that key. Do not include conflicted keys as tags, or this causes the error "Parameter validation failed: Missing required parameter in Tags[ *N* ]: *Value* ".
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCloudFormationProvisionedProductPropsMixin := awscdkmixinspreview.Mixins.NewCfnCloudFormationProvisionedProductPropsMixin(&CfnCloudFormationProvisionedProductMixinProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	NotificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	PathId: jsii.String("pathId"),
//   	PathName: jsii.String("pathName"),
//   	ProductId: jsii.String("productId"),
//   	ProductName: jsii.String("productName"),
//   	ProvisionedProductName: jsii.String("provisionedProductName"),
//   	ProvisioningArtifactId: jsii.String("provisioningArtifactId"),
//   	ProvisioningArtifactName: jsii.String("provisioningArtifactName"),
//   	ProvisioningParameters: []interface{}{
//   		&ProvisioningParameterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ProvisioningPreferences: &ProvisioningPreferencesProperty{
//   		StackSetAccounts: []*string{
//   			jsii.String("stackSetAccounts"),
//   		},
//   		StackSetFailureToleranceCount: jsii.Number(123),
//   		StackSetFailureTolerancePercentage: jsii.Number(123),
//   		StackSetMaxConcurrencyCount: jsii.Number(123),
//   		StackSetMaxConcurrencyPercentage: jsii.Number(123),
//   		StackSetOperationType: jsii.String("stackSetOperationType"),
//   		StackSetRegions: []*string{
//   			jsii.String("stackSetRegions"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html
//
type CfnCloudFormationProvisionedProductPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCloudFormationProvisionedProductMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCloudFormationProvisionedProductPropsMixin
type jsiiProxy_CfnCloudFormationProvisionedProductPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProductPropsMixin) Props() *CfnCloudFormationProvisionedProductMixinProps {
	var returns *CfnCloudFormationProvisionedProductMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProductPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ServiceCatalog::CloudFormationProvisionedProduct`.
func NewCfnCloudFormationProvisionedProductPropsMixin(props *CfnCloudFormationProvisionedProductMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCloudFormationProvisionedProductPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCloudFormationProvisionedProductPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudFormationProvisionedProductPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_servicecatalog.mixins.CfnCloudFormationProvisionedProductPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ServiceCatalog::CloudFormationProvisionedProduct`.
func NewCfnCloudFormationProvisionedProductPropsMixin_Override(c CfnCloudFormationProvisionedProductPropsMixin, props *CfnCloudFormationProvisionedProductMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_servicecatalog.mixins.CfnCloudFormationProvisionedProductPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCloudFormationProvisionedProductPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudFormationProvisionedProductPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_servicecatalog.mixins.CfnCloudFormationProvisionedProductPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudFormationProvisionedProductPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_servicecatalog.mixins.CfnCloudFormationProvisionedProductPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProductPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProductPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

