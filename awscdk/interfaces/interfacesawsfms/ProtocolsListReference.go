package interfacesawsfms


// A reference to a ProtocolsList resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   protocolsListReference := &ProtocolsListReference{
//   	ProtocolsListArn: jsii.String("protocolsListArn"),
//   }
//
type ProtocolsListReference struct {
	// The Arn of the ProtocolsList resource.
	ProtocolsListArn *string `field:"required" json:"protocolsListArn" yaml:"protocolsListArn"`
}

