package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AppSync::DomainName` resource creates a `DomainNameConfig` object to configure a custom domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDomainNamePropsMixin := awscdkcfnpropertymixins.Aws_appsync.NewCfnDomainNamePropsMixin(&CfnDomainNameMixinProps{
//   	CertificateArn: jsii.String("certificateArn"),
//   	Description: jsii.String("description"),
//   	DomainName: jsii.String("domainName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainname.html
//
type CfnDomainNamePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDomainNameMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDomainNamePropsMixin
type jsiiProxy_CfnDomainNamePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDomainNamePropsMixin) Props() *CfnDomainNameMixinProps {
	var returns *CfnDomainNameMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNamePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppSync::DomainName`.
func NewCfnDomainNamePropsMixin(props *CfnDomainNameMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDomainNamePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDomainNamePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDomainNamePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appsync.CfnDomainNamePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppSync::DomainName`.
func NewCfnDomainNamePropsMixin_Override(c CfnDomainNamePropsMixin, props *CfnDomainNameMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appsync.CfnDomainNamePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDomainNamePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomainNamePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_appsync.CfnDomainNamePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainNamePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_appsync.CfnDomainNamePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomainNamePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDomainNamePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

