package awsapigateway


// Properties for defining a `AWS::ApiGatewayV2::Model`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var schema interface{}
//
//   cfnModelV2Props := &cfnModelV2Props{
//   	apiId: jsii.String("apiId"),
//   	name: jsii.String("name"),
//   	schema: schema,
//
//   	// the properties below are optional
//   	contentType: jsii.String("contentType"),
//   	description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnModelV2Props struct {
	// `AWS::ApiGatewayV2::Model.ApiId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-apiid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// `AWS::ApiGatewayV2::Model.Name`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-name
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::ApiGatewayV2::Model.Schema`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-schema
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Schema interface{} `field:"required" json:"schema" yaml:"schema"`
	// `AWS::ApiGatewayV2::Model.ContentType`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-contenttype
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// `AWS::ApiGatewayV2::Model.Description`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-description
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

