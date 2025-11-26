package previewawsmediaconnectmixins


// The details of a VPC interface.
//
// > When configuring VPC interfaces for NDI outputs, keep in mind the following:
// >
// > - VPC interfaces must be defined as nested attributes within the `AWS::MediaConnect::Flow` resource, and not within the top-level `AWS::MediaConnect::FlowVpcInterface` resource.
// > - There's a maximum limit of three VPC interfaces for each flow. If you've already reached this limit, you can't update the flow to use a different VPC interface without first removing an existing one.
// >
// > To update your VPC interfaces in this scenario, you must first remove the VPC interface that’s not being used. Next, add the new VPC interfaces. Lastly, update the `VpcInterfaceAdapter` in the `NDIConfig` property. These changes must be performed as separate manual operations and cannot be done through a single template update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcInterfaceProperty := &VpcInterfaceProperty{
//   	Name: jsii.String("name"),
//   	NetworkInterfaceIds: []*string{
//   		jsii.String("networkInterfaceIds"),
//   	},
//   	NetworkInterfaceType: jsii.String("networkInterfaceType"),
//   	RoleArn: jsii.String("roleArn"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-vpcinterface.html
//
type CfnFlowPropsMixin_VpcInterfaceProperty struct {
	// Immutable and has to be a unique against other VpcInterfaces in this Flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-vpcinterface.html#cfn-mediaconnect-flow-vpcinterface-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// IDs of the network interfaces created in customer's account by MediaConnect .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-vpcinterface.html#cfn-mediaconnect-flow-vpcinterface-networkinterfaceids
	//
	NetworkInterfaceIds *[]*string `field:"optional" json:"networkInterfaceIds" yaml:"networkInterfaceIds"`
	// The type of network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-vpcinterface.html#cfn-mediaconnect-flow-vpcinterface-networkinterfacetype
	//
	NetworkInterfaceType *string `field:"optional" json:"networkInterfaceType" yaml:"networkInterfaceType"`
	// A role Arn MediaConnect can assume to create ENIs in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-vpcinterface.html#cfn-mediaconnect-flow-vpcinterface-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Security Group IDs to be used on ENI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-vpcinterface.html#cfn-mediaconnect-flow-vpcinterface-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Subnet must be in the AZ of the Flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-vpcinterface.html#cfn-mediaconnect-flow-vpcinterface-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

