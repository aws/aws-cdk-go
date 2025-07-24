package awsbedrock


// Settings for generating descriptions of audio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioStandardGenerativeFieldProperty := &AudioStandardGenerativeFieldProperty{
//   	State: jsii.String("state"),
//
//   	// the properties below are optional
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardgenerativefield.html
//
type CfnDataAutomationProject_AudioStandardGenerativeFieldProperty struct {
	// Whether generating descriptions is enabled for audio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardgenerativefield.html#cfn-bedrock-dataautomationproject-audiostandardgenerativefield-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
	// The types of description to generate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardgenerativefield.html#cfn-bedrock-dataautomationproject-audiostandardgenerativefield-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

