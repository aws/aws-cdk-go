package mixinsawsbedrock


// Contains details about the function schema for the action group or the JSON or YAML-formatted payload defining the schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   functionSchemaProperty := &FunctionSchemaProperty{
//   	Functions: []interface{}{
//   		&FunctionProperty{
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   			Parameters: map[string]interface{}{
//   				"parametersKey": &ParameterDetailProperty{
//   					"description": jsii.String("description"),
//   					"required": jsii.Boolean(false),
//   					"type": jsii.String("type"),
//   				},
//   			},
//   			RequireConfirmation: jsii.String("requireConfirmation"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-functionschema.html
//
type CfnAgentPropsMixin_FunctionSchemaProperty struct {
	// A list of functions that each define an action in the action group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-functionschema.html#cfn-bedrock-agent-functionschema-functions
	//
	Functions interface{} `field:"optional" json:"functions" yaml:"functions"`
}

