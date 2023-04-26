// Authorizers for AWS APIGateway V2
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
	// Experimental.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `field:"optional" json:"identitySource" yaml:"identitySource"`
	// The user pool clients that should be used to authorize requests with the user pool.
	// Experimental.
	UserPoolClients *[]awscognito.IUserPoolClient `field:"optional" json:"userPoolClients" yaml:"userPoolClients"`
	// The AWS region in which the user pool is present.
	// Experimental.
	UserPoolRegion *string `field:"optional" json:"userPoolRegion" yaml:"userPoolRegion"`
}

