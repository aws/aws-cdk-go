package awss3outposts


// A reference to a Bucket resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketReference := &BucketReference{
//   	BucketArn: jsii.String("bucketArn"),
//   }
//
type BucketReference struct {
	// The Arn of the Bucket resource.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
}

