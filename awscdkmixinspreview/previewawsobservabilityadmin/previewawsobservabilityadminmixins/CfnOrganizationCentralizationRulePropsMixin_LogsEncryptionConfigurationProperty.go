package previewawsobservabilityadminmixins


// Configuration for encrypting centralized log groups.
//
// This configuration is only applied to destination log groups for which the corresponding source log groups are encrypted using Customer Managed KMS Keys.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logsEncryptionConfigurationProperty := &LogsEncryptionConfigurationProperty{
//   	EncryptionConflictResolutionStrategy: jsii.String("encryptionConflictResolutionStrategy"),
//   	EncryptionStrategy: jsii.String("encryptionStrategy"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration.html
//
type CfnOrganizationCentralizationRulePropsMixin_LogsEncryptionConfigurationProperty struct {
	// Conflict resolution strategy for centralization if the encryption strategy is set to CUSTOMER_MANAGED and the destination log group is encrypted with an AWS_OWNED KMS Key.
	//
	// ALLOW lets centralization go through while SKIP prevents centralization into the destination log group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration-encryptionconflictresolutionstrategy
	//
	EncryptionConflictResolutionStrategy *string `field:"optional" json:"encryptionConflictResolutionStrategy" yaml:"encryptionConflictResolutionStrategy"`
	// Configuration that determines the encryption strategy of the destination log groups.
	//
	// CUSTOMER_MANAGED uses the configured KmsKeyArn to encrypt newly created destination log groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration-encryptionstrategy
	//
	EncryptionStrategy *string `field:"optional" json:"encryptionStrategy" yaml:"encryptionStrategy"`
	// KMS Key ARN belonging to the primary destination account and region, to encrypt newly created central log groups in the primary destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

