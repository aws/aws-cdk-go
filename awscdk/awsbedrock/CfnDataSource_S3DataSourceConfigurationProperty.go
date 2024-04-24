package awsbedrock


// Contains information about the S3 configuration of the data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DataSourceConfigurationProperty := &S3DataSourceConfigurationProperty{
//   	BucketArn: jsii.String("bucketArn"),
//
//   	// the properties below are optional
//   	InclusionPrefixes: []*string{
//   		jsii.String("inclusionPrefixes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-s3datasourceconfiguration.html
//
type CfnDataSource_S3DataSourceConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the bucket that contains the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-s3datasourceconfiguration.html#cfn-bedrock-datasource-s3datasourceconfiguration-bucketarn
	//
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// A list of S3 prefixes that define the object containing the data sources.
	//
	// For more information, see [Organizing objects using prefixes](https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-prefixes.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-s3datasourceconfiguration.html#cfn-bedrock-datasource-s3datasourceconfiguration-inclusionprefixes
	//
	InclusionPrefixes *[]*string `field:"optional" json:"inclusionPrefixes" yaml:"inclusionPrefixes"`
}

