package awsiot


// Properties for defining a `CfnAccountAuditConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountAuditConfigurationProps := &cfnAccountAuditConfigurationProps{
//   	accountId: jsii.String("accountId"),
//   	auditCheckConfigurations: &auditCheckConfigurationsProperty{
//   		authenticatedCognitoRoleOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		caCertificateExpiringCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		caCertificateKeyQualityCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		conflictingClientIdsCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		deviceCertificateExpiringCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		deviceCertificateKeyQualityCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		deviceCertificateSharedCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		intermediateCaRevokedForActiveDeviceCertificatesCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		iotPolicyOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		ioTPolicyPotentialMisConfigurationCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		iotRoleAliasAllowsAccessToUnusedServicesCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		iotRoleAliasOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		loggingDisabledCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		revokedCaCertificateStillActiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		revokedDeviceCertificateStillActiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		unauthenticatedCognitoRoleOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	auditNotificationTargetConfigurations: &auditNotificationTargetConfigurationsProperty{
//   		sns: &auditNotificationTargetProperty{
//   			enabled: jsii.Boolean(false),
//   			roleArn: jsii.String("roleArn"),
//   			targetArn: jsii.String("targetArn"),
//   		},
//   	},
//   }
//
type CfnAccountAuditConfigurationProps struct {
	// The ID of the account.
	//
	// You can use the expression `!Sub "${AWS::AccountId}"` to use your account ID.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Specifies which audit checks are enabled and disabled for this account.
	//
	// Some data collection might start immediately when certain checks are enabled. When a check is disabled, any data collected so far in relation to the check is deleted. To disable a check, set the value of the `Enabled:` key to `false` .
	//
	// If an enabled check is removed from the template, it will also be disabled.
	//
	// You can't disable a check if it's used by any scheduled audit. You must delete the check from the scheduled audit or delete the scheduled audit itself to disable the check.
	//
	// For more information on avialbe auidt checks see [AWS::IoT::AccountAuditConfiguration AuditCheckConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html)
	AuditCheckConfigurations interface{} `field:"required" json:"auditCheckConfigurations" yaml:"auditCheckConfigurations"`
	// The Amazon Resource Name (ARN) of the role that grants permission to AWS IoT to access information about your devices, policies, certificates, and other items as required when performing an audit.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Information about the targets to which audit notifications are sent.
	AuditNotificationTargetConfigurations interface{} `field:"optional" json:"auditNotificationTargetConfigurations" yaml:"auditNotificationTargetConfigurations"`
}

