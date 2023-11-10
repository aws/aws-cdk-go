package awsiot


// The types of audit checks that can be performed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auditCheckConfigurationsProperty := &AuditCheckConfigurationsProperty{
//   	AuthenticatedCognitoRoleOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	CaCertificateExpiringCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	CaCertificateKeyQualityCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	ConflictingClientIdsCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	DeviceCertificateExpiringCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	DeviceCertificateKeyQualityCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	DeviceCertificateSharedCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	IntermediateCaRevokedForActiveDeviceCertificatesCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	IotPolicyOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	IoTPolicyPotentialMisConfigurationCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	IotRoleAliasAllowsAccessToUnusedServicesCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	IotRoleAliasOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	LoggingDisabledCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	RevokedCaCertificateStillActiveCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	RevokedDeviceCertificateStillActiveCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	UnauthenticatedCognitoRoleOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html
//
type CfnAccountAuditConfiguration_AuditCheckConfigurationsProperty struct {
	// Checks the permissiveness of an authenticated Amazon Cognito identity pool role.
	//
	// For this check, AWS IoT Device Defender audits all Amazon Cognito identity pools that have been used to connect to the AWS IoT message broker during the 31 days before the audit is performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-authenticatedcognitoroleoverlypermissivecheck
	//
	AuthenticatedCognitoRoleOverlyPermissiveCheck interface{} `field:"optional" json:"authenticatedCognitoRoleOverlyPermissiveCheck" yaml:"authenticatedCognitoRoleOverlyPermissiveCheck"`
	// Checks if a CA certificate is expiring.
	//
	// This check applies to CA certificates expiring within 30 days or that have expired.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-cacertificateexpiringcheck
	//
	CaCertificateExpiringCheck interface{} `field:"optional" json:"caCertificateExpiringCheck" yaml:"caCertificateExpiringCheck"`
	// Checks the quality of the CA certificate key.
	//
	// The quality checks if the key is in a valid format, not expired, and if the key meets a minimum required size. This check applies to CA certificates that are `ACTIVE` or `PENDING_TRANSFER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-cacertificatekeyqualitycheck
	//
	CaCertificateKeyQualityCheck interface{} `field:"optional" json:"caCertificateKeyQualityCheck" yaml:"caCertificateKeyQualityCheck"`
	// Checks if multiple devices connect using the same client ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-conflictingclientidscheck
	//
	ConflictingClientIdsCheck interface{} `field:"optional" json:"conflictingClientIdsCheck" yaml:"conflictingClientIdsCheck"`
	// Checks if a device certificate is expiring.
	//
	// This check applies to device certificates expiring within 30 days or that have expired.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-devicecertificateexpiringcheck
	//
	DeviceCertificateExpiringCheck interface{} `field:"optional" json:"deviceCertificateExpiringCheck" yaml:"deviceCertificateExpiringCheck"`
	// Checks the quality of the device certificate key.
	//
	// The quality checks if the key is in a valid format, not expired, signed by a registered certificate authority, and if the key meets a minimum required size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-devicecertificatekeyqualitycheck
	//
	DeviceCertificateKeyQualityCheck interface{} `field:"optional" json:"deviceCertificateKeyQualityCheck" yaml:"deviceCertificateKeyQualityCheck"`
	// Checks if multiple concurrent connections use the same X.509 certificate to authenticate with AWS IoT .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-devicecertificatesharedcheck
	//
	DeviceCertificateSharedCheck interface{} `field:"optional" json:"deviceCertificateSharedCheck" yaml:"deviceCertificateSharedCheck"`
	// Checks if device certificates are still active despite being revoked by an intermediate CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-intermediatecarevokedforactivedevicecertificatescheck
	//
	IntermediateCaRevokedForActiveDeviceCertificatesCheck interface{} `field:"optional" json:"intermediateCaRevokedForActiveDeviceCertificatesCheck" yaml:"intermediateCaRevokedForActiveDeviceCertificatesCheck"`
	// Checks the permissiveness of a policy attached to an authenticated Amazon Cognito identity pool role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-iotpolicyoverlypermissivecheck
	//
	IotPolicyOverlyPermissiveCheck interface{} `field:"optional" json:"iotPolicyOverlyPermissiveCheck" yaml:"iotPolicyOverlyPermissiveCheck"`
	// Checks if an AWS IoT policy is potentially misconfigured.
	//
	// Misconfigured policies, including overly permissive policies, can cause security incidents like allowing devices access to unintended resources. This check is a warning for you to make sure that only intended actions are allowed before updating the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-iotpolicypotentialmisconfigurationcheck
	//
	IoTPolicyPotentialMisConfigurationCheck interface{} `field:"optional" json:"ioTPolicyPotentialMisConfigurationCheck" yaml:"ioTPolicyPotentialMisConfigurationCheck"`
	// Checks if a role alias has access to services that haven't been used for the AWS IoT device in the last year.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-iotrolealiasallowsaccesstounusedservicescheck
	//
	IotRoleAliasAllowsAccessToUnusedServicesCheck interface{} `field:"optional" json:"iotRoleAliasAllowsAccessToUnusedServicesCheck" yaml:"iotRoleAliasAllowsAccessToUnusedServicesCheck"`
	// Checks if the temporary credentials provided by AWS IoT role aliases are overly permissive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-iotrolealiasoverlypermissivecheck
	//
	IotRoleAliasOverlyPermissiveCheck interface{} `field:"optional" json:"iotRoleAliasOverlyPermissiveCheck" yaml:"iotRoleAliasOverlyPermissiveCheck"`
	// Checks if AWS IoT logs are disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-loggingdisabledcheck
	//
	LoggingDisabledCheck interface{} `field:"optional" json:"loggingDisabledCheck" yaml:"loggingDisabledCheck"`
	// Checks if a revoked CA certificate is still active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-revokedcacertificatestillactivecheck
	//
	RevokedCaCertificateStillActiveCheck interface{} `field:"optional" json:"revokedCaCertificateStillActiveCheck" yaml:"revokedCaCertificateStillActiveCheck"`
	// Checks if a revoked device certificate is still active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-revokeddevicecertificatestillactivecheck
	//
	RevokedDeviceCertificateStillActiveCheck interface{} `field:"optional" json:"revokedDeviceCertificateStillActiveCheck" yaml:"revokedDeviceCertificateStillActiveCheck"`
	// Checks if policy attached to an unauthenticated Amazon Cognito identity pool role is too permissive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html#cfn-iot-accountauditconfiguration-auditcheckconfigurations-unauthenticatedcognitoroleoverlypermissivecheck
	//
	UnauthenticatedCognitoRoleOverlyPermissiveCheck interface{} `field:"optional" json:"unauthenticatedCognitoRoleOverlyPermissiveCheck" yaml:"unauthenticatedCognitoRoleOverlyPermissiveCheck"`
}

