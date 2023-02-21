package awsssmincidents


// The `ReplicationRegion` property type specifies the Region and KMS key to add to the replication set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationRegionProperty := &replicationRegionProperty{
//   	regionConfiguration: &regionConfigurationProperty{
//   		sseKmsKeyId: jsii.String("sseKmsKeyId"),
//   	},
//   	regionName: jsii.String("regionName"),
//   }
//
type CfnReplicationSet_ReplicationRegionProperty struct {
	// Specifies the Region configuration.
	RegionConfiguration interface{} `field:"optional" json:"regionConfiguration" yaml:"regionConfiguration"`
	// Specifies the region name to add to the replication set.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

