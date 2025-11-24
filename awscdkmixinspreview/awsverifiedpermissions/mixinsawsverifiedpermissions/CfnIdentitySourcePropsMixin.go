package mixinsawsverifiedpermissions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsverifiedpermissions/mixinsawsverifiedpermissions/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates a reference to Amazon Cognito as an external identity provider.
//
// If you are creating a new identity source, then you must specify a `Configuration` . If you are updating an existing identity source, then you must specify an `UpdateConfiguration` .
//
// After you create an identity source, you can use the identities provided by the IdP as proxies for the principal in authorization queries that use the [IsAuthorizedWithToken](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorizedWithToken.html) operation. These identities take the form of tokens that contain claims about the user, such as IDs, attributes and group memberships. Amazon Cognito provides both identity tokens and access tokens, and Verified Permissions can use either or both. Any combination of identity and access tokens results in the same Cedar principal. Verified Permissions automatically translates the information about the identities into the standard Cedar attributes that can be evaluated by your policies. Because the Amazon Cognito identity and access tokens can contain different information, the tokens you choose to use determine the attributes that are available to access in the Cedar principal from your policies.
//
// Amazon Cognito Identity is not available in all of the same AWS Regions as  . Because of this, the `AWS::VerifiedPermissions::IdentitySource` type is not available to create from CloudFormation in Regions where Amazon Cognito Identity is not currently available. Users can still create `AWS::VerifiedPermissions::IdentitySource` in those Regions, but only from the AWS CLI ,  SDK, or from the AWS console.
//
// > To reference a user from this identity source in your Cedar policies, use the following syntax.
// >
// > *IdentityType::"<CognitoUserPoolIdentifier>|<CognitoClientId>*
// >
// > Where `IdentityType` is the string that you provide to the `PrincipalEntityType` parameter for this operation. The `CognitoUserPoolId` and `CognitoClientId` are defined by the Amazon Cognito user pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdentitySourcePropsMixin := awscdkmixinspreview.Mixins.NewCfnIdentitySourcePropsMixin(&CfnIdentitySourceMixinProps{
//   	Configuration: &IdentitySourceConfigurationProperty{
//   		CognitoUserPoolConfiguration: &CognitoUserPoolConfigurationProperty{
//   			ClientIds: []*string{
//   				jsii.String("clientIds"),
//   			},
//   			GroupConfiguration: &CognitoGroupConfigurationProperty{
//   				GroupEntityType: jsii.String("groupEntityType"),
//   			},
//   			UserPoolArn: jsii.String("userPoolArn"),
//   		},
//   		OpenIdConnectConfiguration: &OpenIdConnectConfigurationProperty{
//   			EntityIdPrefix: jsii.String("entityIdPrefix"),
//   			GroupConfiguration: &OpenIdConnectGroupConfigurationProperty{
//   				GroupClaim: jsii.String("groupClaim"),
//   				GroupEntityType: jsii.String("groupEntityType"),
//   			},
//   			Issuer: jsii.String("issuer"),
//   			TokenSelection: &OpenIdConnectTokenSelectionProperty{
//   				AccessTokenOnly: &OpenIdConnectAccessTokenConfigurationProperty{
//   					Audiences: []*string{
//   						jsii.String("audiences"),
//   					},
//   					PrincipalIdClaim: jsii.String("principalIdClaim"),
//   				},
//   				IdentityTokenOnly: &OpenIdConnectIdentityTokenConfigurationProperty{
//   					ClientIds: []*string{
//   						jsii.String("clientIds"),
//   					},
//   					PrincipalIdClaim: jsii.String("principalIdClaim"),
//   				},
//   			},
//   		},
//   	},
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   	PrincipalEntityType: jsii.String("principalEntityType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-identitysource.html
//
type CfnIdentitySourcePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIdentitySourceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIdentitySourcePropsMixin
type jsiiProxy_CfnIdentitySourcePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIdentitySourcePropsMixin) Props() *CfnIdentitySourceMixinProps {
	var returns *CfnIdentitySourceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentitySourcePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::VerifiedPermissions::IdentitySource`.
func NewCfnIdentitySourcePropsMixin(props *CfnIdentitySourceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIdentitySourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIdentitySourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIdentitySourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::VerifiedPermissions::IdentitySource`.
func NewCfnIdentitySourcePropsMixin_Override(c CfnIdentitySourcePropsMixin, props *CfnIdentitySourceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIdentitySourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIdentitySourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdentitySourcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnIdentitySourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdentitySourcePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnIdentitySourcePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

