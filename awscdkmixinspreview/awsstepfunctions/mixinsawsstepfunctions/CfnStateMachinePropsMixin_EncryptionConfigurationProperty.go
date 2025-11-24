package mixinsawsstepfunctions


// Settings to configure server-side encryption for a state machine.
//
// By default, Step Functions provides transparent server-side encryption. With this configuration, you can specify a customer managed AWS  key for encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	KmsDataKeyReusePeriodSeconds: jsii.Number(123),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-encryptionconfiguration.html
//
type CfnStateMachinePropsMixin_EncryptionConfigurationProperty struct {
	// Maximum duration that Step Functions will reuse data keys.
	//
	// When the period expires, Step Functions will call `GenerateDataKey` . Only applies to customer managed keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-encryptionconfiguration.html#cfn-stepfunctions-statemachine-encryptionconfiguration-kmsdatakeyreuseperiodseconds
	//
	KmsDataKeyReusePeriodSeconds *float64 `field:"optional" json:"kmsDataKeyReusePeriodSeconds" yaml:"kmsDataKeyReusePeriodSeconds"`
	// An alias, alias ARN, key ID, or key ARN of a symmetric encryption AWS  key to encrypt data.
	//
	// To specify a AWS  key in a different AWS account, you must use the key ARN or alias ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-encryptionconfiguration.html#cfn-stepfunctions-statemachine-encryptionconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Encryption option for a state machine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-encryptionconfiguration.html#cfn-stepfunctions-statemachine-encryptionconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

