package interfacesawss3vectors


// A reference to a VectorBucket resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vectorBucketReference := &VectorBucketReference{
//   	VectorBucketArn: jsii.String("vectorBucketArn"),
//   }
//
type VectorBucketReference struct {
	// The VectorBucketArn of the VectorBucket resource.
	VectorBucketArn *string `field:"required" json:"vectorBucketArn" yaml:"vectorBucketArn"`
}

