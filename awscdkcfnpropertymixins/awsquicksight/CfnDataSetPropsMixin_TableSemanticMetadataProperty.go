package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   tableSemanticMetadataProperty := &TableSemanticMetadataProperty{
//   	ColumnMetadata: []interface{}{
//   		&SharedColumnSemanticMetadataProperty{
//   			ColumnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
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
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-tablesemanticmetadata.html
//
type CfnDataSetPropsMixin_TableSemanticMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-tablesemanticmetadata.html#cfn-quicksight-dataset-tablesemanticmetadata-columnmetadata
	//
	ColumnMetadata interface{} `field:"optional" json:"columnMetadata" yaml:"columnMetadata"`
}

