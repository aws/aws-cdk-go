package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnClusterSchedulerConfigPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnClusterSchedulerConfigMixinProps := &CfnClusterSchedulerConfigMixinProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	SchedulerConfig: &SchedulerConfigProperty{
//   		FairShare: jsii.String("fairShare"),
//   		IdleResourceSharing: jsii.String("idleResourceSharing"),
//   		PriorityClasses: []interface{}{
//   			&PriorityClassProperty{
//   				Name: jsii.String("name"),
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-clusterschedulerconfig.html
//
type CfnClusterSchedulerConfigMixinProps struct {
	// ARN of the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-clusterschedulerconfig.html#cfn-sagemaker-clusterschedulerconfig-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// Description of the cluster policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-clusterschedulerconfig.html#cfn-sagemaker-clusterschedulerconfig-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name for the cluster policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-clusterschedulerconfig.html#cfn-sagemaker-clusterschedulerconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Cluster policy configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-clusterschedulerconfig.html#cfn-sagemaker-clusterschedulerconfig-schedulerconfig
	//
	SchedulerConfig interface{} `field:"optional" json:"schedulerConfig" yaml:"schedulerConfig"`
	// Tags of the cluster policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-clusterschedulerconfig.html#cfn-sagemaker-clusterschedulerconfig-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

