package awss3tables


// A replication destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   replicationDestinationProperty := &ReplicationDestinationProperty{
//   	DestinationTableBucketArn: jsii.String("destinationTableBucketArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-replicationdestination.html
//
type CfnTableBucketPropsMixin_ReplicationDestinationProperty struct {
	// The ARN of the destination table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-replicationdestination.html#cfn-s3tables-tablebucket-replicationdestination-destinationtablebucketarn
	//
	DestinationTableBucketArn *string `field:"optional" json:"destinationTableBucketArn" yaml:"destinationTableBucketArn"`
}

