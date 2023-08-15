package awselasticloadbalancing

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Construction properties for a LoadBalancer.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   })
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
//   	Vpc: Vpc,
//   })
//   lb.AddListener(&LoadBalancerListener{
//   	ExternalPort: jsii.Number(80),
//   })
//   lb.AddTarget(service.LoadBalancerTarget(&LoadBalancerTargetOptions{
//   	ContainerName: jsii.String("MyContainer"),
//   	ContainerPort: jsii.Number(80),
//   }))
//
type LoadBalancerProps struct {
	// VPC network of the fleet instances.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Enable Loadbalancer access logs Can be used to avoid manual work as aws console Required S3 bucket name , enabled flag Can add interval for pushing log Can set bucket prefix in order to provide folder name inside bucket.
	// Default: - disabled.
	//
	AccessLoggingPolicy *CfnLoadBalancer_AccessLoggingPolicyProperty `field:"optional" json:"accessLoggingPolicy" yaml:"accessLoggingPolicy"`
	// Whether cross zone load balancing is enabled.
	//
	// This controls whether the load balancer evenly distributes requests
	// across each availability zone.
	// Default: true.
	//
	CrossZone *bool `field:"optional" json:"crossZone" yaml:"crossZone"`
	// Health check settings for the load balancing targets.
	//
	// Not required but recommended.
	// Default: - None.
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Whether this is an internet-facing Load Balancer.
	//
	// This controls whether the LB has a public IP address assigned. It does
	// not open up the Load Balancer's security groups to public internet access.
	// Default: false.
	//
	InternetFacing *bool `field:"optional" json:"internetFacing" yaml:"internetFacing"`
	// What listeners to set up for the load balancer.
	//
	// Can also be added by .addListener()
	// Default: -.
	//
	Listeners *[]*LoadBalancerListener `field:"optional" json:"listeners" yaml:"listeners"`
	// Which subnets to deploy the load balancer.
	//
	// Can be used to define a specific set of subnets to deploy the load balancer to.
	// Useful multiple public or private subnets are covering the same availability zone.
	// Default: - Public subnets if internetFacing, Private subnets otherwise.
	//
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// What targets to load balance to.
	//
	// Can also be added by .addTarget()
	// Default: - None.
	//
	Targets *[]ILoadBalancerTarget `field:"optional" json:"targets" yaml:"targets"`
}

