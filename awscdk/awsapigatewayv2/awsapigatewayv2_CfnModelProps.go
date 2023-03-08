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
//   cfnModelProps := &cfnModelProps{
//   	apiId: jsii.String("apiId"),
//   	name: jsii.String("name"),
//   	schema: schema,
//
//   	// the properties below are optional
//   	contentType: jsii.String("contentType"),
//   	description: jsii.String("description"),
//   }
//
type CfnModelProps struct {
	// The API identifier.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The name of the model.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema for the model.
	//
	// For application/json models, this should be JSON schema draft 4 model.
	Schema interface{} `field:"required" json:"schema" yaml:"schema"`
	// The content-type for the model, for example, "application/json".
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The description of the model.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

