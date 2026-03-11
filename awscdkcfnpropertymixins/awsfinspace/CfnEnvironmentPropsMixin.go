package awsfinspace

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsfinspace/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::FinSpace::Environment` resource represents an Amazon FinSpace environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnEnvironmentPropsMixin := awscdkcfnpropertymixins.Aws_finspace.NewCfnEnvironmentPropsMixin(&CfnEnvironmentMixinProps{
//   	DataBundles: []*string{
//   		jsii.String("dataBundles"),
//   	},
//   	Description: jsii.String("description"),
//   	FederationMode: jsii.String("federationMode"),
//   	FederationParameters: &FederationParametersProperty{
//   		ApplicationCallBackUrl: jsii.String("applicationCallBackUrl"),
//   		AttributeMap: []interface{}{
//   			&AttributeMapItemsProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FederationProviderName: jsii.String("federationProviderName"),
//   		FederationUrn: jsii.String("federationUrn"),
//   		SamlMetadataDocument: jsii.String("samlMetadataDocument"),
//   		SamlMetadataUrl: jsii.String("samlMetadataUrl"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Name: jsii.String("name"),
//   	SuperuserParameters: &SuperuserParametersProperty{
//   		EmailAddress: jsii.String("emailAddress"),
//   		FirstName: jsii.String("firstName"),
//   		LastName: jsii.String("lastName"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-finspace-environment.html
//
type CfnEnvironmentPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEnvironmentMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEnvironmentPropsMixin
type jsiiProxy_CfnEnvironmentPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEnvironmentPropsMixin) Props() *CfnEnvironmentMixinProps {
	var returns *CfnEnvironmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::FinSpace::Environment`.
func NewCfnEnvironmentPropsMixin(props *CfnEnvironmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEnvironmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_finspace.CfnEnvironmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::FinSpace::Environment`.
func NewCfnEnvironmentPropsMixin_Override(c CfnEnvironmentPropsMixin, props *CfnEnvironmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_finspace.CfnEnvironmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEnvironmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_finspace.CfnEnvironmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_finspace.CfnEnvironmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEnvironmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

