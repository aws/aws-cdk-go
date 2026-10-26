package awsecs


// You can use early success criteria only with rolling deployment strategy.
//
// The configuration that determines when a rolling update deployment is considered successful. Early success criteria defines the percentage of tasks that must be healthy before a deployment completes. It also controls whether Amazon ECS must remove the previous tasks before a deployment completes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   deploymentEarlySuccessCriteriaProperty := &DeploymentEarlySuccessCriteriaProperty{
//   	Enable: jsii.Boolean(false),
//   	HealthyPercent: jsii.Number(123),
//   	SourceServiceRevisionCleanup: jsii.String("sourceServiceRevisionCleanup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentearlysuccesscriteria.html
//
type CfnServicePropsMixin_DeploymentEarlySuccessCriteriaProperty struct {
	// Specifies whether to use the early success criteria for the service deployment.
	//
	// When set to ``false``, the deployment uses the default behavior, where Amazon ECS considers the deployment successful when the target service revision fully stabilizes and the previous tasks are removed. The default value is ``false``.
	//  When set to ``true``, Amazon ECS monitors the deployment to meet early success criteria. You must also specify ``healthyPercent`` and ``sourceServiceRevisionCleanup``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentearlysuccesscriteria.html#cfn-ecs-service-deploymentearlysuccesscriteria-enable
	//
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// The percentage of healthy tasks that the target service revision must reach before Amazon ECS considers the deployment successful.
	//
	// This percentage is relative to the service's ``desiredCount`` and must be an integer between ``0`` and ``100``. This value must be greater than or equal to the ``minimumHealthyPercent`` value.
	//  After this percentage of tasks is healthy and the bake time elapses, Amazon ECS completes the deployment. Amazon ECS continues to scale the target service revision to 100 percent in the background.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentearlysuccesscriteria.html#cfn-ecs-service-deploymentearlysuccesscriteria-healthypercent
	//
	HealthyPercent *float64 `field:"optional" json:"healthyPercent" yaml:"healthyPercent"`
	// The time when Amazon ECS removes the source revisions' tasks relative to deployment completion.
	//
	// The valid values are:
	//   +  ``BLOCKING``—Amazon ECS removes the previous tasks before it marks the deployment as successful.
	//   +  ``DEFERRED``—Amazon ECS marks the deployment successful, and then removes the previous tasks in the background.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentearlysuccesscriteria.html#cfn-ecs-service-deploymentearlysuccesscriteria-sourceservicerevisioncleanup
	//
	SourceServiceRevisionCleanup *string `field:"optional" json:"sourceServiceRevisionCleanup" yaml:"sourceServiceRevisionCleanup"`
}

