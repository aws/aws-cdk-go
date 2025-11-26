package previewawsdatabrewmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDatasetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDatasetMixinProps := &CfnDatasetMixinProps{
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
//   	Input: &InputProperty{
//   		DatabaseInputDefinition: &DatabaseInputDefinitionProperty{
//   			DatabaseTableName: jsii.String("databaseTableName"),
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//   			QueryString: jsii.String("queryString"),
//   			TempDirectory: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				BucketOwner: jsii.String("bucketOwner"),
//   				Key: jsii.String("key"),
//   			},
//   		},
//   		DataCatalogInputDefinition: &DataCatalogInputDefinitionProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
//   			TempDirectory: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				BucketOwner: jsii.String("bucketOwner"),
//   				Key: jsii.String("key"),
//   			},
//   		},
//   		Metadata: &MetadataProperty{
//   			SourceArn: jsii.String("sourceArn"),
//   		},
//   		S3InputDefinition: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			BucketOwner: jsii.String("bucketOwner"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	PathOptions: &PathOptionsProperty{
//   		FilesLimit: &FilesLimitProperty{
//   			MaxFiles: jsii.Number(123),
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
//   					CreateColumn: jsii.Boolean(false),
//   					DatetimeOptions: &DatetimeOptionsProperty{
//   						Format: jsii.String("format"),
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
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//   				},
//   				PathParameterName: jsii.String("pathParameterName"),
//   			},
//   		},
//   	},
//   	Source: jsii.String("source"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html
//
type CfnDatasetMixinProps struct {
	// The file format of a dataset that is created from an Amazon S3 file or folder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// A set of options that define how DataBrew interprets the data in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-formatoptions
	//
	FormatOptions interface{} `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// Information on how DataBrew can find the dataset, in either the AWS Glue Data Catalog or Amazon S3 .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-input
	//
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// The unique name of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A set of options that defines how DataBrew interprets an Amazon S3 path of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-pathoptions
	//
	PathOptions interface{} `field:"optional" json:"pathOptions" yaml:"pathOptions"`
	// The location of the data for the dataset, either Amazon S3 or the AWS Glue Data Catalog .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Metadata tags that have been applied to the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-dataset.html#cfn-databrew-dataset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

