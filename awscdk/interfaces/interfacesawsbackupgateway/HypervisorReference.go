package interfacesawsbackupgateway


// A reference to a Hypervisor resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hypervisorReference := &HypervisorReference{
//   	HypervisorArn: jsii.String("hypervisorArn"),
//   }
//
type HypervisorReference struct {
	// The HypervisorArn of the Hypervisor resource.
	HypervisorArn *string `field:"required" json:"hypervisorArn" yaml:"hypervisorArn"`
}

