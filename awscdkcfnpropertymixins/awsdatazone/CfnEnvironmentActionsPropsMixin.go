package awsdatazone

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The details about the specified action configured for an environment.
//
// For example, the details of the specified console links for an analytics tool that is available in this environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnEnvironmentActionsPropsMixin := awscdkcfnpropertymixins.Aws_datazone.NewCfnEnvironmentActionsPropsMixin(&CfnEnvironmentActionsMixinProps{
//   	Description: jsii.String("description"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	Identifier: jsii.String("identifier"),
//   	Name: jsii.String("name"),
//   	Parameters: &AwsConsoleLinkParametersProperty{
//   		Uri: jsii.String("uri"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentactions.html
//
type CfnEnvironmentActionsPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEnvironmentActionsMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEnvironmentActionsPropsMixin
type jsiiProxy_CfnEnvironmentActionsPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEnvironmentActionsPropsMixin) Props() *CfnEnvironmentActionsMixinProps {
	var returns *CfnEnvironmentActionsMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentActionsPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataZone::EnvironmentActions`.
func NewCfnEnvironmentActionsPropsMixin(props *CfnEnvironmentActionsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEnvironmentActionsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentActionsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironmentActionsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentActionsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataZone::EnvironmentActions`.
func NewCfnEnvironmentActionsPropsMixin_Override(c CfnEnvironmentActionsPropsMixin, props *CfnEnvironmentActionsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentActionsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEnvironmentActionsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironmentActionsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentActionsPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironmentActionsPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnEnvironmentActionsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironmentActionsPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEnvironmentActionsPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

