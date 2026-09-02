package interfacesawsomics


// A reference to a RunCache resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runCacheReference := &RunCacheReference{
//   	RunCacheArn: jsii.String("runCacheArn"),
//   }
//
type RunCacheReference struct {
	// The Arn of the RunCache resource.
	RunCacheArn *string `field:"required" json:"runCacheArn" yaml:"runCacheArn"`
}

