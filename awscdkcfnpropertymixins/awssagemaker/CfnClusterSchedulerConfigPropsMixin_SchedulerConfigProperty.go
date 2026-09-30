package awssagemaker


// Cluster policy configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   schedulerConfigProperty := &SchedulerConfigProperty{
//   	FairShare: jsii.String("fairShare"),
//   	IdleResourceSharing: jsii.String("idleResourceSharing"),
//   	PriorityClasses: []interface{}{
//   		&PriorityClassProperty{
//   			Name: jsii.String("name"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-clusterschedulerconfig-schedulerconfig.html
//
type CfnClusterSchedulerConfigPropsMixin_SchedulerConfigProperty struct {
	// When enabled, entities borrow idle compute based on assigned FairShareWeight.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-clusterschedulerconfig-schedulerconfig.html#cfn-sagemaker-clusterschedulerconfig-schedulerconfig-fairshare
	//
	FairShare *string `field:"optional" json:"fairShare" yaml:"fairShare"`
	// Configuration for sharing idle compute resources across entities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-clusterschedulerconfig-schedulerconfig.html#cfn-sagemaker-clusterschedulerconfig-schedulerconfig-idleresourcesharing
	//
	IdleResourceSharing *string `field:"optional" json:"idleResourceSharing" yaml:"idleResourceSharing"`
	// List of priority class configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-clusterschedulerconfig-schedulerconfig.html#cfn-sagemaker-clusterschedulerconfig-schedulerconfig-priorityclasses
	//
	PriorityClasses interface{} `field:"optional" json:"priorityClasses" yaml:"priorityClasses"`
}

