package awscodedeploy


// Information about how traffic is rerouted to instances in a replacement environment in a blue/green deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentReadyOptionProperty := &deploymentReadyOptionProperty{
//   	actionOnTimeout: jsii.String("actionOnTimeout"),
//   	waitTimeInMinutes: jsii.Number(123),
//   }
//
type CfnDeploymentGroup_DeploymentReadyOptionProperty struct {
	// Information about when to reroute traffic from an original environment to a replacement environment in a blue/green deployment.
	//
	// - CONTINUE_DEPLOYMENT: Register new instances with the load balancer immediately after the new application revision is installed on the instances in the replacement environment.
	// - STOP_DEPLOYMENT: Do not register new instances with a load balancer unless traffic rerouting is started using [ContinueDeployment](https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_ContinueDeployment.html) . If traffic rerouting is not started before the end of the specified wait period, the deployment status is changed to Stopped.
	ActionOnTimeout *string `field:"optional" json:"actionOnTimeout" yaml:"actionOnTimeout"`
	// The number of minutes to wait before the status of a blue/green deployment is changed to Stopped if rerouting is not started manually.
	//
	// Applies only to the `STOP_DEPLOYMENT` option for `actionOnTimeout` .
	WaitTimeInMinutes *float64 `field:"optional" json:"waitTimeInMinutes" yaml:"waitTimeInMinutes"`
}

