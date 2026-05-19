package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

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
//   				Type: agentcore.SchemaDefinitionType_OBJECT(),
//   				Properties: map[string]SchemaDefinition{
//   					"name": &SchemaDefinition{
//   						"type": agentcore.SchemaDefinitionType_STRING(),
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
type SchemaDefinitionType interface {
	// The type string value.
	Value() *string
	// Returns the string value.
	ToString() *string
}

// The jsii proxy struct for SchemaDefinitionType
type jsiiProxy_SchemaDefinitionType struct {
	_ byte // padding
}

func (j *jsiiProxy_SchemaDefinitionType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom schema definition type not yet defined in this class.
func SchemaDefinitionType_Of(value *string) SchemaDefinitionType {
	_init_.Initialize()

	if err := validateSchemaDefinitionType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns SchemaDefinitionType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.SchemaDefinitionType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func SchemaDefinitionType_ARRAY() SchemaDefinitionType {
	_init_.Initialize()
	var returns SchemaDefinitionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.SchemaDefinitionType",
		"ARRAY",
		&returns,
	)
	return returns
}

func SchemaDefinitionType_BOOLEAN() SchemaDefinitionType {
	_init_.Initialize()
	var returns SchemaDefinitionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.SchemaDefinitionType",
		"BOOLEAN",
		&returns,
	)
	return returns
}

func SchemaDefinitionType_INTEGER() SchemaDefinitionType {
	_init_.Initialize()
	var returns SchemaDefinitionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.SchemaDefinitionType",
		"INTEGER",
		&returns,
	)
	return returns
}

func SchemaDefinitionType_NUMBER() SchemaDefinitionType {
	_init_.Initialize()
	var returns SchemaDefinitionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.SchemaDefinitionType",
		"NUMBER",
		&returns,
	)
	return returns
}

func SchemaDefinitionType_OBJECT() SchemaDefinitionType {
	_init_.Initialize()
	var returns SchemaDefinitionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.SchemaDefinitionType",
		"OBJECT",
		&returns,
	)
	return returns
}

func SchemaDefinitionType_STRING() SchemaDefinitionType {
	_init_.Initialize()
	var returns SchemaDefinitionType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.SchemaDefinitionType",
		"STRING",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SchemaDefinitionType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

