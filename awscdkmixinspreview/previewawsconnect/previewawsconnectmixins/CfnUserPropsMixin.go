package previewawsconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconnect/previewawsconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a user account for an Amazon Connect instance.
//
// For information about how to create user accounts using the Amazon Connect console, see [Add Users](https://docs.aws.amazon.com/connect/latest/adminguide/user-management.html) in the *Amazon Connect Administrator Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserPropsMixin := awscdkmixinspreview.Mixins.NewCfnUserPropsMixin(&CfnUserMixinProps{
//   	AfterContactWorkConfigs: []interface{}{
//   		&AfterContactWorkConfigPerChannelProperty{
//   			AfterContactWorkConfig: &AfterContactWorkConfigProperty{
//   				AfterContactWorkTimeLimit: jsii.Number(123),
//   			},
//   			AgentFirstCallbackAfterContactWorkConfig: &AfterContactWorkConfigProperty{
//   				AfterContactWorkTimeLimit: jsii.Number(123),
//   			},
//   			Channel: jsii.String("channel"),
//   		},
//   	},
//   	AutoAcceptConfigs: []interface{}{
//   		&AutoAcceptConfigProperty{
//   			AgentFirstCallbackAutoAccept: jsii.Boolean(false),
//   			AutoAccept: jsii.Boolean(false),
//   			Channel: jsii.String("channel"),
//   		},
//   	},
//   	DirectoryUserId: jsii.String("directoryUserId"),
//   	HierarchyGroupArn: jsii.String("hierarchyGroupArn"),
//   	IdentityInfo: &UserIdentityInfoProperty{
//   		Email: jsii.String("email"),
//   		FirstName: jsii.String("firstName"),
//   		LastName: jsii.String("lastName"),
//   		Mobile: jsii.String("mobile"),
//   		SecondaryEmail: jsii.String("secondaryEmail"),
//   	},
//   	InstanceArn: jsii.String("instanceArn"),
//   	Password: jsii.String("password"),
//   	PersistentConnectionConfigs: []interface{}{
//   		&PersistentConnectionConfigProperty{
//   			Channel: jsii.String("channel"),
//   			PersistentConnection: jsii.Boolean(false),
//   		},
//   	},
//   	PhoneConfig: &UserPhoneConfigProperty{
//   		AfterContactWorkTimeLimit: jsii.Number(123),
//   		AutoAccept: jsii.Boolean(false),
//   		DeskPhoneNumber: jsii.String("deskPhoneNumber"),
//   		PersistentConnection: jsii.Boolean(false),
//   		PhoneType: jsii.String("phoneType"),
//   	},
//   	PhoneNumberConfigs: []interface{}{
//   		&PhoneNumberConfigProperty{
//   			Channel: jsii.String("channel"),
//   			PhoneNumber: jsii.String("phoneNumber"),
//   			PhoneType: jsii.String("phoneType"),
//   		},
//   	},
//   	RoutingProfileArn: jsii.String("routingProfileArn"),
//   	SecurityProfileArns: []*string{
//   		jsii.String("securityProfileArns"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Username: jsii.String("username"),
//   	UserProficiencies: []interface{}{
//   		&UserProficiencyProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			AttributeValue: jsii.String("attributeValue"),
//   			Level: jsii.Number(123),
//   		},
//   	},
//   	VoiceEnhancementConfigs: []interface{}{
//   		&VoiceEnhancementConfigProperty{
//   			Channel: jsii.String("channel"),
//   			VoiceEnhancementMode: jsii.String("voiceEnhancementMode"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html
//
type CfnUserPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnUserMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserPropsMixin
type jsiiProxy_CfnUserPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnUserPropsMixin) Props() *CfnUserMixinProps {
	var returns *CfnUserMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::User`.
func NewCfnUserPropsMixin(props *CfnUserMixinProps, options *mixins.CfnPropertyMixinOptions) CfnUserPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnUserPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::User`.
func NewCfnUserPropsMixin_Override(c CfnUserPropsMixin, props *CfnUserMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnUserPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnUserPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnUserPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnUserPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnUserPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

