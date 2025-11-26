package previewawsopsworksevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Instance aws.opsworks@OpsWorksInstanceStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   opsWorksInstanceStateChangeProps := &OpsWorksInstanceStateChangeProps{
//   	Ec2InstanceId: []*string{
//   		jsii.String("ec2InstanceId"),
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
//   	Hostname: []*string{
//   		jsii.String("hostname"),
//   	},
//   	InitiatedBy: []*string{
//   		jsii.String("initiatedBy"),
//   	},
//   	InstanceId: []*string{
//   		jsii.String("instanceId"),
//   	},
//   	LayerIds: []*string{
//   		jsii.String("layerIds"),
//   	},
//   	StackId: []*string{
//   		jsii.String("stackId"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_OpsWorksInstanceStateChange_OpsWorksInstanceStateChangeProps struct {
	// ec2-instance-id property.
	//
	// Specify an array of string values to match this event if the actual value of ec2-instance-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Ec2InstanceId *[]*string `field:"optional" json:"ec2InstanceId" yaml:"ec2InstanceId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// hostname property.
	//
	// Specify an array of string values to match this event if the actual value of hostname is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Hostname *[]*string `field:"optional" json:"hostname" yaml:"hostname"`
	// initiated_by property.
	//
	// Specify an array of string values to match this event if the actual value of initiated_by is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InitiatedBy *[]*string `field:"optional" json:"initiatedBy" yaml:"initiatedBy"`
	// instance-id property.
	//
	// Specify an array of string values to match this event if the actual value of instance-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Instance reference.
	//
	// Experimental.
	InstanceId *[]*string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// layer-ids property.
	//
	// Specify an array of string values to match this event if the actual value of layer-ids is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LayerIds *[]*string `field:"optional" json:"layerIds" yaml:"layerIds"`
	// stack-id property.
	//
	// Specify an array of string values to match this event if the actual value of stack-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StackId *[]*string `field:"optional" json:"stackId" yaml:"stackId"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
}

