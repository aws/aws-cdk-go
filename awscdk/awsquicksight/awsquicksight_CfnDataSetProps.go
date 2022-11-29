package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDataSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSetProps := &cfnDataSetProps{
//   	awsAccountId: jsii.String("awsAccountId"),
//   	columnGroups: []interface{}{
//   		&columnGroupProperty{
//   			geoSpatialColumnGroup: &geoSpatialColumnGroupProperty{
//   				columns: []*string{
//   					jsii.String("columns"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				countryCode: jsii.String("countryCode"),
//   			},
//   		},
//   	},
//   	columnLevelPermissionRules: []interface{}{
//   		&columnLevelPermissionRuleProperty{
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			principals: []*string{
//   				jsii.String("principals"),
//   			},
//   		},
//   	},
//   	dataSetId: jsii.String("dataSetId"),
//   	dataSetUsageConfiguration: &dataSetUsageConfigurationProperty{
//   		disableUseAsDirectQuerySource: jsii.Boolean(false),
//   		disableUseAsImportedSource: jsii.Boolean(false),
//   	},
//   	fieldFolders: map[string]interface{}{
//   		"fieldFoldersKey": &FieldFolderProperty{
//   			"columns": []*string{
//   				jsii.String("columns"),
//   			},
//   			"description": jsii.String("description"),
//   		},
//   	},
//   	importMode: jsii.String("importMode"),
//   	ingestionWaitPolicy: &ingestionWaitPolicyProperty{
//   		ingestionWaitTimeInHours: jsii.Number(123),
//   		waitForSpiceIngestion: jsii.Boolean(false),
//   	},
//   	logicalTableMap: map[string]interface{}{
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
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	physicalTableMap: map[string]interface{}{
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
//   	rowLevelPermissionDataSet: &rowLevelPermissionDataSetProperty{
//   		arn: jsii.String("arn"),
//   		permissionPolicy: jsii.String("permissionPolicy"),
//
//   		// the properties below are optional
//   		formatVersion: jsii.String("formatVersion"),
//   		namespace: jsii.String("namespace"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDataSetProps struct {
	// The AWS account ID.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Groupings of columns that work together in certain Amazon QuickSight features.
	//
	// Currently, only geospatial hierarchy is supported.
	ColumnGroups interface{} `field:"optional" json:"columnGroups" yaml:"columnGroups"`
	// A set of one or more definitions of a `ColumnLevelPermissionRule` .
	ColumnLevelPermissionRules interface{} `field:"optional" json:"columnLevelPermissionRules" yaml:"columnLevelPermissionRules"`
	// An ID for the dataset that you want to create.
	//
	// This ID is unique per AWS Region for each AWS account.
	DataSetId *string `field:"optional" json:"dataSetId" yaml:"dataSetId"`
	// `AWS::QuickSight::DataSet.DataSetUsageConfiguration`.
	DataSetUsageConfiguration interface{} `field:"optional" json:"dataSetUsageConfiguration" yaml:"dataSetUsageConfiguration"`
	// The folder that contains fields and nested subfolders for your dataset.
	FieldFolders interface{} `field:"optional" json:"fieldFolders" yaml:"fieldFolders"`
	// Indicates whether you want to import the data into SPICE.
	ImportMode *string `field:"optional" json:"importMode" yaml:"importMode"`
	// The wait policy to use when creating or updating a Dataset.
	//
	// The default is to wait for SPICE ingestion to finish with timeout of 36 hours.
	IngestionWaitPolicy interface{} `field:"optional" json:"ingestionWaitPolicy" yaml:"ingestionWaitPolicy"`
	// Configures the combination and transformation of the data from the physical tables.
	LogicalTableMap interface{} `field:"optional" json:"logicalTableMap" yaml:"logicalTableMap"`
	// The display name for the dataset.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of resource permissions on the dataset.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Declares the physical tables that are available in the underlying data sources.
	PhysicalTableMap interface{} `field:"optional" json:"physicalTableMap" yaml:"physicalTableMap"`
	// The row-level security configuration for the data that you want to create.
	RowLevelPermissionDataSet interface{} `field:"optional" json:"rowLevelPermissionDataSet" yaml:"rowLevelPermissionDataSet"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

