package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataReadCacheConfigurationProperty := &DataReadCacheConfigurationProperty{
//   	SizeGiB: jsii.Number(123),
//   	SizingMode: jsii.String("sizingMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-datareadcacheconfiguration.html
//
type CfnFileSystem_DataReadCacheConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-datareadcacheconfiguration.html#cfn-fsx-filesystem-datareadcacheconfiguration-sizegib
	//
	SizeGiB *float64 `field:"optional" json:"sizeGiB" yaml:"sizeGiB"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-datareadcacheconfiguration.html#cfn-fsx-filesystem-datareadcacheconfiguration-sizingmode
	//
	SizingMode *string `field:"optional" json:"sizingMode" yaml:"sizingMode"`
}

