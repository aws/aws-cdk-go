package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::CloudWatch::DomainAccessGrantForOrganization.
//
// An organization access grant authorizes a principal to administer an organization domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDomainAccessGrantForOrganizationPropsMixin := awscdkcfnpropertymixins.Aws_cloudwatch.NewCfnDomainAccessGrantForOrganizationPropsMixin(&CfnDomainAccessGrantForOrganizationMixinProps{
//   	DomainId: jsii.String("domainId"),
//   	Name: jsii.String("name"),
//   	Permission: jsii.String("permission"),
//   	Principal: &OrganizationAccessGrantPrincipalProperty{
//   		PrincipalAttributes: []interface{}{
//   			&PrincipalAttributeProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		PrincipalId: jsii.String("principalId"),
//   		PrincipalType: jsii.String("principalType"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainaccessgrantfororganization.html
//
type CfnDomainAccessGrantForOrganizationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDomainAccessGrantForOrganizationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDomainAccessGrantForOrganizationPropsMixin
type jsiiProxy_CfnDomainAccessGrantForOrganizationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDomainAccessGrantForOrganizationPropsMixin) Props() *CfnDomainAccessGrantForOrganizationMixinProps {
	var returns *CfnDomainAccessGrantForOrganizationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainAccessGrantForOrganizationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudWatch::DomainAccessGrantForOrganization`.
func NewCfnDomainAccessGrantForOrganizationPropsMixin(props *CfnDomainAccessGrantForOrganizationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDomainAccessGrantForOrganizationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDomainAccessGrantForOrganizationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDomainAccessGrantForOrganizationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainAccessGrantForOrganizationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudWatch::DomainAccessGrantForOrganization`.
func NewCfnDomainAccessGrantForOrganizationPropsMixin_Override(c CfnDomainAccessGrantForOrganizationPropsMixin, props *CfnDomainAccessGrantForOrganizationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainAccessGrantForOrganizationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDomainAccessGrantForOrganizationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomainAccessGrantForOrganizationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainAccessGrantForOrganizationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainAccessGrantForOrganizationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnDomainAccessGrantForOrganizationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomainAccessGrantForOrganizationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDomainAccessGrantForOrganizationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

