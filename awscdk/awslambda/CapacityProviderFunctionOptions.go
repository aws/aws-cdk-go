package awslambda


// Options for creating a function associated with a capacity provider.
//
// Example:
//   var capacityProvider CapacityProvider
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	// Runtime must be equal to or newer than NODEJS_22_X
//   	Runtime: lambda.Runtime_NODEJS_22_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   })
//
//   // Associate the function with the capacity provider
//   capacityProvider.AddFunction(fn, &CapacityProviderFunctionOptions{
//   	PerExecutionEnvironmentMaxConcurrency: jsii.Number(10),
//   	ExecutionEnvironmentMemoryGiBPerVCpu: jsii.Number(4),
//   })
//
type CapacityProviderFunctionOptions struct {
	// Specifies the execution environment memory per VCPU, in GiB.
	// Default: 2.0
	//
	ExecutionEnvironmentMemoryGiBPerVCpu *float64 `field:"optional" json:"executionEnvironmentMemoryGiBPerVCpu" yaml:"executionEnvironmentMemoryGiBPerVCpu"`
	// The scaling options that are applied to the $LATEST.PUBLISHED version.
	// Default: - No scaling limitations are applied to the $LATEST.PUBLISHED version.
	//
	LatestPublishedScalingConfig *LatestPublishedScalingConfig `field:"optional" json:"latestPublishedScalingConfig" yaml:"latestPublishedScalingConfig"`
	// Specifies the maximum number of concurrent invokes a single execution environment can handle.
	// Default: Maximum is set to 10.
	//
	PerExecutionEnvironmentMaxConcurrency *float64 `field:"optional" json:"perExecutionEnvironmentMaxConcurrency" yaml:"perExecutionEnvironmentMaxConcurrency"`
	// A boolean determining whether or not to automatically publish to the $LATEST.PUBLISHED version.
	// Default: - True.
	//
	PublishToLatestPublished *bool `field:"optional" json:"publishToLatestPublished" yaml:"publishToLatestPublished"`
}

