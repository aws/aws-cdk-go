package awssynthetics


// A structure that contains input information for a canary run.
//
// This structure is required.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runConfigProperty := &runConfigProperty{
//   	activeTracing: jsii.Boolean(false),
//   	environmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	memoryInMb: jsii.Number(123),
//   	timeoutInSeconds: jsii.Number(123),
//   }
//
type CfnCanary_RunConfigProperty struct {
	// Specifies whether this canary is to use active AWS X-Ray tracing when it runs.
	//
	// Active tracing enables this canary run to be displayed in the ServiceLens and X-Ray service maps even if the canary does not hit an endpoint that has X-Ray tracing enabled. Using X-Ray tracing incurs charges. For more information, see [Canaries and X-Ray tracing](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_tracing.html) .
	//
	// You can enable active tracing only for canaries that use version `syn-nodejs-2.0` or later for their canary runtime.
	ActiveTracing interface{} `field:"optional" json:"activeTracing" yaml:"activeTracing"`
	// Specifies the keys and values to use for any environment variables used in the canary script.
	//
	// Use the following format:
	//
	// { "key1" : "value1", "key2" : "value2", ...}
	//
	// Keys must start with a letter and be at least two characters. The total size of your environment variables cannot exceed 4 KB. You can't specify any Lambda reserved environment variables as the keys for your environment variables. For more information about reserved keys, see [Runtime environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-runtime) .
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The maximum amount of memory that the canary can use while running.
	//
	// This value must be a multiple of 64. The range is 960 to 3008.
	MemoryInMb *float64 `field:"optional" json:"memoryInMb" yaml:"memoryInMb"`
	// How long the canary is allowed to run before it must stop.
	//
	// You can't set this time to be longer than the frequency of the runs of this canary.
	//
	// If you omit this field, the frequency of the canary is used as this value, up to a maximum of 900 seconds.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

