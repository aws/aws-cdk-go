package awscodedeploy


// The `ELBInfo` property type specifies information about the Elastic Load Balancing load balancer used for an CodeDeploy deployment group.
//
// If you specify the `ELBInfo` property, the `DeploymentStyle.DeploymentOption` property must be set to `WITH_TRAFFIC_CONTROL` for AWS CodeDeploy to route your traffic using the specified load balancers.
//
// `ELBInfo` is a property of the [AWS CodeDeploy DeploymentGroup LoadBalancerInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eLBInfoProperty := &eLBInfoProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnDeploymentGroup_ELBInfoProperty struct {
	// For blue/green deployments, the name of the load balancer that is used to route traffic from original instances to replacement instances in a blue/green deployment.
	//
	// For in-place deployments, the name of the load balancer that instances are deregistered from so they are not serving traffic during a deployment, and then re-registered with after the deployment is complete.
	//
	// > AWS CloudFormation supports blue/green deployments on AWS Lambda compute platforms only.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

