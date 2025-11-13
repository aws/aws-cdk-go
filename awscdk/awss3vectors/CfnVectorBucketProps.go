package awss3vectors


// Properties for defining a `CfnVectorBucket`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVectorBucketProps := &CfnVectorBucketProps{
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SseType: jsii.String("sseType"),
//   	},
//   	VectorBucketName: jsii.String("vectorBucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-vectorbucket.html
//
type CfnVectorBucketProps struct {
	// The encryption configuration for the vector bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-vectorbucket.html#cfn-s3vectors-vectorbucket-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// A name for the vector bucket.
	//
	// The bucket name must contain only lowercase letters, numbers, and hyphens (-). The bucket name must be unique in the same AWS account for each AWS Region. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name.
	//
	// The bucket name must be between 3 and 63 characters long and must not contain uppercase characters or underscores.
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-vectorbucket.html#cfn-s3vectors-vectorbucket-vectorbucketname
	//
	VectorBucketName *string `field:"optional" json:"vectorBucketName" yaml:"vectorBucketName"`
}

