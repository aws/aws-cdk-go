package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharedColumnSemanticMetadataProperty := &SharedColumnSemanticMetadataProperty{
//   	ColumnProperties: []interface{}{
//   		&ColumnSemanticPropertyProperty{
//   			AdditionalNotes: &AdditionalNotesProperty{
//   				Text: jsii.String("text"),
//   			},
//   			Description: &ColumnDescriptionProperty{
//   				Text: jsii.String("text"),
//   			},
//   			SemanticType: &ColumnSemanticTypeProperty{
//   				GeographicalRole: jsii.String("geographicalRole"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ColumnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-sharedcolumnsemanticmetadata.html
//
type CfnDataSet_SharedColumnSemanticMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-sharedcolumnsemanticmetadata.html#cfn-quicksight-dataset-sharedcolumnsemanticmetadata-columnproperties
	//
	ColumnProperties interface{} `field:"required" json:"columnProperties" yaml:"columnProperties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-sharedcolumnsemanticmetadata.html#cfn-quicksight-dataset-sharedcolumnsemanticmetadata-columnnames
	//
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
}

