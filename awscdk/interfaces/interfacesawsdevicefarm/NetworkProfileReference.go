package interfacesawsdevicefarm


// A reference to a NetworkProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkProfileReference := &NetworkProfileReference{
//   	NetworkProfileArn: jsii.String("networkProfileArn"),
//   }
//
type NetworkProfileReference struct {
	// The Arn of the NetworkProfile resource.
	NetworkProfileArn *string `field:"required" json:"networkProfileArn" yaml:"networkProfileArn"`
}

