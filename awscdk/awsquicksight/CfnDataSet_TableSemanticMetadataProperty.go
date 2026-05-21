package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableSemanticMetadataProperty := &TableSemanticMetadataProperty{
//   	ColumnMetadata: []interface{}{
//   		&SharedColumnSemanticMetadataProperty{
//   			ColumnProperties: []interface{}{
//   				&ColumnSemanticPropertyProperty{
//   					AdditionalNotes: &AdditionalNotesProperty{
//   						Text: jsii.String("text"),
//   					},
//   					Description: &ColumnDescriptionProperty{
//   						Text: jsii.String("text"),
//   					},
//   					SemanticType: &ColumnSemanticTypeProperty{
//   						GeographicalRole: jsii.String("geographicalRole"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			ColumnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-tablesemanticmetadata.html
//
type CfnDataSet_TableSemanticMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-tablesemanticmetadata.html#cfn-quicksight-dataset-tablesemanticmetadata-columnmetadata
	//
	ColumnMetadata interface{} `field:"optional" json:"columnMetadata" yaml:"columnMetadata"`
}

