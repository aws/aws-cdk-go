package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoStandardGenerativeFieldProperty := &VideoStandardGenerativeFieldProperty{
//   	State: jsii.String("state"),
//
//   	// the properties below are optional
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videostandardgenerativefield.html
//
type CfnDataAutomationProject_VideoStandardGenerativeFieldProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videostandardgenerativefield.html#cfn-bedrock-dataautomationproject-videostandardgenerativefield-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videostandardgenerativefield.html#cfn-bedrock-dataautomationproject-videostandardgenerativefield-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

