package previewawsquicksightmixins


// A semantic table that represents the final analytical structure of the data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tagRuleConfigurations interface{}
//
//   semanticTableProperty := &SemanticTableProperty{
//   	Alias: jsii.String("alias"),
//   	DestinationTableId: jsii.String("destinationTableId"),
//   	RowLevelPermissionConfiguration: &RowLevelPermissionConfigurationProperty{
//   		RowLevelPermissionDataSet: &RowLevelPermissionDataSetProperty{
//   			Arn: jsii.String("arn"),
//   			FormatVersion: jsii.String("formatVersion"),
//   			Namespace: jsii.String("namespace"),
//   			PermissionPolicy: jsii.String("permissionPolicy"),
//   			Status: jsii.String("status"),
//   		},
//   		TagConfiguration: &RowLevelPermissionTagConfigurationProperty{
//   			Status: jsii.String("status"),
//   			TagRuleConfigurations: tagRuleConfigurations,
//   			TagRules: []interface{}{
//   				&RowLevelPermissionTagRuleProperty{
//   					ColumnName: jsii.String("columnName"),
//   					MatchAllValue: jsii.String("matchAllValue"),
//   					TagKey: jsii.String("tagKey"),
//   					TagMultiValueDelimiter: jsii.String("tagMultiValueDelimiter"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semantictable.html
//
type CfnDataSetPropsMixin_SemanticTableProperty struct {
	// Alias for the semantic table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semantictable.html#cfn-quicksight-dataset-semantictable-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The identifier of the destination table from data preparation that provides data to this semantic table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semantictable.html#cfn-quicksight-dataset-semantictable-destinationtableid
	//
	DestinationTableId *string `field:"optional" json:"destinationTableId" yaml:"destinationTableId"`
	// Configuration for row level security that control data access for this semantic table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semantictable.html#cfn-quicksight-dataset-semantictable-rowlevelpermissionconfiguration
	//
	RowLevelPermissionConfiguration interface{} `field:"optional" json:"rowLevelPermissionConfiguration" yaml:"rowLevelPermissionConfiguration"`
}

