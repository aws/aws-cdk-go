package awsbedrock


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-splitterconfiguration.html#cfn-bedrock-dataautomationproject-splitterconfiguration-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

