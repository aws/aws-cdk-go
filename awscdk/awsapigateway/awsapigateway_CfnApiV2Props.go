package awsapigateway


// Properties for defining a `AWS::ApiGatewayV2::Api`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var body interface{}
//   var tags interface{}
//
//   cfnApiV2Props := &cfnApiV2Props{
//   	apiKeySelectionExpression: jsii.String("apiKeySelectionExpression"),
//   	basePath: jsii.String("basePath"),
//   	body: body,
//   	bodyS3Location: &bodyS3LocationProperty{
//   		bucket: jsii.String("bucket"),
//   		etag: jsii.String("etag"),
//   		key: jsii.String("key"),
//   		version: jsii.String("version"),
//   	},
//   	corsConfiguration: &corsProperty{
//   		allowCredentials: jsii.Boolean(false),
//   		allowHeaders: []*string{
//   			jsii.String("allowHeaders"),
//   		},
//   		allowMethods: []*string{
//   			jsii.String("allowMethods"),
//   		},
//   		allowOrigins: []*string{
//   			jsii.String("allowOrigins"),
//   		},
//   		exposeHeaders: []*string{
//   			jsii.String("exposeHeaders"),
//   		},
//   		maxAge: jsii.Number(123),
//   	},
//   	credentialsArn: jsii.String("credentialsArn"),
//   	description: jsii.String("description"),
//   	disableSchemaValidation: jsii.Boolean(false),
//   	failOnWarnings: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	protocolType: jsii.String("protocolType"),
//   	routeKey: jsii.String("routeKey"),
//   	routeSelectionExpression: jsii.String("routeSelectionExpression"),
//   	tags: tags,
//   	target: jsii.String("target"),
//   	version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnApiV2Props struct {
	// `AWS::ApiGatewayV2::Api.ApiKeySelectionExpression`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-apikeyselectionexpression
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiKeySelectionExpression *string `field:"optional" json:"apiKeySelectionExpression" yaml:"apiKeySelectionExpression"`
	// `AWS::ApiGatewayV2::Api.BasePath`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-basepath
	//
	// Deprecated: moved to package aws-apigatewayv2.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// `AWS::ApiGatewayV2::Api.Body`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-body
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// `AWS::ApiGatewayV2::Api.BodyS3Location`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-bodys3location
	//
	// Deprecated: moved to package aws-apigatewayv2.
	BodyS3Location interface{} `field:"optional" json:"bodyS3Location" yaml:"bodyS3Location"`
	// `AWS::ApiGatewayV2::Api.CorsConfiguration`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-corsconfiguration
	//
	// Deprecated: moved to package aws-apigatewayv2.
	CorsConfiguration interface{} `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// `AWS::ApiGatewayV2::Api.CredentialsArn`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-credentialsarn
	//
	// Deprecated: moved to package aws-apigatewayv2.
	CredentialsArn *string `field:"optional" json:"credentialsArn" yaml:"credentialsArn"`
	// `AWS::ApiGatewayV2::Api.Description`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-description
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::ApiGatewayV2::Api.DisableSchemaValidation`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-disableschemavalidation
	//
	// Deprecated: moved to package aws-apigatewayv2.
	DisableSchemaValidation interface{} `field:"optional" json:"disableSchemaValidation" yaml:"disableSchemaValidation"`
	// `AWS::ApiGatewayV2::Api.FailOnWarnings`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-failonwarnings
	//
	// Deprecated: moved to package aws-apigatewayv2.
	FailOnWarnings interface{} `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// `AWS::ApiGatewayV2::Api.Name`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-name
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::ApiGatewayV2::Api.ProtocolType`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-protocoltype
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// `AWS::ApiGatewayV2::Api.RouteKey`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-routekey
	//
	// Deprecated: moved to package aws-apigatewayv2.
	RouteKey *string `field:"optional" json:"routeKey" yaml:"routeKey"`
	// `AWS::ApiGatewayV2::Api.RouteSelectionExpression`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-routeselectionexpression
	//
	// Deprecated: moved to package aws-apigatewayv2.
	RouteSelectionExpression *string `field:"optional" json:"routeSelectionExpression" yaml:"routeSelectionExpression"`
	// `AWS::ApiGatewayV2::Api.Tags`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-tags
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::ApiGatewayV2::Api.Target`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-target
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// `AWS::ApiGatewayV2::Api.Version`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-version
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

