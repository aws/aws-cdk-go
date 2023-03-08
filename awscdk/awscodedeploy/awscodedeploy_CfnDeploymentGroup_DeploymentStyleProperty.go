package awscodedeploy


// Information about the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentStyleProperty := &deploymentStyleProperty{
//   	deploymentOption: jsii.String("deploymentOption"),
//   	deploymentType: jsii.String("deploymentType"),
//   }
//
type CfnDeploymentGroup_DeploymentStyleProperty struct {
	// Indicates whether to route deployment traffic behind a load balancer.
	//
	// > An Amazon EC2 Application Load Balancer or Network Load Balancer is required for an Amazon ECS deployment.
	DeploymentOption *string `field:"optional" json:"deploymentOption" yaml:"deploymentOption"`
	// Indicates whether to run an in-place or blue/green deployment.
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
}

