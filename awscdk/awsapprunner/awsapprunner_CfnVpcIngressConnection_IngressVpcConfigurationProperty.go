package awsapprunner


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressVpcConfigurationProperty := &ingressVpcConfigurationProperty{
//   	vpcEndpointId: jsii.String("vpcEndpointId"),
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnVpcIngressConnection_IngressVpcConfigurationProperty struct {
	// `CfnVpcIngressConnection.IngressVpcConfigurationProperty.VpcEndpointId`.
	VpcEndpointId *string `field:"required" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// `CfnVpcIngressConnection.IngressVpcConfigurationProperty.VpcId`.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

