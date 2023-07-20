package awsecs


// One of the methods which provide a way for you to quickly identify when a deployment has failed, and then to optionally roll back the failure to the last working deployment.
//
// When the alarms are generated, Amazon ECS sets the service deployment to failed. Set the rollback parameter to have Amazon ECS to roll back your service to the last completed deployment after a failure.
//
// You can only use the `DeploymentAlarms` method to detect failures when the `DeploymentController` is set to `ECS` (rolling update).
//
// For more information, see [Rolling update](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html) in the **Amazon Elastic Container Service Developer Guide** .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentAlarmsProperty := &DeploymentAlarmsProperty{
//   	AlarmNames: []*string{
//   		jsii.String("alarmNames"),
//   	},
//   	Enable: jsii.Boolean(false),
//   	Rollback: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentalarms.html
//
type CfnService_DeploymentAlarmsProperty struct {
	// One or more CloudWatch alarm names.
	//
	// Use a "," to separate the alarms.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentalarms.html#cfn-ecs-service-deploymentalarms-alarmnames
	//
	AlarmNames *[]*string `field:"required" json:"alarmNames" yaml:"alarmNames"`
	// Determines whether to use the CloudWatch alarm option in the service deployment process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentalarms.html#cfn-ecs-service-deploymentalarms-enable
	//
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Determines whether to configure Amazon ECS to roll back the service if a service deployment fails.
	//
	// If rollback is used, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentalarms.html#cfn-ecs-service-deploymentalarms-rollback
	//
	Rollback interface{} `field:"required" json:"rollback" yaml:"rollback"`
}

