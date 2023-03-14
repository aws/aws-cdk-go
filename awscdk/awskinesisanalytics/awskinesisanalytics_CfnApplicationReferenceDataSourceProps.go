package awskinesisanalytics


// Properties for defining a `CfnApplicationReferenceDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationReferenceDataSourceProps := &cfnApplicationReferenceDataSourceProps{
//   	applicationName: jsii.String("applicationName"),
//   	referenceDataSource: &referenceDataSourceProperty{
//   		referenceSchema: &referenceSchemaProperty{
//   			recordColumns: []interface{}{
//   				&recordColumnProperty{
//   					name: jsii.String("name"),
//   					sqlType: jsii.String("sqlType"),
//
//   					// the properties below are optional
//   					mapping: jsii.String("mapping"),
//   				},
//   			},
//   			recordFormat: &recordFormatProperty{
//   				recordFormatType: jsii.String("recordFormatType"),
//
//   				// the properties below are optional
//   				mappingParameters: &mappingParametersProperty{
//   					csvMappingParameters: &cSVMappingParametersProperty{
//   						recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   						recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   					},
//   					jsonMappingParameters: &jSONMappingParametersProperty{
//   						recordRowPath: jsii.String("recordRowPath"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			recordEncoding: jsii.String("recordEncoding"),
//   		},
//
//   		// the properties below are optional
//   		s3ReferenceDataSource: &s3ReferenceDataSourceProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			fileKey: jsii.String("fileKey"),
//   			referenceRoleArn: jsii.String("referenceRoleArn"),
//   		},
//   		tableName: jsii.String("tableName"),
//   	},
//   }
//
type CfnApplicationReferenceDataSourceProps struct {
	// Name of an existing application.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The reference data source can be an object in your Amazon S3 bucket.
	//
	// Amazon Kinesis Analytics reads the object and copies the data into the in-application table that is created. You provide an S3 bucket, object key name, and the resulting in-application table that is created. You must also provide an IAM role with the necessary permissions that Amazon Kinesis Analytics can assume to read the object from your S3 bucket on your behalf.
	ReferenceDataSource interface{} `field:"required" json:"referenceDataSource" yaml:"referenceDataSource"`
}

