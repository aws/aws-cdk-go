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
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
//   	}),
//   	Instances: jsii.Number(2),
//   	InstanceProps: &InstanceProps{
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
//   		VpcSubnets: &SubnetSelection{
//   			SubnetType: ec2.SubnetType_PUBLIC,
//   		},
//   		Vpc: *Vpc,
//   	},
//   })
//
type SubnetSelection struct {
	// Select subnets only in the given AZs.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// If true, return at most one subnet per AZ.
	OnePerAz *bool `field:"optional" json:"onePerAz" yaml:"onePerAz"`
	// List of provided subnet filters.
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
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// Explicitly select individual subnets.
	//
	// Use this if you don't want to automatically use all subnets in
	// a group, but have a need to control selection down to
	// individual subnets.
	//
	// Cannot be specified together with `subnetType` or `subnetGroupName`.
	Subnets *[]ISubnet `field:"optional" json:"subnets" yaml:"subnets"`
	// Select all subnets of the given type.
	//
	// At most one of `subnetType` and `subnetGroupName` can be supplied.
	SubnetType SubnetType `field:"optional" json:"subnetType" yaml:"subnetType"`
}

