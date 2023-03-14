package awselasticsearch


// Specifies options for cold storage. For more information, see [Cold storage for Amazon Elasticsearch Service](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/cold-storage.html) .
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coldStorageOptionsProperty := &coldStorageOptionsProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnDomain_ColdStorageOptionsProperty struct {
	// Whether to enable or disable cold storage on the domain.
	//
	// You must enable UltraWarm storage in order to enable cold storage.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

