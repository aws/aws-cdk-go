package awsec2


// A reference to a EC2Fleet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eC2FleetReference := &EC2FleetReference{
//   	FleetId: jsii.String("fleetId"),
//   }
//
type EC2FleetReference struct {
	// The FleetId of the EC2Fleet resource.
	FleetId *string `field:"required" json:"fleetId" yaml:"fleetId"`
}

