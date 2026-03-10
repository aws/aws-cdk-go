package previewawscodedeployevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.codedeploy@CodeDeployInstanceStateChangeNotification event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeDeployInstanceStateChangeNotificationProps := &CodeDeployInstanceStateChangeNotificationProps{
//   	Application: []*string{
//   		jsii.String("application"),
//   	},
//   	DeploymentGroup: []*string{
//   		jsii.String("deploymentGroup"),
//   	},
//   	DeploymentId: []*string{
//   		jsii.String("deploymentId"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	InstanceGroupId: []*string{
//   		jsii.String("instanceGroupId"),
//   	},
//   	InstanceId: []*string{
//   		jsii.String("instanceId"),
//   	},
//   	Region: []*string{
//   		jsii.String("region"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   }
//
// Experimental.
type CodeDeployInstanceStateChangeNotification_CodeDeployInstanceStateChangeNotificationProps struct {
	// application property.
	//
	// Specify an array of string values to match this event if the actual value of application is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Application *[]*string `field:"optional" json:"application" yaml:"application"`
	// deploymentGroup property.
	//
	// Specify an array of string values to match this event if the actual value of deploymentGroup is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeploymentGroup *[]*string `field:"optional" json:"deploymentGroup" yaml:"deploymentGroup"`
	// deploymentId property.
	//
	// Specify an array of string values to match this event if the actual value of deploymentId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeploymentId *[]*string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// instanceGroupId property.
	//
	// Specify an array of string values to match this event if the actual value of instanceGroupId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceGroupId *[]*string `field:"optional" json:"instanceGroupId" yaml:"instanceGroupId"`
	// instanceId property.
	//
	// Specify an array of string values to match this event if the actual value of instanceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceId *[]*string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// region property.
	//
	// Specify an array of string values to match this event if the actual value of region is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Region *[]*string `field:"optional" json:"region" yaml:"region"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
}

