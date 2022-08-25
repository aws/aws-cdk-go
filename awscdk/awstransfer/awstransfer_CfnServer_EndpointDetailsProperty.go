package awstransfer


// The virtual private cloud (VPC) endpoint settings that are configured for your server.
//
// When you host your endpoint within your VPC, you can make it accessible only to resources within your VPC, or you can attach Elastic IP addresses and make it accessible to clients over the internet. Your VPC's default security groups are automatically assigned to your endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointDetailsProperty := &endpointDetailsProperty{
//   	addressAllocationIds: []*string{
//   		jsii.String("addressAllocationIds"),
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	vpcEndpointId: jsii.String("vpcEndpointId"),
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnServer_EndpointDetailsProperty struct {
	// A list of address allocation IDs that are required to attach an Elastic IP address to your server's endpoint.
	//
	// > This property can only be set when `EndpointType` is set to `VPC` and it is only valid in the `UpdateServer` API.
	AddressAllocationIds *[]*string `field:"optional" json:"addressAllocationIds" yaml:"addressAllocationIds"`
	// A list of security groups IDs that are available to attach to your server's endpoint.
	//
	// > This property can only be set when `EndpointType` is set to `VPC` .
	// >
	// > You can edit the `SecurityGroupIds` property in the [UpdateServer](https://docs.aws.amazon.com/transfer/latest/userguide/API_UpdateServer.html) API only if you are changing the `EndpointType` from `PUBLIC` or `VPC_ENDPOINT` to `VPC` . To change security groups associated with your server's VPC endpoint after creation, use the Amazon EC2 [ModifyVpcEndpoint](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyVpcEndpoint.html) API.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of subnet IDs that are required to host your server endpoint in your VPC.
	//
	// > This property can only be set when `EndpointType` is set to `VPC` .
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC endpoint.
	//
	// > This property can only be set when `EndpointType` is set to `VPC_ENDPOINT` .
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The VPC ID of the virtual private cloud in which the server's endpoint will be hosted.
	//
	// > This property can only be set when `EndpointType` is set to `VPC` .
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

