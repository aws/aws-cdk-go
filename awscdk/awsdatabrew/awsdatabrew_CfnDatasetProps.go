package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDataset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatasetProps := &cfnDatasetProps{
//   	input: &inputProperty{
//   		databaseInputDefinition: &databaseInputDefinitionProperty{
//   			glueConnectionName: jsii.String("glueConnectionName"),
//
//   			// the properties below are optional
//   			databaseTableName: jsii.String("databaseTableName"),
//   			queryString: jsii.String("queryString"),
//   			tempDirectory: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   			},
//   		},
//   		dataCatalogInputDefinition: &dataCatalogInputDefinitionProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			tableName: jsii.String("tableName"),
//   			tempDirectory: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   			},
//   		},
//   		metadata: &metadataProperty{
//   			sourceArn: jsii.String("sourceArn"),
//   		},
//   		s3InputDefinition: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	format: jsii.String("format"),
//   	formatOptions: &formatOptionsProperty{
//   		csv: &csvOptionsProperty{
//   			delimiter: jsii.String("delimiter"),
//   			headerRow: jsii.Boolean(false),
//   		},
//   		excel: &excelOptionsProperty{
//   			headerRow: jsii.Boolean(false),
//   			sheetIndexes: []interface{}{
//   				jsii.Number(123),
//   			},
//   			sheetNames: []*string{
//   				jsii.String("sheetNames"),
//   			},
//   		},
//   		json: &jsonOptionsProperty{
//   			multiLine: jsii.Boolean(false),
//   		},
//   	},
//   	pathOptions: &pathOptionsProperty{
//   		filesLimit: &filesLimitProperty{
//   			maxFiles: jsii.Number(123),
//
//   			// the properties below are optional
//   			order: jsii.String("order"),
//   			orderedBy: jsii.String("orderedBy"),
//   		},
//   		lastModifiedDateCondition: &filterExpressionProperty{
//   			expression: jsii.String("expression"),
//   			valuesMap: []interface{}{
//   				&filterValueProperty{
//   					value: jsii.String("value"),
//   					valueReference: jsii.String("valueReference"),
//   				},
//   			},
//   		},
//   		parameters: []interface{}{
//   			&pathParameterProperty{
//   				datasetParameter: &datasetParameterProperty{
//   					name: jsii.String("name"),
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					createColumn: jsii.Boolean(false),
//   					datetimeOptions: &datetimeOptionsProperty{
//   						format: jsii.String("format"),
//
//   						// the properties below are optional
//   						localeCode: jsii.String("localeCode"),
//   						timezoneOffset: jsii.String("timezoneOffset"),
//   					},
//   					filter: &filterExpressionProperty{
//   						expression: jsii.String("expression"),
//   						valuesMap: []interface{}{
//   							&filterValueProperty{
//   								value: jsii.String("value"),
//   								valueReference: jsii.String("valueReference"),
//   							},
//   						},
//   					},
//   				},
//   				pathParameterName: jsii.String("pathParameterName"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDatasetProps struct {
	// Information on how DataBrew can find the dataset, in either the AWS Glue Data Catalog or Amazon S3 .
	Input interface{} `field:"required" json:"input" yaml:"input"`
	// The unique name of the dataset.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The file format of a dataset that is created from an Amazon S3 file or folder.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// A set of options that define how DataBrew interprets the data in the dataset.
	FormatOptions interface{} `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// A set of options that defines how DataBrew interprets an Amazon S3 path of the dataset.
	PathOptions interface{} `field:"optional" json:"pathOptions" yaml:"pathOptions"`
	// Metadata tags that have been applied to the dataset.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

