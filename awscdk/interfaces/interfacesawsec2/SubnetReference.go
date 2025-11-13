package interfacesawsec2


// A reference to a Subnet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetReference := &SubnetReference{
//   	SubnetId: jsii.String("subnetId"),
//   }
//
type SubnetReference struct {
	// The SubnetId of the Subnet resource.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

