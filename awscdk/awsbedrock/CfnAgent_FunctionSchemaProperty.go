package awsbedrock


// Contains details about the function schema for the action group or the JSON or YAML-formatted payload defining the schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionSchemaProperty := &FunctionSchemaProperty{
//   	Functions: []interface{}{
//   		&FunctionProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			Parameters: map[string]interface{}{
//   				"parametersKey": &ParameterDetailProperty{
//   					"type": jsii.String("type"),
//
//   					// the properties below are optional
//   					"description": jsii.String("description"),
//   					"required": jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-functionschema.html
//
type CfnAgent_FunctionSchemaProperty struct {
	// A list of functions that each define an action in the action group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-functionschema.html#cfn-bedrock-agent-functionschema-functions
	//
	Functions interface{} `field:"required" json:"functions" yaml:"functions"`
}

