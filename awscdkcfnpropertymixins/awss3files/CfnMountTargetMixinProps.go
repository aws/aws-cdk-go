package awss3files


// Properties for CfnMountTargetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnMountTargetMixinProps := &CfnMountTargetMixinProps{
//   	FileSystemId: jsii.String("fileSystemId"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	Ipv4Address: jsii.String("ipv4Address"),
//   	Ipv6Address: jsii.String("ipv6Address"),
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-mounttarget.html
//
type CfnMountTargetMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-mounttarget.html#cfn-s3files-mounttarget-filesystemid
	//
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-mounttarget.html#cfn-s3files-mounttarget-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-mounttarget.html#cfn-s3files-mounttarget-ipv4address
	//
	Ipv4Address *string `field:"optional" json:"ipv4Address" yaml:"ipv4Address"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-mounttarget.html#cfn-s3files-mounttarget-ipv6address
	//
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-mounttarget.html#cfn-s3files-mounttarget-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-mounttarget.html#cfn-s3files-mounttarget-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

