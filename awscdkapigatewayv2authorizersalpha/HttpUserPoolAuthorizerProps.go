package awscdkapigatewayv2authorizersalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
)

// Properties to initialize HttpUserPoolAuthorizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_authorizers_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userPoolClient userPoolClient
//
//   httpUserPoolAuthorizerProps := &HttpUserPoolAuthorizerProps{
//   	AuthorizerName: jsii.String("authorizerName"),
//   	IdentitySource: []*string{
//   		jsii.String("identitySource"),
//   	},
//   	UserPoolClients: []iUserPoolClient{
//   		userPoolClient,
//   	},
//   	UserPoolRegion: jsii.String("userPoolRegion"),
//   }
//
// Experimental.
type HttpUserPoolAuthorizerProps struct {
	// Friendly name of the authorizer.
	// Default: - same value as `id` passed in the constructor.
	//
	// Experimental.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// The identity source for which authorization is requested.
	// Default: ['$request.header.Authorization']
	//
	// Experimental.
	IdentitySource *[]*string `field:"optional" json:"identitySource" yaml:"identitySource"`
	// The user pool clients that should be used to authorize requests with the user pool.
	// Default: - a new client will be created for the given user pool.
	//
	// Experimental.
	UserPoolClients *[]awscognito.IUserPoolClient `field:"optional" json:"userPoolClients" yaml:"userPoolClients"`
	// The AWS region in which the user pool is present.
	// Default: - same region as the Route the authorizer is attached to.
	//
	// Experimental.
	UserPoolRegion *string `field:"optional" json:"userPoolRegion" yaml:"userPoolRegion"`
}

