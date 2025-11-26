package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawseks/previewawseksmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Associates an identity provider configuration to a cluster.
//
// If you want to authenticate identities using an identity provider, you can create an identity provider configuration and associate it to your cluster. After configuring authentication to your cluster you can create Kubernetes `Role` and `ClusterRole` objects, assign permissions to them, and then bind them to the identities using Kubernetes `RoleBinding` and `ClusterRoleBinding` objects. For more information see [Using RBAC Authorization](https://docs.aws.amazon.com/https://kubernetes.io/docs/reference/access-authn-authz/rbac/) in the Kubernetes documentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdentityProviderConfigPropsMixin := awscdkmixinspreview.Mixins.NewCfnIdentityProviderConfigPropsMixin(&CfnIdentityProviderConfigMixinProps{
//   	ClusterName: jsii.String("clusterName"),
//   	IdentityProviderConfigName: jsii.String("identityProviderConfigName"),
//   	Oidc: &OidcIdentityProviderConfigProperty{
//   		ClientId: jsii.String("clientId"),
//   		GroupsClaim: jsii.String("groupsClaim"),
//   		GroupsPrefix: jsii.String("groupsPrefix"),
//   		IssuerUrl: jsii.String("issuerUrl"),
//   		RequiredClaims: []interface{}{
//   			&RequiredClaimProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		UsernameClaim: jsii.String("usernameClaim"),
//   		UsernamePrefix: jsii.String("usernamePrefix"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-identityproviderconfig.html
//
type CfnIdentityProviderConfigPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIdentityProviderConfigMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIdentityProviderConfigPropsMixin
type jsiiProxy_CfnIdentityProviderConfigPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIdentityProviderConfigPropsMixin) Props() *CfnIdentityProviderConfigMixinProps {
	var returns *CfnIdentityProviderConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityProviderConfigPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EKS::IdentityProviderConfig`.
func NewCfnIdentityProviderConfigPropsMixin(props *CfnIdentityProviderConfigMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIdentityProviderConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIdentityProviderConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIdentityProviderConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnIdentityProviderConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EKS::IdentityProviderConfig`.
func NewCfnIdentityProviderConfigPropsMixin_Override(c CfnIdentityProviderConfigPropsMixin, props *CfnIdentityProviderConfigMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnIdentityProviderConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIdentityProviderConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIdentityProviderConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnIdentityProviderConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdentityProviderConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnIdentityProviderConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdentityProviderConfigPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnIdentityProviderConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

