package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizers interface{}
//
//   authProperty := &AuthProperty{
//   	AddDefaultAuthorizerToCorsPreflight: jsii.Boolean(false),
//   	Authorizers: authorizers,
//   	DefaultAuthorizer: jsii.String("defaultAuthorizer"),
//   }
//
type CfnApi_AuthProperty struct {
	// `CfnApi.AuthProperty.AddDefaultAuthorizerToCorsPreflight`.
	AddDefaultAuthorizerToCorsPreflight interface{} `field:"optional" json:"addDefaultAuthorizerToCorsPreflight" yaml:"addDefaultAuthorizerToCorsPreflight"`
	// `CfnApi.AuthProperty.Authorizers`.
	Authorizers interface{} `field:"optional" json:"authorizers" yaml:"authorizers"`
	// `CfnApi.AuthProperty.DefaultAuthorizer`.
	DefaultAuthorizer *string `field:"optional" json:"defaultAuthorizer" yaml:"defaultAuthorizer"`
}

