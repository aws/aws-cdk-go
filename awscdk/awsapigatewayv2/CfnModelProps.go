package awsapigatewayv2


// Properties for defining a `CfnModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var schema interface{}
//
//   cfnModelProps := &CfnModelProps{
//   	ApiId: jsii.String("apiId"),
//   	Name: jsii.String("name"),
//   	Schema: schema,
//
//   	// the properties below are optional
//   	ContentType: jsii.String("contentType"),
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html
//
type CfnModelProps struct {
	// The API identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-apiid
	//
	ApiId interface{} `field:"required" json:"apiId" yaml:"apiId"`
	// The name of the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema for the model.
	//
	// For application/json models, this should be JSON schema draft 4 model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-schema
	//
	Schema interface{} `field:"required" json:"schema" yaml:"schema"`
	// The content-type for the model, for example, "application/json".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The description of the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

