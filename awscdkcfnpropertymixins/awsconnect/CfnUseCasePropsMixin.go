package awsconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for a use case associated with an Amazon Connect integration association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnUseCasePropsMixin := awscdkcfnpropertymixins.Aws_connect.NewCfnUseCasePropsMixin(&CfnUseCaseMixinProps{
//   	InstanceId: jsii.String("instanceId"),
//   	IntegrationAssociationId: jsii.String("integrationAssociationId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseCaseType: jsii.String("useCaseType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-usecase.html
//
type CfnUseCasePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnUseCaseMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUseCasePropsMixin
type jsiiProxy_CfnUseCasePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnUseCasePropsMixin) Props() *CfnUseCaseMixinProps {
	var returns *CfnUseCaseMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUseCasePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::UseCase`.
func NewCfnUseCasePropsMixin(props *CfnUseCaseMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnUseCasePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUseCasePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUseCasePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnUseCasePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::UseCase`.
func NewCfnUseCasePropsMixin_Override(c CfnUseCasePropsMixin, props *CfnUseCaseMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnUseCasePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnUseCasePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUseCasePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnUseCasePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUseCasePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnUseCasePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUseCasePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnUseCasePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

