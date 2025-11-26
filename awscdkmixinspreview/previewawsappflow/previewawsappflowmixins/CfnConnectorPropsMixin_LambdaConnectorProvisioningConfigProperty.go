package previewawsappflowmixins


// Contains information about the configuration of the lambda which is being registered as the connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lambdaConnectorProvisioningConfigProperty := &LambdaConnectorProvisioningConfigProperty{
//   	LambdaArn: jsii.String("lambdaArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connector-lambdaconnectorprovisioningconfig.html
//
type CfnConnectorPropsMixin_LambdaConnectorProvisioningConfigProperty struct {
	// Lambda ARN of the connector being registered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connector-lambdaconnectorprovisioningconfig.html#cfn-appflow-connector-lambdaconnectorprovisioningconfig-lambdaarn
	//
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
}

