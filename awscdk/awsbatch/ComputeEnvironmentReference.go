package awsbatch


// A reference to a ComputeEnvironment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeEnvironmentReference := &ComputeEnvironmentReference{
//   	ComputeEnvironmentArn: jsii.String("computeEnvironmentArn"),
//   }
//
type ComputeEnvironmentReference struct {
	// The ComputeEnvironmentArn of the ComputeEnvironment resource.
	ComputeEnvironmentArn *string `field:"required" json:"computeEnvironmentArn" yaml:"computeEnvironmentArn"`
}

