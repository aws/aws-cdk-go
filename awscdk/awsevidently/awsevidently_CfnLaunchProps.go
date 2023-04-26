package awsevidently

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLaunch`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunchProps := &CfnLaunchProps{
//   	Groups: []interface{}{
//   		&LaunchGroupObjectProperty{
//   			Feature: jsii.String("feature"),
//   			GroupName: jsii.String("groupName"),
//   			Variation: jsii.String("variation"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Project: jsii.String("project"),
//   	ScheduledSplitsConfig: []interface{}{
//   		&StepConfigProperty{
//   			GroupWeights: []interface{}{
//   				&GroupToWeightProperty{
//   					GroupName: jsii.String("groupName"),
//   					SplitWeight: jsii.Number(123),
//   				},
//   			},
//   			StartTime: jsii.String("startTime"),
//
//   			// the properties below are optional
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
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ExecutionStatus: &ExecutionStatusObjectProperty{
//   		Status: jsii.String("status"),
//
//   		// the properties below are optional
//   		DesiredState: jsii.String("desiredState"),
//   		Reason: jsii.String("reason"),
//   	},
//   	MetricMonitors: []interface{}{
//   		&MetricDefinitionObjectProperty{
//   			EntityIdKey: jsii.String("entityIdKey"),
//   			MetricName: jsii.String("metricName"),
//   			ValueKey: jsii.String("valueKey"),
//
//   			// the properties below are optional
//   			EventPattern: jsii.String("eventPattern"),
//   			UnitLabel: jsii.String("unitLabel"),
//   		},
//   	},
//   	RandomizationSalt: jsii.String("randomizationSalt"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLaunchProps struct {
	// An array of structures that contains the feature and variations that are to be used for the launch.
	//
	// You can up to five launch groups in a launch.
	Groups interface{} `field:"required" json:"groups" yaml:"groups"`
	// The name for the launch.
	//
	// It can include up to 127 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name or ARN of the project that you want to create the launch in.
	Project *string `field:"required" json:"project" yaml:"project"`
	// An array of structures that define the traffic allocation percentages among the feature variations during each step of the launch.
	ScheduledSplitsConfig interface{} `field:"required" json:"scheduledSplitsConfig" yaml:"scheduledSplitsConfig"`
	// An optional description for the launch.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A structure that you can use to start and stop the launch.
	ExecutionStatus interface{} `field:"optional" json:"executionStatus" yaml:"executionStatus"`
	// An array of structures that define the metrics that will be used to monitor the launch performance.
	//
	// You can have up to three metric monitors in the array.
	MetricMonitors interface{} `field:"optional" json:"metricMonitors" yaml:"metricMonitors"`
	// When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served.
	//
	// This randomization ID is a combination of the entity ID and `randomizationSalt` . If you omit `randomizationSalt` , Evidently uses the launch name as the `randomizationsSalt` .
	RandomizationSalt *string `field:"optional" json:"randomizationSalt" yaml:"randomizationSalt"`
	// Assigns one or more tags (key-value pairs) to the launch.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with a launch.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

