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
//   			"properties": &CloudWatchEventEventProperty{
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
type CfnStateMachineProps struct {
	// `AWS::Serverless::StateMachine.Definition`.
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// `AWS::Serverless::StateMachine.DefinitionSubstitutions`.
	DefinitionSubstitutions interface{} `field:"optional" json:"definitionSubstitutions" yaml:"definitionSubstitutions"`
	// `AWS::Serverless::StateMachine.DefinitionUri`.
	DefinitionUri interface{} `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// `AWS::Serverless::StateMachine.Events`.
	Events interface{} `field:"optional" json:"events" yaml:"events"`
	// `AWS::Serverless::StateMachine.Logging`.
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// `AWS::Serverless::StateMachine.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Serverless::StateMachine.PermissionsBoundaries`.
	PermissionsBoundaries *string `field:"optional" json:"permissionsBoundaries" yaml:"permissionsBoundaries"`
	// `AWS::Serverless::StateMachine.Policies`.
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// `AWS::Serverless::StateMachine.Role`.
	Role *string `field:"optional" json:"role" yaml:"role"`
	// `AWS::Serverless::StateMachine.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::Serverless::StateMachine.Tracing`.
	Tracing interface{} `field:"optional" json:"tracing" yaml:"tracing"`
	// `AWS::Serverless::StateMachine.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

