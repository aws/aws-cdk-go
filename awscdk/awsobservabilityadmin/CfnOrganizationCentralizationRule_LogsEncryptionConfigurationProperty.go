package awsobservabilityadmin


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logsEncryptionConfigurationProperty := &LogsEncryptionConfigurationProperty{
//   	EncryptionStrategy: jsii.String("encryptionStrategy"),
//
//   	// the properties below are optional
//   	EncryptionConflictResolutionStrategy: jsii.String("encryptionConflictResolutionStrategy"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration.html
//
type CfnOrganizationCentralizationRule_LogsEncryptionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration-encryptionstrategy
	//
	EncryptionStrategy *string `field:"required" json:"encryptionStrategy" yaml:"encryptionStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration-encryptionconflictresolutionstrategy
	//
	EncryptionConflictResolutionStrategy *string `field:"optional" json:"encryptionConflictResolutionStrategy" yaml:"encryptionConflictResolutionStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-logsencryptionconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

