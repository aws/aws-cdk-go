package awsec2


// A reference to a Host resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostReference := &HostReference{
//   	HostId: jsii.String("hostId"),
//   }
//
type HostReference struct {
	// The HostId of the Host resource.
	HostId *string `field:"required" json:"hostId" yaml:"hostId"`
}

