package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   readCacheConfigurationProperty := &ReadCacheConfigurationProperty{
//   	SizeGiB: jsii.Number(123),
//   	SizingMode: jsii.String("sizingMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-readcacheconfiguration.html
//
type CfnFileSystem_ReadCacheConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-readcacheconfiguration.html#cfn-fsx-filesystem-readcacheconfiguration-sizegib
	//
	SizeGiB *float64 `field:"optional" json:"sizeGiB" yaml:"sizeGiB"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-readcacheconfiguration.html#cfn-fsx-filesystem-readcacheconfiguration-sizingmode
	//
	SizingMode *string `field:"optional" json:"sizingMode" yaml:"sizingMode"`
}

