package awsbedrock


// Document splitter settings.
//
// If a document is too large to be processed in one pass, the document splitter splits it into smaller documents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   splitterConfigurationProperty := &SplitterConfigurationProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-splitterconfiguration.html
//
type CfnDataAutomationProject_SplitterConfigurationProperty struct {
	// Whether document splitter is enabled for a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-splitterconfiguration.html#cfn-bedrock-dataautomationproject-splitterconfiguration-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

