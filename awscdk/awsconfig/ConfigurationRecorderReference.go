package awsconfig


// A reference to a ConfigurationRecorder resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationRecorderReference := &ConfigurationRecorderReference{
//   	ConfigurationRecorderId: jsii.String("configurationRecorderId"),
//   }
//
type ConfigurationRecorderReference struct {
	// The Id of the ConfigurationRecorder resource.
	ConfigurationRecorderId *string `field:"required" json:"configurationRecorderId" yaml:"configurationRecorderId"`
}

