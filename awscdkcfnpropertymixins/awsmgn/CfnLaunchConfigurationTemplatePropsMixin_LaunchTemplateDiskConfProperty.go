package awsmgn


// Launch template disk configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   launchTemplateDiskConfProperty := &LaunchTemplateDiskConfProperty{
//   	Iops: jsii.Number(123),
//   	Throughput: jsii.Number(123),
//   	VolumeType: jsii.String("volumeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-launchtemplatediskconf.html
//
type CfnLaunchConfigurationTemplatePropsMixin_LaunchTemplateDiskConfProperty struct {
	// Launch template disk IOPS configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-launchtemplatediskconf.html#cfn-mgn-launchconfigurationtemplate-launchtemplatediskconf-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Launch template disk throughput configuration, in MiB/s.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-launchtemplatediskconf.html#cfn-mgn-launchconfigurationtemplate-launchtemplatediskconf-throughput
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Launch template disk volume type configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-launchtemplatediskconf.html#cfn-mgn-launchconfigurationtemplate-launchtemplatediskconf-volumetype
	//
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

