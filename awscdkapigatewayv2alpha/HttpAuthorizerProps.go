package awscdkapigatewayv2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties to initialize an instance of `HttpAuthorizer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var httpApi httpApi
//
//   httpAuthorizerProps := &HttpAuthorizerProps{
//   	HttpApi: httpApi,
//   	IdentitySource: []*string{
//   		jsii.String("identitySource"),
//   	},
//   	Type: apigatewayv2_alpha.HttpAuthorizerType_IAM,
//
//   	// the properties below are optional
//   	AuthorizerName: jsii.String("authorizerName"),
//   	AuthorizerUri: jsii.String("authorizerUri"),
//   	EnableSimpleResponses: jsii.Boolean(false),
//   	JwtAudience: []*string{
//   		jsii.String("jwtAudience"),
//   	},
//   	JwtIssuer: jsii.String("jwtIssuer"),
//   	PayloadFormatVersion: apigatewayv2_alpha.AuthorizerPayloadVersion_VERSION_1_0,
//   	ResultsCacheTtl: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
// Experimental.
type HttpAuthorizerProps struct {
	// HTTP Api to attach the authorizer to.
	// Experimental.
	HttpApi IHttpApi `field:"required" json:"httpApi" yaml:"httpApi"`
	// The identity source for which authorization is requested.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-identitysource
	//
	// Experimental.
	IdentitySource *[]*string `field:"required" json:"identitySource" yaml:"identitySource"`
	// The type of authorizer.
	// Experimental.
	Type HttpAuthorizerType `field:"required" json:"type" yaml:"type"`
	// Name of the authorizer.
	// Default: - id of the HttpAuthorizer construct.
	//
	// Experimental.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// The authorizer's Uniform Resource Identifier (URI).
	//
	// For REQUEST authorizers, this must be a well-formed Lambda function URI.
	// Default: - required for Request authorizer types.
	//
	// Experimental.
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
	// Specifies whether a Lambda authorizer returns a response in a simple format.
	//
	// If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
	// Default: - The lambda authorizer must return an IAM policy as its response.
	//
	// Experimental.
	EnableSimpleResponses *bool `field:"optional" json:"enableSimpleResponses" yaml:"enableSimpleResponses"`
	// A list of the intended recipients of the JWT.
	//
	// A valid JWT must provide an aud that matches at least one entry in this list.
	// Default: - required for JWT authorizer typess.
	//
	// Experimental.
	JwtAudience *[]*string `field:"optional" json:"jwtAudience" yaml:"jwtAudience"`
	// The base domain of the identity provider that issues JWT.
	// Default: - required for JWT authorizer types.
	//
	// Experimental.
	JwtIssuer *string `field:"optional" json:"jwtIssuer" yaml:"jwtIssuer"`
	// Specifies the format of the payload sent to an HTTP API Lambda authorizer.
	// Default: AuthorizerPayloadVersion.VERSION_2_0 if the authorizer type is HttpAuthorizerType.LAMBDA
	//
	// Experimental.
	PayloadFormatVersion AuthorizerPayloadVersion `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// How long APIGateway should cache the results.
	//
	// Max 1 hour.
	// Default: - API Gateway will not cache authorizer responses.
	//
	// Experimental.
	ResultsCacheTtl awscdk.Duration `field:"optional" json:"resultsCacheTtl" yaml:"resultsCacheTtl"`
}

