package awsrobomaker


// Properties for defining a `CfnRobotApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRobotApplicationProps := &CfnRobotApplicationProps{
//   	RobotSoftwareSuite: &RobotSoftwareSuiteProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Version: jsii.String("version"),
//   	},
//
//   	// the properties below are optional
//   	CurrentRevisionId: jsii.String("currentRevisionId"),
//   	Environment: jsii.String("environment"),
//   	Name: jsii.String("name"),
//   	Sources: []interface{}{
//   		&SourceConfigProperty{
//   			Architecture: jsii.String("architecture"),
//   			S3Bucket: jsii.String("s3Bucket"),
//   			S3Key: jsii.String("s3Key"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplication.html
//
type CfnRobotApplicationProps struct {
	// The robot software suite used by the robot application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplication.html#cfn-robomaker-robotapplication-robotsoftwaresuite
	//
	RobotSoftwareSuite interface{} `field:"required" json:"robotSoftwareSuite" yaml:"robotSoftwareSuite"`
	// The current revision id.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplication.html#cfn-robomaker-robotapplication-currentrevisionid
	//
	CurrentRevisionId *string `field:"optional" json:"currentRevisionId" yaml:"currentRevisionId"`
	// The environment of the robot application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplication.html#cfn-robomaker-robotapplication-environment
	//
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The name of the robot application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplication.html#cfn-robomaker-robotapplication-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sources of the robot application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplication.html#cfn-robomaker-robotapplication-sources
	//
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// A map that contains tag keys and tag values that are attached to the robot application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplication.html#cfn-robomaker-robotapplication-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

