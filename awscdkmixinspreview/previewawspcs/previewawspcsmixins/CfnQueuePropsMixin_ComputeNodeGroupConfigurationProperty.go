package previewawspcsmixins


// The compute node group configuration for a queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   computeNodeGroupConfigurationProperty := &ComputeNodeGroupConfigurationProperty{
//   	ComputeNodeGroupId: jsii.String("computeNodeGroupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-queue-computenodegroupconfiguration.html
//
type CfnQueuePropsMixin_ComputeNodeGroupConfigurationProperty struct {
	// The compute node group ID for the compute node group configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-queue-computenodegroupconfiguration.html#cfn-pcs-queue-computenodegroupconfiguration-computenodegroupid
	//
	ComputeNodeGroupId *string `field:"optional" json:"computeNodeGroupId" yaml:"computeNodeGroupId"`
}

