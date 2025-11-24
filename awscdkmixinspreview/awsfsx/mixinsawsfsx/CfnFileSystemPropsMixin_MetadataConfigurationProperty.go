package mixinsawsfsx


// The configuration that allows you to specify the performance of metadata operations for an FSx for Lustre file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metadataConfigurationProperty := &MetadataConfigurationProperty{
//   	Iops: jsii.Number(123),
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-metadataconfiguration.html
//
type CfnFileSystemPropsMixin_MetadataConfigurationProperty struct {
	// The number of Metadata IOPS provisioned for the file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-metadataconfiguration.html#cfn-fsx-filesystem-metadataconfiguration-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Specifies whether the file system is using the AUTOMATIC setting of metadata IOPS or if it is using a USER_PROVISIONED value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-metadataconfiguration.html#cfn-fsx-filesystem-metadataconfiguration-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

