package awskinesisanalyticsv2


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
//   		},
//   		tableName: jsii.String("tableName"),
//   	},
//   }
//
type CfnApplicationReferenceDataSourceProps struct {
	// The name of the application.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// For a SQL-based Kinesis Data Analytics application, describes the reference data source by providing the source information (Amazon S3 bucket name and object key name), the resulting in-application table name that is created, and the necessary schema to map the data elements in the Amazon S3 object to the in-application table.
	ReferenceDataSource interface{} `field:"required" json:"referenceDataSource" yaml:"referenceDataSource"`
}

