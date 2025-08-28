package awsiot


// Properties for defining a `CfnEncryptionConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEncryptionConfigurationProps := &CfnEncryptionConfigurationProps{
//   	EncryptionType: jsii.String("encryptionType"),
//
//   	// the properties below are optional
//   	KmsAccessRoleArn: jsii.String("kmsAccessRoleArn"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-encryptionconfiguration.html
//
type CfnEncryptionConfigurationProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-encryptionconfiguration.html#cfn-iot-encryptionconfiguration-encryptiontype
	//
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-encryptionconfiguration.html#cfn-iot-encryptionconfiguration-kmsaccessrolearn
	//
	KmsAccessRoleArn *string `field:"optional" json:"kmsAccessRoleArn" yaml:"kmsAccessRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-encryptionconfiguration.html#cfn-iot-encryptionconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

