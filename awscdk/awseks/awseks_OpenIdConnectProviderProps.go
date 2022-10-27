package awseks


// Initialization properties for `OpenIdConnectProvider`.
//
// Example:
//   // or create a new one using an existing issuer url
//   var issuerUrl string
//   // you can import an existing provider
//   provider := eks.openIdConnectProvider.fromOpenIdConnectProviderArn(this, jsii.String("Provider"), jsii.String("arn:aws:iam::123456:oidc-provider/oidc.eks.eu-west-1.amazonaws.com/id/AB123456ABC"))
//   provider2 := eks.NewOpenIdConnectProvider(this, jsii.String("Provider"), &openIdConnectProviderProps{
//   	url: issuerUrl,
//   })
//
//   cluster := eks.cluster.fromClusterAttributes(this, jsii.String("MyCluster"), &clusterAttributes{
//   	clusterName: jsii.String("Cluster"),
//   	openIdConnectProvider: provider,
//   	kubectlRoleArn: jsii.String("arn:aws:iam::123456:role/service-role/k8sservicerole"),
//   })
//
//   serviceAccount := cluster.addServiceAccount(jsii.String("MyServiceAccount"))
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   bucket.grantReadWrite(serviceAccount)
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
}

