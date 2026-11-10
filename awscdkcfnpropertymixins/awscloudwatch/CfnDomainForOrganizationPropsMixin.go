package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::CloudWatch::DomainForOrganization.
//
// An organization domain provides a single sign-in URL and centralized identity provider configuration for your AWS Organization, removing the need for each account to configure its own domain. Only the management account can create it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDomainForOrganizationPropsMixin := awscdkcfnpropertymixins.Aws_cloudwatch.NewCfnDomainForOrganizationPropsMixin(&CfnDomainForOrganizationMixinProps{
//   	DomainAccessRoleArn: jsii.String("domainAccessRoleArn"),
//   	IdentityProviderConfiguration: &IdentityProviderConfigurationProperty{
//   		IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   			IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   		},
//   	},
//   	IdentityProviders: []*string{
//   		jsii.String("identityProviders"),
//   	},
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainfororganization.html
//
type CfnDomainForOrganizationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDomainForOrganizationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDomainForOrganizationPropsMixin
type jsiiProxy_CfnDomainForOrganizationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDomainForOrganizationPropsMixin) Props() *CfnDomainForOrganizationMixinProps {
	var returns *CfnDomainForOrganizationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainForOrganizationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudWatch::DomainForOrganization`.
func NewCfnDomainForOrganizationPropsMixin(props *CfnDomainForOrganizationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDomainForOrganizationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDomainForOrganizationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDomainForOrganizationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainForOrganizationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudWatch::DomainForOrganization`.
func NewCfnDomainForOrganizationPropsMixin_Override(c CfnDomainForOrganizationPropsMixin, props *CfnDomainForOrganizationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainForOrganizationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDomainForOrganizationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomainForOrganizationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainForOrganizationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainForOrganizationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainForOrganizationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomainForOrganizationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDomainForOrganizationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

