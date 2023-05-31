package awskinesisanalyticsv2


// Properties for defining a `CfnApplicationReferenceDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationReferenceDataSourceProps := &CfnApplicationReferenceDataSourceProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	ReferenceDataSource: &ReferenceDataSourceProperty{
//   		ReferenceSchema: &ReferenceSchemaProperty{
//   			RecordColumns: []interface{}{
//   				&RecordColumnProperty{
//   					Name: jsii.String("name"),
//   					SqlType: jsii.String("sqlType"),
//
//   					// the properties below are optional
//   					Mapping: jsii.String("mapping"),
//   				},
//   			},
//   			RecordFormat: &RecordFormatProperty{
//   				RecordFormatType: jsii.String("recordFormatType"),
//
//   				// the properties below are optional
//   				MappingParameters: &MappingParametersProperty{
//   					CsvMappingParameters: &CSVMappingParametersProperty{
//   						RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   						RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   					},
//   					JsonMappingParameters: &JSONMappingParametersProperty{
//   						RecordRowPath: jsii.String("recordRowPath"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			RecordEncoding: jsii.String("recordEncoding"),
//   		},
//
//   		// the properties below are optional
//   		S3ReferenceDataSource: &S3ReferenceDataSourceProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			FileKey: jsii.String("fileKey"),
//   		},
//   		TableName: jsii.String("tableName"),
//   	},
//   }
//
type CfnApplicationReferenceDataSourceProps struct {
	// The name of the application.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// For a SQL-based Kinesis Data Analytics application, describes the reference data source by providing the source information (Amazon S3 bucket name and object key name), the resulting in-application table name that is created, and the necessary schema to map the data elements in the Amazon S3 object to the in-application table.
	ReferenceDataSource interface{} `field:"required" json:"referenceDataSource" yaml:"referenceDataSource"`
}

