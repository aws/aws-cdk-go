package awselasticsearch


// Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service key to use.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionAtRestOptionsProperty := &encryptionAtRestOptionsProperty{
//   	enabled: jsii.Boolean(false),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnDomain_EncryptionAtRestOptionsProperty struct {
	// Specify `true` to enable encryption at rest.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The KMS key ID.
	//
	// Takes the form `1a2a3a4-1a2a-3a4a-5a6a-1a2a3a4a5a6a` . Required if you enable encryption at rest.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

