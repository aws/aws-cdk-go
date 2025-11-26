package previewawsbedrockmixins


// Settings for generating descriptions of video.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   videoStandardGenerativeFieldProperty := &VideoStandardGenerativeFieldProperty{
//   	State: jsii.String("state"),
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videostandardgenerativefield.html
//
type CfnDataAutomationProjectPropsMixin_VideoStandardGenerativeFieldProperty struct {
	// Whether generating descriptions is enabled for video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videostandardgenerativefield.html#cfn-bedrock-dataautomationproject-videostandardgenerativefield-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// The types of description to generate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videostandardgenerativefield.html#cfn-bedrock-dataautomationproject-videostandardgenerativefield-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

