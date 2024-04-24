package awsbedrock


// Contains details about how a data source is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceConfigurationProperty := &DataSourceConfigurationProperty{
//   	S3Configuration: &S3DataSourceConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//
//   		// the properties below are optional
//   		InclusionPrefixes: []*string{
//   			jsii.String("inclusionPrefixes"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html
//
type CfnDataSource_DataSourceConfigurationProperty struct {
	// Contains details about the configuration of the S3 object containing the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html#cfn-bedrock-datasource-datasourceconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// The type of storage for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html#cfn-bedrock-datasource-datasourceconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

