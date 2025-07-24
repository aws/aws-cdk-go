package awsecs


// The task placement strategy for a task or service.
//
// For more information, see [Task placement strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementStrategyProperty := &PlacementStrategyProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Field: jsii.String("field"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html
//
type CfnService_PlacementStrategyProperty struct {
	// The type of placement strategy.
	//
	// The `random` placement strategy randomly places tasks on available candidates. The `spread` placement strategy spreads placement across available candidates evenly based on the `field` parameter. The `binpack` strategy places tasks on available candidates that have the least available amount of the resource that's specified with the `field` parameter. For example, if you binpack on memory, a task is placed on the instance with the least amount of remaining memory but still enough to run the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html#cfn-ecs-service-placementstrategy-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The field to apply the placement strategy against.
	//
	// For the `spread` placement strategy, valid values are `instanceId` (or `host` , which has the same effect), or any platform or custom attribute that's applied to a container instance, such as `attribute:ecs.availability-zone` . For the `binpack` placement strategy, valid values are `cpu` and `memory` . For the `random` placement strategy, this field is not used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html#cfn-ecs-service-placementstrategy-field
	//
	Field *string `field:"optional" json:"field" yaml:"field"`
}

