package previewawss3mixins


// This resource contains the details of the bucket where the Amazon S3 Storage Lens metrics export will be placed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var sses3 interface{}
//
//   s3BucketDestinationProperty := &S3BucketDestinationProperty{
//   	AccountId: jsii.String("accountId"),
//   	Arn: jsii.String("arn"),
//   	Encryption: &EncryptionProperty{
//   		Ssekms: &SSEKMSProperty{
//   			KeyId: jsii.String("keyId"),
//   		},
//   		Sses3: sses3,
//   	},
//   	Format: jsii.String("format"),
//   	OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-s3bucketdestination.html
//
type CfnStorageLensPropsMixin_S3BucketDestinationProperty struct {
	// This property contains the details of the AWS account ID of the S3 Storage Lens export bucket destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-s3bucketdestination.html#cfn-s3-storagelens-s3bucketdestination-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// This property contains the details of the ARN of the bucket destination of the S3 Storage Lens export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-s3bucketdestination.html#cfn-s3-storagelens-s3bucketdestination-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// This property contains the details of the encryption of the bucket destination of the Amazon S3 Storage Lens metrics export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-s3bucketdestination.html#cfn-s3-storagelens-s3bucketdestination-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// This property contains the details of the format of the S3 Storage Lens export bucket destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-s3bucketdestination.html#cfn-s3-storagelens-s3bucketdestination-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// This property contains the details of the output schema version of the S3 Storage Lens export bucket destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-s3bucketdestination.html#cfn-s3-storagelens-s3bucketdestination-outputschemaversion
	//
	OutputSchemaVersion *string `field:"optional" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
	// This property contains the details of the prefix of the bucket destination of the S3 Storage Lens export .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-s3bucketdestination.html#cfn-s3-storagelens-s3bucketdestination-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

