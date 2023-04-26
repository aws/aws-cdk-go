package awss3


// The container element for a bucket's policy status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyStatusProperty := &PolicyStatusProperty{
//   	IsPublic: jsii.String("isPublic"),
//   }
//
type CfnAccessPoint_PolicyStatusProperty struct {
	// The policy status for this bucket.
	//
	// `TRUE` indicates that this bucket is public. `FALSE` indicates that the bucket is not public.
	IsPublic *string `field:"optional" json:"isPublic" yaml:"isPublic"`
}

