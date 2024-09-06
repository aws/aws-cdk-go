package awselasticloadbalancingv2


// What kind of addresses to allocate to the load balancer.
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
//   service := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("myService"), &ApplicationLoadBalancedFargateServiceProps{
//   	Cluster: Cluster,
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	IpAddressType: elbv2.IpAddressType_DUAL_STACK,
//   })
//
//   applicationLoadBalancedEc2Service := ecsPatterns.NewApplicationLoadBalancedEc2Service(this, jsii.String("myService"), &ApplicationLoadBalancedEc2ServiceProps{
//   	Cluster: Cluster,
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_*FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	IpAddressType: elbv2.IpAddressType_DUAL_STACK,
//   })
//
type IpAddressType string

const (
	// Allocate IPv4 addresses.
	IpAddressType_IPV4 IpAddressType = "IPV4"
	// Allocate both IPv4 and IPv6 addresses.
	IpAddressType_DUAL_STACK IpAddressType = "DUAL_STACK"
)

