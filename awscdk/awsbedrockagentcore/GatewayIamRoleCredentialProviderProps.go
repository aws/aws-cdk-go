package awsbedrockagentcore


// Properties for configuring the IAM role credential provider.
//
// When a target is authenticated via the gateway's execution role, the gateway signs
// outbound requests with SigV4. These properties let you tell the gateway which AWS
// service and region to sign for, instead of letting it infer them from the target
// endpoint.
//
// Example:
//   agentcore.GatewayCredentialProvider_FromIamRole(&GatewayIamRoleCredentialProviderProps{
//   	Service: jsii.String("bedrock-runtime"),
//   	 // SigV4 signing name (typically the endpoint prefix); see the AWS service authorization reference
//   	Region: jsii.String("us-east-1"),
//   })
//
type GatewayIamRoleCredentialProviderProps struct {
	// The AWS Region used for SigV4 signing of outbound requests. Can be up to 32 characters long.
	//
	// Pattern: ^[a-zA-Z0-9-]+$.
	// Default: - Gateway's own Region.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The AWS service name used for SigV4 signing of outbound requests.
	//
	// Use the SigV4 signing name (typically the endpoint prefix), e.g.
	// `bedrock-runtime`, `s3`, `execute-api`, `dynamodb`.
	// Can be up to 64 characters long.
	//
	// Pattern: ^[a-zA-Z0-9._-]+$
	// See: https://docs.aws.amazon.com/service-authorization/latest/reference/reference_policies_actions-resources-contextkeys.html
	//
	// Default: - Gateway infers the service from the target endpoint.
	//
	Service *string `field:"optional" json:"service" yaml:"service"`
}

