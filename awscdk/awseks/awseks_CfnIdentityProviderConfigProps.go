package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnIdentityProviderConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIdentityProviderConfigProps := &CfnIdentityProviderConfigProps{
//   	ClusterName: jsii.String("clusterName"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	IdentityProviderConfigName: jsii.String("identityProviderConfigName"),
//   	Oidc: &OidcIdentityProviderConfigProperty{
//   		ClientId: jsii.String("clientId"),
//   		IssuerUrl: jsii.String("issuerUrl"),
//
//   		// the properties below are optional
//   		GroupsClaim: jsii.String("groupsClaim"),
//   		GroupsPrefix: jsii.String("groupsPrefix"),
//   		RequiredClaims: []interface{}{
//   			&RequiredClaimProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		UsernameClaim: jsii.String("usernameClaim"),
//   		UsernamePrefix: jsii.String("usernamePrefix"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIdentityProviderConfigProps struct {
	// The cluster that the configuration is associated to.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The type of the identity provider configuration.
	//
	// The only type available is `oidc` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of the configuration.
	IdentityProviderConfigName *string `field:"optional" json:"identityProviderConfigName" yaml:"identityProviderConfigName"`
	// An object representing an OpenID Connect (OIDC) identity provider configuration.
	Oidc interface{} `field:"optional" json:"oidc" yaml:"oidc"`
	// The metadata to apply to the provider configuration to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

