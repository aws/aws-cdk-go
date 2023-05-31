package awsapigateway


// Properties for defining a `AWS::ApiGatewayV2::Authorizer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAuthorizerV2Props := &CfnAuthorizerV2Props{
//   	ApiId: jsii.String("apiId"),
//   	AuthorizerType: jsii.String("authorizerType"),
//   	IdentitySource: []*string{
//   		jsii.String("identitySource"),
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AuthorizerCredentialsArn: jsii.String("authorizerCredentialsArn"),
//   	AuthorizerResultTtlInSeconds: jsii.Number(123),
//   	AuthorizerUri: jsii.String("authorizerUri"),
//   	IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   	JwtConfiguration: &JWTConfigurationProperty{
//   		Audience: []*string{
//   			jsii.String("audience"),
//   		},
//   		Issuer: jsii.String("issuer"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnAuthorizerV2Props struct {
	// `AWS::ApiGatewayV2::Authorizer.ApiId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-apiid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// `AWS::ApiGatewayV2::Authorizer.AuthorizerType`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-authorizertype
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AuthorizerType *string `field:"required" json:"authorizerType" yaml:"authorizerType"`
	// `AWS::ApiGatewayV2::Authorizer.IdentitySource`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-identitysource
	//
	// Deprecated: moved to package aws-apigatewayv2.
	IdentitySource *[]*string `field:"required" json:"identitySource" yaml:"identitySource"`
	// `AWS::ApiGatewayV2::Authorizer.Name`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-name
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::ApiGatewayV2::Authorizer.AuthorizerCredentialsArn`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-authorizercredentialsarn
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AuthorizerCredentialsArn *string `field:"optional" json:"authorizerCredentialsArn" yaml:"authorizerCredentialsArn"`
	// `AWS::ApiGatewayV2::Authorizer.AuthorizerResultTtlInSeconds`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-authorizerresultttlinseconds
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// `AWS::ApiGatewayV2::Authorizer.AuthorizerUri`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-authorizeruri
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
	// `AWS::ApiGatewayV2::Authorizer.IdentityValidationExpression`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-identityvalidationexpression
	//
	// Deprecated: moved to package aws-apigatewayv2.
	IdentityValidationExpression *string `field:"optional" json:"identityValidationExpression" yaml:"identityValidationExpression"`
	// `AWS::ApiGatewayV2::Authorizer.JwtConfiguration`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-jwtconfiguration
	//
	// Deprecated: moved to package aws-apigatewayv2.
	JwtConfiguration interface{} `field:"optional" json:"jwtConfiguration" yaml:"jwtConfiguration"`
}

