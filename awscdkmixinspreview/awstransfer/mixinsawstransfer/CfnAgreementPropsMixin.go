package mixinsawstransfer

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awstransfer/mixinsawstransfer/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an agreement.
//
// An agreement is a bilateral trading partner agreement, or partnership, between an AWS Transfer Family server and an AS2 process. The agreement defines the file and message transfer relationship between the server and the AS2 process. To define an agreement, Transfer Family combines a server, local profile, partner profile, certificate, and other attributes.
//
// The partner is identified with the `PartnerProfileId` , and the AS2 process is identified with the `LocalProfileId` .
//
// > Specify *either* `BaseDirectory` or `CustomDirectories` , but not both. Specifying both causes the command to fail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgreementPropsMixin := awscdkmixinspreview.Mixins.NewCfnAgreementPropsMixin(&CfnAgreementMixinProps{
//   	AccessRole: jsii.String("accessRole"),
//   	BaseDirectory: jsii.String("baseDirectory"),
//   	CustomDirectories: &CustomDirectoriesProperty{
//   		FailedFilesDirectory: jsii.String("failedFilesDirectory"),
//   		MdnFilesDirectory: jsii.String("mdnFilesDirectory"),
//   		PayloadFilesDirectory: jsii.String("payloadFilesDirectory"),
//   		StatusFilesDirectory: jsii.String("statusFilesDirectory"),
//   		TemporaryFilesDirectory: jsii.String("temporaryFilesDirectory"),
//   	},
//   	Description: jsii.String("description"),
//   	EnforceMessageSigning: jsii.String("enforceMessageSigning"),
//   	LocalProfileId: jsii.String("localProfileId"),
//   	PartnerProfileId: jsii.String("partnerProfileId"),
//   	PreserveFilename: jsii.String("preserveFilename"),
//   	ServerId: jsii.String("serverId"),
//   	Status: jsii.String("status"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-agreement.html
//
type CfnAgreementPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAgreementMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAgreementPropsMixin
type jsiiProxy_CfnAgreementPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAgreementPropsMixin) Props() *CfnAgreementMixinProps {
	var returns *CfnAgreementMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgreementPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Transfer::Agreement`.
func NewCfnAgreementPropsMixin(props *CfnAgreementMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAgreementPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAgreementPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAgreementPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnAgreementPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Transfer::Agreement`.
func NewCfnAgreementPropsMixin_Override(c CfnAgreementPropsMixin, props *CfnAgreementMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnAgreementPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAgreementPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgreementPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnAgreementPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAgreementPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnAgreementPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAgreementPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAgreementPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

