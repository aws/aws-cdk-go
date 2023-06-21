package awscognito


// OpenID Connect endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oidcEndpoints := &OidcEndpoints{
//   	Authorization: jsii.String("authorization"),
//   	JwksUri: jsii.String("jwksUri"),
//   	Token: jsii.String("token"),
//   	UserInfo: jsii.String("userInfo"),
//   }
//
type OidcEndpoints struct {
	// Authorization endpoint.
	Authorization *string `field:"required" json:"authorization" yaml:"authorization"`
	// Jwks_uri endpoint.
	JwksUri *string `field:"required" json:"jwksUri" yaml:"jwksUri"`
	// Token endpoint.
	Token *string `field:"required" json:"token" yaml:"token"`
	// UserInfo endpoint.
	UserInfo *string `field:"required" json:"userInfo" yaml:"userInfo"`
}

