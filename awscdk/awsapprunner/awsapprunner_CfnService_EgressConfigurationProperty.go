package awsapprunner


// Describes configuration settings related to outbound network traffic of an AWS App Runner service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   egressConfigurationProperty := &egressConfigurationProperty{
//   	egressType: jsii.String("egressType"),
//
//   	// the properties below are optional
//   	vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   }
//
type CfnService_EgressConfigurationProperty struct {
	// The type of egress configuration.
	//
	// Set to `DEFAULT` for access to resources hosted on public networks.
	//
	// Set to `VPC` to associate your service to a custom VPC specified by `VpcConnectorArn` .
	EgressType *string `field:"required" json:"egressType" yaml:"egressType"`
	// The Amazon Resource Name (ARN) of the App Runner VPC connector that you want to associate with your App Runner service.
	//
	// Only valid when `EgressType = VPC` .
	VpcConnectorArn *string `field:"optional" json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
}

