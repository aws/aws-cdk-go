package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   fsrmConfigurationProperty := &FsrmConfigurationProperty{
//   	EventLogDestination: jsii.String("eventLogDestination"),
//   	FsrmServiceEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-fsrmconfiguration.html
//
type CfnFileSystemPropsMixin_FsrmConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-fsrmconfiguration.html#cfn-fsx-filesystem-fsrmconfiguration-eventlogdestination
	//
	EventLogDestination *string `field:"optional" json:"eventLogDestination" yaml:"eventLogDestination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-fsrmconfiguration.html#cfn-fsx-filesystem-fsrmconfiguration-fsrmserviceenabled
	//
	FsrmServiceEnabled interface{} `field:"optional" json:"fsrmServiceEnabled" yaml:"fsrmServiceEnabled"`
}

