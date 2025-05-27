package awsbedrock


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowalias-flowaliasconcurrencyconfiguration.html#cfn-bedrock-flowalias-flowaliasconcurrencyconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Number of nodes executed concurrently at a time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowalias-flowaliasconcurrencyconfiguration.html#cfn-bedrock-flowalias-flowaliasconcurrencyconfiguration-maxconcurrency
	//
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
}

