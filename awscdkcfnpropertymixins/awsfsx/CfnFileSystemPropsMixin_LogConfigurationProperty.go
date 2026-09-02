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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-logconfiguration.html
//
type CfnFileSystemPropsMixin_LogConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-logconfiguration.html#cfn-fsx-filesystem-logconfiguration-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-logconfiguration.html#cfn-fsx-filesystem-logconfiguration-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
}

