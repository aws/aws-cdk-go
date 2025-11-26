package previewawssecuritylakemixins


// Provides lifecycle details of Amazon Security Lake object.
//
// To manage your data so that it is stored cost effectively, you can configure retention settings for the data. You can specify your preferred Amazon S3 storage class and the time period for Amazon S3 objects to stay in that storage class before they transition to a different storage class or expire. For more information about Amazon S3 Lifecycle configurations, see [Managing your storage lifecycle](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html) in the *Amazon Simple Storage Service User Guide* .
//
// In Security Lake , you specify retention settings at the Region level. For example, you might choose to transition all S3 objects in a specific AWS Region to the `S3 Standard-IA` storage class 30 days after they're written to the data lake. The default Amazon S3 storage class is S3 Standard.
//
// > Security Lake doesn't support Amazon S3 Object Lock. When the data lake buckets are created, S3 Object Lock is disabled by default. Enabling S3 Object Lock with default retention mode interrupts the delivery of normalized log data to the data lake.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lifecycleConfigurationProperty := &LifecycleConfigurationProperty{
//   	Expiration: &ExpirationProperty{
//   		Days: jsii.Number(123),
//   	},
//   	Transitions: []interface{}{
//   		&TransitionsProperty{
//   			Days: jsii.Number(123),
//   			StorageClass: jsii.String("storageClass"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-lifecycleconfiguration.html
//
type CfnDataLakePropsMixin_LifecycleConfigurationProperty struct {
	// Provides data expiration details of the Amazon Security Lake object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-lifecycleconfiguration.html#cfn-securitylake-datalake-lifecycleconfiguration-expiration
	//
	Expiration interface{} `field:"optional" json:"expiration" yaml:"expiration"`
	// Provides data storage transition details of Amazon Security Lake object.
	//
	// By configuring these settings, you can specify your preferred Amazon S3 storage class and the time period for S3 objects to stay in that storage class before they transition to a different storage class.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-lifecycleconfiguration.html#cfn-securitylake-datalake-lifecycleconfiguration-transitions
	//
	Transitions interface{} `field:"optional" json:"transitions" yaml:"transitions"`
}

