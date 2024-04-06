package awsec2


// Properties for looking up an existing VPC.
//
// The combination of properties must specify filter down to exactly one
// non-default VPC, otherwise an error is raised.
//
// Example:
//   vpc := ec2.Vpc_FromLookup(this, jsii.String("Vpc"), &VpcLookupOptions{
//   	IsDefault: jsii.Boolean(true),
//   })
//   cluster := ecs.NewCluster(this, jsii.String("ECSCluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &TaskDefinitionProps{
//   	Compatibility: ecs.Compatibility_EC2,
//   })
//
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("foo/bar")),
//   	MemoryLimitMiB: jsii.Number(256),
//   })
//
//   runTask := tasks.NewEcsRunTask(this, jsii.String("Run"), &EcsRunTaskProps{
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	LaunchTarget: tasks.NewEcsEc2LaunchTarget(),
//   	EnableExecuteCommand: jsii.Boolean(true),
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

