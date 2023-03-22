package awskinesisanalyticsv2


// For a SQL-based Kinesis Data Analytics application, describes the reference data source by providing the source information (Amazon S3 bucket name and object key name), the resulting in-application table name that is created, and the necessary schema to map the data elements in the Amazon S3 object to the in-application table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceDataSourceProperty := &referenceDataSourceProperty{
//   	referenceSchema: &referenceSchemaProperty{
//   		recordColumns: []interface{}{
//   			&recordColumnProperty{
//   				name: jsii.String("name"),
//   				sqlType: jsii.String("sqlType"),
//
//   				// the properties below are optional
//   				mapping: jsii.String("mapping"),
//   			},
//   		},
//   		recordFormat: &recordFormatProperty{
//   			recordFormatType: jsii.String("recordFormatType"),
//
//   			// the properties below are optional
//   			mappingParameters: &mappingParametersProperty{
//   				csvMappingParameters: &cSVMappingParametersProperty{
//   					recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   					recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   				},
//   				jsonMappingParameters: &jSONMappingParametersProperty{
//   					recordRowPath: jsii.String("recordRowPath"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		recordEncoding: jsii.String("recordEncoding"),
//   	},
//
//   	// the properties below are optional
//   	s3ReferenceDataSource: &s3ReferenceDataSourceProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		fileKey: jsii.String("fileKey"),
//   	},
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnApplicationReferenceDataSource_ReferenceDataSourceProperty struct {
	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns created in the in-application stream.
	ReferenceSchema interface{} `field:"required" json:"referenceSchema" yaml:"referenceSchema"`
	// Identifies the S3 bucket and object that contains the reference data.
	//
	// A Kinesis Data Analytics application loads reference data only once. If the data changes, you call the [UpdateApplication](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_UpdateApplication.html) operation to trigger reloading of data into your application.
	S3ReferenceDataSource interface{} `field:"optional" json:"s3ReferenceDataSource" yaml:"s3ReferenceDataSource"`
	// The name of the in-application table to create.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

