package awsiot


// A reference to a ResourceSpecificLogging resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceSpecificLoggingReference := &ResourceSpecificLoggingReference{
//   	TargetId: jsii.String("targetId"),
//   }
//
type ResourceSpecificLoggingReference struct {
	// The TargetId of the ResourceSpecificLogging resource.
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
}

