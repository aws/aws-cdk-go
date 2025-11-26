package previewawsapigatewayv2mixins


// Properties for CfnModelPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var schema interface{}
//
//   cfnModelMixinProps := &CfnModelMixinProps{
//   	ApiId: jsii.String("apiId"),
//   	ContentType: jsii.String("contentType"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Schema: schema,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html
//
type CfnModelMixinProps struct {
	// The API identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-apiid
	//
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// The content-type for the model, for example, "application/json".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The description of the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The schema for the model.
	//
	// For application/json models, this should be JSON schema draft 4 model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html#cfn-apigatewayv2-model-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

