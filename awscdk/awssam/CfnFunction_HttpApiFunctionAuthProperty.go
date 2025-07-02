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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-httpapifunctionauth.html
//
type CfnFunction_HttpApiFunctionAuthProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-httpapifunctionauth.html#cfn-serverless-function-httpapifunctionauth-authorizationscopes
	//
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-httpapifunctionauth.html#cfn-serverless-function-httpapifunctionauth-authorizer
	//
	Authorizer *string `field:"optional" json:"authorizer" yaml:"authorizer"`
}

