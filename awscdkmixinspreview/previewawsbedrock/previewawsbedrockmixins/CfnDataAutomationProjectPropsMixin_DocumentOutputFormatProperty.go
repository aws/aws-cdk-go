package previewawsbedrockmixins


// A document output format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnDataAutomationProjectPropsMixin_DocumentOutputFormatProperty struct {
	// Output settings for additional file formats.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoutputformat.html#cfn-bedrock-dataautomationproject-documentoutputformat-additionalfileformat
	//
	AdditionalFileFormat interface{} `field:"optional" json:"additionalFileFormat" yaml:"additionalFileFormat"`
	// An output text format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoutputformat.html#cfn-bedrock-dataautomationproject-documentoutputformat-textformat
	//
	TextFormat interface{} `field:"optional" json:"textFormat" yaml:"textFormat"`
}

