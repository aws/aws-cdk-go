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
//   cfnLaunchProps := &cfnLaunchProps{
//   	groups: []interface{}{
//   		&launchGroupObjectProperty{
//   			feature: jsii.String("feature"),
//   			groupName: jsii.String("groupName"),
//   			variation: jsii.String("variation"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	project: jsii.String("project"),
//   	scheduledSplitsConfig: []interface{}{
//   		&stepConfigProperty{
//   			groupWeights: []interface{}{
//   				&groupToWeightProperty{
//   					groupName: jsii.String("groupName"),
//   					splitWeight: jsii.Number(123),
//   				},
//   			},
//   			startTime: jsii.String("startTime"),
//
//   			// the properties below are optional
//   			segmentOverrides: []interface{}{
//   				&segmentOverrideProperty{
//   					evaluationOrder: jsii.Number(123),
//   					segment: jsii.String("segment"),
//   					weights: []interface{}{
//   						&groupToWeightProperty{
//   							groupName: jsii.String("groupName"),
//   							splitWeight: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	executionStatus: &executionStatusObjectProperty{
//   		status: jsii.String("status"),
//
//   		// the properties below are optional
//   		desiredState: jsii.String("desiredState"),
//   		reason: jsii.String("reason"),
//   	},
//   	metricMonitors: []interface{}{
//   		&metricDefinitionObjectProperty{
//   			entityIdKey: jsii.String("entityIdKey"),
//   			metricName: jsii.String("metricName"),
//   			valueKey: jsii.String("valueKey"),
//
//   			// the properties below are optional
//   			eventPattern: jsii.String("eventPattern"),
//   			unitLabel: jsii.String("unitLabel"),
//   		},
//   	},
//   	randomizationSalt: jsii.String("randomizationSalt"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

