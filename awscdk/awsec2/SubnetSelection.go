package awsec2


// Customize subnets that are selected for placement of ENIs.
//
// Constructs that allow customization of VPC placement use parameters of this
// type to provide placement settings.
//
// By default, the instances are placed in the private subnets.
//
// Example:
//   var vpc vpc
//
//
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	MasterUser: &Login{
//   		Username: jsii.String("myuser"),
//   	},
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   	Vpc: Vpc,
//   	RemovalPolicy: awscdk.RemovalPolicy_SNAPSHOT,
//   	InstanceRemovalPolicy: awscdk.RemovalPolicy_RETAIN,
//   })
//
type SubnetSelection struct {
	// Select subnets only in the given AZs.
	// Default: no filtering on AZs is done.
	//
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// If true, return at most one subnet per AZ.
	// Default: false.
	//
	OnePerAz *bool `field:"optional" json:"onePerAz" yaml:"onePerAz"`
	// List of provided subnet filters.
	// Default: - none.
	//
	SubnetFilters *[]SubnetFilter `field:"optional" json:"subnetFilters" yaml:"subnetFilters"`
	// Select the subnet group with the given name.
	//
	// Select the subnet group with the given name. This only needs
	// to be used if you have multiple subnet groups of the same type
	// and you need to distinguish between them. Otherwise, prefer
	// `subnetType`.
	//
	// This field does not select individual subnets, it selects all subnets that
	// share the given subnet group name. This is the name supplied in
	// `subnetConfiguration`.
	//
	// At most one of `subnetType` and `subnetGroupName` can be supplied.
	// Default: - Selection by type instead of by name.
	//
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// Explicitly select individual subnets.
	//
	// Use this if you don't want to automatically use all subnets in
	// a group, but have a need to control selection down to
	// individual subnets.
	//
	// Cannot be specified together with `subnetType` or `subnetGroupName`.
	// Default: - Use all subnets in a selected group (all private subnets by default).
	//
	Subnets *[]ISubnet `field:"optional" json:"subnets" yaml:"subnets"`
	// Select all subnets of the given type.
	//
	// At most one of `subnetType` and `subnetGroupName` can be supplied.
	// Default: SubnetType.PRIVATE_WITH_EGRESS (or ISOLATED or PUBLIC if there are no PRIVATE_WITH_EGRESS subnets)
	//
	SubnetType SubnetType `field:"optional" json:"subnetType" yaml:"subnetType"`
}

