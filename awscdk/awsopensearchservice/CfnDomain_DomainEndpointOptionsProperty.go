package awsopensearchservice


// Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainEndpointOptionsProperty := &DomainEndpointOptionsProperty{
//   	CustomEndpoint: jsii.String("customEndpoint"),
//   	CustomEndpointCertificateArn: jsii.String("customEndpointCertificateArn"),
//   	CustomEndpointEnabled: jsii.Boolean(false),
//   	EnforceHttps: jsii.Boolean(false),
//   	TlsSecurityPolicy: jsii.String("tlsSecurityPolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html
//
type CfnDomain_DomainEndpointOptionsProperty struct {
	// The fully qualified URL for your custom endpoint.
	//
	// Required if you enabled a custom endpoint for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html#cfn-opensearchservice-domain-domainendpointoptions-customendpoint
	//
	CustomEndpoint *string `field:"optional" json:"customEndpoint" yaml:"customEndpoint"`
	// The AWS Certificate Manager ARN for your domain's SSL/TLS certificate.
	//
	// Required if you enabled a custom endpoint for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html#cfn-opensearchservice-domain-domainendpointoptions-customendpointcertificatearn
	//
	CustomEndpointCertificateArn *string `field:"optional" json:"customEndpointCertificateArn" yaml:"customEndpointCertificateArn"`
	// True to enable a custom endpoint for the domain.
	//
	// If enabled, you must also provide values for `CustomEndpoint` and `CustomEndpointCertificateArn` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html#cfn-opensearchservice-domain-domainendpointoptions-customendpointenabled
	//
	CustomEndpointEnabled interface{} `field:"optional" json:"customEndpointEnabled" yaml:"customEndpointEnabled"`
	// True to require that all traffic to the domain arrive over HTTPS.
	//
	// Required if you enable fine-grained access control in [AdvancedSecurityOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html#cfn-opensearchservice-domain-domainendpointoptions-enforcehttps
	//
	EnforceHttps interface{} `field:"optional" json:"enforceHttps" yaml:"enforceHttps"`
	// The minimum TLS version required for traffic to the domain. Valid values are TLS 1.3 (recommended) or 1.2:.
	//
	// - `Policy-Min-TLS-1-0-2019-07`
	// - `Policy-Min-TLS-1-2-2019-07`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html#cfn-opensearchservice-domain-domainendpointoptions-tlssecuritypolicy
	//
	TlsSecurityPolicy *string `field:"optional" json:"tlsSecurityPolicy" yaml:"tlsSecurityPolicy"`
}

