package mixinsawskinesisanalytics


// Properties for CfnApplicationReferenceDataSourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationReferenceDataSourceMixinProps := &CfnApplicationReferenceDataSourceMixinProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	ReferenceDataSource: &ReferenceDataSourceProperty{
//   		ReferenceSchema: &ReferenceSchemaProperty{
//   			RecordColumns: []interface{}{
//   				&RecordColumnProperty{
//   					Mapping: jsii.String("mapping"),
//   					Name: jsii.String("name"),
//   					SqlType: jsii.String("sqlType"),
//   				},
//   			},
//   			RecordEncoding: jsii.String("recordEncoding"),
//   			RecordFormat: &RecordFormatProperty{
//   				MappingParameters: &MappingParametersProperty{
//   					CsvMappingParameters: &CSVMappingParametersProperty{
//   						RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   						RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   					},
//   					JsonMappingParameters: &JSONMappingParametersProperty{
//   						RecordRowPath: jsii.String("recordRowPath"),
//   					},
//   				},
//   				RecordFormatType: jsii.String("recordFormatType"),
//   			},
//   		},
//   		S3ReferenceDataSource: &S3ReferenceDataSourceProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			FileKey: jsii.String("fileKey"),
//   			ReferenceRoleArn: jsii.String("referenceRoleArn"),
//   		},
//   		TableName: jsii.String("tableName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationreferencedatasource.html
//
type CfnApplicationReferenceDataSourceMixinProps struct {
	// Name of an existing application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationreferencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-applicationname
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The reference data source can be an object in your Amazon S3 bucket.
	//
	// Amazon Kinesis Analytics reads the object and copies the data into the in-application table that is created. You provide an S3 bucket, object key name, and the resulting in-application table that is created. You must also provide an IAM role with the necessary permissions that Amazon Kinesis Analytics can assume to read the object from your S3 bucket on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationreferencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource
	//
	ReferenceDataSource interface{} `field:"optional" json:"referenceDataSource" yaml:"referenceDataSource"`
}

