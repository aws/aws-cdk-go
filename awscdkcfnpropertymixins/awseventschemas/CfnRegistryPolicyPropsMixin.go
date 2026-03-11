package awseventschemas

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awseventschemas/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::EventSchemas::RegistryPolicy` resource to specify resource-based policies for an EventBridge Schema Registry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var policy interface{}
//
//   cfnRegistryPolicyPropsMixin := awscdkcfnpropertymixins.Aws_eventschemas.NewCfnRegistryPolicyPropsMixin(&CfnRegistryPolicyMixinProps{
//   	Policy: policy,
//   	RegistryName: jsii.String("registryName"),
//   	RevisionId: jsii.String("revisionId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-registrypolicy.html
//
type CfnRegistryPolicyPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRegistryPolicyMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRegistryPolicyPropsMixin
type jsiiProxy_CfnRegistryPolicyPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRegistryPolicyPropsMixin) Props() *CfnRegistryPolicyMixinProps {
	var returns *CfnRegistryPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistryPolicyPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EventSchemas::RegistryPolicy`.
func NewCfnRegistryPolicyPropsMixin(props *CfnRegistryPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRegistryPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRegistryPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRegistryPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_eventschemas.CfnRegistryPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EventSchemas::RegistryPolicy`.
func NewCfnRegistryPolicyPropsMixin_Override(c CfnRegistryPolicyPropsMixin, props *CfnRegistryPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_eventschemas.CfnRegistryPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRegistryPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRegistryPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_eventschemas.CfnRegistryPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRegistryPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_eventschemas.CfnRegistryPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRegistryPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRegistryPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

