package awss3tables


// A reference to a TableBucket resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableBucketReference := &TableBucketReference{
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   }
//
type TableBucketReference struct {
	// The TableBucketARN of the TableBucket resource.
	TableBucketArn *string `field:"required" json:"tableBucketArn" yaml:"tableBucketArn"`
}

