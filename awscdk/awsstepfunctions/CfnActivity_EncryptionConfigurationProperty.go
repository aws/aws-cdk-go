package awsstepfunctions


// Settings to configure server-side encryption for an activity.
//
// By default, Step Functions provides transparent server-side encryption. With this configuration, you can specify a customer managed AWS KMS key for encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	KmsDataKeyReusePeriodSeconds: jsii.Number(123),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-activity-encryptionconfiguration.html
//
type CfnActivity_EncryptionConfigurationProperty struct {
	// Encryption option for an activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-activity-encryptionconfiguration.html#cfn-stepfunctions-activity-encryptionconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Maximum duration that Step Functions will reuse data keys.
	//
	// When the period expires, Step Functions will call `GenerateDataKey` . Only applies to customer managed keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-activity-encryptionconfiguration.html#cfn-stepfunctions-activity-encryptionconfiguration-kmsdatakeyreuseperiodseconds
	//
	KmsDataKeyReusePeriodSeconds *float64 `field:"optional" json:"kmsDataKeyReusePeriodSeconds" yaml:"kmsDataKeyReusePeriodSeconds"`
	// An alias, alias ARN, key ID, or key ARN of a symmetric encryption AWS KMS key to encrypt data.
	//
	// To specify a AWS KMS key in a different AWS account, you must use the key ARN or alias ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-activity-encryptionconfiguration.html#cfn-stepfunctions-activity-encryptionconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

