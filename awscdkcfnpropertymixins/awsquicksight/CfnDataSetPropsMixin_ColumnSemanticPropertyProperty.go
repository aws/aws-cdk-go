package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   columnSemanticPropertyProperty := &ColumnSemanticPropertyProperty{
//   	AdditionalNotes: &AdditionalNotesProperty{
//   		Text: jsii.String("text"),
//   	},
//   	Description: &ColumnDescriptionProperty{
//   		Text: jsii.String("text"),
//   	},
//   	SemanticType: &ColumnSemanticTypeProperty{
//   		GeographicalRole: jsii.String("geographicalRole"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columnsemanticproperty.html
//
type CfnDataSetPropsMixin_ColumnSemanticPropertyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columnsemanticproperty.html#cfn-quicksight-dataset-columnsemanticproperty-additionalnotes
	//
	AdditionalNotes interface{} `field:"optional" json:"additionalNotes" yaml:"additionalNotes"`
	// <p>Metadata that contains a description for a column.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columnsemanticproperty.html#cfn-quicksight-dataset-columnsemanticproperty-description
	//
	Description interface{} `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columnsemanticproperty.html#cfn-quicksight-dataset-columnsemanticproperty-semantictype
	//
	SemanticType interface{} `field:"optional" json:"semanticType" yaml:"semanticType"`
}

