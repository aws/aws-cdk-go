package awsappflow


// Contains information about the configuration of the connector being registered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorProvisioningConfigProperty := &ConnectorProvisioningConfigProperty{
//   	Lambda: &LambdaConnectorProvisioningConfigProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   	},
//   }
//
type CfnConnector_ConnectorProvisioningConfigProperty struct {
	// Contains information about the configuration of the lambda which is being registered as the connector.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
}

