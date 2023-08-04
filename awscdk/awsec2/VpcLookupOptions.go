package awsec2


// Properties for looking up an existing VPC.
//
// The combination of properties must specify filter down to exactly one
// non-default VPC, otherwise an error is raised.
//
// Example:
//   // create a cloud9 ec2 environment in a new VPC
//   vpc := ec2.NewVpc(this, jsii.String("VPC"), &VpcProps{
//   	MaxAzs: jsii.Number(3),
//   })
//   cloud9.NewEc2Environment(this, jsii.String("Cloud9Env"), &Ec2EnvironmentProps{
//   	Vpc: Vpc,
//   	ImageId: cloud9.ImageId_AMAZON_LINUX_2,
//   })
//
//   // or create the cloud9 environment in the default VPC with specific instanceType
//   defaultVpc := ec2.Vpc_FromLookup(this, jsii.String("DefaultVPC"), &VpcLookupOptions{
//   	IsDefault: jsii.Boolean(true),
//   })
//   cloud9.NewEc2Environment(this, jsii.String("Cloud9Env2"), &Ec2EnvironmentProps{
//   	Vpc: defaultVpc,
//   	InstanceType: ec2.NewInstanceType(jsii.String("t3.large")),
//   	ImageId: cloud9.ImageId_AMAZON_LINUX_2,
//   })
//
//   // or specify in a different subnetSelection
//   c9env := cloud9.NewEc2Environment(this, jsii.String("Cloud9Env3"), &Ec2EnvironmentProps{
//   	Vpc: Vpc,
//   	SubnetSelection: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	},
//   	ImageId: cloud9.ImageId_AMAZON_LINUX_2,
//   })
//
//   // print the Cloud9 IDE URL in the output
//   // print the Cloud9 IDE URL in the output
//   awscdk.NewCfnOutput(this, jsii.String("URL"), &CfnOutputProps{
//   	Value: c9env.IdeUrl,
//   })
//
type VpcLookupOptions struct {
	// Whether to match the default VPC.
	// Default: Don't care whether we return the default VPC.
	//
	IsDefault *bool `field:"optional" json:"isDefault" yaml:"isDefault"`
	// The ID of the AWS account that owns the VPC.
	// Default: the account id of the parent stack.
	//
	OwnerAccountId *string `field:"optional" json:"ownerAccountId" yaml:"ownerAccountId"`
	// Optional to override inferred region.
	// Default: Current stack's environment region.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Whether to look up whether a VPN Gateway is attached to the looked up VPC.
	//
	// You can set this to `false` if you know the VPC does not have a VPN Gateway
	// attached, in order to avoid an API call.
	//
	// If you change this property from `false` to `true` or undefined, you may
	// need to clear the corresponding context entry in `cdk.context.json` in
	// order to trigger a new lookup.
	// Default: true.
	//
	ReturnVpnGateways *bool `field:"optional" json:"returnVpnGateways" yaml:"returnVpnGateways"`
	// Optional tag for subnet group name.
	//
	// If not provided, we'll look at the aws-cdk:subnet-name tag.
	// If the subnet does not have the specified tag,
	// we'll use its type as the name.
	// Default: aws-cdk:subnet-name.
	//
	SubnetGroupNameTag *string `field:"optional" json:"subnetGroupNameTag" yaml:"subnetGroupNameTag"`
	// Tags on the VPC.
	//
	// The VPC must have all of these tags.
	// Default: Don't filter on tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC.
	//
	// If given, will import exactly this VPC.
	// Default: Don't filter on vpcId.
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// The name of the VPC.
	//
	// If given, will import the VPC with this name.
	// Default: Don't filter on vpcName.
	//
	VpcName *string `field:"optional" json:"vpcName" yaml:"vpcName"`
}

