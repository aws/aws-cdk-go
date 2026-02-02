package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Initialization properties for `OidcProviderNative`.
//
// Example:
//   // or create a new one using an existing issuer url
//   var issuerUrl string
//   // you can import an existing provider
//   provider := eks.OidcProviderNative_FromOidcProviderArn(this, jsii.String("Provider"), jsii.String("arn:aws:iam::123456:oidc-provider/oidc.eks.eu-west-1.amazonaws.com/id/AB123456ABC"))
//   provider2 := eks.NewOidcProviderNative(this, jsii.String("Provider"), &OidcProviderNativeProps{
//   	Url: issuerUrl,
//   })
//
//   cluster := eks.Cluster_FromClusterAttributes(this, jsii.String("MyCluster"), &ClusterAttributes{
//   	ClusterName: jsii.String("Cluster"),
//   	OpenIdConnectProvider: provider,
//   	KubectlRoleArn: jsii.String("arn:aws:iam::123456:role/service-role/k8sservicerole"),
//   })
//
//   serviceAccount := cluster.AddServiceAccount(jsii.String("MyServiceAccount"))
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   bucket.GrantReadWrite(serviceAccount)
//
type OidcProviderNativeProps struct {
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

