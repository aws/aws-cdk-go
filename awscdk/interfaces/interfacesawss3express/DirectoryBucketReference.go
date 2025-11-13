package interfacesawss3express


// A reference to a DirectoryBucket resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   directoryBucketReference := &DirectoryBucketReference{
//   	BucketName: jsii.String("bucketName"),
//   	DirectoryBucketArn: jsii.String("directoryBucketArn"),
//   }
//
type DirectoryBucketReference struct {
	// The BucketName of the DirectoryBucket resource.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The ARN of the DirectoryBucket resource.
	DirectoryBucketArn *string `field:"required" json:"directoryBucketArn" yaml:"directoryBucketArn"`
}

