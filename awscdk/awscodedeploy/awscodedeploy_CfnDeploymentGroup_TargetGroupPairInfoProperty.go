package awscodedeploy


// Information about two target groups and how traffic is routed during an Amazon ECS deployment.
//
// An optional test traffic route can be specified.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupPairInfoProperty := &targetGroupPairInfoProperty{
//   	prodTrafficRoute: &trafficRouteProperty{
//   		listenerArns: []*string{
//   			jsii.String("listenerArns"),
//   		},
//   	},
//   	targetGroups: []interface{}{
//   		&targetGroupInfoProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	testTrafficRoute: &trafficRouteProperty{
//   		listenerArns: []*string{
//   			jsii.String("listenerArns"),
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_TargetGroupPairInfoProperty struct {
	// The path used by a load balancer to route production traffic when an Amazon ECS deployment is complete.
	ProdTrafficRoute interface{} `field:"optional" json:"prodTrafficRoute" yaml:"prodTrafficRoute"`
	// One pair of target groups.
	//
	// One is associated with the original task set. The second is associated with the task set that serves traffic after the deployment is complete.
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// An optional path used by a load balancer to route test traffic after an Amazon ECS deployment.
	//
	// Validation can occur while test traffic is served during a deployment.
	TestTrafficRoute interface{} `field:"optional" json:"testTrafficRoute" yaml:"testTrafficRoute"`
}

