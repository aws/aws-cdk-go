package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Initialization properties for `OpenIdConnectProvider`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   // Step 1: Add retain policy to existing provider
//   existingProvider := eks.NewOpenIdConnectProvider(this, jsii.String("Provider"), &OpenIdConnectProviderProps{
//   	Url: jsii.String("https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLE"),
//   	RemovalPolicy: cdk.RemovalPolicy_RETAIN,
//   })
//
type OpenIdConnectProviderProps struct {
	// The URL of the identity provider.
	//
	// The URL must begin with https:// and
	// should correspond to the iss claim in the provider's OpenID Connect ID
	// tokens. Per the OIDC standard, path components are allowed but query
	// parameters are not. Typically the URL consists of only a hostname, like
	// https://server.example.org or https://example.com.
	//
	// You can find your OIDC Issuer URL by:
	// aws eks describe-cluster --name %cluster_name% --query "cluster.identity.oidc.issuer" --output text
	Url *string `field:"required" json:"url" yaml:"url"`
	// The removal policy to apply to the OpenID Connect Provider.
	// Default: - RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

