package awsappflow


// A reference to a ConnectorProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorProfileReference := &ConnectorProfileReference{
//   	ConnectorProfileArn: jsii.String("connectorProfileArn"),
//   	ConnectorProfileName: jsii.String("connectorProfileName"),
//   }
//
type ConnectorProfileReference struct {
	// The ARN of the ConnectorProfile resource.
	ConnectorProfileArn *string `field:"required" json:"connectorProfileArn" yaml:"connectorProfileArn"`
	// The ConnectorProfileName of the ConnectorProfile resource.
	ConnectorProfileName *string `field:"required" json:"connectorProfileName" yaml:"connectorProfileName"`
}

