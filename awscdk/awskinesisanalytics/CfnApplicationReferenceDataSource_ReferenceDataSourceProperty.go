package awskinesisanalytics


// Describes the reference data source by providing the source information (S3 bucket name and object key name), the resulting in-application table name that is created, and the necessary schema to map the data elements in the Amazon S3 object to the in-application table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceDataSourceProperty := &ReferenceDataSourceProperty{
//   	ReferenceSchema: &ReferenceSchemaProperty{
//   		RecordColumns: []interface{}{
//   			&RecordColumnProperty{
//   				Name: jsii.String("name"),
//   				SqlType: jsii.String("sqlType"),
//
//   				// the properties below are optional
//   				Mapping: jsii.String("mapping"),
//   			},
//   		},
//   		RecordFormat: &RecordFormatProperty{
//   			RecordFormatType: jsii.String("recordFormatType"),
//
//   			// the properties below are optional
//   			MappingParameters: &MappingParametersProperty{
//   				CsvMappingParameters: &CSVMappingParametersProperty{
//   					RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   					RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   				},
//   				JsonMappingParameters: &JSONMappingParametersProperty{
//   					RecordRowPath: jsii.String("recordRowPath"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		RecordEncoding: jsii.String("recordEncoding"),
//   	},
//
//   	// the properties below are optional
//   	S3ReferenceDataSource: &S3ReferenceDataSourceProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		FileKey: jsii.String("fileKey"),
//   		ReferenceRoleArn: jsii.String("referenceRoleArn"),
//   	},
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html
//
type CfnApplicationReferenceDataSource_ReferenceDataSourceProperty struct {
	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns created in the in-application stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource-referenceschema
	//
	ReferenceSchema interface{} `field:"required" json:"referenceSchema" yaml:"referenceSchema"`
	// Identifies the S3 bucket and object that contains the reference data.
	//
	// Also identifies the IAM role Amazon Kinesis Analytics can assume to read this object on your behalf. An Amazon Kinesis Analytics application loads reference data only once. If the data changes, you call the `UpdateApplication` operation to trigger reloading of data into your application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource-s3referencedatasource
	//
	S3ReferenceDataSource interface{} `field:"optional" json:"s3ReferenceDataSource" yaml:"s3ReferenceDataSource"`
	// Name of the in-application table to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

