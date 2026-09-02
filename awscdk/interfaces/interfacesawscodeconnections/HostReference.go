package interfacesawscodeconnections


// A reference to a Host resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostReference := &HostReference{
//   	HostArn: jsii.String("hostArn"),
//   }
//
type HostReference struct {
	// The HostArn of the Host resource.
	HostArn *string `field:"required" json:"hostArn" yaml:"hostArn"`
}

