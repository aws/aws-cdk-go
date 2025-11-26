package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var authorizers interface{}
//
//   httpApiAuthProperty := &HttpApiAuthProperty{
//   	Authorizers: authorizers,
//   	DefaultAuthorizer: jsii.String("defaultAuthorizer"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-httpapiauth.html
//
type CfnHttpApiPropsMixin_HttpApiAuthProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-httpapiauth.html#cfn-serverless-httpapi-httpapiauth-authorizers
	//
	Authorizers interface{} `field:"optional" json:"authorizers" yaml:"authorizers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-httpapiauth.html#cfn-serverless-httpapi-httpapiauth-defaultauthorizer
	//
	DefaultAuthorizer *string `field:"optional" json:"defaultAuthorizer" yaml:"defaultAuthorizer"`
}

