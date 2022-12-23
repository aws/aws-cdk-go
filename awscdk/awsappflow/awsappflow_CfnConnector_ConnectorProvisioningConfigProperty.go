package awsappflow


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorProvisioningConfigProperty := &connectorProvisioningConfigProperty{
//   	lambda: &lambdaConnectorProvisioningConfigProperty{
//   		lambdaArn: jsii.String("lambdaArn"),
//   	},
//   }
//
type CfnConnector_ConnectorProvisioningConfigProperty struct {
	// `CfnConnector.ConnectorProvisioningConfigProperty.Lambda`.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
}

