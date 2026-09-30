package interfacesawseventsv2


// A reference to a ResourcePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePolicyReference := &ResourcePolicyReference{
//   	EventBusArn: jsii.String("eventBusArn"),
//   }
//
type ResourcePolicyReference struct {
	// The EventBusArn of the ResourcePolicy resource.
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
}

