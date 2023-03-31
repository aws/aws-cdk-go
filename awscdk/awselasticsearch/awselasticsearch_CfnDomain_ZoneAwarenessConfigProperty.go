package awselasticsearch


// Specifies zone awareness configuration options. Only use if `ZoneAwarenessEnabled` is `true` .
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zoneAwarenessConfigProperty := &zoneAwarenessConfigProperty{
//   	availabilityZoneCount: jsii.Number(123),
//   }
//
type CfnDomain_ZoneAwarenessConfigProperty struct {
	// If you enabled multiple Availability Zones (AZs), the number of AZs that you want the domain to use.
	//
	// Valid values are `2` and `3` . Default is 2.
	AvailabilityZoneCount *float64 `field:"optional" json:"availabilityZoneCount" yaml:"availabilityZoneCount"`
}

