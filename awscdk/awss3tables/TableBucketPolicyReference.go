package awss3tables


// A reference to a TableBucketPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableBucketPolicyReference := &TableBucketPolicyReference{
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   }
//
type TableBucketPolicyReference struct {
	// The TableBucketARN of the TableBucketPolicy resource.
	TableBucketArn *string `field:"required" json:"tableBucketArn" yaml:"tableBucketArn"`
}

