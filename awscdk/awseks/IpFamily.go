package awseks


// EKS cluster IP family.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv33"
//   var vpc vpc
//
//
//   func associateSubnetWithV6Cidr(vpc *vpc, count *f64, subnet iSubnet) {
//   	cfnSubnet := *subnet.Node.defaultChild.(cfnSubnet)
//   	cfnSubnet.Ipv6CidrBlock = awscdk.Fn_Select(count, awscdk.Fn_Cidr(awscdk.Fn_Select(jsii.Number(0), vpc.VpcIpv6CidrBlocks), jsii.Number(256), (jsii.Number(128 - 64)).toString()))
//   	cfnSubnet.AssignIpv6AddressOnCreation = true
//   }
//
//   // make an ipv6 cidr
//   ipv6cidr := ec2.NewCfnVPCCidrBlock(this, jsii.String("CIDR6"), &CfnVPCCidrBlockProps{
//   	VpcId: vpc.VpcId,
//   	AmazonProvidedIpv6CidrBlock: jsii.Boolean(true),
//   })
//
//   // connect the ipv6 cidr to all vpc subnets
//   subnetcount := 0
//   subnets := vpc.PublicSubnets.concat(vpc.PrivateSubnets)
//   for _, subnet := range subnets {
//   	// Wait for the ipv6 cidr to complete
//   	subnet.Node.AddDependency(ipv6cidr)
//   	associateSubnetWithV6Cidr(vpc, subnetcount, subnet)
//   	subnetcount = subnetcount + 1
//   }
//
//   cluster := eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_33(),
//   	Vpc: vpc,
//   	IpFamily: eks.IpFamily_IP_V6,
//   	VpcSubnets: []subnetSelection{
//   		&subnetSelection{
//   			Subnets: vpc.*PublicSubnets,
//   		},
//   	},
//   	KubectlLayer: kubectlv33.NewKubectlV33Layer(this, jsii.String("kubectl")),
//   })
//
type IpFamily string

const (
	// Use IPv4 for pods and services in your cluster.
	IpFamily_IP_V4 IpFamily = "IP_V4"
	// Use IPv6 for pods and services in your cluster.
	IpFamily_IP_V6 IpFamily = "IP_V6"
)

