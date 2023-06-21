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
type JsonSchemaType string

const (
	JsonSchemaType_NULL JsonSchemaType = "NULL"
	JsonSchemaType_BOOLEAN JsonSchemaType = "BOOLEAN"
	JsonSchemaType_OBJECT JsonSchemaType = "OBJECT"
	JsonSchemaType_ARRAY JsonSchemaType = "ARRAY"
	JsonSchemaType_NUMBER JsonSchemaType = "NUMBER"
	JsonSchemaType_INTEGER JsonSchemaType = "INTEGER"
	JsonSchemaType_STRING JsonSchemaType = "STRING"
)

