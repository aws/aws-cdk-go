package awselasticsearch


// Specifies whether node-to-node encryption is enabled.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeToNodeEncryptionOptionsProperty := &nodeToNodeEncryptionOptionsProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnDomain_NodeToNodeEncryptionOptionsProperty struct {
	// Specifies whether node-to-node encryption is enabled, as a Boolean.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

