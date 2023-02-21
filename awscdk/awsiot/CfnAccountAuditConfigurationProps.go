package awsiot


// Properties for defining a `CfnAccountAuditConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountAuditConfigurationProps := &CfnAccountAuditConfigurationProps{
//   	AccountId: jsii.String("accountId"),
//   	AuditCheckConfigurations: &AuditCheckConfigurationsProperty{
//   		AuthenticatedCognitoRoleOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		CaCertificateExpiringCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		CaCertificateKeyQualityCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		ConflictingClientIdsCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		DeviceCertificateExpiringCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		DeviceCertificateKeyQualityCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		DeviceCertificateSharedCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		IntermediateCaRevokedForActiveDeviceCertificatesCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		IotPolicyOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		IoTPolicyPotentialMisConfigurationCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		IotRoleAliasAllowsAccessToUnusedServicesCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		IotRoleAliasOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		LoggingDisabledCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		RevokedCaCertificateStillActiveCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		RevokedDeviceCertificateStillActiveCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		UnauthenticatedCognitoRoleOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AuditNotificationTargetConfigurations: &AuditNotificationTargetConfigurationsProperty{
//   		Sns: &AuditNotificationTargetProperty{
//   			Enabled: jsii.Boolean(false),
//   			RoleArn: jsii.String("roleArn"),
//   			TargetArn: jsii.String("targetArn"),
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

