package awsec2


// Specifies the Classic Load Balancers and target groups to attach to a Spot Fleet request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancersConfigProperty := &LoadBalancersConfigProperty{
//   	ClassicLoadBalancersConfig: &ClassicLoadBalancersConfigProperty{
//   		ClassicLoadBalancers: []interface{}{
//   			&ClassicLoadBalancerProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	TargetGroupsConfig: &TargetGroupsConfigProperty{
//   		TargetGroups: []interface{}{
//   			&TargetGroupProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   		},
//   	},
//   }
//
type CfnSpotFleet_LoadBalancersConfigProperty struct {
	// The Classic Load Balancers.
	ClassicLoadBalancersConfig interface{} `field:"optional" json:"classicLoadBalancersConfig" yaml:"classicLoadBalancersConfig"`
	// The target groups.
	TargetGroupsConfig interface{} `field:"optional" json:"targetGroupsConfig" yaml:"targetGroupsConfig"`
}

