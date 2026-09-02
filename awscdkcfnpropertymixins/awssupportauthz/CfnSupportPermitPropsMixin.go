package awssupportauthz

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssupportauthz/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::SupportAuthZ::SupportPermit.
//
// Represents a support permit that grants AWS support time-bounded access to one or more resources for a set of actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var allActions interface{}
//   var allResourcesInRegion interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnSupportPermitPropsMixin := awscdkcfnpropertymixins.Aws_supportauthz.NewCfnSupportPermitPropsMixin(&CfnSupportPermitMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Permit: &PermitProperty{
//   		Actions: &ActionSetProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			AllActions: allActions,
//   		},
//   		Conditions: []interface{}{
//   			&ConditionProperty{
//   				AllowAfter: jsii.String("allowAfter"),
//   				AllowBefore: jsii.String("allowBefore"),
//   			},
//   		},
//   		Resources: &ResourceSetProperty{
//   			AllResourcesInRegion: allResourcesInRegion,
//   			Resources: []*string{
//   				jsii.String("resources"),
//   			},
//   		},
//   	},
//   	SigningKeyInfo: &SigningKeyInfoProperty{
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	SupportCaseDisplayId: jsii.String("supportCaseDisplayId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportauthz-supportpermit.html
//
type CfnSupportPermitPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSupportPermitMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSupportPermitPropsMixin
type jsiiProxy_CfnSupportPermitPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSupportPermitPropsMixin) Props() *CfnSupportPermitMixinProps {
	var returns *CfnSupportPermitMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSupportPermitPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SupportAuthZ::SupportPermit`.
func NewCfnSupportPermitPropsMixin(props *CfnSupportPermitMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSupportPermitPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSupportPermitPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSupportPermitPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_supportauthz.CfnSupportPermitPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SupportAuthZ::SupportPermit`.
func NewCfnSupportPermitPropsMixin_Override(c CfnSupportPermitPropsMixin, props *CfnSupportPermitMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_supportauthz.CfnSupportPermitPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSupportPermitPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSupportPermitPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_supportauthz.CfnSupportPermitPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSupportPermitPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_supportauthz.CfnSupportPermitPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSupportPermitPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSupportPermitPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

