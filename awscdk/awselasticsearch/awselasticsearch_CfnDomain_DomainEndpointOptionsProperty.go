package awselasticsearch


// Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainEndpointOptionsProperty := &domainEndpointOptionsProperty{
//   	customEndpoint: jsii.String("customEndpoint"),
//   	customEndpointCertificateArn: jsii.String("customEndpointCertificateArn"),
//   	customEndpointEnabled: jsii.Boolean(false),
//   	enforceHttps: jsii.Boolean(false),
//   	tlsSecurityPolicy: jsii.String("tlsSecurityPolicy"),
//   }
//
type CfnDomain_DomainEndpointOptionsProperty struct {
	// The fully qualified URL for your custom endpoint.
	//
	// Required if you enabled a custom endpoint for the domain.
	CustomEndpoint *string `field:"optional" json:"customEndpoint" yaml:"customEndpoint"`
	// The AWS Certificate Manager ARN for your domain's SSL/TLS certificate.
	//
	// Required if you enabled a custom endpoint for the domain.
	CustomEndpointCertificateArn *string `field:"optional" json:"customEndpointCertificateArn" yaml:"customEndpointCertificateArn"`
	// True to enable a custom endpoint for the domain.
	//
	// If enabled, you must also provide values for `CustomEndpoint` and `CustomEndpointCertificateArn` .
	CustomEndpointEnabled interface{} `field:"optional" json:"customEndpointEnabled" yaml:"customEndpointEnabled"`
	// True to require that all traffic to the domain arrive over HTTPS.
	EnforceHttps interface{} `field:"optional" json:"enforceHttps" yaml:"enforceHttps"`
	// The minimum TLS version required for traffic to the domain. Valid values are TLS 1.0 (default) or 1.2:.
	//
	// - `Policy-Min-TLS-1-0-2019-07`
	// - `Policy-Min-TLS-1-2-2019-07`.
	TlsSecurityPolicy *string `field:"optional" json:"tlsSecurityPolicy" yaml:"tlsSecurityPolicy"`
}

