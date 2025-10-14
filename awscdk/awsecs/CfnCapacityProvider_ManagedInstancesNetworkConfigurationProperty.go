package awsecs


// The network configuration for Amazon ECS Managed Instances.
//
// This specifies the VPC subnets and security groups that instances use for network connectivity. Amazon ECS Managed Instances support multiple network modes including `awsvpc` (instances receive ENIs for task isolation), `host` (instances share network namespace with tasks), and `none` (no external network connectivity), ensuring backward compatibility for migrating workloads from Fargate or Amazon EC2.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedInstancesNetworkConfigurationProperty := &ManagedInstancesNetworkConfigurationProperty{
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesnetworkconfiguration.html
//
type CfnCapacityProvider_ManagedInstancesNetworkConfigurationProperty struct {
	// The list of subnet IDs where Amazon ECS can launch Amazon ECS Managed Instances.
	//
	// Instances are distributed across the specified subnets for high availability. All subnets must be in the same VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesnetworkconfiguration.html#cfn-ecs-capacityprovider-managedinstancesnetworkconfiguration-subnets
	//
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// The list of security group IDs to apply to Amazon ECS Managed Instances.
	//
	// These security groups control the network traffic allowed to and from the instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesnetworkconfiguration.html#cfn-ecs-capacityprovider-managedinstancesnetworkconfiguration-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

