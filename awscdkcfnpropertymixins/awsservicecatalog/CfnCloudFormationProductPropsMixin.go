package awsservicecatalog

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a product.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var info interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnCloudFormationProductPropsMixin := awscdkcfnpropertymixins.Aws_servicecatalog.NewCfnCloudFormationProductPropsMixin(&CfnCloudFormationProductMixinProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	Description: jsii.String("description"),
//   	Distributor: jsii.String("distributor"),
//   	Name: jsii.String("name"),
//   	Owner: jsii.String("owner"),
//   	ProductType: jsii.String("productType"),
//   	ProvisioningArtifactParameters: []interface{}{
//   		&ProvisioningArtifactPropertiesProperty{
//   			Description: jsii.String("description"),
//   			DisableTemplateValidation: jsii.Boolean(false),
//   			Info: info,
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ReplaceProvisioningArtifacts: jsii.Boolean(false),
//   	SourceConnection: &SourceConnectionProperty{
//   		ConnectionParameters: &ConnectionParametersProperty{
//   			CodeStar: &CodeStarParametersProperty{
//   				ArtifactPath: jsii.String("artifactPath"),
//   				Branch: jsii.String("branch"),
//   				ConnectionArn: jsii.String("connectionArn"),
//   				Repository: jsii.String("repository"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	SupportDescription: jsii.String("supportDescription"),
//   	SupportEmail: jsii.String("supportEmail"),
//   	SupportUrl: jsii.String("supportUrl"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html
//
type CfnCloudFormationProductPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnCloudFormationProductMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCloudFormationProductPropsMixin
type jsiiProxy_CfnCloudFormationProductPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnCloudFormationProductPropsMixin) Props() *CfnCloudFormationProductMixinProps {
	var returns *CfnCloudFormationProductMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProductPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ServiceCatalog::CloudFormationProduct`.
func NewCfnCloudFormationProductPropsMixin(props *CfnCloudFormationProductMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnCloudFormationProductPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCloudFormationProductPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudFormationProductPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_servicecatalog.CfnCloudFormationProductPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ServiceCatalog::CloudFormationProduct`.
func NewCfnCloudFormationProductPropsMixin_Override(c CfnCloudFormationProductPropsMixin, props *CfnCloudFormationProductMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_servicecatalog.CfnCloudFormationProductPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCloudFormationProductPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudFormationProductPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_servicecatalog.CfnCloudFormationProductPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudFormationProductPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_servicecatalog.CfnCloudFormationProductPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudFormationProductPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCloudFormationProductPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

