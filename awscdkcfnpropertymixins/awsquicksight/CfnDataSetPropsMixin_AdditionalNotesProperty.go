package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   additionalNotesProperty := &AdditionalNotesProperty{
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-additionalnotes.html
//
type CfnDataSetPropsMixin_AdditionalNotesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-additionalnotes.html#cfn-quicksight-dataset-additionalnotes-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
}

