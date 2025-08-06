package awsbedrock


// Determines how multiple nodes in a flow can run in parallel.
//
// Running nodes concurrently can improve your flow's performance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowAliasConcurrencyConfigurationProperty := &FlowAliasConcurrencyConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	MaxConcurrency: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowalias-flowaliasconcurrencyconfiguration.html
//
type CfnFlowAlias_FlowAliasConcurrencyConfigurationProperty struct {
	// The type of concurrency to use for parallel node execution. Specify one of the following options:.
	//
	// - `Automatic` - Amazon Bedrock determines which nodes can be executed in parallel based on the flow definition and its dependencies.
	// - `Manual` - You specify which nodes can be executed in parallel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowalias-flowaliasconcurrencyconfiguration.html#cfn-bedrock-flowalias-flowaliasconcurrencyconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The maximum number of nodes that can be executed concurrently in the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowalias-flowaliasconcurrencyconfiguration.html#cfn-bedrock-flowalias-flowaliasconcurrencyconfiguration-maxconcurrency
	//
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
}

