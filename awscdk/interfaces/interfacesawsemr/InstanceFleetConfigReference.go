package interfacesawsemr


// A reference to a InstanceFleetConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceFleetConfigReference := &InstanceFleetConfigReference{
//   	ClusterId: jsii.String("clusterId"),
//   	InstanceFleetId: jsii.String("instanceFleetId"),
//   }
//
type InstanceFleetConfigReference struct {
	// The ClusterId of the InstanceFleetConfig resource.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The InstanceFleetId of the InstanceFleetConfig resource.
	InstanceFleetId *string `field:"required" json:"instanceFleetId" yaml:"instanceFleetId"`
}

