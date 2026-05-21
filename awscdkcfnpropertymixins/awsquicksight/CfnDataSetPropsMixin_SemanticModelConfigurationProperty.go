package awsquicksight


// Configuration for the semantic model that defines how prepared data is structured for analysis and reporting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var tagRuleConfigurations interface{}
//
//   semanticModelConfigurationProperty := &SemanticModelConfigurationProperty{
//   	SemanticMetadata: []interface{}{
//   		&DataSetSemanticMetadataProperty{
//   			CustomInstructions: []interface{}{
//   				&CustomInstructionProperty{
//   					InlineCustomInstruction: &InlineCustomInstructionProperty{
//   						InstructionText: jsii.String("instructionText"),
//   						UploadedDocumentMetadata: &UploadedDocumentMetadataProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   			Description: &DataSetSemanticDescriptionProperty{
//   				Text: jsii.String("text"),
//   			},
//   		},
//   	},
//   	TableMap: map[string]interface{}{
//   		"tableMapKey": &SemanticTableProperty{
//   			"alias": jsii.String("alias"),
//   			"destinationTableId": jsii.String("destinationTableId"),
//   			"rowLevelPermissionConfiguration": &RowLevelPermissionConfigurationProperty{
//   				"rowLevelPermissionDataSet": &RowLevelPermissionDataSetProperty{
//   					"arn": jsii.String("arn"),
//   					"formatVersion": jsii.String("formatVersion"),
//   					"namespace": jsii.String("namespace"),
//   					"permissionPolicy": jsii.String("permissionPolicy"),
//   					"status": jsii.String("status"),
//   				},
//   				"tagConfiguration": &RowLevelPermissionTagConfigurationProperty{
//   					"status": jsii.String("status"),
//   					"tagRuleConfigurations": tagRuleConfigurations,
//   					"tagRules": []interface{}{
//   						&RowLevelPermissionTagRuleProperty{
//   							"columnName": jsii.String("columnName"),
//   							"matchAllValue": jsii.String("matchAllValue"),
//   							"tagKey": jsii.String("tagKey"),
//   							"tagMultiValueDelimiter": jsii.String("tagMultiValueDelimiter"),
//   						},
//   					},
//   				},
//   			},
//   			"semanticMetadata": &TableSemanticMetadataProperty{
//   				"columnMetadata": []interface{}{
//   					&SharedColumnSemanticMetadataProperty{
//   						"columnNames": []*string{
//   							jsii.String("columnNames"),
//   						},
//   						"columnProperties": []interface{}{
//   							&ColumnSemanticPropertyProperty{
//   								"additionalNotes": &AdditionalNotesProperty{
//   									"text": jsii.String("text"),
//   								},
//   								"description": &ColumnDescriptionProperty{
//   									"text": jsii.String("text"),
//   								},
//   								"semanticType": &ColumnSemanticTypeProperty{
//   									"geographicalRole": jsii.String("geographicalRole"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semanticmodelconfiguration.html
//
type CfnDataSetPropsMixin_SemanticModelConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semanticmodelconfiguration.html#cfn-quicksight-dataset-semanticmodelconfiguration-semanticmetadata
	//
	SemanticMetadata interface{} `field:"optional" json:"semanticMetadata" yaml:"semanticMetadata"`
	// A map of semantic tables that define the analytical structure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semanticmodelconfiguration.html#cfn-quicksight-dataset-semanticmodelconfiguration-tablemap
	//
	TableMap interface{} `field:"optional" json:"tableMap" yaml:"tableMap"`
}

