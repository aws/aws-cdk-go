package awssagemaker


// Use this to specify the AWS Key Management Service (KMS) Key ID, or `KMSKeyId` , for at rest data encryption.
//
// You can turn `OnlineStore` on or off by specifying the `EnableOnlineStore` flag at General Assembly.
//
// The default value is `False` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onlineStoreConfigProperty := &OnlineStoreConfigProperty{
//   	EnableOnlineStore: jsii.Boolean(false),
//   	SecurityConfig: &OnlineStoreSecurityConfigProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	StorageType: jsii.String("storageType"),
//   	TtlDuration: &TtlDurationProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-onlinestoreconfig.html
//
type CfnFeatureGroup_OnlineStoreConfigProperty struct {
	// Turn `OnlineStore` off by specifying `False` for the `EnableOnlineStore` flag.
	//
	// Turn `OnlineStore` on by specifying `True` for the `EnableOnlineStore` flag.
	//
	// The default value is `False` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-onlinestoreconfig.html#cfn-sagemaker-featuregroup-onlinestoreconfig-enableonlinestore
	//
	EnableOnlineStore interface{} `field:"optional" json:"enableOnlineStore" yaml:"enableOnlineStore"`
	// Use to specify KMS Key ID ( `KMSKeyId` ) for at-rest encryption of your `OnlineStore` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-onlinestoreconfig.html#cfn-sagemaker-featuregroup-onlinestoreconfig-securityconfig
	//
	SecurityConfig interface{} `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// Option for different tiers of low latency storage for real-time data retrieval.
	//
	// - `Standard` : A managed low latency data store for feature groups.
	// - `InMemory` : A managed data store for feature groups that supports very low latency retrieval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-onlinestoreconfig.html#cfn-sagemaker-featuregroup-onlinestoreconfig-storagetype
	//
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Time to live duration, where the record is hard deleted after the expiration time is reached;
	//
	// `ExpiresAt` = `EventTime` + `TtlDuration` . For information on HardDelete, see the [DeleteRecord](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_feature_store_DeleteRecord.html) API in the Amazon SageMaker API Reference guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-onlinestoreconfig.html#cfn-sagemaker-featuregroup-onlinestoreconfig-ttlduration
	//
	TtlDuration interface{} `field:"optional" json:"ttlDuration" yaml:"ttlDuration"`
}

