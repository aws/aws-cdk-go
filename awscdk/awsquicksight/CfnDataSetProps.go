package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tagRuleConfigurations interface{}
//
//   cfnDataSetProps := &CfnDataSetProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	ColumnGroups: []interface{}{
//   		&ColumnGroupProperty{
//   			GeoSpatialColumnGroup: &GeoSpatialColumnGroupProperty{
//   				Columns: []*string{
//   					jsii.String("columns"),
//   				},
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				CountryCode: jsii.String("countryCode"),
//   			},
//   		},
//   	},
//   	ColumnLevelPermissionRules: []interface{}{
//   		&ColumnLevelPermissionRuleProperty{
//   			ColumnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			Principals: []*string{
//   				jsii.String("principals"),
//   			},
//   		},
//   	},
//   	DataSetId: jsii.String("dataSetId"),
//   	DatasetParameters: []interface{}{
//   		&DatasetParameterProperty{
//   			DateTimeDatasetParameter: &DateTimeDatasetParameterProperty{
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   				ValueType: jsii.String("valueType"),
//
//   				// the properties below are optional
//   				DefaultValues: &DateTimeDatasetParameterDefaultValuesProperty{
//   					StaticValues: []*string{
//   						jsii.String("staticValues"),
//   					},
//   				},
//   				TimeGranularity: jsii.String("timeGranularity"),
//   			},
//   			DecimalDatasetParameter: &DecimalDatasetParameterProperty{
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   				ValueType: jsii.String("valueType"),
//
//   				// the properties below are optional
//   				DefaultValues: &DecimalDatasetParameterDefaultValuesProperty{
//   					StaticValues: []interface{}{
//   						jsii.Number(123),
//   					},
//   				},
//   			},
//   			IntegerDatasetParameter: &IntegerDatasetParameterProperty{
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   				ValueType: jsii.String("valueType"),
//
//   				// the properties below are optional
//   				DefaultValues: &IntegerDatasetParameterDefaultValuesProperty{
//   					StaticValues: []interface{}{
//   						jsii.Number(123),
//   					},
//   				},
//   			},
//   			StringDatasetParameter: &StringDatasetParameterProperty{
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   				ValueType: jsii.String("valueType"),
//
//   				// the properties below are optional
//   				DefaultValues: &StringDatasetParameterDefaultValuesProperty{
//   					StaticValues: []*string{
//   						jsii.String("staticValues"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	DataSetRefreshProperties: &DataSetRefreshPropertiesProperty{
//   		RefreshConfiguration: &RefreshConfigurationProperty{
//   			IncrementalRefresh: &IncrementalRefreshProperty{
//   				LookbackWindow: &LookbackWindowProperty{
//   					ColumnName: jsii.String("columnName"),
//   					Size: jsii.Number(123),
//   					SizeUnit: jsii.String("sizeUnit"),
//   				},
//   			},
//   		},
//   	},
//   	DataSetUsageConfiguration: &DataSetUsageConfigurationProperty{
//   		DisableUseAsDirectQuerySource: jsii.Boolean(false),
//   		DisableUseAsImportedSource: jsii.Boolean(false),
//   	},
//   	FieldFolders: map[string]interface{}{
//   		"fieldFoldersKey": &FieldFolderProperty{
//   			"columns": []*string{
//   				jsii.String("columns"),
//   			},
//   			"description": jsii.String("description"),
//   		},
//   	},
//   	ImportMode: jsii.String("importMode"),
//   	IngestionWaitPolicy: &IngestionWaitPolicyProperty{
//   		IngestionWaitTimeInHours: jsii.Number(123),
//   		WaitForSpiceIngestion: jsii.Boolean(false),
//   	},
//   	LogicalTableMap: map[string]interface{}{
//   		"logicalTableMapKey": &LogicalTableProperty{
//   			"alias": jsii.String("alias"),
//   			"source": &LogicalTableSourceProperty{
//   				"dataSetArn": jsii.String("dataSetArn"),
//   				"joinInstruction": &JoinInstructionProperty{
//   					"leftOperand": jsii.String("leftOperand"),
//   					"onClause": jsii.String("onClause"),
//   					"rightOperand": jsii.String("rightOperand"),
//   					"type": jsii.String("type"),
//
//   					// the properties below are optional
//   					"leftJoinKeyProperties": &JoinKeyPropertiesProperty{
//   						"uniqueKey": jsii.Boolean(false),
//   					},
//   					"rightJoinKeyProperties": &JoinKeyPropertiesProperty{
//   						"uniqueKey": jsii.Boolean(false),
//   					},
//   				},
//   				"physicalTableId": jsii.String("physicalTableId"),
//   			},
//
//   			// the properties below are optional
//   			"dataTransforms": []interface{}{
//   				&TransformOperationProperty{
//   					"castColumnTypeOperation": &CastColumnTypeOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"newColumnType": jsii.String("newColumnType"),
//
//   						// the properties below are optional
//   						"format": jsii.String("format"),
//   					},
//   					"createColumnsOperation": &CreateColumnsOperationProperty{
//   						"columns": []interface{}{
//   							&CalculatedColumnProperty{
//   								"columnId": jsii.String("columnId"),
//   								"columnName": jsii.String("columnName"),
//   								"expression": jsii.String("expression"),
//   							},
//   						},
//   					},
//   					"filterOperation": &FilterOperationProperty{
//   						"conditionExpression": jsii.String("conditionExpression"),
//   					},
//   					"overrideDatasetParameterOperation": &OverrideDatasetParameterOperationProperty{
//   						"parameterName": jsii.String("parameterName"),
//
//   						// the properties below are optional
//   						"newDefaultValues": &NewDefaultValuesProperty{
//   							"dateTimeStaticValues": []*string{
//   								jsii.String("dateTimeStaticValues"),
//   							},
//   							"decimalStaticValues": []interface{}{
//   								jsii.Number(123),
//   							},
//   							"integerStaticValues": []interface{}{
//   								jsii.Number(123),
//   							},
//   							"stringStaticValues": []*string{
//   								jsii.String("stringStaticValues"),
//   							},
//   						},
//   						"newParameterName": jsii.String("newParameterName"),
//   					},
//   					"projectOperation": &ProjectOperationProperty{
//   						"projectedColumns": []*string{
//   							jsii.String("projectedColumns"),
//   						},
//   					},
//   					"renameColumnOperation": &RenameColumnOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"newColumnName": jsii.String("newColumnName"),
//   					},
//   					"tagColumnOperation": &TagColumnOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"tags": []ColumnTagProperty{
//   							&ColumnTagProperty{
//   								"columnDescription": &ColumnDescriptionProperty{
//   									"text": jsii.String("text"),
//   								},
//   								"columnGeographicRole": jsii.String("columnGeographicRole"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   		},
//   	},
//   	PhysicalTableMap: map[string]interface{}{
//   		"physicalTableMapKey": &PhysicalTableProperty{
//   			"customSql": &CustomSqlProperty{
//   				"columns": []interface{}{
//   					&InputColumnProperty{
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//   					},
//   				},
//   				"dataSourceArn": jsii.String("dataSourceArn"),
//   				"name": jsii.String("name"),
//   				"sqlQuery": jsii.String("sqlQuery"),
//   			},
//   			"relationalTable": &RelationalTableProperty{
//   				"dataSourceArn": jsii.String("dataSourceArn"),
//   				"inputColumns": []interface{}{
//   					&InputColumnProperty{
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//   					},
//   				},
//   				"name": jsii.String("name"),
//
//   				// the properties below are optional
//   				"catalog": jsii.String("catalog"),
//   				"schema": jsii.String("schema"),
//   			},
//   			"s3Source": &S3SourceProperty{
//   				"dataSourceArn": jsii.String("dataSourceArn"),
//   				"inputColumns": []interface{}{
//   					&InputColumnProperty{
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				"uploadSettings": &UploadSettingsProperty{
//   					"containsHeader": jsii.Boolean(false),
//   					"delimiter": jsii.String("delimiter"),
//   					"format": jsii.String("format"),
//   					"startFromRow": jsii.Number(123),
//   					"textQualifier": jsii.String("textQualifier"),
//   				},
//   			},
//   		},
//   	},
//   	RowLevelPermissionDataSet: &RowLevelPermissionDataSetProperty{
//   		Arn: jsii.String("arn"),
//   		PermissionPolicy: jsii.String("permissionPolicy"),
//
//   		// the properties below are optional
//   		FormatVersion: jsii.String("formatVersion"),
//   		Namespace: jsii.String("namespace"),
//   		Status: jsii.String("status"),
//   	},
//   	RowLevelPermissionTagConfiguration: &RowLevelPermissionTagConfigurationProperty{
//   		TagRules: []interface{}{
//   			&RowLevelPermissionTagRuleProperty{
//   				ColumnName: jsii.String("columnName"),
//   				TagKey: jsii.String("tagKey"),
//
//   				// the properties below are optional
//   				MatchAllValue: jsii.String("matchAllValue"),
//   				TagMultiValueDelimiter: jsii.String("tagMultiValueDelimiter"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Status: jsii.String("status"),
//   		TagRuleConfigurations: tagRuleConfigurations,
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html
//
type CfnDataSetProps struct {
	// The AWS account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Groupings of columns that work together in certain Amazon QuickSight features.
	//
	// Currently, only geospatial hierarchy is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-columngroups
	//
	ColumnGroups interface{} `field:"optional" json:"columnGroups" yaml:"columnGroups"`
	// A set of one or more definitions of a `ColumnLevelPermissionRule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-columnlevelpermissionrules
	//
	ColumnLevelPermissionRules interface{} `field:"optional" json:"columnLevelPermissionRules" yaml:"columnLevelPermissionRules"`
	// An ID for the dataset that you want to create.
	//
	// This ID is unique per AWS Region for each AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-datasetid
	//
	DataSetId *string `field:"optional" json:"dataSetId" yaml:"dataSetId"`
	// The parameters that are declared in a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-datasetparameters
	//
	DatasetParameters interface{} `field:"optional" json:"datasetParameters" yaml:"datasetParameters"`
	// The refresh properties of a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-datasetrefreshproperties
	//
	DataSetRefreshProperties interface{} `field:"optional" json:"dataSetRefreshProperties" yaml:"dataSetRefreshProperties"`
	// The usage configuration to apply to child datasets that reference this dataset as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-datasetusageconfiguration
	//
	DataSetUsageConfiguration interface{} `field:"optional" json:"dataSetUsageConfiguration" yaml:"dataSetUsageConfiguration"`
	// The folder that contains fields and nested subfolders for your dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-fieldfolders
	//
	FieldFolders interface{} `field:"optional" json:"fieldFolders" yaml:"fieldFolders"`
	// Indicates whether you want to import the data into SPICE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-importmode
	//
	ImportMode *string `field:"optional" json:"importMode" yaml:"importMode"`
	// The wait policy to use when creating or updating a Dataset.
	//
	// The default is to wait for SPICE ingestion to finish with timeout of 36 hours.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-ingestionwaitpolicy
	//
	IngestionWaitPolicy interface{} `field:"optional" json:"ingestionWaitPolicy" yaml:"ingestionWaitPolicy"`
	// Configures the combination and transformation of the data from the physical tables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-logicaltablemap
	//
	LogicalTableMap interface{} `field:"optional" json:"logicalTableMap" yaml:"logicalTableMap"`
	// The display name for the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of resource permissions on the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Declares the physical tables that are available in the underlying data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-physicaltablemap
	//
	PhysicalTableMap interface{} `field:"optional" json:"physicalTableMap" yaml:"physicalTableMap"`
	// The row-level security configuration for the data that you want to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-rowlevelpermissiondataset
	//
	RowLevelPermissionDataSet interface{} `field:"optional" json:"rowLevelPermissionDataSet" yaml:"rowLevelPermissionDataSet"`
	// The element you can use to define tags for row-level security.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-rowlevelpermissiontagconfiguration
	//
	RowLevelPermissionTagConfiguration interface{} `field:"optional" json:"rowLevelPermissionTagConfiguration" yaml:"rowLevelPermissionTagConfiguration"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

