package awsapigatewayv2


// Reference to an http authorizer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   authorizerId := awscdk.Fn_ImportValue(jsii.String("authorizerId"))
//   authorizerType := awscdk.Fn_ImportValue(jsii.String("authorizerType"))
//
//   authorizer := awscdk.HttpAuthorizer_FromHttpAuthorizerAttributes(this, jsii.String("HttpAuthorizer"), &HttpAuthorizerAttributes{
//   	AuthorizerId: jsii.String(AuthorizerId),
//   	AuthorizerType: jsii.String(AuthorizerType),
//   })
//
type HttpAuthorizerAttributes struct {
	// Id of the Authorizer.
	AuthorizerId *string `field:"required" json:"authorizerId" yaml:"authorizerId"`
	// Type of authorizer.
	//
	// Possible values are:
	// - JWT - JSON Web Token Authorizer
	// - CUSTOM - Lambda Authorizer
	// - NONE - No Authorization.
	AuthorizerType *string `field:"required" json:"authorizerType" yaml:"authorizerType"`
}

