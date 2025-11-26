package previewawstransfermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawstransfer/previewawstransfermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Transfer::User` resource creates a user and associates them with an existing server.
//
// You can only create and associate users with servers that have the `IdentityProviderType` set to `SERVICE_MANAGED` . Using parameters for `CreateUser` , you can specify the user name, set the home directory, store the user's public key, and assign the user's AWS Identity and Access Management (IAM) role. You can also optionally add a session policy, and assign metadata with tags that can be used to group and search for users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserPropsMixin := awscdkmixinspreview.Mixins.NewCfnUserPropsMixin(&CfnUserMixinProps{
//   	HomeDirectory: jsii.String("homeDirectory"),
//   	HomeDirectoryMappings: []interface{}{
//   		&HomeDirectoryMapEntryProperty{
//   			Entry: jsii.String("entry"),
//   			Target: jsii.String("target"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	HomeDirectoryType: jsii.String("homeDirectoryType"),
//   	Policy: jsii.String("policy"),
//   	PosixProfile: &PosixProfileProperty{
//   		Gid: jsii.Number(123),
//   		SecondaryGids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		Uid: jsii.Number(123),
//   	},
//   	Role: jsii.String("role"),
//   	ServerId: jsii.String("serverId"),
//   	SshPublicKeys: []*string{
//   		jsii.String("sshPublicKeys"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserName: jsii.String("userName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html
//
type CfnUserPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnUserMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
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


// Create a mixin to apply properties to `AWS::Transfer::User`.
func NewCfnUserPropsMixin(props *CfnUserMixinProps, options *mixins.CfnPropertyMixinOptions) CfnUserPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnUserPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Transfer::User`.
func NewCfnUserPropsMixin_Override(c CfnUserPropsMixin, props *CfnUserMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnUserPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnUserPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnUserPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

