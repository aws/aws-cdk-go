package awsbedrockagentcore


// Resource configuration for an agent runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   resourceConfigurationProperty := &ResourceConfigurationProperty{
//   	MinimumMemoryInMiB: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-resourceconfiguration.html
//
type CfnRuntimePropsMixin_ResourceConfigurationProperty struct {
	// The minimum amount of memory in MiB to allocate to the agent runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-resourceconfiguration.html#cfn-bedrockagentcore-runtime-resourceconfiguration-minimummemoryinmib
	//
	MinimumMemoryInMiB *float64 `field:"optional" json:"minimumMemoryInMiB" yaml:"minimumMemoryInMiB"`
}

