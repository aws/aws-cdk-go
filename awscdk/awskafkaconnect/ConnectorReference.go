package awskafkaconnect


// A reference to a Connector resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorReference := &ConnectorReference{
//   	ConnectorArn: jsii.String("connectorArn"),
//   }
//
type ConnectorReference struct {
	// The ConnectorArn of the Connector resource.
	ConnectorArn *string `field:"required" json:"connectorArn" yaml:"connectorArn"`
}

