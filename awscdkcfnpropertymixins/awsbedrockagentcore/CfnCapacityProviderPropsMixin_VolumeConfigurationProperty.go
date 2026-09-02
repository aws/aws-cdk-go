package awsbedrockagentcore


// Configuration for a persistent volume attached to a capacity provider session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   volumeConfigurationProperty := &VolumeConfigurationProperty{
//   	EbsConfiguration: &EbsVolumeConfigurationProperty{
//   		Encrypted: jsii.Boolean(false),
//   		Iops: jsii.Number(123),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		Name: jsii.String("name"),
//   		SizeGiB: jsii.Number(123),
//   		SnapshotId: jsii.String("snapshotId"),
//   		Throughput: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-volumeconfiguration.html
//
type CfnCapacityProviderPropsMixin_VolumeConfigurationProperty struct {
	// Configuration for an EBS-backed persistent volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-volumeconfiguration.html#cfn-bedrockagentcore-capacityprovider-volumeconfiguration-ebsconfiguration
	//
	EbsConfiguration interface{} `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
}

