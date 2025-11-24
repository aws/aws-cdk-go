package mixinsawss3


// > We recommend that you create your S3 Metadata configurations by using the V2 [MetadataConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-metadataconfiguration.html) resource type. We no longer recommend using the V1 `MetadataTableConfiguration` resource type. >  > If you created your S3 Metadata configuration before July 15, 2025, we recommend that you delete and re-create your configuration by using the [MetadataConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-metadataconfiguration.html) resource type so that you can expire journal table records and create a live inventory table.
//
// Creates a V1 S3 Metadata configuration for a general purpose bucket. For more information, see [Accelerating data discovery with S3 Metadata](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metadataTableConfigurationProperty := &MetadataTableConfigurationProperty{
//   	S3TablesDestination: &S3TablesDestinationProperty{
//   		TableArn: jsii.String("tableArn"),
//   		TableBucketArn: jsii.String("tableBucketArn"),
//   		TableName: jsii.String("tableName"),
//   		TableNamespace: jsii.String("tableNamespace"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadatatableconfiguration.html
//
type CfnBucketPropsMixin_MetadataTableConfigurationProperty struct {
	// The destination information for the metadata table configuration.
	//
	// The destination table bucket must be in the same Region and AWS account as the general purpose bucket. The specified metadata table name must be unique within the `aws_s3_metadata` namespace in the destination table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadatatableconfiguration.html#cfn-s3-bucket-metadatatableconfiguration-s3tablesdestination
	//
	S3TablesDestination interface{} `field:"optional" json:"s3TablesDestination" yaml:"s3TablesDestination"`
}

