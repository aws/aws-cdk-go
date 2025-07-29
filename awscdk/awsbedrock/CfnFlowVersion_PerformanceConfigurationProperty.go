package awsbedrock


// Performance settings for a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   performanceConfigurationProperty := &PerformanceConfigurationProperty{
//   	Latency: jsii.String("latency"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-performanceconfiguration.html
//
type CfnFlowVersion_PerformanceConfigurationProperty struct {
	// To use a latency-optimized version of the model, set to `optimized` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-performanceconfiguration.html#cfn-bedrock-flowversion-performanceconfiguration-latency
	//
	Latency *string `field:"optional" json:"latency" yaml:"latency"`
}

