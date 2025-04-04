package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataConfigurationProperty := &MetadataConfigurationProperty{
//   	Iops: jsii.Number(123),
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-metadataconfiguration.html
//
type CfnFileSystem_MetadataConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-metadataconfiguration.html#cfn-fsx-filesystem-metadataconfiguration-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-metadataconfiguration.html#cfn-fsx-filesystem-metadataconfiguration-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

