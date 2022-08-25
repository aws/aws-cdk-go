package awsecs


// Kernel parameters to set in the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   systemControl := &systemControl{
//   	namespace: jsii.String("namespace"),
//   	value: jsii.String("value"),
//   }
//
type SystemControl struct {
	// The namespaced kernel parameter for which to set a value.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The value for the namespaced kernel parameter specified in namespace.
	Value *string `field:"required" json:"value" yaml:"value"`
}

