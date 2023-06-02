package awsapprunner


// Specifications for the customer’s VPC and related PrivateLink VPC endpoint that are used to associate with the VPC Ingress Connection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressVpcConfigurationProperty := &IngressVpcConfigurationProperty{
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
type CfnVpcIngressConnection_IngressVpcConfigurationProperty struct {
	// The ID of the VPC endpoint that your App Runner service connects to.
	VpcEndpointId *string `field:"required" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The ID of the VPC that is used for the VPC endpoint.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

