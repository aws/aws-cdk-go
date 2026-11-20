package awskinesisanalytics


// Details about the source (application and/or S3 location) the snapshot was taken from.
//
// Populated by DescribeSnapshot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapshotSourceDescriptionProperty := &SnapshotSourceDescriptionProperty{
//   	S3PrefixLocation: &S3PrefixLocationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		KeyPrefix: jsii.String("keyPrefix"),
//   	},
//   	SourceApplicationDescription: &SourceApplicationDescriptionProperty{
//   		ApplicationCreationTimestamp: jsii.String("applicationCreationTimestamp"),
//   		ApplicationName: jsii.String("applicationName"),
//   		ApplicationVersionId: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-snapshotsourcedescription.html
//
type CfnSnapshotV2_SnapshotSourceDescriptionProperty struct {
	// An Amazon S3 prefix location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-snapshotsourcedescription.html#cfn-kinesisanalyticsv2-snapshot-snapshotsourcedescription-s3prefixlocation
	//
	S3PrefixLocation interface{} `field:"optional" json:"s3PrefixLocation" yaml:"s3PrefixLocation"`
	// Details about the application the snapshot was taken from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-snapshotsourcedescription.html#cfn-kinesisanalyticsv2-snapshot-snapshotsourcedescription-sourceapplicationdescription
	//
	SourceApplicationDescription interface{} `field:"optional" json:"sourceApplicationDescription" yaml:"sourceApplicationDescription"`
}

