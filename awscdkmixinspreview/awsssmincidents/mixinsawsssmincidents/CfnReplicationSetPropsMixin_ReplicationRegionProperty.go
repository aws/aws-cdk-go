package mixinsawsssmincidents


// The `ReplicationRegion` property type specifies the Region and AWS Key Management Service key to add to the replication set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   replicationRegionProperty := &ReplicationRegionProperty{
//   	RegionConfiguration: &RegionConfigurationProperty{
//   		SseKmsKeyId: jsii.String("sseKmsKeyId"),
//   	},
//   	RegionName: jsii.String("regionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-replicationset-replicationregion.html
//
type CfnReplicationSetPropsMixin_ReplicationRegionProperty struct {
	// Specifies the Region configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-replicationset-replicationregion.html#cfn-ssmincidents-replicationset-replicationregion-regionconfiguration
	//
	RegionConfiguration interface{} `field:"optional" json:"regionConfiguration" yaml:"regionConfiguration"`
	// Specifies the region name to add to the replication set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-replicationset-replicationregion.html#cfn-ssmincidents-replicationset-replicationregion-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

