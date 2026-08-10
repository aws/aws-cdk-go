package awssam


// Properties for defining a `CfnStateMachine`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var definition interface{}
//
//   cfnStateMachineProps := &CfnStateMachineProps{
//   	Definition: definition,
//   	DefinitionSubstitutions: map[string]*string{
//   		"definitionSubstitutionsKey": jsii.String("definitionSubstitutions"),
//   	},
//   	DefinitionUri: jsii.String("definitionUri"),
//   	Events: map[string]interface{}{
//   		"eventsKey": &EventSourceProperty{
//   			"properties": &ApiEventProperty{
//   				"method": jsii.String("method"),
//   				"path": jsii.String("path"),
//
//   				// the properties below are optional
//   				"restApiId": jsii.String("restApiId"),
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	Logging: &LoggingConfigurationProperty{
//   		Destinations: []interface{}{
//   			&LogDestinationProperty{
//   				CloudWatchLogsLogGroup: &CloudWatchLogsLogGroupProperty{
//   					LogGroupArn: jsii.String("logGroupArn"),
//   				},
//   			},
//   		},
//   		IncludeExecutionData: jsii.Boolean(false),
//   		Level: jsii.String("level"),
//   	},
//   	Name: jsii.String("name"),
//   	PermissionsBoundaries: jsii.String("permissionsBoundaries"),
//   	Policies: jsii.String("policies"),
//   	Role: jsii.String("role"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Tracing: &TracingConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html
//
type CfnStateMachineProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html#cfn-serverless-statemachine-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html#cfn-serverless-statemachine-definitionsubstitutions
	//
	DefinitionSubstitutions interface{} `field:"optional" json:"definitionSubstitutions" yaml:"definitionSubstitutions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html#cfn-serverless-statemachine-definitionuri
	//
	DefinitionUri interface{} `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html#cfn-serverless-statemachine-events
	//
	Events interface{} `field:"optional" json:"events" yaml:"events"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html#cfn-serverless-statemachine-logging
	//
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html#cfn-serverless-statemachine-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html#cfn-serverless-statemachine-permissionsboundaries
	//
	PermissionsBoundaries *string `field:"optional" json:"permissionsBoundaries" yaml:"permissionsBoundaries"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html#cfn-serverless-statemachine-policies
	//
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html#cfn-serverless-statemachine-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html#cfn-serverless-statemachine-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html#cfn-serverless-statemachine-tracing
	//
	Tracing interface{} `field:"optional" json:"tracing" yaml:"tracing"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-statemachine.html#cfn-serverless-statemachine-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

