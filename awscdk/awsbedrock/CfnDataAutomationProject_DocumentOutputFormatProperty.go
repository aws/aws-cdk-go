package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentOutputFormatProperty := &DocumentOutputFormatProperty{
//   	AdditionalFileFormat: &DocumentOutputAdditionalFileFormatProperty{
//   		State: jsii.String("state"),
//   	},
//   	TextFormat: &DocumentOutputTextFormatProperty{
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoutputformat.html
//
type CfnDataAutomationProject_DocumentOutputFormatProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoutputformat.html#cfn-bedrock-dataautomationproject-documentoutputformat-additionalfileformat
	//
	AdditionalFileFormat interface{} `field:"required" json:"additionalFileFormat" yaml:"additionalFileFormat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoutputformat.html#cfn-bedrock-dataautomationproject-documentoutputformat-textformat
	//
	TextFormat interface{} `field:"required" json:"textFormat" yaml:"textFormat"`
}

