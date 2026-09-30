package interfacesawsglue


// A reference to a ConnectionType resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionTypeReference := &ConnectionTypeReference{
//   	ConnectionTypeArn: jsii.String("connectionTypeArn"),
//   }
//
type ConnectionTypeReference struct {
	// The ConnectionTypeArn of the ConnectionType resource.
	ConnectionTypeArn *string `field:"required" json:"connectionTypeArn" yaml:"connectionTypeArn"`
}

