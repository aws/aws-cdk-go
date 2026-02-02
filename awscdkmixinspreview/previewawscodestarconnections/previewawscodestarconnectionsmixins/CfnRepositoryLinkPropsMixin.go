package previewawscodestarconnectionsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscodestarconnections/previewawscodestarconnectionsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Information about the repository link resource, such as the repository link ARN, the associated connection ARN, encryption key ARN, and owner ID.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRepositoryLinkPropsMixin := awscdkmixinspreview.Mixins.NewCfnRepositoryLinkPropsMixin(&CfnRepositoryLinkMixinProps{
//   	ConnectionArn: jsii.String("connectionArn"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	OwnerId: jsii.String("ownerId"),
//   	RepositoryName: jsii.String("repositoryName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html
//
type CfnRepositoryLinkPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRepositoryLinkMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRepositoryLinkPropsMixin
type jsiiProxy_CfnRepositoryLinkPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRepositoryLinkPropsMixin) Props() *CfnRepositoryLinkMixinProps {
	var returns *CfnRepositoryLinkMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepositoryLinkPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodeStarConnections::RepositoryLink`.
func NewCfnRepositoryLinkPropsMixin(props *CfnRepositoryLinkMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRepositoryLinkPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRepositoryLinkPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRepositoryLinkPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codestarconnections.mixins.CfnRepositoryLinkPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodeStarConnections::RepositoryLink`.
func NewCfnRepositoryLinkPropsMixin_Override(c CfnRepositoryLinkPropsMixin, props *CfnRepositoryLinkMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codestarconnections.mixins.CfnRepositoryLinkPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRepositoryLinkPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRepositoryLinkPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codestarconnections.mixins.CfnRepositoryLinkPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRepositoryLinkPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codestarconnections.mixins.CfnRepositoryLinkPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRepositoryLinkPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRepositoryLinkPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

