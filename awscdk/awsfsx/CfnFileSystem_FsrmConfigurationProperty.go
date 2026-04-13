package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fsrmConfigurationProperty := &FsrmConfigurationProperty{
//   	FsrmServiceEnabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	EventLogDestination: jsii.String("eventLogDestination"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-fsrmconfiguration.html
//
type CfnFileSystem_FsrmConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-fsrmconfiguration.html#cfn-fsx-filesystem-fsrmconfiguration-fsrmserviceenabled
	//
	FsrmServiceEnabled interface{} `field:"required" json:"fsrmServiceEnabled" yaml:"fsrmServiceEnabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-fsrmconfiguration.html#cfn-fsx-filesystem-fsrmconfiguration-eventlogdestination
	//
	EventLogDestination *string `field:"optional" json:"eventLogDestination" yaml:"eventLogDestination"`
}

