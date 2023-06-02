package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customStatements interface{}
//
//   apiEventProperty := &ApiEventProperty{
//   	Method: jsii.String("method"),
//   	Path: jsii.String("path"),
//
//   	// the properties below are optional
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
//   	RequestModel: &RequestModelProperty{
//   		Model: jsii.String("model"),
//
//   		// the properties below are optional
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
type CfnFunction_ApiEventProperty struct {
	// `CfnFunction.ApiEventProperty.Method`.
	Method *string `field:"required" json:"method" yaml:"method"`
	// `CfnFunction.ApiEventProperty.Path`.
	Path *string `field:"required" json:"path" yaml:"path"`
	// `CfnFunction.ApiEventProperty.Auth`.
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// `CfnFunction.ApiEventProperty.RequestModel`.
	RequestModel interface{} `field:"optional" json:"requestModel" yaml:"requestModel"`
	// `CfnFunction.ApiEventProperty.RequestParameters`.
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// `CfnFunction.ApiEventProperty.RestApiId`.
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
}

