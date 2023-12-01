package awss3express


// Properties for defining a `CfnDirectoryBucket`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDirectoryBucketProps := &CfnDirectoryBucketProps{
//   	DataRedundancy: jsii.String("dataRedundancy"),
//   	LocationName: jsii.String("locationName"),
//
//   	// the properties below are optional
//   	BucketName: jsii.String("bucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-directorybucket.html
//
type CfnDirectoryBucketProps struct {
	// Specifies the number of Avilability Zone that's used for redundancy for the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-directorybucket.html#cfn-s3express-directorybucket-dataredundancy
	//
	DataRedundancy *string `field:"required" json:"dataRedundancy" yaml:"dataRedundancy"`
	// Specifies the AZ ID of the Availability Zone where the directory bucket will be created.
	//
	// An example AZ ID value is 'use1-az5'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-directorybucket.html#cfn-s3express-directorybucket-locationname
	//
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
	// Specifies a name for the bucket.
	//
	// The bucket name must contain only lowercase letters, numbers, dots (.), and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az2--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-directorybucket.html#cfn-s3express-directorybucket-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
}

