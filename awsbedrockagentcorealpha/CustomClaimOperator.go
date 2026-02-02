package awsbedrockagentcorealpha


// Custom claim match operator.
//
// Shared by Runtime and Gateway custom claim implementations.
//
// Example:
//   var userPool UserPool
//   var userPoolClient UserPoolClient
//
//
//   // Optional: Create custom claims (CustomClaimOperator and GatewayCustomClaim from agentcore)
//   customClaims := []GatewayCustomClaim{
//   	agentcore.GatewayCustomClaim_WithStringValue(jsii.String("department"), jsii.String("engineering")),
//   	agentcore.GatewayCustomClaim_WithStringArrayValue(jsii.String("roles"), []*string{
//   		jsii.String("admin"),
//   	}, agentcore.CustomClaimOperator_CONTAINS),
//   	agentcore.GatewayCustomClaim_WithStringArrayValue(jsii.String("permissions"), []*string{
//   		jsii.String("read"),
//   		jsii.String("write"),
//   	}, agentcore.CustomClaimOperator_CONTAINS_ANY),
//   }
//
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   	AuthorizerConfiguration: agentcore.GatewayAuthorizer_UsingCognito(&CognitoAuthorizerProps{
//   		UserPool: userPool,
//   		AllowedClients: []IUserPoolClient{
//   			userPoolClient,
//   		},
//   		AllowedAudiences: []*string{
//   			jsii.String("audience1"),
//   		},
//   		AllowedScopes: []*string{
//   			jsii.String("read"),
//   			jsii.String("write"),
//   		},
//   		CustomClaims: customClaims,
//   	}),
//   })
//
// Experimental.
type CustomClaimOperator string

const (
	// Equals operator - used for STRING type claims.
	// Experimental.
	CustomClaimOperator_EQUALS CustomClaimOperator = "EQUALS"
	// Contains operator - used for STRING_ARRAY type claims.
	//
	// Checks if the claim array contains a specific string value.
	// Experimental.
	CustomClaimOperator_CONTAINS CustomClaimOperator = "CONTAINS"
	// ContainsAny operator - used for STRING_ARRAY type claims.
	//
	// Checks if the claim array contains any of the provided string values.
	// Experimental.
	CustomClaimOperator_CONTAINS_ANY CustomClaimOperator = "CONTAINS_ANY"
)

