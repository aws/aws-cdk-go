package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var customStatements interface{}
//
//   apiEventProperty := &ApiEventProperty{
//   	Auth: &AuthProperty{
//   		ApiKeyRequired: jsii.Boolean(false),
//   		AuthorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		Authorizer: jsii.String("authorizer"),
//   		ResourcePolicy: &AuthResourcePolicyProperty{
//   			AwsAccountBlacklist: []*string{
//   				jsii.String("awsAccountBlacklist"),
//   			},
//   			AwsAccountWhitelist: []*string{
//   				jsii.String("awsAccountWhitelist"),
//   			},
//   			CustomStatements: []interface{}{
//   				customStatements,
//   			},
//   			IntrinsicVpcBlacklist: []*string{
//   				jsii.String("intrinsicVpcBlacklist"),
//   			},
//   			IntrinsicVpceBlacklist: []*string{
//   				jsii.String("intrinsicVpceBlacklist"),
//   			},
//   			IntrinsicVpceWhitelist: []*string{
//   				jsii.String("intrinsicVpceWhitelist"),
//   			},
//   			IntrinsicVpcWhitelist: []*string{
//   				jsii.String("intrinsicVpcWhitelist"),
//   			},
//   			IpRangeBlacklist: []*string{
//   				jsii.String("ipRangeBlacklist"),
//   			},
//   			IpRangeWhitelist: []*string{
//   				jsii.String("ipRangeWhitelist"),
//   			},
//   			SourceVpcBlacklist: []*string{
//   				jsii.String("sourceVpcBlacklist"),
//   			},
//   			SourceVpcWhitelist: []*string{
//   				jsii.String("sourceVpcWhitelist"),
//   			},
//   		},
//   	},
//   	Method: jsii.String("method"),
//   	Path: jsii.String("path"),
//   	RequestModel: &RequestModelProperty{
//   		Model: jsii.String("model"),
//   		Required: jsii.Boolean(false),
//   		ValidateBody: jsii.Boolean(false),
//   		ValidateParameters: jsii.Boolean(false),
//   	},
//   	RequestParameters: []interface{}{
//   		jsii.String("requestParameters"),
//   	},
//   	RestApiId: jsii.String("restApiId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-apievent.html
//
type CfnFunctionPropsMixin_ApiEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-apievent.html#cfn-serverless-function-apievent-auth
	//
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-apievent.html#cfn-serverless-function-apievent-method
	//
	Method *string `field:"optional" json:"method" yaml:"method"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-apievent.html#cfn-serverless-function-apievent-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-apievent.html#cfn-serverless-function-apievent-requestmodel
	//
	RequestModel interface{} `field:"optional" json:"requestModel" yaml:"requestModel"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-apievent.html#cfn-serverless-function-apievent-requestparameters
	//
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-apievent.html#cfn-serverless-function-apievent-restapiid
	//
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
}

