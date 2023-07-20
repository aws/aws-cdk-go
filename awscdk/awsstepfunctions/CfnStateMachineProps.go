package awsstepfunctions


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
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Definition: definition,
//   	DefinitionS3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//
//   		// the properties below are optional
//   		Version: jsii.String("version"),
//   	},
//   	DefinitionString: jsii.String("definitionString"),
//   	DefinitionSubstitutions: map[string]*string{
//   		"definitionSubstitutionsKey": jsii.String("definitionSubstitutions"),
//   	},
//   	LoggingConfiguration: &LoggingConfigurationProperty{
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
//   	StateMachineName: jsii.String("stateMachineName"),
//   	StateMachineType: jsii.String("stateMachineType"),
//   	Tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TracingConfiguration: &TracingConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html
//
type CfnStateMachineProps struct {
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon States Language definition of the state machine.
	//
	// The state machine definition must be in JSON or YAML, and the format of the object must match the format of your CloudFormation template file. See [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// The name of the S3 bucket where the state machine definition is stored.
	//
	// The state machine definition must be a JSON or YAML file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitions3location
	//
	DefinitionS3Location interface{} `field:"optional" json:"definitionS3Location" yaml:"definitionS3Location"`
	// The Amazon States Language definition of the state machine.
	//
	// The state machine definition must be in JSON. See [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionstring
	//
	DefinitionString *string `field:"optional" json:"definitionString" yaml:"definitionString"`
	// A map (string to string) that specifies the mappings for placeholder variables in the state machine definition.
	//
	// This enables the customer to inject values obtained at runtime, for example from intrinsic functions, in the state machine definition. Variables can be template parameter names, resource logical IDs, resource attributes, or a variable in a key-value map.
	//
	// Substitutions must follow the syntax: `${key_name}` or `${variable_1,variable_2,...}` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionsubstitutions
	//
	DefinitionSubstitutions interface{} `field:"optional" json:"definitionSubstitutions" yaml:"definitionSubstitutions"`
	// Defines what execution history events are logged and where they are logged.
	//
	// > By default, the `level` is set to `OFF` . For more information see [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-loggingconfiguration
	//
	LoggingConfiguration interface{} `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// The name of the state machine.
	//
	// A name must *not* contain:
	//
	// - white space
	// - brackets `< > { } [ ]`
	// - wildcard characters `? *`
	// - special characters `" # % \ ^ | ~ ` $ & , ; : /`
	// - control characters ( `U+0000-001F` , `U+007F-009F` )
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinename
	//
	StateMachineName *string `field:"optional" json:"stateMachineName" yaml:"stateMachineName"`
	// Determines whether a `STANDARD` or `EXPRESS` state machine is created.
	//
	// The default is `STANDARD` . You cannot update the `type` of a state machine once it has been created. For more information on `STANDARD` and `EXPRESS` workflows, see [Standard Versus Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-standard-vs-express.html) in the AWS Step Functions Developer Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinetype
	//
	StateMachineType *string `field:"optional" json:"stateMachineType" yaml:"stateMachineType"`
	// The list of tags to add to a resource.
	//
	// Tags may only contain Unicode letters, digits, white space, or these symbols: `_ . : / = + -
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tags
	//
	Tags *[]*CfnStateMachine_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
	// Selects whether or not the state machine's AWS X-Ray tracing is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tracingconfiguration
	//
	TracingConfiguration interface{} `field:"optional" json:"tracingConfiguration" yaml:"tracingConfiguration"`
}

