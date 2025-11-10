package awss3vectors


// A reference to a VectorBucketPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vectorBucketPolicyReference := &VectorBucketPolicyReference{
//   	VectorBucketArn: jsii.String("vectorBucketArn"),
//   }
//
type VectorBucketPolicyReference struct {
	// The VectorBucketArn of the VectorBucketPolicy resource.
	VectorBucketArn *string `field:"required" json:"vectorBucketArn" yaml:"vectorBucketArn"`
}

