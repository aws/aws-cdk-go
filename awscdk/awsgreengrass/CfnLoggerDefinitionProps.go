package awsgreengrass


// Properties for defining a `CfnLoggerDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnLoggerDefinitionProps := &CfnLoggerDefinitionProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	InitialVersion: &LoggerDefinitionVersionProperty{
//   		Loggers: []interface{}{
//   			&LoggerProperty{
//   				Component: jsii.String("component"),
//   				Id: jsii.String("id"),
//   				Level: jsii.String("level"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Space: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinition.html
//
type CfnLoggerDefinitionProps struct {
	// The name of the logger definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinition.html#cfn-greengrass-loggerdefinition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The logger definition version to include when the logger definition is created.
	//
	// A logger definition version contains a list of [`logger`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-loggerdefinition-logger.html) property types.
	//
	// > To associate a logger definition version after the logger definition is created, create an [`AWS::Greengrass::LoggerDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html) resource and specify the ID of this logger definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinition.html#cfn-greengrass-loggerdefinition-initialversion
	//
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the logger definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/v1/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinition.html#cfn-greengrass-loggerdefinition-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

