package previewawsdirectoryservicemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdirectoryservice/previewawsdirectoryservicemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DirectoryService::MicrosoftAD` resource specifies a Microsoft Active Directory in AWS so that your directory users and groups can access the AWS Management Console and AWS applications using their existing credentials.
//
// For more information, see [AWS Managed Microsoft AD](https://docs.aws.amazon.com/directoryservice/latest/admin-guide/directory_microsoft_ad.html) in the *Directory Service Admin Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMicrosoftADPropsMixin := awscdkmixinspreview.Mixins.NewCfnMicrosoftADPropsMixin(&CfnMicrosoftADMixinProps{
//   	CreateAlias: jsii.Boolean(false),
//   	Edition: jsii.String("edition"),
//   	EnableSso: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	Password: jsii.String("password"),
//   	ShortName: jsii.String("shortName"),
//   	VpcSettings: &VpcSettingsProperty{
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html
//
type CfnMicrosoftADPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMicrosoftADMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMicrosoftADPropsMixin
type jsiiProxy_CfnMicrosoftADPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMicrosoftADPropsMixin) Props() *CfnMicrosoftADMixinProps {
	var returns *CfnMicrosoftADMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftADPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DirectoryService::MicrosoftAD`.
func NewCfnMicrosoftADPropsMixin(props *CfnMicrosoftADMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMicrosoftADPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMicrosoftADPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMicrosoftADPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_directoryservice.mixins.CfnMicrosoftADPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DirectoryService::MicrosoftAD`.
func NewCfnMicrosoftADPropsMixin_Override(c CfnMicrosoftADPropsMixin, props *CfnMicrosoftADMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_directoryservice.mixins.CfnMicrosoftADPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMicrosoftADPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMicrosoftADPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_directoryservice.mixins.CfnMicrosoftADPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMicrosoftADPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_directoryservice.mixins.CfnMicrosoftADPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMicrosoftADPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMicrosoftADPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

