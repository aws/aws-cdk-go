// An experiment to bundle the entire CDK into a single module
package awscdk


// The service to retrieve the dynamic reference from.
// Experimental.
type CfnDynamicReferenceService string

const (
	// Plaintext value stored in AWS Systems Manager Parameter Store.
	// Experimental.
	CfnDynamicReferenceService_SSM CfnDynamicReferenceService = "SSM"
	// Secure string stored in AWS Systems Manager Parameter Store.
	// Experimental.
	CfnDynamicReferenceService_SSM_SECURE CfnDynamicReferenceService = "SSM_SECURE"
	// Secret stored in AWS Secrets Manager.
	// Experimental.
	CfnDynamicReferenceService_SECRETS_MANAGER CfnDynamicReferenceService = "SECRETS_MANAGER"
)

