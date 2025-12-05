package interfacesawssecurityhub


// A reference to a ConnectorV2 resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorV2Reference := &ConnectorV2Reference{
//   	ConnectorArn: jsii.String("connectorArn"),
//   }
//
type ConnectorV2Reference struct {
	// The ConnectorArn of the ConnectorV2 resource.
	ConnectorArn *string `field:"required" json:"connectorArn" yaml:"connectorArn"`
}

