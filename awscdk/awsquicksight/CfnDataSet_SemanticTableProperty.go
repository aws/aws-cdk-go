package awsquicksight


// A semantic table that represents the final analytical structure of the data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tagRuleConfigurations interface{}
//
//   semanticTableProperty := &SemanticTableProperty{
//   	Alias: jsii.String("alias"),
//   	DestinationTableId: jsii.String("destinationTableId"),
//
//   	// the properties below are optional
//   	RowLevelPermissionConfiguration: &RowLevelPermissionConfigurationProperty{
//   		RowLevelPermissionDataSet: &RowLevelPermissionDataSetProperty{
//   			Arn: jsii.String("arn"),
//   			PermissionPolicy: jsii.String("permissionPolicy"),
//
//   			// the properties below are optional
//   			FormatVersion: jsii.String("formatVersion"),
//   			Namespace: jsii.String("namespace"),
//   			Status: jsii.String("status"),
//   		},
//   		TagConfiguration: &RowLevelPermissionTagConfigurationProperty{
//   			TagRules: []interface{}{
//   				&RowLevelPermissionTagRuleProperty{
//   					ColumnName: jsii.String("columnName"),
//   					TagKey: jsii.String("tagKey"),
//
//   					// the properties below are optional
//   					MatchAllValue: jsii.String("matchAllValue"),
//   					TagMultiValueDelimiter: jsii.String("tagMultiValueDelimiter"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Status: jsii.String("status"),
//   			TagRuleConfigurations: tagRuleConfigurations,
//   		},
//   	},
//   	SemanticMetadata: &TableSemanticMetadataProperty{
//   		ColumnMetadata: []interface{}{
//   			&SharedColumnSemanticMetadataProperty{
//   				ColumnProperties: []interface{}{
//   					&ColumnSemanticPropertyProperty{
//   						AdditionalNotes: &AdditionalNotesProperty{
//   							Text: jsii.String("text"),
//   						},
//   						Description: &ColumnDescriptionProperty{
//   							Text: jsii.String("text"),
//   						},
//   						SemanticType: &ColumnSemanticTypeProperty{
//   							GeographicalRole: jsii.String("geographicalRole"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				ColumnNames: []*string{
//   					jsii.String("columnNames"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semantictable.html
//
type CfnDataSet_SemanticTableProperty struct {
	// Alias for the semantic table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semantictable.html#cfn-quicksight-dataset-semantictable-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The identifier of the destination table from data preparation that provides data to this semantic table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semantictable.html#cfn-quicksight-dataset-semantictable-destinationtableid
	//
	DestinationTableId *string `field:"required" json:"destinationTableId" yaml:"destinationTableId"`
	// Configuration for row level security that control data access for this semantic table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semantictable.html#cfn-quicksight-dataset-semantictable-rowlevelpermissionconfiguration
	//
	RowLevelPermissionConfiguration interface{} `field:"optional" json:"rowLevelPermissionConfiguration" yaml:"rowLevelPermissionConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semantictable.html#cfn-quicksight-dataset-semantictable-semanticmetadata
	//
	SemanticMetadata interface{} `field:"optional" json:"semanticMetadata" yaml:"semanticMetadata"`
}

