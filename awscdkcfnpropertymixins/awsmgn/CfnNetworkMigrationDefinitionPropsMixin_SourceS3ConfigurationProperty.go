package awsmgn


// S3 configuration for source network data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sourceS3ConfigurationProperty := &SourceS3ConfigurationProperty{
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3BucketOwner: jsii.String("s3BucketOwner"),
//   	S3Key: jsii.String("s3Key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-sources3configuration.html
//
type CfnNetworkMigrationDefinitionPropsMixin_SourceS3ConfigurationProperty struct {
	// The name of the S3 bucket containing source data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-sources3configuration.html#cfn-mgn-networkmigrationdefinition-sources3configuration-s3bucket
	//
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The AWS account ID of the S3 bucket owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-sources3configuration.html#cfn-mgn-networkmigrationdefinition-sources3configuration-s3bucketowner
	//
	S3BucketOwner *string `field:"optional" json:"s3BucketOwner" yaml:"s3BucketOwner"`
	// The S3 key (path) for the source data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-sources3configuration.html#cfn-mgn-networkmigrationdefinition-sources3configuration-s3key
	//
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
}

