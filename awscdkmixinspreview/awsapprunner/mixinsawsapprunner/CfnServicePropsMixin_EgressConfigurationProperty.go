package mixinsawsapprunner


// Describes configuration settings related to outbound network traffic of an AWS App Runner service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   egressConfigurationProperty := &EgressConfigurationProperty{
//   	EgressType: jsii.String("egressType"),
//   	VpcConnectorArn: jsii.String("vpcConnectorArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-egressconfiguration.html
//
type CfnServicePropsMixin_EgressConfigurationProperty struct {
	// The type of egress configuration.
	//
	// Set to `DEFAULT` for access to resources hosted on public networks.
	//
	// Set to `VPC` to associate your service to a custom VPC specified by `VpcConnectorArn` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-egressconfiguration.html#cfn-apprunner-service-egressconfiguration-egresstype
	//
	EgressType *string `field:"optional" json:"egressType" yaml:"egressType"`
	// The Amazon Resource Name (ARN) of the App Runner VPC connector that you want to associate with your App Runner service.
	//
	// Only valid when `EgressType = VPC` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-egressconfiguration.html#cfn-apprunner-service-egressconfiguration-vpcconnectorarn
	//
	VpcConnectorArn *string `field:"optional" json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
}

