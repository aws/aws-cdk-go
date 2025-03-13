package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageStandardGenerativeFieldProperty := &ImageStandardGenerativeFieldProperty{
//   	State: jsii.String("state"),
//
//   	// the properties below are optional
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imagestandardgenerativefield.html
//
type CfnDataAutomationProject_ImageStandardGenerativeFieldProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imagestandardgenerativefield.html#cfn-bedrock-dataautomationproject-imagestandardgenerativefield-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imagestandardgenerativefield.html#cfn-bedrock-dataautomationproject-imagestandardgenerativefield-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

