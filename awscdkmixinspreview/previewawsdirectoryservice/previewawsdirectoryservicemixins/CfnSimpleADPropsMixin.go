package previewawsdirectoryservicemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdirectoryservice/previewawsdirectoryservicemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DirectoryService::SimpleAD` resource specifies an Directory Service Simple Active Directory ( Simple AD ) in AWS so that your directory users and groups can access the the console and AWS applications using their existing credentials.
//
// Simple AD is a Microsoft Active Directory–compatible directory. For more information, see [Simple Active Directory](https://docs.aws.amazon.com/directoryservice/latest/admin-guide/directory_simple_ad.html) in the *Directory Service Admin Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSimpleADPropsMixin := awscdkmixinspreview.Mixins.NewCfnSimpleADPropsMixin(&CfnSimpleADMixinProps{
//   	CreateAlias: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	EnableSso: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	Password: jsii.String("password"),
//   	ShortName: jsii.String("shortName"),
//   	Size: jsii.String("size"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html
//
type CfnSimpleADPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSimpleADMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSimpleADPropsMixin
type jsiiProxy_CfnSimpleADPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSimpleADPropsMixin) Props() *CfnSimpleADMixinProps {
	var returns *CfnSimpleADMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleADPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DirectoryService::SimpleAD`.
func NewCfnSimpleADPropsMixin(props *CfnSimpleADMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSimpleADPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSimpleADPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSimpleADPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_directoryservice.mixins.CfnSimpleADPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DirectoryService::SimpleAD`.
func NewCfnSimpleADPropsMixin_Override(c CfnSimpleADPropsMixin, props *CfnSimpleADMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_directoryservice.mixins.CfnSimpleADPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSimpleADPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSimpleADPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_directoryservice.mixins.CfnSimpleADPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSimpleADPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_directoryservice.mixins.CfnSimpleADPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSimpleADPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSimpleADPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

