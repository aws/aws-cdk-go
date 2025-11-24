package mixinsawsbedrock


// Settings for generating categorical data from images.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageExtractionCategoryProperty := &ImageExtractionCategoryProperty{
//   	State: jsii.String("state"),
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imageextractioncategory.html
//
type CfnDataAutomationProjectPropsMixin_ImageExtractionCategoryProperty struct {
	// Whether generating categorical data from images is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imageextractioncategory.html#cfn-bedrock-dataautomationproject-imageextractioncategory-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// The types of data to generate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imageextractioncategory.html#cfn-bedrock-dataautomationproject-imageextractioncategory-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

