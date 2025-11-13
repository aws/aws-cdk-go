package interfacesawsbatch


// A reference to a ConsumableResource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   consumableResourceReference := &ConsumableResourceReference{
//   	ConsumableResourceArn: jsii.String("consumableResourceArn"),
//   }
//
type ConsumableResourceReference struct {
	// The ConsumableResourceArn of the ConsumableResource resource.
	ConsumableResourceArn *string `field:"required" json:"consumableResourceArn" yaml:"consumableResourceArn"`
}

