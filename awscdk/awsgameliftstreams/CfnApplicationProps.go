package awsgameliftstreams


// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	ApplicationSourceUri: jsii.String("applicationSourceUri"),
//   	Description: jsii.String("description"),
//   	ExecutablePath: jsii.String("executablePath"),
//   	RuntimeEnvironment: &RuntimeEnvironmentProperty{
//   		Type: jsii.String("type"),
//   		Version: jsii.String("version"),
//   	},
//
//   	// the properties below are optional
//   	ApplicationLogOutputUri: jsii.String("applicationLogOutputUri"),
//   	ApplicationLogPaths: []*string{
//   		jsii.String("applicationLogPaths"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html
//
type CfnApplicationProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-applicationsourceuri
	//
	ApplicationSourceUri *string `field:"required" json:"applicationSourceUri" yaml:"applicationSourceUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-executablepath
	//
	ExecutablePath *string `field:"required" json:"executablePath" yaml:"executablePath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-runtimeenvironment
	//
	RuntimeEnvironment interface{} `field:"required" json:"runtimeEnvironment" yaml:"runtimeEnvironment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-applicationlogoutputuri
	//
	ApplicationLogOutputUri *string `field:"optional" json:"applicationLogOutputUri" yaml:"applicationLogOutputUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-applicationlogpaths
	//
	ApplicationLogPaths *[]*string `field:"optional" json:"applicationLogPaths" yaml:"applicationLogPaths"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

