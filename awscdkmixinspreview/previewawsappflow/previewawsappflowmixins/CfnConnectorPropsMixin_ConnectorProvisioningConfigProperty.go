package previewawsappflowmixins


// Contains information about the configuration of the connector being registered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   connectorProvisioningConfigProperty := &ConnectorProvisioningConfigProperty{
//   	Lambda: &LambdaConnectorProvisioningConfigProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connector-connectorprovisioningconfig.html
//
type CfnConnectorPropsMixin_ConnectorProvisioningConfigProperty struct {
	// Contains information about the configuration of the lambda which is being registered as the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connector-connectorprovisioningconfig.html#cfn-appflow-connector-connectorprovisioningconfig-lambda
	//
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
}

