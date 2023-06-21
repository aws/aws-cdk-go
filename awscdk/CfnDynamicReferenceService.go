package awscdk


// The service to retrieve the dynamic reference from.
//
// Example:
//   awscdk.NewCfnDynamicReference(awscdk.CfnDynamicReferenceService_SECRETS_MANAGER, jsii.String("secret-id:secret-string:json-key:version-stage:version-id"))
//
type CfnDynamicReferenceService string

const (
	// Plaintext value stored in AWS Systems Manager Parameter Store.
	CfnDynamicReferenceService_SSM CfnDynamicReferenceService = "SSM"
	// Secure string stored in AWS Systems Manager Parameter Store.
	CfnDynamicReferenceService_SSM_SECURE CfnDynamicReferenceService = "SSM_SECURE"
	// Secret stored in AWS Secrets Manager.
	CfnDynamicReferenceService_SECRETS_MANAGER CfnDynamicReferenceService = "SECRETS_MANAGER"
)

