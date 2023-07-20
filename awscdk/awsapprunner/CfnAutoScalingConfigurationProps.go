package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAutoScalingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAutoScalingConfigurationProps := &CfnAutoScalingConfigurationProps{
//   	AutoScalingConfigurationName: jsii.String("autoScalingConfigurationName"),
//   	MaxConcurrency: jsii.Number(123),
//   	MaxSize: jsii.Number(123),
//   	MinSize: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-autoscalingconfiguration.html
//
type CfnAutoScalingConfigurationProps struct {
	// The customer-provided auto scaling configuration name.
	//
	// It can be used in multiple revisions of a configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-autoscalingconfiguration.html#cfn-apprunner-autoscalingconfiguration-autoscalingconfigurationname
	//
	AutoScalingConfigurationName *string `field:"optional" json:"autoScalingConfigurationName" yaml:"autoScalingConfigurationName"`
	// The maximum number of concurrent requests that an instance processes.
	//
	// If the number of concurrent requests exceeds this limit, App Runner scales the service up.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-autoscalingconfiguration.html#cfn-apprunner-autoscalingconfiguration-maxconcurrency
	//
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The maximum number of instances that a service scales up to.
	//
	// At most `MaxSize` instances actively serve traffic for your service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-autoscalingconfiguration.html#cfn-apprunner-autoscalingconfiguration-maxsize
	//
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of instances that App Runner provisions for a service.
	//
	// The service always has at least `MinSize` provisioned instances. Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.
	//
	// App Runner temporarily doubles the number of provisioned instances during deployments, to maintain the same capacity for both old and new code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-autoscalingconfiguration.html#cfn-apprunner-autoscalingconfiguration-minsize
	//
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// A list of metadata items that you can associate with your auto scaling configuration resource.
	//
	// A tag is a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-autoscalingconfiguration.html#cfn-apprunner-autoscalingconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

