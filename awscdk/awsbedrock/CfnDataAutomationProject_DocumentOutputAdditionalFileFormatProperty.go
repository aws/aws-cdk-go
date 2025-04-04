package awsbedrock


// Output settings for additional file formats.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentOutputAdditionalFileFormatProperty := &DocumentOutputAdditionalFileFormatProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoutputadditionalfileformat.html
//
type CfnDataAutomationProject_DocumentOutputAdditionalFileFormatProperty struct {
	// Whether additional file formats are enabled for a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoutputadditionalfileformat.html#cfn-bedrock-dataautomationproject-documentoutputadditionalfileformat-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
}

