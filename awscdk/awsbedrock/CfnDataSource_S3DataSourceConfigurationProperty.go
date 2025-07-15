package awsbedrock


// The configuration information to connect to Amazon S3 as your data source.
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
//   	BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   	InclusionPrefixes: []*string{
//   		jsii.String("inclusionPrefixes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-s3datasourceconfiguration.html
//
type CfnDataSource_S3DataSourceConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the S3 bucket that contains your data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-s3datasourceconfiguration.html#cfn-bedrock-datasource-s3datasourceconfiguration-bucketarn
	//
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The account ID for the owner of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-s3datasourceconfiguration.html#cfn-bedrock-datasource-s3datasourceconfiguration-bucketowneraccountid
	//
	BucketOwnerAccountId *string `field:"optional" json:"bucketOwnerAccountId" yaml:"bucketOwnerAccountId"`
	// A list of S3 prefixes to include certain files or content.
	//
	// This field is an array with a maximum of one item, which can contain a string that has a maximum length of 300 characters. For more information, see [Organizing objects using prefixes](https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-prefixes.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-s3datasourceconfiguration.html#cfn-bedrock-datasource-s3datasourceconfiguration-inclusionprefixes
	//
	InclusionPrefixes *[]*string `field:"optional" json:"inclusionPrefixes" yaml:"inclusionPrefixes"`
}

