package awsbedrockagentcorealpha


// Schema definition types.
//
// Example:
//   // Create a gateway first
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_22_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String(`
//   	    exports.handler = async (event) => {
//   	      return {
//   	        statusCode: 200,
//   	        body: JSON.stringify({ message: 'Hello from Lambda!' })
//   	      };
//   	    };
//   	  `)),
//   })
//
//   lambdaTarget := gateway.AddLambdaTarget(jsii.String("MyLambdaTarget"), &AddLambdaTargetOptions{
//   	GatewayTargetName: jsii.String("my-lambda-target"),
//   	Description: jsii.String("Lambda function target"),
//   	LambdaFunction: lambdaFunction,
//   	ToolSchema: agentcore.ToolSchema_FromInline([]ToolDefinition{
//   		&ToolDefinition{
//   			Name: jsii.String("hello_world"),
//   			Description: jsii.String("A simple hello world tool"),
//   			InputSchema: &SchemaDefinition{
//   				Type: agentcore.SchemaDefinitionType_OBJECT,
//   				Properties: map[string]SchemaDefinition{
//   					"name": &SchemaDefinition{
//   						"type": agentcore.SchemaDefinitionType_STRING,
//   						"description": jsii.String("The name to greet"),
//   					},
//   				},
//   				Required: []*string{
//   					jsii.String("name"),
//   				},
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type SchemaDefinitionType string

const (
	// String type.
	// Experimental.
	SchemaDefinitionType_STRING SchemaDefinitionType = "STRING"
	// Number type.
	// Experimental.
	SchemaDefinitionType_NUMBER SchemaDefinitionType = "NUMBER"
	// Object type.
	// Experimental.
	SchemaDefinitionType_OBJECT SchemaDefinitionType = "OBJECT"
	// Array type.
	// Experimental.
	SchemaDefinitionType_ARRAY SchemaDefinitionType = "ARRAY"
	// Boolean type.
	// Experimental.
	SchemaDefinitionType_BOOLEAN SchemaDefinitionType = "BOOLEAN"
	// Integer type.
	// Experimental.
	SchemaDefinitionType_INTEGER SchemaDefinitionType = "INTEGER"
)

