package awss3


// Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.
//
// For information about the S3 Intelligent-Tiering storage class, see [Storage class for automatically optimizing frequently and infrequently accessed objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intelligentTieringConfigurationProperty := &IntelligentTieringConfigurationProperty{
//   	Id: jsii.String("id"),
//   	Status: jsii.String("status"),
//   	Tierings: []interface{}{
//   		&TieringProperty{
//   			AccessTier: jsii.String("accessTier"),
//   			Days: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   	TagFilters: []interface{}{
//   		&TagFilterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-intelligenttieringconfiguration.html
//
type CfnBucket_IntelligentTieringConfigurationProperty struct {
	// The ID used to identify the S3 Intelligent-Tiering configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-intelligenttieringconfiguration.html#cfn-s3-bucket-intelligenttieringconfiguration-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// Specifies the status of the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-intelligenttieringconfiguration.html#cfn-s3-bucket-intelligenttieringconfiguration-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// Specifies a list of S3 Intelligent-Tiering storage class tiers in the configuration.
	//
	// At least one tier must be defined in the list. At most, you can specify two tiers in the list, one for each available AccessTier: `ARCHIVE_ACCESS` and `DEEP_ARCHIVE_ACCESS` .
	//
	// > You only need Intelligent Tiering Configuration enabled on a bucket if you want to automatically move objects stored in the Intelligent-Tiering storage class to Archive Access or Deep Archive Access tiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-intelligenttieringconfiguration.html#cfn-s3-bucket-intelligenttieringconfiguration-tierings
	//
	Tierings interface{} `field:"required" json:"tierings" yaml:"tierings"`
	// An object key name prefix that identifies the subset of objects to which the rule applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-intelligenttieringconfiguration.html#cfn-s3-bucket-intelligenttieringconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// A container for a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-intelligenttieringconfiguration.html#cfn-s3-bucket-intelligenttieringconfiguration-tagfilters
	//
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

