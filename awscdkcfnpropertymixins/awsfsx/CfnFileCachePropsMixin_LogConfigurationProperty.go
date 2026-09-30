package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	Destination: jsii.String("destination"),
//   	Level: jsii.String("level"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-logconfiguration.html
//
type CfnFileCachePropsMixin_LogConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-logconfiguration.html#cfn-fsx-filecache-logconfiguration-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-logconfiguration.html#cfn-fsx-filecache-logconfiguration-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
}

