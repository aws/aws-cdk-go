package previewawsnetworkmanagermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsnetworkmanager/previewawsnetworkmanagermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new, empty global network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGlobalNetworkPropsMixin := awscdkmixinspreview.Mixins.NewCfnGlobalNetworkPropsMixin(&CfnGlobalNetworkMixinProps{
//   	CreatedAt: jsii.String("createdAt"),
//   	Description: jsii.String("description"),
//   	State: jsii.String("state"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-globalnetwork.html
//
type CfnGlobalNetworkPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGlobalNetworkMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGlobalNetworkPropsMixin
type jsiiProxy_CfnGlobalNetworkPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGlobalNetworkPropsMixin) Props() *CfnGlobalNetworkMixinProps {
	var returns *CfnGlobalNetworkMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalNetworkPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkManager::GlobalNetwork`.
func NewCfnGlobalNetworkPropsMixin(props *CfnGlobalNetworkMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGlobalNetworkPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGlobalNetworkPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGlobalNetworkPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnGlobalNetworkPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkManager::GlobalNetwork`.
func NewCfnGlobalNetworkPropsMixin_Override(c CfnGlobalNetworkPropsMixin, props *CfnGlobalNetworkMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnGlobalNetworkPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGlobalNetworkPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGlobalNetworkPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnGlobalNetworkPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGlobalNetworkPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnGlobalNetworkPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGlobalNetworkPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalNetworkPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

