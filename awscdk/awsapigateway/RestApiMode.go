package awsapigateway


// The Mode that determines how API Gateway handles resource updates when importing an OpenAPI definition.
//
// Example:
//   api := apigateway.NewSpecRestApi(this, jsii.String("books-api"), &SpecRestApiProps{
//   	ApiDefinition: apigateway.ApiDefinition_FromAsset(jsii.String("path-to-file.json")),
//   	Mode: apigateway.RestApiMode_MERGE,
//   })
//
type RestApiMode string

const (
	// The new API definition replaces the existing one.
	RestApiMode_OVERWRITE RestApiMode = "OVERWRITE"
	// The new API definition is merged with the existing API.
	RestApiMode_MERGE RestApiMode = "MERGE"
)

