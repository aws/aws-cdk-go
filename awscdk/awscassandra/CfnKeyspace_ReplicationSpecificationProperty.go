package awscassandra


// You can use `ReplicationSpecification` to configure the `ReplicationStrategy` of a keyspace in Amazon Keyspaces.
//
// The `ReplicationSpecification` property is `CreateOnly` and cannot be changed after the keyspace has been created. This property applies automatically to all tables in the keyspace.
//
// For more information, see [Multi-Region Replication](https://docs.aws.amazon.com/keyspaces/latest/devguide/multiRegion-replication.html) in the *Amazon Keyspaces Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationSpecificationProperty := &ReplicationSpecificationProperty{
//   	RegionList: []*string{
//   		jsii.String("regionList"),
//   	},
//   	ReplicationStrategy: jsii.String("replicationStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-keyspace-replicationspecification.html
//
type CfnKeyspace_ReplicationSpecificationProperty struct {
	// Specifies the AWS Regions that the keyspace is replicated in.
	//
	// You must specify at least two and up to six Regions, including the Region that the keyspace is being created in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-keyspace-replicationspecification.html#cfn-cassandra-keyspace-replicationspecification-regionlist
	//
	RegionList *[]*string `field:"optional" json:"regionList" yaml:"regionList"`
	// The options are:.
	//
	// - `SINGLE_REGION` (optional)
	// - `MULTI_REGION`
	//
	// If no value is specified, the default is `SINGLE_REGION` . If `MULTI_REGION` is specified, `RegionList` is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-keyspace-replicationspecification.html#cfn-cassandra-keyspace-replicationspecification-replicationstrategy
	//
	ReplicationStrategy *string `field:"optional" json:"replicationStrategy" yaml:"replicationStrategy"`
}

