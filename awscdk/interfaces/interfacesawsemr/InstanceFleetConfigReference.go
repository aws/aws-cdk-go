package interfacesawsemr


// A reference to a InstanceFleetConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceFleetConfigReference := &InstanceFleetConfigReference{
//   	InstanceFleetConfigId: jsii.String("instanceFleetConfigId"),
//   }
//
type InstanceFleetConfigReference struct {
	// The Id of the InstanceFleetConfig resource.
	InstanceFleetConfigId *string `field:"required" json:"instanceFleetConfigId" yaml:"instanceFleetConfigId"`
}

