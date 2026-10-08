package awsredshift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Redshift::RedshiftIdcApplication.
//
// Creates an Amazon Redshift application for use with IAM Identity Center.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRedshiftIdcApplicationPropsMixin := awscdkcfnpropertymixins.Aws_redshift.NewCfnRedshiftIdcApplicationPropsMixin(&CfnRedshiftIdcApplicationMixinProps{
//   	ApplicationType: jsii.String("applicationType"),
//   	AuthorizedTokenIssuerList: []interface{}{
//   		&AuthorizedTokenIssuerProperty{
//   			AuthorizedAudiencesList: []*string{
//   				jsii.String("authorizedAudiencesList"),
//   			},
//   			TrustedTokenIssuerArn: jsii.String("trustedTokenIssuerArn"),
//   		},
//   	},
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	IdcDisplayName: jsii.String("idcDisplayName"),
//   	IdcInstanceArn: jsii.String("idcInstanceArn"),
//   	IdentityNamespace: jsii.String("identityNamespace"),
//   	RedshiftIdcApplicationName: jsii.String("redshiftIdcApplicationName"),
//   	ServiceIntegrations: []interface{}{
//   		&ServiceIntegrationsUnionProperty{
//   			LakeFormation: []interface{}{
//   				&LakeFormationScopeUnionProperty{
//   					LakeFormationQuery: &LakeFormationQueryProperty{
//   						Authorization: jsii.String("authorization"),
//   					},
//   				},
//   			},
//   			Redshift: []interface{}{
//   				&RedshiftScopeUnionProperty{
//   					Connect: &ConnectProperty{
//   						Authorization: jsii.String("authorization"),
//   					},
//   				},
//   			},
//   			S3AccessGrants: []interface{}{
//   				&S3AccessGrantsScopeUnionProperty{
//   					ReadWriteAccess: &ReadWriteAccessProperty{
//   						Authorization: jsii.String("authorization"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	SsoTagKeys: []*string{
//   		jsii.String("ssoTagKeys"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-redshiftidcapplication.html
//
type CfnRedshiftIdcApplicationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRedshiftIdcApplicationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRedshiftIdcApplicationPropsMixin
type jsiiProxy_CfnRedshiftIdcApplicationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRedshiftIdcApplicationPropsMixin) Props() *CfnRedshiftIdcApplicationMixinProps {
	var returns *CfnRedshiftIdcApplicationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRedshiftIdcApplicationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Redshift::RedshiftIdcApplication`.
func NewCfnRedshiftIdcApplicationPropsMixin(props *CfnRedshiftIdcApplicationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRedshiftIdcApplicationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRedshiftIdcApplicationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRedshiftIdcApplicationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnRedshiftIdcApplicationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Redshift::RedshiftIdcApplication`.
func NewCfnRedshiftIdcApplicationPropsMixin_Override(c CfnRedshiftIdcApplicationPropsMixin, props *CfnRedshiftIdcApplicationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnRedshiftIdcApplicationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRedshiftIdcApplicationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRedshiftIdcApplicationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnRedshiftIdcApplicationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRedshiftIdcApplicationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnRedshiftIdcApplicationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRedshiftIdcApplicationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRedshiftIdcApplicationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

