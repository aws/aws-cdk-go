package mixinsawsevidently

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLaunchPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLaunchMixinProps := &CfnLaunchMixinProps{
//   	Description: jsii.String("description"),
//   	ExecutionStatus: &ExecutionStatusObjectProperty{
//   		DesiredState: jsii.String("desiredState"),
//   		Reason: jsii.String("reason"),
//   		Status: jsii.String("status"),
//   	},
//   	Groups: []interface{}{
//   		&LaunchGroupObjectProperty{
//   			Description: jsii.String("description"),
//   			Feature: jsii.String("feature"),
//   			GroupName: jsii.String("groupName"),
//   			Variation: jsii.String("variation"),
//   		},
//   	},
//   	MetricMonitors: []interface{}{
//   		&MetricDefinitionObjectProperty{
//   			EntityIdKey: jsii.String("entityIdKey"),
//   			EventPattern: jsii.String("eventPattern"),
//   			MetricName: jsii.String("metricName"),
//   			UnitLabel: jsii.String("unitLabel"),
//   			ValueKey: jsii.String("valueKey"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Project: jsii.String("project"),
//   	RandomizationSalt: jsii.String("randomizationSalt"),
//   	ScheduledSplitsConfig: []interface{}{
//   		&StepConfigProperty{
//   			GroupWeights: []interface{}{
//   				&GroupToWeightProperty{
//   					GroupName: jsii.String("groupName"),
//   					SplitWeight: jsii.Number(123),
//   				},
//   			},
//   			SegmentOverrides: []interface{}{
//   				&SegmentOverrideProperty{
//   					EvaluationOrder: jsii.Number(123),
//   					Segment: jsii.String("segment"),
//   					Weights: []interface{}{
//   						&GroupToWeightProperty{
//   							GroupName: jsii.String("groupName"),
//   							SplitWeight: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   			StartTime: jsii.String("startTime"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html
//
type CfnLaunchMixinProps struct {
	// An optional description for the launch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A structure that you can use to start and stop the launch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-executionstatus
	//
	ExecutionStatus interface{} `field:"optional" json:"executionStatus" yaml:"executionStatus"`
	// An array of structures that contains the feature and variations that are to be used for the launch.
	//
	// You can up to five launch groups in a launch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-groups
	//
	Groups interface{} `field:"optional" json:"groups" yaml:"groups"`
	// An array of structures that define the metrics that will be used to monitor the launch performance.
	//
	// You can have up to three metric monitors in the array.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-metricmonitors
	//
	MetricMonitors interface{} `field:"optional" json:"metricMonitors" yaml:"metricMonitors"`
	// The name for the launch.
	//
	// It can include up to 127 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name or ARN of the project that you want to create the launch in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-project
	//
	Project *string `field:"optional" json:"project" yaml:"project"`
	// When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served.
	//
	// This randomization ID is a combination of the entity ID and `randomizationSalt` . If you omit `randomizationSalt` , Evidently uses the launch name as the `randomizationsSalt` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-randomizationsalt
	//
	RandomizationSalt *string `field:"optional" json:"randomizationSalt" yaml:"randomizationSalt"`
	// An array of structures that define the traffic allocation percentages among the feature variations during each step of the launch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-scheduledsplitsconfig
	//
	ScheduledSplitsConfig interface{} `field:"optional" json:"scheduledSplitsConfig" yaml:"scheduledSplitsConfig"`
	// Assigns one or more tags (key-value pairs) to the launch.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with a launch.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

