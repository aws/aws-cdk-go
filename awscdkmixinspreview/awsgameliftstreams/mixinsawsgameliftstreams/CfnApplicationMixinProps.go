package mixinsawsgameliftstreams


// Properties for CfnApplicationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationMixinProps := &CfnApplicationMixinProps{
//   	ApplicationLogOutputUri: jsii.String("applicationLogOutputUri"),
//   	ApplicationLogPaths: []*string{
//   		jsii.String("applicationLogPaths"),
//   	},
//   	ApplicationSourceUri: jsii.String("applicationSourceUri"),
//   	Description: jsii.String("description"),
//   	ExecutablePath: jsii.String("executablePath"),
//   	RuntimeEnvironment: &RuntimeEnvironmentProperty{
//   		Type: jsii.String("type"),
//   		Version: jsii.String("version"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html
//
type CfnApplicationMixinProps struct {
	// An Amazon S3 URI to a bucket where you would like Amazon GameLift Streams to save application logs.
	//
	// Required if you specify one or more `ApplicationLogPaths` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-applicationlogoutputuri
	//
	ApplicationLogOutputUri *string `field:"optional" json:"applicationLogOutputUri" yaml:"applicationLogOutputUri"`
	// Locations of log files that your content generates during a stream session.
	//
	// Enter path values that are relative to the `ApplicationSourceUri` location. You can specify up to 10 log paths. Amazon GameLift Streams uploads designated log files to the Amazon S3 bucket that you specify in `ApplicationLogOutputUri` at the end of a stream session. To retrieve stored log files, call [GetStreamSession](https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetStreamSession.html) and get the `LogFileLocationUri` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-applicationlogpaths
	//
	ApplicationLogPaths *[]*string `field:"optional" json:"applicationLogPaths" yaml:"applicationLogPaths"`
	// The location of the content that you want to stream.
	//
	// Enter an Amazon S3 URI to a bucket that contains your game or other application. The location can have a multi-level prefix structure, but it must include all the files needed to run the content. Amazon GameLift Streams copies everything under the specified location.
	//
	// This value is immutable. To designate a different content location, create a new application.
	//
	// > The Amazon S3 bucket and the Amazon GameLift Streams application must be in the same AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-applicationsourceuri
	//
	ApplicationSourceUri *string `field:"optional" json:"applicationSourceUri" yaml:"applicationSourceUri"`
	// A human-readable label for the application.
	//
	// You can update this value later.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The relative path and file name of the executable file that Amazon GameLift Streams will stream.
	//
	// Specify a path relative to the location set in `ApplicationSourceUri` . The file must be contained within the application's root folder. For Windows applications, the file must be a valid Windows executable or batch file with a filename ending in .exe, .cmd, or .bat. For Linux applications, the file must be a valid Linux binary executable or a script that contains an initial interpreter line starting with a shebang (' `#!` ').
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-executablepath
	//
	ExecutablePath *string `field:"optional" json:"executablePath" yaml:"executablePath"`
	// A set of configuration settings to run the application on a stream group.
	//
	// This configures the operating system, and can include compatibility layers and other drivers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-runtimeenvironment
	//
	RuntimeEnvironment interface{} `field:"optional" json:"runtimeEnvironment" yaml:"runtimeEnvironment"`
	// A list of labels to assign to the new application resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-application.html#cfn-gameliftstreams-application-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

