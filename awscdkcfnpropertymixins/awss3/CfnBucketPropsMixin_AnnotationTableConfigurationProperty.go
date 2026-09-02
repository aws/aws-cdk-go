package awss3


// The annotation table configuration for an S3 Metadata configuration.
//
// The annotation table tracks all annotations on objects in your bucket so that you can query annotation data at scale.
//   If you've disabled your annotation table configuration and now want to re-enable it, you must first manually delete the old annotation table from your AWS managed table bucket. Otherwise, the newly re-enabled annotation table configuration will enter a failed state because the annotation table already exists in the table bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   annotationTableConfigurationProperty := &AnnotationTableConfigurationProperty{
//   	ConfigurationState: jsii.String("configurationState"),
//   	EncryptionConfiguration: &MetadataTableEncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SseAlgorithm: jsii.String("sseAlgorithm"),
//   	},
//   	Role: jsii.String("role"),
//   	TableArn: jsii.String("tableArn"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-annotationtableconfiguration.html
//
type CfnBucketPropsMixin_AnnotationTableConfigurationProperty struct {
	// Specifies whether the annotation table configuration is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-annotationtableconfiguration.html#cfn-s3-bucket-annotationtableconfiguration-configurationstate
	//
	ConfigurationState *string `field:"optional" json:"configurationState" yaml:"configurationState"`
	// The encryption settings for an S3 Metadata journal table or inventory table configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-annotationtableconfiguration.html#cfn-s3-bucket-annotationtableconfiguration-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The ARN of the IAM role that grants Amazon S3 Metadata permission to read annotations from your bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-annotationtableconfiguration.html#cfn-s3-bucket-annotationtableconfiguration-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
	// The Amazon Resource Name (ARN) for the annotation table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-annotationtableconfiguration.html#cfn-s3-bucket-annotationtableconfiguration-tablearn
	//
	TableArn *string `field:"optional" json:"tableArn" yaml:"tableArn"`
	// The name of the annotation table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-annotationtableconfiguration.html#cfn-s3-bucket-annotationtableconfiguration-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

