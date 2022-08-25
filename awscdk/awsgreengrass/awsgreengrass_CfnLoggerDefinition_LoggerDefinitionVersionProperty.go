package awsgreengrass


// A logger definition version contains a list of [loggers](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-loggerdefinition-logger.html) .
//
// > After you create a logger definition version that contains the loggers you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an AWS CloudFormation template, `LoggerDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::LoggerDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggerDefinitionVersionProperty := &loggerDefinitionVersionProperty{
//   	loggers: []interface{}{
//   		&loggerProperty{
//   			component: jsii.String("component"),
//   			id: jsii.String("id"),
//   			level: jsii.String("level"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			space: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnLoggerDefinition_LoggerDefinitionVersionProperty struct {
	// The loggers in this version.
	Loggers interface{} `field:"required" json:"loggers" yaml:"loggers"`
}

