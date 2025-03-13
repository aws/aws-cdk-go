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
type JsonSchemaVersion string

const (
	// In API Gateway models are defined using the JSON schema draft 4.
	// See: https://tools.ietf.org/html/draft-zyp-json-schema-04
	//
	JsonSchemaVersion_DRAFT4 JsonSchemaVersion = "DRAFT4"
	JsonSchemaVersion_DRAFT7 JsonSchemaVersion = "DRAFT7"
)

