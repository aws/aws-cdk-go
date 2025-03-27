package awsbedrock


// Additional settings for a project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentOverrideConfigurationProperty := &DocumentOverrideConfigurationProperty{
//   	Splitter: &SplitterConfigurationProperty{
//   		State: jsii.String("state"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoverrideconfiguration.html
//
type CfnDataAutomationProject_DocumentOverrideConfigurationProperty struct {
	// Whether document splitter is enabled for a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoverrideconfiguration.html#cfn-bedrock-dataautomationproject-documentoverrideconfiguration-splitter
	//
	Splitter interface{} `field:"optional" json:"splitter" yaml:"splitter"`
}

