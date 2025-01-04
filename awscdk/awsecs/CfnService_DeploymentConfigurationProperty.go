package awsecs


// Optional deployment parameters that control how many tasks run during a deployment and the ordering of stopping and starting tasks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentConfigurationProperty := &DeploymentConfigurationProperty{
//   	Alarms: &DeploymentAlarmsProperty{
//   		AlarmNames: []*string{
//   			jsii.String("alarmNames"),
//   		},
//   		Enable: jsii.Boolean(false),
//   		Rollback: jsii.Boolean(false),
//   	},
//   	DeploymentCircuitBreaker: &DeploymentCircuitBreakerProperty{
//   		Enable: jsii.Boolean(false),
//   		Rollback: jsii.Boolean(false),
//   	},
//   	MaximumPercent: jsii.Number(123),
//   	MinimumHealthyPercent: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html
//
type CfnService_DeploymentConfigurationProperty struct {
	// Information about the CloudWatch alarms.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html#cfn-ecs-service-deploymentconfiguration-alarms
	//
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// > The deployment circuit breaker can only be used for services using the rolling update ( `ECS` ) deployment type.
	//
	// The *deployment circuit breaker* determines whether a service deployment will fail if the service can't reach a steady state. If you use the deployment circuit breaker, a service deployment will transition to a failed state and stop launching new tasks. If you use the rollback option, when a service deployment fails, the service is rolled back to the last deployment that completed successfully. For more information, see [Rolling update](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html) in the *Amazon Elastic Container Service Developer Guide*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html#cfn-ecs-service-deploymentconfiguration-deploymentcircuitbreaker
	//
	DeploymentCircuitBreaker interface{} `field:"optional" json:"deploymentCircuitBreaker" yaml:"deploymentCircuitBreaker"`
	// If a service is using the rolling update ( `ECS` ) deployment type, the `maximumPercent` parameter represents an upper limit on the number of your service's tasks that are allowed in the `RUNNING` or `PENDING` state during a deployment, as a percentage of the `desiredCount` (rounded down to the nearest integer).
	//
	// This parameter enables you to define the deployment batch size. For example, if your service is using the `REPLICA` service scheduler and has a `desiredCount` of four tasks and a `maximumPercent` value of 200%, the scheduler may start four new tasks before stopping the four older tasks (provided that the cluster resources required to do this are available). The default `maximumPercent` value for a service using the `REPLICA` service scheduler is 200%.
	//
	// The Amazon ECS scheduler uses this parameter to replace unhealthy tasks by starting replacement tasks first and then stopping the unhealthy tasks, as long as cluster resources for starting replacement tasks are available. For more information about how the scheduler replaces unhealthy tasks, see [Amazon ECS services](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html) .
	//
	// If a service is using either the blue/green ( `CODE_DEPLOY` ) or `EXTERNAL` deployment types, and tasks in the service use the EC2 launch type, the *maximum percent* value is set to the default value. The *maximum percent* value is used to define the upper limit on the number of the tasks in the service that remain in the `RUNNING` state while the container instances are in the `DRAINING` state.
	//
	// > You can't specify a custom `maximumPercent` value for a service that uses either the blue/green ( `CODE_DEPLOY` ) or `EXTERNAL` deployment types and has tasks that use the EC2 launch type.
	//
	// If the service uses either the blue/green ( `CODE_DEPLOY` ) or `EXTERNAL` deployment types, and the tasks in the service use the Fargate launch type, the maximum percent value is not used. The value is still returned when describing your service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html#cfn-ecs-service-deploymentconfiguration-maximumpercent
	//
	MaximumPercent *float64 `field:"optional" json:"maximumPercent" yaml:"maximumPercent"`
	// If a service is using the rolling update ( `ECS` ) deployment type, the `minimumHealthyPercent` represents a lower limit on the number of your service's tasks that must remain in the `RUNNING` state during a deployment, as a percentage of the `desiredCount` (rounded up to the nearest integer).
	//
	// This parameter enables you to deploy without using additional cluster capacity. For example, if your service has a `desiredCount` of four tasks and a `minimumHealthyPercent` of 50%, the service scheduler may stop two existing tasks to free up cluster capacity before starting two new tasks.
	//
	// If any tasks are unhealthy and if `maximumPercent` doesn't allow the Amazon ECS scheduler to start replacement tasks, the scheduler stops the unhealthy tasks one-by-one — using the `minimumHealthyPercent` as a constraint — to clear up capacity to launch replacement tasks. For more information about how the scheduler replaces unhealthy tasks, see [Amazon ECS services](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html) .
	//
	// For services that *do not* use a load balancer, the following should be noted:
	//
	// - A service is considered healthy if all essential containers within the tasks in the service pass their health checks.
	// - If a task has no essential containers with a health check defined, the service scheduler will wait for 40 seconds after a task reaches a `RUNNING` state before the task is counted towards the minimum healthy percent total.
	// - If a task has one or more essential containers with a health check defined, the service scheduler will wait for the task to reach a healthy status before counting it towards the minimum healthy percent total. A task is considered healthy when all essential containers within the task have passed their health checks. The amount of time the service scheduler can wait for is determined by the container health check settings.
	//
	// For services that *do* use a load balancer, the following should be noted:
	//
	// - If a task has no essential containers with a health check defined, the service scheduler will wait for the load balancer target group health check to return a healthy status before counting the task towards the minimum healthy percent total.
	// - If a task has an essential container with a health check defined, the service scheduler will wait for both the task to reach a healthy status and the load balancer target group health check to return a healthy status before counting the task towards the minimum healthy percent total.
	//
	// The default value for a replica service for `minimumHealthyPercent` is 100%. The default `minimumHealthyPercent` value for a service using the `DAEMON` service schedule is 0% for the AWS CLI , the AWS SDKs, and the APIs and 50% for the AWS Management Console.
	//
	// The minimum number of healthy tasks during a deployment is the `desiredCount` multiplied by the `minimumHealthyPercent` /100, rounded up to the nearest integer value.
	//
	// If a service is using either the blue/green ( `CODE_DEPLOY` ) or `EXTERNAL` deployment types and is running tasks that use the EC2 launch type, the *minimum healthy percent* value is set to the default value. The *minimum healthy percent* value is used to define the lower limit on the number of the tasks in the service that remain in the `RUNNING` state while the container instances are in the `DRAINING` state.
	//
	// > You can't specify a custom `minimumHealthyPercent` value for a service that uses either the blue/green ( `CODE_DEPLOY` ) or `EXTERNAL` deployment types and has tasks that use the EC2 launch type.
	//
	// If a service is using either the blue/green ( `CODE_DEPLOY` ) or `EXTERNAL` deployment types and is running tasks that use the Fargate launch type, the minimum healthy percent value is not used, although it is returned when describing your service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html#cfn-ecs-service-deploymentconfiguration-minimumhealthypercent
	//
	MinimumHealthyPercent *float64 `field:"optional" json:"minimumHealthyPercent" yaml:"minimumHealthyPercent"`
}

