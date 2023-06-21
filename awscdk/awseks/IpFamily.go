package awseks


// EKS cluster IP family.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//
//
//   // make an ipv6 cidr
//   ipv6cidr := ec2.NewCfnVPCCidrBlock(this, jsii.String("CIDR6"), &CfnVPCCidrBlockProps{
//   	VpcId: vpc.VpcId,
//   	AmazonProvidedIpv6CidrBlock: jsii.Boolean(true),
//   })
//
//   // connect the ipv6 cidr to all vpc subnets
//   subnetcount := 0
//   subnets := []iSubnet{
//   	(SpreadElement ...vpc.publicSubnets
//   			vpc.PublicSubnets),
//   	(SpreadElement ...vpc.privateSubnets
//   			vpc.PrivateSubnets),
//   }
//   for _, subnet := range subnets {
//   	// Wait for the ipv6 cidr to complete
//   	subnet.Node.AddDependency(ipv6cidr)
//   	this._associate_subnet_with_v6_cidr(subnetcount, subnet)
//   	subnetcount++
//   }
//
//   cluster := eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
//   	Vpc: vpc,
//   	IpFamily: eks.IpFamily_IP_V6,
//   	VpcSubnets: []subnetSelection{
//   		&subnetSelection{
//   			Subnets: []*iSubnet{
//   				(SpreadElement ...vpc.publicSubnets
//   						vpc.*PublicSubnets),
//   			},
//   		},
//   	},
//   })
//
type IpFamily string

const (
	// Use IPv4 for pods and services in your cluster.
	IpFamily_IP_V4 IpFamily = "IP_V4"
	// Use IPv6 for pods and services in your cluster.
	IpFamily_IP_V6 IpFamily = "IP_V6"
)

