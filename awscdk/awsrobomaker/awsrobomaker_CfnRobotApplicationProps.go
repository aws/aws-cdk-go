package awsrobomaker


// Properties for defining a `CfnRobotApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRobotApplicationProps := &cfnRobotApplicationProps{
//   	robotSoftwareSuite: &robotSoftwareSuiteProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		version: jsii.String("version"),
//   	},
//
//   	// the properties below are optional
//   	currentRevisionId: jsii.String("currentRevisionId"),
//   	environment: jsii.String("environment"),
//   	name: jsii.String("name"),
//   	sources: []interface{}{
//   		&sourceConfigProperty{
//   			architecture: jsii.String("architecture"),
//   			s3Bucket: jsii.String("s3Bucket"),
//   			s3Key: jsii.String("s3Key"),
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnRobotApplicationProps struct {
	// The robot software suite used by the robot application.
	RobotSoftwareSuite interface{} `field:"required" json:"robotSoftwareSuite" yaml:"robotSoftwareSuite"`
	// The current revision id.
	CurrentRevisionId *string `field:"optional" json:"currentRevisionId" yaml:"currentRevisionId"`
	// The environment of the robot application.
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The name of the robot application.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sources of the robot application.
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// A map that contains tag keys and tag values that are attached to the robot application.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

