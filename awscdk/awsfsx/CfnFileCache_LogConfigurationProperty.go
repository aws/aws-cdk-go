package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	Level: jsii.String("level"),
//
//   	// the properties below are optional
//   	Destination: jsii.String("destination"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-logconfiguration.html
//
type CfnFileCache_LogConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-logconfiguration.html#cfn-fsx-filecache-logconfiguration-level
	//
	Level *string `field:"required" json:"level" yaml:"level"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-logconfiguration.html#cfn-fsx-filecache-logconfiguration-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

