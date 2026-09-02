package awsec2


// Describes a storage location in Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageLocationProperty := &StorageLocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fpgaimage-storagelocation.html
//
type CfnFpgaImage_StorageLocationProperty struct {
	// The name of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fpgaimage-storagelocation.html#cfn-ec2-fpgaimage-storagelocation-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fpgaimage-storagelocation.html#cfn-ec2-fpgaimage-storagelocation-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
}

