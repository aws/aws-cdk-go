package awsbedrock


// An output text format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentOutputTextFormatProperty := &DocumentOutputTextFormatProperty{
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoutputtextformat.html
//
type CfnDataAutomationProject_DocumentOutputTextFormatProperty struct {
	// The types of output text to generate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoutputtextformat.html#cfn-bedrock-dataautomationproject-documentoutputtextformat-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

