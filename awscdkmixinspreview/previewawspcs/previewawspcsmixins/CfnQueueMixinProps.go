package previewawspcsmixins


// Properties for CfnQueuePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnQueueMixinProps := &CfnQueueMixinProps{
//   	ClusterId: jsii.String("clusterId"),
//   	ComputeNodeGroupConfigurations: []interface{}{
//   		&ComputeNodeGroupConfigurationProperty{
//   			ComputeNodeGroupId: jsii.String("computeNodeGroupId"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	SlurmConfiguration: &SlurmConfigurationProperty{
//   		SlurmCustomSettings: []interface{}{
//   			&SlurmCustomSettingProperty{
//   				ParameterName: jsii.String("parameterName"),
//   				ParameterValue: jsii.String("parameterValue"),
//   			},
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-queue.html
//
type CfnQueueMixinProps struct {
	// The ID of the cluster of the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-queue.html#cfn-pcs-queue-clusterid
	//
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// The list of compute node group configurations associated with the queue.
	//
	// Queues assign jobs to associated compute node groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-queue.html#cfn-pcs-queue-computenodegroupconfigurations
	//
	ComputeNodeGroupConfigurations interface{} `field:"optional" json:"computeNodeGroupConfigurations" yaml:"computeNodeGroupConfigurations"`
	// The name that identifies the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-queue.html#cfn-pcs-queue-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Additional options related to the Slurm scheduler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-queue.html#cfn-pcs-queue-slurmconfiguration
	//
	SlurmConfiguration interface{} `field:"optional" json:"slurmConfiguration" yaml:"slurmConfiguration"`
	// 1 or more tags added to the resource.
	//
	// Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcs-queue.html#cfn-pcs-queue-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

