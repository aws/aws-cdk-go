package awsiotfleetwise


// A reference to a Vehicle resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vehicleReference := &VehicleReference{
//   	VehicleArn: jsii.String("vehicleArn"),
//   	VehicleName: jsii.String("vehicleName"),
//   }
//
type VehicleReference struct {
	// The ARN of the Vehicle resource.
	VehicleArn *string `field:"required" json:"vehicleArn" yaml:"vehicleArn"`
	// The Name of the Vehicle resource.
	VehicleName *string `field:"required" json:"vehicleName" yaml:"vehicleName"`
}

