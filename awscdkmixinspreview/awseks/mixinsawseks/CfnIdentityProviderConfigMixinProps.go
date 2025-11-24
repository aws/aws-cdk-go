package mixinsawseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIdentityProviderConfigPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdentityProviderConfigMixinProps := &CfnIdentityProviderConfigMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-identityproviderconfig.html
//
type CfnIdentityProviderConfigMixinProps struct {
	// The name of your cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-identityproviderconfig.html#cfn-eks-identityproviderconfig-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The name of the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-identityproviderconfig.html#cfn-eks-identityproviderconfig-identityproviderconfigname
	//
	IdentityProviderConfigName *string `field:"optional" json:"identityProviderConfigName" yaml:"identityProviderConfigName"`
	// An object representing an OpenID Connect (OIDC) identity provider configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-identityproviderconfig.html#cfn-eks-identityproviderconfig-oidc
	//
	Oidc interface{} `field:"optional" json:"oidc" yaml:"oidc"`
	// Metadata that assists with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Tags don't propagate to any other cluster or AWS resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-identityproviderconfig.html#cfn-eks-identityproviderconfig-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of the identity provider configuration.
	//
	// The only type available is `oidc` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-identityproviderconfig.html#cfn-eks-identityproviderconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

