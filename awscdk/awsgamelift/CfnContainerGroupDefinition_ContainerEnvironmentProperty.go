package awsgamelift


// *This data type is used with the Amazon GameLift containers feature, which is currently in public preview.*.
//
// An environment variable to set inside a container, in the form of a key-value pair.
//
// *Related data type:* `ContainerDefinition$Environment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerEnvironmentProperty := &ContainerEnvironmentProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerenvironment.html
//
type CfnContainerGroupDefinition_ContainerEnvironmentProperty struct {
	// The environment variable name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerenvironment.html#cfn-gamelift-containergroupdefinition-containerenvironment-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The environment variable value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerenvironment.html#cfn-gamelift-containergroupdefinition-containerenvironment-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

