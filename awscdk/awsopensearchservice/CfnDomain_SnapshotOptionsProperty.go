package awsopensearchservice


// *DEPRECATED* .
//
// This setting is only relevant to domains running legacy Elasticsearch OSS versions earlier than 5.3. It does not apply to OpenSearch domains.
//
// The automated snapshot configuration for the OpenSearch Service domain indexes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapshotOptionsProperty := &SnapshotOptionsProperty{
//   	AutomatedSnapshotStartHour: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-snapshotoptions.html
//
type CfnDomain_SnapshotOptionsProperty struct {
	// The hour in UTC during which the service takes an automated daily snapshot of the indexes in the OpenSearch Service domain.
	//
	// For example, if you specify 0, OpenSearch Service takes an automated snapshot everyday between midnight and 1 am. You can specify a value between 0 and 23.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-snapshotoptions.html#cfn-opensearchservice-domain-snapshotoptions-automatedsnapshotstarthour
	//
	AutomatedSnapshotStartHour *float64 `field:"optional" json:"automatedSnapshotStartHour" yaml:"automatedSnapshotStartHour"`
}

