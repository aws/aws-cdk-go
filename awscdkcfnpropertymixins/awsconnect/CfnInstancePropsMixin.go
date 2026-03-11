package awsconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// *This is a preview release for Amazon Connect . It is subject to change.*.
//
// Initiates an Amazon Connect instance with all the supported channels enabled. It does not attach any storage, such as Amazon Simple Storage Service (Amazon S3) or Amazon Kinesis.
//
// Amazon Connect enforces a limit on the total number of instances that you can create or delete in 30 days. If you exceed this limit, you will get an error message indicating there has been an excessive number of attempts at creating or deleting instances. You must wait 30 days before you can restart creating and deleting instances in your account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnInstancePropsMixin := awscdkcfnpropertymixins.Aws_connect.NewCfnInstancePropsMixin(&CfnInstanceMixinProps{
//   	Attributes: &AttributesProperty{
//   		AutoResolveBestVoices: jsii.Boolean(false),
//   		ContactflowLogs: jsii.Boolean(false),
//   		ContactLens: jsii.Boolean(false),
//   		EarlyMedia: jsii.Boolean(false),
//   		EnhancedChatMonitoring: jsii.Boolean(false),
//   		EnhancedContactMonitoring: jsii.Boolean(false),
//   		HighVolumeOutBound: jsii.Boolean(false),
//   		InboundCalls: jsii.Boolean(false),
//   		MessageStreaming: jsii.Boolean(false),
//   		MultiPartyChatConference: jsii.Boolean(false),
//   		MultiPartyConference: jsii.Boolean(false),
//   		OutboundCalls: jsii.Boolean(false),
//   		UseCustomTtsVoices: jsii.Boolean(false),
//   	},
//   	DirectoryId: jsii.String("directoryId"),
//   	IdentityManagementType: jsii.String("identityManagementType"),
//   	InstanceAlias: jsii.String("instanceAlias"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instance.html
//
type CfnInstancePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnInstanceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInstancePropsMixin
type jsiiProxy_CfnInstancePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnInstancePropsMixin) Props() *CfnInstanceMixinProps {
	var returns *CfnInstanceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstancePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::Instance`.
func NewCfnInstancePropsMixin(props *CfnInstanceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnInstancePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInstancePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstancePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnInstancePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::Instance`.
func NewCfnInstancePropsMixin_Override(c CfnInstancePropsMixin, props *CfnInstanceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnInstancePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnInstancePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstancePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnInstancePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstancePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnInstancePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstancePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnInstancePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

