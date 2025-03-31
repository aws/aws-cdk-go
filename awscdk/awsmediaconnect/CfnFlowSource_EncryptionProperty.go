package awsmediaconnect


// Information about the encryption of the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionProperty := &EncryptionProperty{
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Algorithm: jsii.String("algorithm"),
//   	ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   	DeviceId: jsii.String("deviceId"),
//   	KeyType: jsii.String("keyType"),
//   	Region: jsii.String("region"),
//   	ResourceId: jsii.String("resourceId"),
//   	SecretArn: jsii.String("secretArn"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html
//
type CfnFlowSource_EncryptionProperty struct {
	// The ARN of the role that you created during setup (when you set up MediaConnect as a trusted entity).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The type of algorithm that is used for static key encryption (such as aes128, aes192, or aes256).
	//
	// If you are using SPEKE or SRT-password encryption, this property must be left blank.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-algorithm
	//
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content.
	//
	// This parameter is not valid for static key encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-constantinitializationvector
	//
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// The value of one of the devices that you configured with your digital rights management (DRM) platform key provider.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-deviceid
	//
	DeviceId *string `field:"optional" json:"deviceId" yaml:"deviceId"`
	// The type of key that is used for the encryption.
	//
	// If you don't specify a `keyType` value, the service uses the default setting ( `static-key` ). Valid key types are: `static-key` , `speke` , and `srt-password` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-keytype
	//
	// Default: - "static-key".
	//
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// The AWS Region that the API Gateway proxy endpoint was created in.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// An identifier for the content.
	//
	// The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The ARN of the secret that you created in AWS Secrets Manager to store the encryption key.
	//
	// This parameter is required for static key encryption and is not valid for SPEKE encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The URL from the API Gateway proxy that you set up to talk to your key server.
	//
	// This parameter is required for SPEKE encryption and is not valid for static key encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}

