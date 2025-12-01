package awslambda


// The scaling configuration that will be applied to the $LATEST.PUBLISHED version.
//
// Example:
//   var capacityProvider CapacityProvider
//   var fn Function
//
//
//   capacityProvider.AddFunction(fn, &CapacityProviderFunctionOptions{
//   	LatestPublishedScalingConfig: &LatestPublishedScalingConfig{
//   		MinExecutionEnvironments: jsii.Number(5),
//   		MaxExecutionEnvironments: jsii.Number(25),
//   	},
//   })
//
type LatestPublishedScalingConfig struct {
	// The maximum number of execution environments allowed for the $LATEST.PUBLISHED version when published into a capacity provider.
	//
	// This setting limits the total number of execution environments that can be created
	// to handle concurrent invocations of this specific version.
	// Default: - No maximum specified.
	//
	MaxExecutionEnvironments *float64 `field:"optional" json:"maxExecutionEnvironments" yaml:"maxExecutionEnvironments"`
	// The minimum number of execution environments to maintain for the $LATEST.PUBLISHED version when published into a capacity provider.
	//
	// This setting ensures that at least this many execution environments are always
	// available to handle function invocations for this specific version, reducing cold start latency.
	// Default: - 3 execution environments are set to be the minimum.
	//
	MinExecutionEnvironments *float64 `field:"optional" json:"minExecutionEnvironments" yaml:"minExecutionEnvironments"`
}

