package awsapigateway


// Example:
//   var api restApi
//
//
//   // We define the JSON Schema for the transformed valid response
//   responseModel := api.AddModel(jsii.String("ResponseModel"), &ModelOptions{
//   	ContentType: jsii.String("application/json"),
//   	ModelName: jsii.String("ResponseModel"),
//   	Schema: &JsonSchema{
//   		Schema: apigateway.JsonSchemaVersion_DRAFT4,
//   		Title: jsii.String("pollResponse"),
//   		Type: apigateway.JsonSchemaType_OBJECT,
//   		Properties: map[string]jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.JsonSchemaType_STRING,
//   			},
//   			"greeting": &jsonSchema{
//   				"type": apigateway.JsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
//   // We define the JSON Schema for the transformed error response
//   errorResponseModel := api.AddModel(jsii.String("ErrorResponseModel"), &ModelOptions{
//   	ContentType: jsii.String("application/json"),
//   	ModelName: jsii.String("ErrorResponseModel"),
//   	Schema: &jsonSchema{
//   		Schema: apigateway.JsonSchemaVersion_DRAFT4,
//   		Title: jsii.String("errorResponse"),
//   		Type: apigateway.JsonSchemaType_OBJECT,
//   		Properties: map[string]*jsonSchema{
//   			"state": &jsonSchema{
//   				"type": apigateway.JsonSchemaType_STRING,
//   			},
//   			"message": &jsonSchema{
//   				"type": apigateway.JsonSchemaType_STRING,
//   			},
//   		},
//   	},
//   })
//
type ModelOptions struct {
	// The schema to use to transform data to one or more output formats.
	//
	// Specify null ({}) if you don't want to specify a schema.
	Schema *JsonSchema `field:"required" json:"schema" yaml:"schema"`
	// The content type for the model.
	//
	// You can also force a
	// content type in the request or response model mapping.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A description that identifies this model.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the model.
	//
	// Important
	//  If you specify a name, you cannot perform updates that
	//  require replacement of this resource. You can perform
	//  updates that require no or some interruption. If you
	//  must replace the resource, specify a new name.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
}

