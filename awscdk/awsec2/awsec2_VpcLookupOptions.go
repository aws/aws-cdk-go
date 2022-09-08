package awsec2


// Properties for looking up an existing VPC.
//
// The combination of properties must specify filter down to exactly one
// non-default VPC, otherwise an error is raised.
//
// Example:
//   // create a cloud9 ec2 environment in a new VPC
//   vpc := ec2.NewVpc(this, jsii.String("VPC"), &vpcProps{
//   	maxAzs: jsii.Number(3),
//   })
//   cloud9.NewEc2Environment(this, jsii.String("Cloud9Env"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	imageId: cloud9.imageId_AMAZON_LINUX_2,
//   })
//
//   // or create the cloud9 environment in the default VPC with specific instanceType
//   defaultVpc := ec2.vpc.fromLookup(this, jsii.String("DefaultVPC"), &vpcLookupOptions{
//   	isDefault: jsii.Boolean(true),
//   })
//   cloud9.NewEc2Environment(this, jsii.String("Cloud9Env2"), &ec2EnvironmentProps{
//   	vpc: defaultVpc,
//   	instanceType: ec2.NewInstanceType(jsii.String("t3.large")),
//   	imageId: cloud9.*imageId_AMAZON_LINUX_2,
//   })
//
//   // or specify in a different subnetSelection
//   c9env := cloud9.NewEc2Environment(this, jsii.String("Cloud9Env3"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	subnetSelection: &subnetSelection{
//   		subnetType: ec2.subnetType_PRIVATE_WITH_EGRESS,
//   	},
//   	imageId: cloud9.*imageId_AMAZON_LINUX_2,
//   })
//
//   // print the Cloud9 IDE URL in the output
//   // print the Cloud9 IDE URL in the output
//   awscdk.NewCfnOutput(this, jsii.String("URL"), &cfnOutputProps{
//   	value: c9env.ideUrl,
//   })
//
type VpcLookupOptions struct {
	// Whether to match the default VPC.
	IsDefault *bool `field:"optional" json:"isDefault" yaml:"isDefault"`
	// Optional to override inferred region.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Optional tag for subnet group name.
	//
	// If not provided, we'll look at the aws-cdk:subnet-name tag.
	// If the subnet does not have the specified tag,
	// we'll use its type as the name.
	SubnetGroupNameTag *string `field:"optional" json:"subnetGroupNameTag" yaml:"subnetGroupNameTag"`
	// Tags on the VPC.
	//
	// The VPC must have all of these tags.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC.
	//
	// If given, will import exactly this VPC.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// The name of the VPC.
	//
	// If given, will import the VPC with this name.
	VpcName *string `field:"optional" json:"vpcName" yaml:"vpcName"`
}

