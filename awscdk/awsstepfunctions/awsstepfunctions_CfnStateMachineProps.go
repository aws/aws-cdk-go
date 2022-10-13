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
//   cfnStateMachineProps := &cfnStateMachineProps{
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	definition: definition,
//   	definitionS3Location: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//
//   		// the properties below are optional
//   		version: jsii.String("version"),
//   	},
//   	definitionString: jsii.String("definitionString"),
//   	definitionSubstitutions: map[string]*string{
//   		"definitionSubstitutionsKey": jsii.String("definitionSubstitutions"),
//   	},
//   	loggingConfiguration: &loggingConfigurationProperty{
//   		destinations: []interface{}{
//   			&logDestinationProperty{
//   				cloudWatchLogsLogGroup: &cloudWatchLogsLogGroupProperty{
//   					logGroupArn: jsii.String("logGroupArn"),
//   				},
//   			},
//   		},
//   		includeExecutionData: jsii.Boolean(false),
//   		level: jsii.String("level"),
//   	},
//   	stateMachineName: jsii.String("stateMachineName"),
//   	stateMachineType: jsii.String("stateMachineType"),
//   	tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tracingConfiguration: &tracingConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnStateMachineProps struct {
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon States Language definition of the state machine.
	//
	// The state machine definition must be in JSON or YAML, and the format of the object must match the format of your AWS Step Functions template file. See [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) .
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// The name of the S3 bucket where the state machine definition is stored.
	//
	// The state machine definition must be a JSON or YAML file.
	DefinitionS3Location interface{} `field:"optional" json:"definitionS3Location" yaml:"definitionS3Location"`
	// The Amazon States Language definition of the state machine.
	//
	// The state machine definition must be in JSON. See [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) .
	DefinitionString *string `field:"optional" json:"definitionString" yaml:"definitionString"`
	// A map (string to string) that specifies the mappings for placeholder variables in the state machine definition.
	//
	// This enables the customer to inject values obtained at runtime, for example from intrinsic functions, in the state machine definition. Variables can be template parameter names, resource logical IDs, resource attributes, or a variable in a key-value map.
	DefinitionSubstitutions interface{} `field:"optional" json:"definitionSubstitutions" yaml:"definitionSubstitutions"`
	// Defines what execution history events are logged and where they are logged.
	//
	// > By default, the `level` is set to `OFF` . For more information see [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
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
	StateMachineName *string `field:"optional" json:"stateMachineName" yaml:"stateMachineName"`
	// Determines whether a `STANDARD` or `EXPRESS` state machine is created.
	//
	// The default is `STANDARD` . You cannot update the `type` of a state machine once it has been created. For more information on `STANDARD` and `EXPRESS` workflows, see [Standard Versus Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-standard-vs-express.html) in the AWS Step Functions Developer Guide.
	StateMachineType *string `field:"optional" json:"stateMachineType" yaml:"stateMachineType"`
	// The list of tags to add to a resource.
	//
	// Tags may only contain Unicode letters, digits, white space, or these symbols: `_ . : / = + - @` .
	Tags *[]*CfnStateMachine_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
	// Selects whether or not the state machine's AWS X-Ray tracing is enabled.
	TracingConfiguration interface{} `field:"optional" json:"tracingConfiguration" yaml:"tracingConfiguration"`
}

