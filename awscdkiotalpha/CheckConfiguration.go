package awscdkiotalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The types of audit checks.
//
// Example:
//   iot.NewAccountAuditConfiguration(this, jsii.String("AuditConfiguration"), &AccountAuditConfigurationProps{
//   	CheckConfiguration: &CheckConfiguration{
//   		// enabled
//   		AuthenticatedCognitoRoleOverlyPermissiveCheck: jsii.Boolean(true),
//   		// enabled by default
//   		CaCertificateExpiringCheck: undefined,
//   		// disabled
//   		CaCertificateKeyQualityCheck: jsii.Boolean(false),
//   		ConflictingClientIdsCheck: jsii.Boolean(false),
//   		DeviceCertificateAgeCheck: jsii.Boolean(false),
//   		DeviceCertificateExpiringCheck: jsii.Boolean(false),
//   		DeviceCertificateKeyQualityCheck: jsii.Boolean(false),
//   		DeviceCertificateSharedCheck: jsii.Boolean(false),
//   		IntermediateCaRevokedForActiveDeviceCertificatesCheck: jsii.Boolean(false),
//   		IoTPolicyPotentialMisConfigurationCheck: jsii.Boolean(false),
//   		IotPolicyOverlyPermissiveCheck: jsii.Boolean(false),
//   		IotRoleAliasAllowsAccessToUnusedServicesCheck: jsii.Boolean(false),
//   		IotRoleAliasOverlyPermissiveCheck: jsii.Boolean(false),
//   		LoggingDisabledCheck: jsii.Boolean(false),
//   		RevokedCaCertificateStillActiveCheck: jsii.Boolean(false),
//   		RevokedDeviceCertificateStillActiveCheck: jsii.Boolean(false),
//   		UnauthenticatedCognitoRoleOverlyPermissiveCheck: jsii.Boolean(false),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/iot-device-defender/latest/devguide/device-defender-audit-checks.html
//
// Experimental.
type CheckConfiguration struct {
	// Checks the permissiveness of an authenticated Amazon Cognito identity pool role.
	//
	// For this check, AWS IoT Device Defender audits all Amazon Cognito identity pools that have been used to connect to the AWS IoT message broker
	// during the 31 days before the audit is performed.
	// Default: true.
	//
	// Experimental.
	AuthenticatedCognitoRoleOverlyPermissiveCheck *bool `field:"optional" json:"authenticatedCognitoRoleOverlyPermissiveCheck" yaml:"authenticatedCognitoRoleOverlyPermissiveCheck"`
	// Checks if a CA certificate is expiring.
	//
	// This check applies to CA certificates expiring within 30 days or that have expired.
	// Default: true.
	//
	// Experimental.
	CaCertificateExpiringCheck *bool `field:"optional" json:"caCertificateExpiringCheck" yaml:"caCertificateExpiringCheck"`
	// Checks the quality of the CA certificate key.
	//
	// The quality checks if the key is in a valid format, not expired, and if the key meets a minimum required size.
	//
	// This check applies to CA certificates that are ACTIVE or PENDING_TRANSFER.
	// Default: true.
	//
	// Experimental.
	CaCertificateKeyQualityCheck *bool `field:"optional" json:"caCertificateKeyQualityCheck" yaml:"caCertificateKeyQualityCheck"`
	// Checks if multiple devices connect using the same client ID.
	// Default: true.
	//
	// Experimental.
	ConflictingClientIdsCheck *bool `field:"optional" json:"conflictingClientIdsCheck" yaml:"conflictingClientIdsCheck"`
	// Checks when a device certificate has been active for a number of days greater than or equal to the number you specify.
	// Default: true.
	//
	// Experimental.
	DeviceCertificateAgeCheck *bool `field:"optional" json:"deviceCertificateAgeCheck" yaml:"deviceCertificateAgeCheck"`
	// The duration used to check if a device certificate has been active for a number of days greater than or equal to the number you specify.
	//
	// Valid values range from 30 days (minimum) to 3650 days (10 years, maximum).
	//
	// You cannot specify a value for this check if `deviceCertificateAgeCheck` is set to `false`.
	// Default: - 365 days.
	//
	// Experimental.
	DeviceCertificateAgeCheckDuration awscdk.Duration `field:"optional" json:"deviceCertificateAgeCheckDuration" yaml:"deviceCertificateAgeCheckDuration"`
	// Checks if a device certificate is expiring.
	//
	// This check applies to device certificates expiring within 30 days or that have expired.
	// Default: true.
	//
	// Experimental.
	DeviceCertificateExpiringCheck *bool `field:"optional" json:"deviceCertificateExpiringCheck" yaml:"deviceCertificateExpiringCheck"`
	// Checks the quality of the device certificate key.
	//
	// The quality checks if the key is in a valid format, not expired, signed by a registered certificate authority,
	// and if the key meets a minimum required size.
	// Default: true.
	//
	// Experimental.
	DeviceCertificateKeyQualityCheck *bool `field:"optional" json:"deviceCertificateKeyQualityCheck" yaml:"deviceCertificateKeyQualityCheck"`
	// Checks if multiple concurrent connections use the same X.509 certificate to authenticate with AWS IoT.
	// Default: true.
	//
	// Experimental.
	DeviceCertificateSharedCheck *bool `field:"optional" json:"deviceCertificateSharedCheck" yaml:"deviceCertificateSharedCheck"`
	// Checks if device certificates are still active despite being revoked by an intermediate CA.
	// Default: true.
	//
	// Experimental.
	IntermediateCaRevokedForActiveDeviceCertificatesCheck *bool `field:"optional" json:"intermediateCaRevokedForActiveDeviceCertificatesCheck" yaml:"intermediateCaRevokedForActiveDeviceCertificatesCheck"`
	// Checks the permissiveness of a policy attached to an authenticated Amazon Cognito identity pool role.
	// Default: true.
	//
	// Experimental.
	IotPolicyOverlyPermissiveCheck *bool `field:"optional" json:"iotPolicyOverlyPermissiveCheck" yaml:"iotPolicyOverlyPermissiveCheck"`
	// Checks if an AWS IoT policy is potentially misconfigured.
	//
	// Misconfigured policies, including overly permissive policies, can cause security incidents like allowing devices access to unintended resources.
	//
	// This check is a warning for you to make sure that only intended actions are allowed before updating the policy.
	// Default: true.
	//
	// Experimental.
	IoTPolicyPotentialMisConfigurationCheck *bool `field:"optional" json:"ioTPolicyPotentialMisConfigurationCheck" yaml:"ioTPolicyPotentialMisConfigurationCheck"`
	// Checks if a role alias has access to services that haven't been used for the AWS IoT device in the last year.
	// Default: true.
	//
	// Experimental.
	IotRoleAliasAllowsAccessToUnusedServicesCheck *bool `field:"optional" json:"iotRoleAliasAllowsAccessToUnusedServicesCheck" yaml:"iotRoleAliasAllowsAccessToUnusedServicesCheck"`
	// Checks if the temporary credentials provided by AWS IoT role aliases are overly permissive.
	// Default: true.
	//
	// Experimental.
	IotRoleAliasOverlyPermissiveCheck *bool `field:"optional" json:"iotRoleAliasOverlyPermissiveCheck" yaml:"iotRoleAliasOverlyPermissiveCheck"`
	// Checks if AWS IoT logs are disabled.
	// Default: true.
	//
	// Experimental.
	LoggingDisabledCheck *bool `field:"optional" json:"loggingDisabledCheck" yaml:"loggingDisabledCheck"`
	// Checks if a revoked CA certificate is still active.
	// Default: true.
	//
	// Experimental.
	RevokedCaCertificateStillActiveCheck *bool `field:"optional" json:"revokedCaCertificateStillActiveCheck" yaml:"revokedCaCertificateStillActiveCheck"`
	// Checks if a revoked device certificate is still active.
	// Default: true.
	//
	// Experimental.
	RevokedDeviceCertificateStillActiveCheck *bool `field:"optional" json:"revokedDeviceCertificateStillActiveCheck" yaml:"revokedDeviceCertificateStillActiveCheck"`
	// Checks if policy attached to an unauthenticated Amazon Cognito identity pool role is too permissive.
	// Default: true.
	//
	// Experimental.
	UnauthenticatedCognitoRoleOverlyPermissiveCheck *bool `field:"optional" json:"unauthenticatedCognitoRoleOverlyPermissiveCheck" yaml:"unauthenticatedCognitoRoleOverlyPermissiveCheck"`
}

