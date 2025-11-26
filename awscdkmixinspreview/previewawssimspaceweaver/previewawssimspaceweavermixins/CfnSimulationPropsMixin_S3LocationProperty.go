package previewawssimspaceweavermixins


// A location in Amazon Simple Storage Service ( Amazon S3 ) where SimSpace Weaver stores simulation data, such as your app .zip files and schema file. For more information about Amazon S3 , see the [*Amazon Simple Storage Service User Guide*](https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LocationProperty := &S3LocationProperty{
//   	BucketName: jsii.String("bucketName"),
//   	ObjectKey: jsii.String("objectKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simspaceweaver-simulation-s3location.html
//
type CfnSimulationPropsMixin_S3LocationProperty struct {
	// The name of an Amazon S3 bucket.
	//
	// For more information about buckets, see [Creating, configuring, and working with Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-buckets-s3.html) in the *Amazon Simple Storage Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simspaceweaver-simulation-s3location.html#cfn-simspaceweaver-simulation-s3location-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The key name of an object in Amazon S3.
	//
	// For more information about Amazon S3 objects and object keys, see [Uploading, downloading, and working with objects in Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/userguide/uploading-downloading-objects.html) in the *Amazon Simple Storage Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simspaceweaver-simulation-s3location.html#cfn-simspaceweaver-simulation-s3location-objectkey
	//
	ObjectKey *string `field:"optional" json:"objectKey" yaml:"objectKey"`
}

