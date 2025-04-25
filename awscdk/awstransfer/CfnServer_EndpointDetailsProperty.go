package awstransfer


// The virtual private cloud (VPC) endpoint settings that are configured for your server.
//
// When you host your endpoint within your VPC, you can make your endpoint accessible only to resources within your VPC, or you can attach Elastic IP addresses and make your endpoint accessible to clients over the internet. Your VPC's default security groups are automatically assigned to your endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointDetailsProperty := &EndpointDetailsProperty{
//   	AddressAllocationIds: []*string{
//   		jsii.String("addressAllocationIds"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html
//
type CfnServer_EndpointDetailsProperty struct {
	// A list of address allocation IDs that are required to attach an Elastic IP address to your server's endpoint.
	//
	// An address allocation ID corresponds to the allocation ID of an Elastic IP address. This value can be retrieved from the `allocationId` field from the Amazon EC2 [Address](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Address.html) data type. One way to retrieve this value is by calling the EC2 [DescribeAddresses](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAddresses.html) API.
	//
	// This parameter is optional. Set this parameter if you want to make your VPC endpoint public-facing. For details, see [Create an internet-facing endpoint for your server](https://docs.aws.amazon.com/transfer/latest/userguide/create-server-in-vpc.html#create-internet-facing-endpoint) .
	//
	// > This property can only be set as follows:
	// >
	// > - `EndpointType` must be set to `VPC`
	// > - The Transfer Family server must be offline.
	// > - You cannot set this parameter for Transfer Family servers that use the FTP protocol.
	// > - The server must already have `SubnetIds` populated ( `SubnetIds` and `AddressAllocationIds` cannot be updated simultaneously).
	// > - `AddressAllocationIds` can't contain duplicates, and must be equal in length to `SubnetIds` . For example, if you have three subnet IDs, you must also specify three address allocation IDs.
	// > - Call the `UpdateServer` API to set or change this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-addressallocationids
	//
	AddressAllocationIds *[]*string `field:"optional" json:"addressAllocationIds" yaml:"addressAllocationIds"`
	// A list of security groups IDs that are available to attach to your server's endpoint.
	//
	// > This property can only be set when `EndpointType` is set to `VPC` .
	// >
	// > You can edit the `SecurityGroupIds` property in the [UpdateServer](https://docs.aws.amazon.com/transfer/latest/userguide/API_UpdateServer.html) API only if you are changing the `EndpointType` from `PUBLIC` or `VPC_ENDPOINT` to `VPC` . To change security groups associated with your server's VPC endpoint after creation, use the Amazon EC2 [ModifyVpcEndpoint](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyVpcEndpoint.html) API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of subnet IDs that are required to host your server endpoint in your VPC.
	//
	// > This property can only be set when `EndpointType` is set to `VPC` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC endpoint.
	//
	// > This property can only be set when `EndpointType` is set to `VPC_ENDPOINT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-vpcendpointid
	//
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The VPC ID of the virtual private cloud in which the server's endpoint will be hosted.
	//
	// > This property can only be set when `EndpointType` is set to `VPC` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

