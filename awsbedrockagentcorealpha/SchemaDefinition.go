package awsbedrockagentcorealpha


// Schema definition for tool input/output.
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
type SchemaDefinition struct {
	// The type of the schema definition.
	//
	// This field specifies the data type of the schema.
	// Experimental.
	Type SchemaDefinitionType `field:"required" json:"type" yaml:"type"`
	// The description of the schema definition.
	//
	// This description provides information about the purpose and usage of the schema.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The items in the schema definition.
	//
	// This field is used for array types to define the structure of the array elements.
	// Default: - No items definition.
	//
	// Experimental.
	Items **SchemaDefinition `field:"optional" json:"items" yaml:"items"`
	// The properties of the schema definition.
	//
	// These properties define the fields in the schema.
	// Default: - No properties.
	//
	// Experimental.
	Properties *map[string]*SchemaDefinition `field:"optional" json:"properties" yaml:"properties"`
	// The required fields in the schema definition.
	//
	// These fields must be provided when using the schema.
	// Default: - No required fields.
	//
	// Experimental.
	Required *[]*string `field:"optional" json:"required" yaml:"required"`
}

