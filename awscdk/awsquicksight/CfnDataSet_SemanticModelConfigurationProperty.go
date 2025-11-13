package awsquicksight


// Configuration for the semantic model that defines how prepared data is structured for analysis and reporting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tagRuleConfigurations interface{}
//
//   semanticModelConfigurationProperty := &SemanticModelConfigurationProperty{
//   	TableMap: map[string]interface{}{
//   		"tableMapKey": &SemanticTableProperty{
//   			"alias": jsii.String("alias"),
//   			"destinationTableId": jsii.String("destinationTableId"),
//
//   			// the properties below are optional
//   			"rowLevelPermissionConfiguration": &RowLevelPermissionConfigurationProperty{
//   				"rowLevelPermissionDataSet": &RowLevelPermissionDataSetProperty{
//   					"arn": jsii.String("arn"),
//   					"permissionPolicy": jsii.String("permissionPolicy"),
//
//   					// the properties below are optional
//   					"formatVersion": jsii.String("formatVersion"),
//   					"namespace": jsii.String("namespace"),
//   					"status": jsii.String("status"),
//   				},
//   				"tagConfiguration": &RowLevelPermissionTagConfigurationProperty{
//   					"tagRules": []interface{}{
//   						&RowLevelPermissionTagRuleProperty{
//   							"columnName": jsii.String("columnName"),
//   							"tagKey": jsii.String("tagKey"),
//
//   							// the properties below are optional
//   							"matchAllValue": jsii.String("matchAllValue"),
//   							"tagMultiValueDelimiter": jsii.String("tagMultiValueDelimiter"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					"status": jsii.String("status"),
//   					"tagRuleConfigurations": tagRuleConfigurations,
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semanticmodelconfiguration.html
//
type CfnDataSet_SemanticModelConfigurationProperty struct {
	// A map of semantic tables that define the analytical structure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-semanticmodelconfiguration.html#cfn-quicksight-dataset-semanticmodelconfiguration-tablemap
	//
	TableMap interface{} `field:"optional" json:"tableMap" yaml:"tableMap"`
}

