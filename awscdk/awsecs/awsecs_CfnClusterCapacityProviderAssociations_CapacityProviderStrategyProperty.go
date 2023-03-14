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
//   capacityProviderStrategyProperty := &capacityProviderStrategyProperty{
//   	capacityProvider: jsii.String("capacityProvider"),
//
//   	// the properties below are optional
//   	base: jsii.Number(123),
//   	weight: jsii.Number(123),
//   }
//
type CfnClusterCapacityProviderAssociations_CapacityProviderStrategyProperty struct {
	// The short name of the capacity provider.
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// The *base* value designates how many tasks, at a minimum, to run on the specified capacity provider.
	//
	// Only one capacity provider in a capacity provider strategy can have a *base* defined. If no value is specified, the default value of `0` is used.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// The *weight* value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The `weight` value is taken into consideration after the `base` value, if defined, is satisfied.
	//
	// If no `weight` value is specified, the default value of `0` is used. When multiple capacity providers are specified within a capacity provider strategy, at least one of the capacity providers must have a weight value greater than zero and any capacity providers with a weight of `0` will not be used to place tasks. If you specify multiple capacity providers in a strategy that all have a weight of `0` , any `RunTask` or `CreateService` actions using the capacity provider strategy will fail.
	//
	// An example scenario for using weights is defining a strategy that contains two capacity providers and both have a weight of `1` , then when the `base` is satisfied, the tasks will be split evenly across the two capacity providers. Using that same logic, if you specify a weight of `1` for *capacityProviderA* and a weight of `4` for *capacityProviderB* , then for every one task that is run using *capacityProviderA* , four tasks would use *capacityProviderB* .
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

