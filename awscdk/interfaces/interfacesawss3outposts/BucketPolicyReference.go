package interfacesawss3outposts


// A reference to a BucketPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketPolicyReference := &BucketPolicyReference{
//   	Bucket: jsii.String("bucket"),
//   }
//
type BucketPolicyReference struct {
	// The Bucket of the BucketPolicy resource.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
}

