package previewawsemrevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.emr@EMRInstanceFleetStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eMRInstanceFleetStateChangeProps := &EMRInstanceFleetStateChangeProps{
//   	ClusterId: []*string{
//   		jsii.String("clusterId"),
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
//   	InstanceFleetId: []*string{
//   		jsii.String("instanceFleetId"),
//   	},
//   	InstanceFleetType: []*string{
//   		jsii.String("instanceFleetType"),
//   	},
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	ProvisionedOnDemandCapacity: []*string{
//   		jsii.String("provisionedOnDemandCapacity"),
//   	},
//   	ProvisionedSpotCapacity: []*string{
//   		jsii.String("provisionedSpotCapacity"),
//   	},
//   	Severity: []*string{
//   		jsii.String("severity"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   	TargetOnDemandCapacity: []*string{
//   		jsii.String("targetOnDemandCapacity"),
//   	},
//   	TargetSpotCapacity: []*string{
//   		jsii.String("targetSpotCapacity"),
//   	},
//   }
//
// Experimental.
type EMRInstanceFleetStateChange_EMRInstanceFleetStateChangeProps struct {
	// clusterId property.
	//
	// Specify an array of string values to match this event if the actual value of clusterId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClusterId *[]*string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// instanceFleetId property.
	//
	// Specify an array of string values to match this event if the actual value of instanceFleetId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceFleetId *[]*string `field:"optional" json:"instanceFleetId" yaml:"instanceFleetId"`
	// instanceFleetType property.
	//
	// Specify an array of string values to match this event if the actual value of instanceFleetType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceFleetType *[]*string `field:"optional" json:"instanceFleetType" yaml:"instanceFleetType"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// provisionedOnDemandCapacity property.
	//
	// Specify an array of string values to match this event if the actual value of provisionedOnDemandCapacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProvisionedOnDemandCapacity *[]*string `field:"optional" json:"provisionedOnDemandCapacity" yaml:"provisionedOnDemandCapacity"`
	// provisionedSpotCapacity property.
	//
	// Specify an array of string values to match this event if the actual value of provisionedSpotCapacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProvisionedSpotCapacity *[]*string `field:"optional" json:"provisionedSpotCapacity" yaml:"provisionedSpotCapacity"`
	// severity property.
	//
	// Specify an array of string values to match this event if the actual value of severity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Severity *[]*string `field:"optional" json:"severity" yaml:"severity"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
	// targetOnDemandCapacity property.
	//
	// Specify an array of string values to match this event if the actual value of targetOnDemandCapacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetOnDemandCapacity *[]*string `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// targetSpotCapacity property.
	//
	// Specify an array of string values to match this event if the actual value of targetSpotCapacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetSpotCapacity *[]*string `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}

