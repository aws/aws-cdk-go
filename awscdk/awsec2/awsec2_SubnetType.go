package awsec2


// The type of Subnet.
//
// Example:
//   var vpc vpc
//
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	MasterUser: &Login{
//   		Username: jsii.String("myuser"),
//   		 // NOTE: 'admin' is reserved by DocumentDB
//   		ExcludeCharacters: jsii.String("\"@/:"),
//   		 // optional, defaults to the set "\"@/" and is also used for eventually created rotations
//   		SecretName: jsii.String("/myapp/mydocdb/masteruser"),
//   	},
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R5, ec2.InstanceSize_LARGE),
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   	Vpc: Vpc,
//   })
//
// Experimental.
type SubnetType string

const (
	// Isolated Subnets do not route traffic to the Internet (in this VPC), and as such, do not require NAT gateways.
	//
	// Isolated subnets can only connect to or be connected to from other
	// instances in the same VPC. A default VPC configuration will not include
	// isolated subnets.
	//
	// This can be good for subnets with RDS or Elasticache instances,
	// or which route Internet traffic through a peer VPC.
	// Deprecated: use `SubnetType.PRIVATE_ISOLATED`
	SubnetType_ISOLATED SubnetType = "ISOLATED"
	// Isolated Subnets do not route traffic to the Internet (in this VPC), and as such, do not require NAT gateways.
	//
	// Isolated subnets can only connect to or be connected to from other
	// instances in the same VPC. A default VPC configuration will not include
	// isolated subnets.
	//
	// This can be good for subnets with RDS or Elasticache instances,
	// or which route Internet traffic through a peer VPC.
	// Experimental.
	SubnetType_PRIVATE_ISOLATED SubnetType = "PRIVATE_ISOLATED"
	// Subnet that routes to the internet, but not vice versa.
	//
	// Instances in a private subnet can connect to the Internet, but will not
	// allow connections to be initiated from the Internet. NAT Gateway(s) are
	// required with this subnet type to route the Internet traffic through.
	// If a NAT Gateway is not required or desired, use `SubnetType.PRIVATE_ISOLATED` instead.
	//
	// By default, a NAT gateway is created in every public subnet for maximum availability.
	// Be aware that you will be charged for NAT gateways.
	//
	// Normally a Private subnet will use a NAT gateway in the same AZ, but
	// if `natGateways` is used to reduce the number of NAT gateways, a NAT
	// gateway from another AZ will be used instead.
	// Deprecated: use `PRIVATE_WITH_NAT`.
	SubnetType_PRIVATE SubnetType = "PRIVATE"
	// Subnet that routes to the internet (via a NAT gateway), but not vice versa.
	//
	// Instances in a private subnet can connect to the Internet, but will not
	// allow connections to be initiated from the Internet. NAT Gateway(s) are
	// required with this subnet type to route the Internet traffic through.
	// If a NAT Gateway is not required or desired, use `SubnetType.PRIVATE_ISOLATED` instead.
	//
	// By default, a NAT gateway is created in every public subnet for maximum availability.
	// Be aware that you will be charged for NAT gateways.
	//
	// Normally a Private subnet will use a NAT gateway in the same AZ, but
	// if `natGateways` is used to reduce the number of NAT gateways, a NAT
	// gateway from another AZ will be used instead.
	// Experimental.
	SubnetType_PRIVATE_WITH_NAT SubnetType = "PRIVATE_WITH_NAT"
	// Subnet connected to the Internet.
	//
	// Instances in a Public subnet can connect to the Internet and can be
	// connected to from the Internet as long as they are launched with public
	// IPs (controlled on the AutoScalingGroup or other constructs that launch
	// instances).
	//
	// Public subnets route outbound traffic via an Internet Gateway.
	// Experimental.
	SubnetType_PUBLIC SubnetType = "PUBLIC"
)

