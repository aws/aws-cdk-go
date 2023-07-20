package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatasetProps := &CfnDatasetProps{
//   	Input: &InputProperty{
//   		DatabaseInputDefinition: &DatabaseInputDefinitionProperty{
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//
//   			// the properties below are optional
//   			DatabaseTableName: jsii.String("databaseTableName"),
//   			QueryString: jsii.String("queryString"),
//   			TempDirectory: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				Key: jsii.String("key"),
//   			},
//   		},
//   		DataCatalogInputDefinition: &DataCatalogInputDefinitionProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
//   			TempDirectory: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				Key: jsii.String("key"),
//   			},
//   		},
//   		Metadata: &MetadataProperty{
//   			SourceArn: jsii.String("sourceArn"),
//   		},
//   		S3InputDefinition: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Format: jsii.String("format"),
//   	FormatOptions: &FormatOptionsProperty{
//   		Csv: &CsvOptionsProperty{
//   			Delimiter: jsii.String("delimiter"),
//   			HeaderRow: jsii.Boolean(false),
//   		},
//   		Excel: &ExcelOptionsProperty{
//   			HeaderRow: jsii.Boolean(false),
//   			SheetIndexes: []interface{}{
//   				jsii.Number(123),
//   			},
//   			SheetNames: []*string{
//   				jsii.String("sheetNames"),
//   			},
//   		},
//   		Json: &JsonOptionsProperty{
//   			MultiLine: jsii.Boolean(false),
//   		},
//   	},
//   	PathOptions: &PathOptionsProperty{
//   		FilesLimit: &FilesLimitProperty{
//   			MaxFiles: jsii.Number(123),
//
//   			// the properties below are optional
//   			Order: jsii.String("order"),
//   			OrderedBy: jsii.String("orderedBy"),
//   		},
//   		LastModifiedDateCondition: &FilterExpressionProperty{
//   			Expression: jsii.String("expression"),
//   			ValuesMap: []interface{}{
//   				&FilterValueProperty{
//   					Value: jsii.String("value"),
//   					ValueReference: jsii.String("valueReference"),
//   				},
//   			},
//   		},
//   		Parameters: []interface{}{
//   			&PathParameterProperty{
//   				DatasetParameter: &DatasetParameterProperty{
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					CreateColumn: jsii.Boolean(false),
//   					DatetimeOptions: &DatetimeOptionsProperty{
//   						Format: jsii.String("format"),
//
//   						// the properties below are optional
//   						LocaleCode: jsii.String("localeCode"),
//   						TimezoneOffset: jsii.String("timezoneOffset"),
//   					},
//   					Filter: &FilterExpressionProperty{
//   						Expression: jsii.String("expression"),
//   						ValuesMap: []interface{}{
//   							&FilterValueProperty{
//   								Value: jsii.String("value"),
//   								ValueReference: jsii.String("valueReference"),
//   							},
//   						},
//   					},
//   				},
//   				PathParameterName: jsii.String("pathParameterName"),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html
//
type CfnDatasetProps struct {
	// Information on how DataBrew can find the dataset, in either the AWS Glue Data Catalog or Amazon S3 .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-input
	//
	Input interface{} `field:"required" json:"input" yaml:"input"`
	// The unique name of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The file format of a dataset that is created from an Amazon S3 file or folder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// A set of options that define how DataBrew interprets the data in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-formatoptions
	//
	FormatOptions interface{} `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// A set of options that defines how DataBrew interprets an Amazon S3 path of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-pathoptions
	//
	PathOptions interface{} `field:"optional" json:"pathOptions" yaml:"pathOptions"`
	// Metadata tags that have been applied to the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

