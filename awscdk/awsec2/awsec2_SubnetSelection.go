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
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	masterUser: &login{
//   		username: jsii.String("myuser"),
//   		 // NOTE: 'admin' is reserved by DocumentDB
//   		excludeCharacters: jsii.String("\"@/:"),
//   		 // optional, defaults to the set "\"@/" and is also used for eventually created rotations
//   		secretName: jsii.String("/myapp/mydocdb/masteruser"),
//   	},
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_R5, ec2.instanceSize_LARGE),
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	},
//   	vpc: vpc,
//   })
//
// Experimental.
type SubnetSelection struct {
	// Select subnets only in the given AZs.
	// Experimental.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// If true, return at most one subnet per AZ.
	// Experimental.
	OnePerAz *bool `field:"optional" json:"onePerAz" yaml:"onePerAz"`
	// List of provided subnet filters.
	// Experimental.
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
	// Experimental.
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// Alias for `subnetGroupName`.
	//
	// Select the subnet group with the given name. This only needs
	// to be used if you have multiple subnet groups of the same type
	// and you need to distinguish between them.
	// Deprecated: Use `subnetGroupName` instead.
	SubnetName *string `field:"optional" json:"subnetName" yaml:"subnetName"`
	// Explicitly select individual subnets.
	//
	// Use this if you don't want to automatically use all subnets in
	// a group, but have a need to control selection down to
	// individual subnets.
	//
	// Cannot be specified together with `subnetType` or `subnetGroupName`.
	// Experimental.
	Subnets *[]ISubnet `field:"optional" json:"subnets" yaml:"subnets"`
	// Select all subnets of the given type.
	//
	// At most one of `subnetType` and `subnetGroupName` can be supplied.
	// Experimental.
	SubnetType SubnetType `field:"optional" json:"subnetType" yaml:"subnetType"`
}

