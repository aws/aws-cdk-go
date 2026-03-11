package awscases

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscases/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a domain, which is a container for all case data, such as cases, fields, templates and layouts.
//
// Each Amazon Connect instance can be associated with only one Cases domain.
//
// > This will not associate your connect instance to Cases domain. Instead, use the Amazon Connect [CreateIntegrationAssociation](https://docs.aws.amazon.com/connect/latest/APIReference/API_CreateIntegrationAssociation.html) API. You need specific IAM permissions to successfully associate the Cases domain. For more information, see [Onboard to Cases](https://docs.aws.amazon.com/connect/latest/adminguide/required-permissions-iam-cases.html#onboard-cases-iam) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDomainPropsMixin := awscdkcfnpropertymixins.Aws_cases.NewCfnDomainPropsMixin(&CfnDomainMixinProps{
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-domain.html
//
type CfnDomainPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDomainMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDomainPropsMixin
type jsiiProxy_CfnDomainPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDomainPropsMixin) Props() *CfnDomainMixinProps {
	var returns *CfnDomainMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cases::Domain`.
func NewCfnDomainPropsMixin(props *CfnDomainMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDomainPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDomainPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDomainPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cases.CfnDomainPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cases::Domain`.
func NewCfnDomainPropsMixin_Override(c CfnDomainPropsMixin, props *CfnDomainMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cases.CfnDomainPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDomainPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomainPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cases.CfnDomainPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cases.CfnDomainPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomainPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDomainPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

