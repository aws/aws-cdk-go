package awsecs


// The `PlacementStrategy` property specifies the task placement strategy for a task or service.
//
// For more information, see [Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementStrategyProperty := &placementStrategyProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	field: jsii.String("field"),
//   }
//
type CfnService_PlacementStrategyProperty struct {
	// The type of placement strategy.
	//
	// The `random` placement strategy randomly places tasks on available candidates. The `spread` placement strategy spreads placement across available candidates evenly based on the `field` parameter. The `binpack` strategy places tasks on available candidates that have the least available amount of the resource that's specified with the `field` parameter. For example, if you binpack on memory, a task is placed on the instance with the least amount of remaining memory but still enough to run the task.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The field to apply the placement strategy against.
	//
	// For the `spread` placement strategy, valid values are `instanceId` (or `host` , which has the same effect), or any platform or custom attribute that's applied to a container instance, such as `attribute:ecs.availability-zone` . For the `binpack` placement strategy, valid values are `cpu` and `memory` . For the `random` placement strategy, this field is not used.
	Field *string `field:"optional" json:"field" yaml:"field"`
}

