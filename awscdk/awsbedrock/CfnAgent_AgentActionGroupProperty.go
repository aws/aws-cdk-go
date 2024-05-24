package awsbedrock


// Contains details about an action group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentActionGroupProperty := &AgentActionGroupProperty{
//   	ActionGroupName: jsii.String("actionGroupName"),
//
//   	// the properties below are optional
//   	ActionGroupExecutor: &ActionGroupExecutorProperty{
//   		CustomControl: jsii.String("customControl"),
//   		Lambda: jsii.String("lambda"),
//   	},
//   	ActionGroupState: jsii.String("actionGroupState"),
//   	ApiSchema: &APISchemaProperty{
//   		Payload: jsii.String("payload"),
//   		S3: &S3IdentifierProperty{
//   			S3BucketName: jsii.String("s3BucketName"),
//   			S3ObjectKey: jsii.String("s3ObjectKey"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	FunctionSchema: &FunctionSchemaProperty{
//   		Functions: []interface{}{
//   			&FunctionProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   				Parameters: map[string]interface{}{
//   					"parametersKey": &ParameterDetailProperty{
//   						"type": jsii.String("type"),
//
//   						// the properties below are optional
//   						"description": jsii.String("description"),
//   						"required": jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ParentActionGroupSignature: jsii.String("parentActionGroupSignature"),
//   	SkipResourceInUseCheckOnDelete: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentactiongroup.html
//
type CfnAgent_AgentActionGroupProperty struct {
	// The name of the action group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentactiongroup.html#cfn-bedrock-agent-agentactiongroup-actiongroupname
	//
	ActionGroupName *string `field:"required" json:"actionGroupName" yaml:"actionGroupName"`
	// The Amazon Resource Name (ARN) of the Lambda function containing the business logic that is carried out upon invoking the action or the custom control method for handling the information elicited from the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentactiongroup.html#cfn-bedrock-agent-agentactiongroup-actiongroupexecutor
	//
	ActionGroupExecutor interface{} `field:"optional" json:"actionGroupExecutor" yaml:"actionGroupExecutor"`
	// Specifies whether the action group is available for the agent to invoke or not when sending an [InvokeAgent](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html) request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentactiongroup.html#cfn-bedrock-agent-agentactiongroup-actiongroupstate
	//
	ActionGroupState *string `field:"optional" json:"actionGroupState" yaml:"actionGroupState"`
	// Contains either details about the S3 object containing the OpenAPI schema for the action group or the JSON or YAML-formatted payload defining the schema.
	//
	// For more information, see [Action group OpenAPI schemas](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-api-schema.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentactiongroup.html#cfn-bedrock-agent-agentactiongroup-apischema
	//
	ApiSchema interface{} `field:"optional" json:"apiSchema" yaml:"apiSchema"`
	// The description of the action group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentactiongroup.html#cfn-bedrock-agent-agentactiongroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Defines functions that each define parameters that the agent needs to invoke from the user.
	//
	// Each function represents an action in an action group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentactiongroup.html#cfn-bedrock-agent-agentactiongroup-functionschema
	//
	FunctionSchema interface{} `field:"optional" json:"functionSchema" yaml:"functionSchema"`
	// If this field is set as `AMAZON.UserInput` , the agent can request the user for additional information when trying to complete a task. The `description` , `apiSchema` , and `actionGroupExecutor` fields must be blank for this action group.
	//
	// During orchestration, if the agent determines that it needs to invoke an API in an action group, but doesn't have enough information to complete the API request, it will invoke this action group instead and return an [Observation](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_Observation.html) reprompting the user for more information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentactiongroup.html#cfn-bedrock-agent-agentactiongroup-parentactiongroupsignature
	//
	ParentActionGroupSignature *string `field:"optional" json:"parentActionGroupSignature" yaml:"parentActionGroupSignature"`
	// Specifies whether to delete the resource even if it's in use.
	//
	// By default, this value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-agentactiongroup.html#cfn-bedrock-agent-agentactiongroup-skipresourceinusecheckondelete
	//
	// Default: - false.
	//
	SkipResourceInUseCheckOnDelete interface{} `field:"optional" json:"skipResourceInUseCheckOnDelete" yaml:"skipResourceInUseCheckOnDelete"`
}

