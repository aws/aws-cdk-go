package awss3tables


// Specifies replication configuration for the table bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   replicationConfigurationProperty := &ReplicationConfigurationProperty{
//   	Role: jsii.String("role"),
//   	Rules: []interface{}{
//   		&ReplicationRuleProperty{
//   			Destinations: []interface{}{
//   				&ReplicationDestinationProperty{
//   					DestinationTableBucketArn: jsii.String("destinationTableBucketArn"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-replicationconfiguration.html
//
type CfnTableBucketPropsMixin_ReplicationConfigurationProperty struct {
	// The ARN of the IAM role to use for replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-replicationconfiguration.html#cfn-s3tables-tablebucket-replicationconfiguration-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
	// List of replication rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-replicationconfiguration.html#cfn-s3tables-tablebucket-replicationconfiguration-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

