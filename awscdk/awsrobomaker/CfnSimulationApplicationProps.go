package awsrobomaker


// Properties for defining a `CfnSimulationApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSimulationApplicationProps := &CfnSimulationApplicationProps{
//   	RobotSoftwareSuite: &RobotSoftwareSuiteProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Version: jsii.String("version"),
//   	},
//   	SimulationSoftwareSuite: &SimulationSoftwareSuiteProperty{
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
//   	RenderingEngine: &RenderingEngineProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
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
type CfnSimulationApplicationProps struct {
	// The robot software suite used by the simulation application.
	RobotSoftwareSuite interface{} `field:"required" json:"robotSoftwareSuite" yaml:"robotSoftwareSuite"`
	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite interface{} `field:"required" json:"simulationSoftwareSuite" yaml:"simulationSoftwareSuite"`
	// The current revision id.
	CurrentRevisionId *string `field:"optional" json:"currentRevisionId" yaml:"currentRevisionId"`
	// The environment of the simulation application.
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The name of the simulation application.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The rendering engine for the simulation application.
	RenderingEngine interface{} `field:"optional" json:"renderingEngine" yaml:"renderingEngine"`
	// The sources of the simulation application.
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// A map that contains tag keys and tag values that are attached to the simulation application.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

