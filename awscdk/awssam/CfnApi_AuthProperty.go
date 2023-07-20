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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-auth.html
//
type CfnApi_AuthProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-auth.html#cfn-serverless-api-auth-adddefaultauthorizertocorspreflight
	//
	AddDefaultAuthorizerToCorsPreflight interface{} `field:"optional" json:"addDefaultAuthorizerToCorsPreflight" yaml:"addDefaultAuthorizerToCorsPreflight"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-auth.html#cfn-serverless-api-auth-authorizers
	//
	Authorizers interface{} `field:"optional" json:"authorizers" yaml:"authorizers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-auth.html#cfn-serverless-api-auth-defaultauthorizer
	//
	DefaultAuthorizer *string `field:"optional" json:"defaultAuthorizer" yaml:"defaultAuthorizer"`
}

