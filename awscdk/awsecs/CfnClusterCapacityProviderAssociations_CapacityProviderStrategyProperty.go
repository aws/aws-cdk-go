package awsecs


// The `CapacityProviderStrategy` property specifies the details of the default capacity provider strategy for the cluster.
//
// When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderStrategyProperty := &CapacityProviderStrategyProperty{
//   	CapacityProvider: jsii.String("capacityProvider"),
//
//   	// the properties below are optional
//   	Base: jsii.Number(123),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-clustercapacityproviderassociations-capacityproviderstrategy.html
//
type CfnClusterCapacityProviderAssociations_CapacityProviderStrategyProperty struct {
	// The short name of the capacity provider.
	//
	// This can be either an AWS managed capacity provider ( `FARGATE` or `FARGATE_SPOT` ) or the name of a custom capacity provider that you created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-clustercapacityproviderassociations-capacityproviderstrategy.html#cfn-ecs-clustercapacityproviderassociations-capacityproviderstrategy-capacityprovider
	//
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// The *base* value designates how many tasks, at a minimum, to run on the specified capacity provider for each service.
	//
	// Only one capacity provider in a capacity provider strategy can have a *base* defined. If no value is specified, the default value of `0` is used.
	//
	// Base value characteristics:
	//
	// - Only one capacity provider in a strategy can have a base defined
	// - The default value is `0` if not specified
	// - The valid range is 0 to 100,000
	// - Base requirements are satisfied first before weight distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-clustercapacityproviderassociations-capacityproviderstrategy.html#cfn-ecs-clustercapacityproviderassociations-capacityproviderstrategy-base
	//
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// The *weight* value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The `weight` value is taken into consideration after the `base` value, if defined, is satisfied.
	//
	// If no `weight` value is specified, the default value of `0` is used. When multiple capacity providers are specified within a capacity provider strategy, at least one of the capacity providers must have a weight value greater than zero and any capacity providers with a weight of `0` can't be used to place tasks. If you specify multiple capacity providers in a strategy that all have a weight of `0` , any `RunTask` or `CreateService` actions using the capacity provider strategy will fail.
	//
	// Weight value characteristics:
	//
	// - Weight is considered after the base value is satisfied
	// - The default value is `0` if not specified
	// - The valid range is 0 to 1,000
	// - At least one capacity provider must have a weight greater than zero
	// - Capacity providers with weight of `0` cannot place tasks
	//
	// Task distribution logic:
	//
	// - Base satisfaction: The minimum number of tasks specified by the base value are placed on that capacity provider
	// - Weight distribution: After base requirements are met, additional tasks are distributed according to weight ratios
	//
	// Examples:
	//
	// Equal Distribution: Two capacity providers both with weight `1` will split tasks evenly after base requirements are met.
	//
	// Weighted Distribution: If capacityProviderA has weight `1` and capacityProviderB has weight `4` , then for every 1 task on A, 4 tasks will run on B.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-clustercapacityproviderassociations-capacityproviderstrategy.html#cfn-ecs-clustercapacityproviderassociations-capacityproviderstrategy-weight
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

