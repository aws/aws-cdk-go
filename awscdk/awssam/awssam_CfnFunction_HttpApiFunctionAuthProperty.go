package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpApiFunctionAuthProperty := &HttpApiFunctionAuthProperty{
//   	AuthorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	Authorizer: jsii.String("authorizer"),
//   }
//
type CfnFunction_HttpApiFunctionAuthProperty struct {
	// `CfnFunction.HttpApiFunctionAuthProperty.AuthorizationScopes`.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// `CfnFunction.HttpApiFunctionAuthProperty.Authorizer`.
	Authorizer *string `field:"optional" json:"authorizer" yaml:"authorizer"`
}

