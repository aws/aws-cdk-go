package awsec2


// The types of IP addresses provisioned in the VPC.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // The VPC and subnet must have associated IPv6 CIDR blocks.
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	IpProtocol: ec2.IpProtocol_DUAL_STACK,
//   })
//   cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   networkLoadbalancedFargateService := ecsPatterns.NewNetworkLoadBalancedFargateService(this, jsii.String("NlbFargateService"), &NetworkLoadBalancedFargateServiceProps{
//   	Cluster: Cluster,
//   	TaskImageOptions: &NetworkLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	IpAddressType: elbv2.IpAddressType_DUAL_STACK,
//   })
//
//   networkLoadbalancedEc2Service := ecsPatterns.NewNetworkLoadBalancedEc2Service(this, jsii.String("NlbEc2Service"), &NetworkLoadBalancedEc2ServiceProps{
//   	Cluster: Cluster,
//   	TaskImageOptions: &NetworkLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_*FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	IpAddressType: elbv2.IpAddressType_DUAL_STACK,
//   })
//
type IpProtocol string

const (
	// The vpc will be configured with only IPv4 addresses.
	//
	// This is the default protocol if unset.
	IpProtocol_IPV4_ONLY IpProtocol = "IPV4_ONLY"
	// The vpc will have both IPv4 and IPv6 addresses.
	//
	// Unless specified, public IPv4 addresses will not be auto assigned,
	// an egress only internet gateway (EIGW) will be created and configured,
	// and NATs and internet gateways (IGW) will be configured with IPv6 addresses.
	IpProtocol_DUAL_STACK IpProtocol = "DUAL_STACK"
)

