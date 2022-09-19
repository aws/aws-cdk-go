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
//   oAuthAuthorizationProps := &oAuthAuthorizationProps{
//   	authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	clientId: jsii.String("clientId"),
//   	clientSecret: secretValue,
//   	httpMethod: awscdk.Aws_events.httpMethod_POST,
//
//   	// the properties below are optional
//   	bodyParameters: map[string]*httpParameter{
//   		"bodyParametersKey": httpParameter,
//   	},
//   	headerParameters: map[string]*httpParameter{
//   		"headerParametersKey": httpParameter,
//   	},
//   	queryStringParameters: map[string]*httpParameter{
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
	BodyParameters *map[string]HttpParameter `field:"optional" json:"bodyParameters" yaml:"bodyParameters"`
	// Additional string parameters to add to the OAuth request header.
	HeaderParameters *map[string]HttpParameter `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// Additional string parameters to add to the OAuth request query string.
	QueryStringParameters *map[string]HttpParameter `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

