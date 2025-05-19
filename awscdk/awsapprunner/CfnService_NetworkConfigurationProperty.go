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
//   networkConfigurationProperty := &NetworkConfigurationProperty{
//   	EgressConfiguration: &EgressConfigurationProperty{
//   		EgressType: jsii.String("egressType"),
//
//   		// the properties below are optional
//   		VpcConnectorArn: jsii.String("vpcConnectorArn"),
//   	},
//   	IngressConfiguration: &IngressConfigurationProperty{
//   		IsPubliclyAccessible: jsii.Boolean(false),
//   	},
//   	IpAddressType: jsii.String("ipAddressType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-networkconfiguration.html
//
type CfnService_NetworkConfigurationProperty struct {
	// Network configuration settings for outbound message traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-networkconfiguration.html#cfn-apprunner-service-networkconfiguration-egressconfiguration
	//
	EgressConfiguration interface{} `field:"optional" json:"egressConfiguration" yaml:"egressConfiguration"`
	// Network configuration settings for inbound message traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-networkconfiguration.html#cfn-apprunner-service-networkconfiguration-ingressconfiguration
	//
	IngressConfiguration interface{} `field:"optional" json:"ingressConfiguration" yaml:"ingressConfiguration"`
	// App Runner provides you with the option to choose between *Internet Protocol version 4 (IPv4)* and *dual stack* (IPv4 and IPv6) for your incoming public network configuration.
	//
	// This is an optional parameter. If you do not specify an `IpAddressType` , it defaults to select IPv4.
	//
	// > Currently, App Runner supports dual stack for only Public endpoint. Only IPv4 is supported for Private endpoint. If you update a service that's using dual-stack Public endpoint to a Private endpoint, your App Runner service will default to support only IPv4 for Private endpoint and fail to receive traffic originating from IPv6 endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-networkconfiguration.html#cfn-apprunner-service-networkconfiguration-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
}

