package awsapprunner


// Describes configuration settings related to network traffic of an AWS App Runner service.
//
// Consists of embedded objects for each configurable network feature.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	egressConfiguration: &egressConfigurationProperty{
//   		egressType: jsii.String("egressType"),
//
//   		// the properties below are optional
//   		vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   	},
//   	ingressConfiguration: &ingressConfigurationProperty{
//   		isPubliclyAccessible: jsii.Boolean(false),
//   	},
//   }
//
type CfnService_NetworkConfigurationProperty struct {
	// Network configuration settings for outbound message traffic.
	EgressConfiguration interface{} `field:"optional" json:"egressConfiguration" yaml:"egressConfiguration"`
	// `CfnService.NetworkConfigurationProperty.IngressConfiguration`.
	IngressConfiguration interface{} `field:"optional" json:"ingressConfiguration" yaml:"ingressConfiguration"`
}

