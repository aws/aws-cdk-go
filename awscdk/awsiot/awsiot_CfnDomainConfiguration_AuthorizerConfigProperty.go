package awsiot


// An object that specifies the authorization service for a domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizerConfigProperty := &authorizerConfigProperty{
//   	allowAuthorizerOverride: jsii.Boolean(false),
//   	defaultAuthorizerName: jsii.String("defaultAuthorizerName"),
//   }
//
type CfnDomainConfiguration_AuthorizerConfigProperty struct {
	// A Boolean that specifies whether the domain configuration's authorization service can be overridden.
	AllowAuthorizerOverride interface{} `field:"optional" json:"allowAuthorizerOverride" yaml:"allowAuthorizerOverride"`
	// The name of the authorization service for a domain configuration.
	DefaultAuthorizerName *string `field:"optional" json:"defaultAuthorizerName" yaml:"defaultAuthorizerName"`
}

