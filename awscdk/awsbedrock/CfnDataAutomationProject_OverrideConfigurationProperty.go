package awsbedrock


// Override configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   overrideConfigurationProperty := &OverrideConfigurationProperty{
//   	Document: &DocumentOverrideConfigurationProperty{
//   		Splitter: &SplitterConfigurationProperty{
//   			State: jsii.String("state"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-overrideconfiguration.html
//
type CfnDataAutomationProject_OverrideConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-overrideconfiguration.html#cfn-bedrock-dataautomationproject-overrideconfiguration-document
	//
	Document interface{} `field:"optional" json:"document" yaml:"document"`
}

