package previewawsecsmixins


// The details of a capacity provider strategy.
//
// A capacity provider strategy can be set when using the `RunTask` or `CreateService` APIs or as the default capacity provider strategy for a cluster with the `CreateCluster` API.
//
// Only capacity providers that are already associated with a cluster and have an `ACTIVE` or `UPDATING` status can be used in a capacity provider strategy. The `PutClusterCapacityProviders` API is used to associate a capacity provider with a cluster.
//
// If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must already be created. New Auto Scaling group capacity providers can be created with the `CreateCapacityProvider` API operation.
//
// To use an AWS Fargate capacity provider, specify either the `FARGATE` or `FARGATE_SPOT` capacity providers. The AWS Fargate capacity providers are available to all accounts and only need to be associated with a cluster to be used in a capacity provider strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   capacityProviderStrategyItemProperty := &CapacityProviderStrategyItemProperty{
//   	Base: jsii.Number(123),
//   	CapacityProvider: jsii.String("capacityProvider"),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-capacityproviderstrategyitem.html
//
type CfnServicePropsMixin_CapacityProviderStrategyItemProperty struct {
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-capacityproviderstrategyitem.html#cfn-ecs-service-capacityproviderstrategyitem-base
	//
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// The short name of the capacity provider.
	//
	// This can be either an AWS managed capacity provider ( `FARGATE` or `FARGATE_SPOT` ) or the name of a custom capacity provider that you created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-capacityproviderstrategyitem.html#cfn-ecs-service-capacityproviderstrategyitem-capacityprovider
	//
	CapacityProvider *string `field:"optional" json:"capacityProvider" yaml:"capacityProvider"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-capacityproviderstrategyitem.html#cfn-ecs-service-capacityproviderstrategyitem-weight
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

