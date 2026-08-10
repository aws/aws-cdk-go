package awsmediapackage


// Configuration parameters for where in an S3 bucket to place the harvested content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3DestinationProperty := &S3DestinationProperty{
//   	BucketName: jsii.String("bucketName"),
//   	ManifestKey: jsii.String("manifestKey"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-harvestjob-s3destination.html
//
type CfnHarvestJobPropsMixin_S3DestinationProperty struct {
	// The name of an S3 bucket within which harvested content will be exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-harvestjob-s3destination.html#cfn-mediapackage-harvestjob-s3destination-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The key in the specified S3 bucket where the harvested top-level manifest will be placed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-harvestjob-s3destination.html#cfn-mediapackage-harvestjob-s3destination-manifestkey
	//
	ManifestKey *string `field:"optional" json:"manifestKey" yaml:"manifestKey"`
	// The IAM role used to write to the specified S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-harvestjob-s3destination.html#cfn-mediapackage-harvestjob-s3destination-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

