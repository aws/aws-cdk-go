package awss3tables


// A replication rule for the table bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   replicationRuleProperty := &ReplicationRuleProperty{
//   	Destinations: []interface{}{
//   		&ReplicationDestinationProperty{
//   			DestinationTableBucketArn: jsii.String("destinationTableBucketArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-replicationrule.html
//
type CfnTableBucketPropsMixin_ReplicationRuleProperty struct {
	// List of replication destinations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-replicationrule.html#cfn-s3tables-tablebucket-replicationrule-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
}

