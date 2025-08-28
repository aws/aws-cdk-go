package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for `Authorization.oauth()`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var httpParameter httpParameter
//   var secretValue secretValue
//
//   oAuthAuthorizationProps := &OAuthAuthorizationProps{
//   	AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: secretValue,
//   	HttpMethod: awscdk.Aws_events.HttpMethod_POST,
//
//   	// the properties below are optional
//   	BodyParameters: map[string]*httpParameter{
//   		"bodyParametersKey": httpParameter,
//   	},
//   	HeaderParameters: map[string]*httpParameter{
//   		"headerParametersKey": httpParameter,
//   	},
//   	QueryStringParameters: map[string]*httpParameter{
//   		"queryStringParametersKey": httpParameter,
//   	},
//   }
//
type OAuthAuthorizationProps struct {
	// The URL to the authorization endpoint.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The client ID to use for OAuth authorization for the connection.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret associated with the client ID to use for OAuth authorization for the connection.
	ClientSecret awscdk.SecretValue `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The method to use for the authorization request.
	//
	// (Can only choose POST, GET or PUT).
	HttpMethod HttpMethod `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// Additional string parameters to add to the OAuth request body.
	// Default: - No additional parameters.
	//
	BodyParameters *map[string]HttpParameter `field:"optional" json:"bodyParameters" yaml:"bodyParameters"`
	// Additional string parameters to add to the OAuth request header.
	// Default: - No additional parameters.
	//
	HeaderParameters *map[string]HttpParameter `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// Additional string parameters to add to the OAuth request query string.
	// Default: - No additional parameters.
	//
	QueryStringParameters *map[string]HttpParameter `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

