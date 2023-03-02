package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customStatements interface{}
//
//   apiEventProperty := &apiEventProperty{
//   	method: jsii.String("method"),
//   	path: jsii.String("path"),
//
//   	// the properties below are optional
//   	auth: &authProperty{
//   		apiKeyRequired: jsii.Boolean(false),
//   		authorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		authorizer: jsii.String("authorizer"),
//   		resourcePolicy: &authResourcePolicyProperty{
//   			awsAccountBlacklist: []*string{
//   				jsii.String("awsAccountBlacklist"),
//   			},
//   			awsAccountWhitelist: []*string{
//   				jsii.String("awsAccountWhitelist"),
//   			},
//   			customStatements: []interface{}{
//   				customStatements,
//   			},
//   			intrinsicVpcBlacklist: []*string{
//   				jsii.String("intrinsicVpcBlacklist"),
//   			},
//   			intrinsicVpceBlacklist: []*string{
//   				jsii.String("intrinsicVpceBlacklist"),
//   			},
//   			intrinsicVpceWhitelist: []*string{
//   				jsii.String("intrinsicVpceWhitelist"),
//   			},
//   			intrinsicVpcWhitelist: []*string{
//   				jsii.String("intrinsicVpcWhitelist"),
//   			},
//   			ipRangeBlacklist: []*string{
//   				jsii.String("ipRangeBlacklist"),
//   			},
//   			ipRangeWhitelist: []*string{
//   				jsii.String("ipRangeWhitelist"),
//   			},
//   			sourceVpcBlacklist: []*string{
//   				jsii.String("sourceVpcBlacklist"),
//   			},
//   			sourceVpcWhitelist: []*string{
//   				jsii.String("sourceVpcWhitelist"),
//   			},
//   		},
//   	},
//   	requestModel: &requestModelProperty{
//   		model: jsii.String("model"),
//
//   		// the properties below are optional
//   		required: jsii.Boolean(false),
//   		validateBody: jsii.Boolean(false),
//   		validateParameters: jsii.Boolean(false),
//   	},
//   	requestParameters: []interface{}{
//   		jsii.String("requestParameters"),
//   	},
//   	restApiId: jsii.String("restApiId"),
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

