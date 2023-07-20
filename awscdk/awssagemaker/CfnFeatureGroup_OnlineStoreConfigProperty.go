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
}

