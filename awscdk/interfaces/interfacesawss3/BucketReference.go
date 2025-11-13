package interfacesawss3


// A reference to a Bucket resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketReference := &BucketReference{
//   	BucketArn: jsii.String("bucketArn"),
//   	BucketName: jsii.String("bucketName"),
//   }
//
type BucketReference struct {
	// The ARN of the Bucket resource.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The BucketName of the Bucket resource.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

