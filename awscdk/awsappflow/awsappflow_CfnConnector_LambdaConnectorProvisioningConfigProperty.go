package awsappflow


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConnectorProvisioningConfigProperty := &lambdaConnectorProvisioningConfigProperty{
//   	lambdaArn: jsii.String("lambdaArn"),
//   }
//
type CfnConnector_LambdaConnectorProvisioningConfigProperty struct {
	// `CfnConnector.LambdaConnectorProvisioningConfigProperty.LambdaArn`.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
}

