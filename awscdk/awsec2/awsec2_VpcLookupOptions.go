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
//   })
//
//   // or create the cloud9 environment in the default VPC with specific instanceType
//   defaultVpc := ec2.Vpc_FromLookup(this, jsii.String("DefaultVPC"), &VpcLookupOptions{
//   	IsDefault: jsii.Boolean(true),
//   })
//   cloud9.NewEc2Environment(this, jsii.String("Cloud9Env2"), &Ec2EnvironmentProps{
//   	Vpc: defaultVpc,
//   	InstanceType: ec2.NewInstanceType(jsii.String("t3.large")),
//   })
//
//   // or specify in a different subnetSelection
//   c9env := cloud9.NewEc2Environment(this, jsii.String("Cloud9Env3"), &Ec2EnvironmentProps{
//   	Vpc: Vpc,
//   	SubnetSelection: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_NAT,
//   	},
//   })
//
//   // print the Cloud9 IDE URL in the output
//   // print the Cloud9 IDE URL in the output
//   awscdk.NewCfnOutput(this, jsii.String("URL"), &CfnOutputProps{
//   	Value: c9env.IdeUrl,
//   })
//
// Experimental.
type VpcLookupOptions struct {
	// Whether to match the default VPC.
	// Experimental.
	IsDefault *bool `field:"optional" json:"isDefault" yaml:"isDefault"`
	// Optional to override inferred region.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Optional tag for subnet group name.
	//
	// If not provided, we'll look at the aws-cdk:subnet-name tag.
	// If the subnet does not have the specified tag,
	// we'll use its type as the name.
	// Experimental.
	SubnetGroupNameTag *string `field:"optional" json:"subnetGroupNameTag" yaml:"subnetGroupNameTag"`
	// Tags on the VPC.
	//
	// The VPC must have all of these tags.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC.
	//
	// If given, will import exactly this VPC.
	// Experimental.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// The name of the VPC.
	//
	// If given, will import the VPC with this name.
	// Experimental.
	VpcName *string `field:"optional" json:"vpcName" yaml:"vpcName"`
}

