package previewawsguarddutyevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Detector aws.guardduty@GuardDutyFinding event.
//
// Example:
//
//
// Experimental.
type DetectorEvents_GuardDutyFinding_GuardDutyFindingProps struct {
	// accountId property.
	//
	// Specify an array of string values to match this event if the actual value of accountId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
	// createdAt property.
	//
	// Specify an array of string values to match this event if the actual value of createdAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreatedAt *[]*string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// id property.
	//
	// Specify an array of string values to match this event if the actual value of id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
	// partition property.
	//
	// Specify an array of string values to match this event if the actual value of partition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Partition *[]*string `field:"optional" json:"partition" yaml:"partition"`
	// region property.
	//
	// Specify an array of string values to match this event if the actual value of region is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Region *[]*string `field:"optional" json:"region" yaml:"region"`
	// resource property.
	//
	// Specify an array of string values to match this event if the actual value of resource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Resource *DetectorEvents_GuardDutyFinding_Resource `field:"optional" json:"resource" yaml:"resource"`
	// schemaVersion property.
	//
	// Specify an array of string values to match this event if the actual value of schemaVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SchemaVersion *[]*string `field:"optional" json:"schemaVersion" yaml:"schemaVersion"`
	// service property.
	//
	// Specify an array of string values to match this event if the actual value of service is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Service *DetectorEvents_GuardDutyFinding_Service `field:"optional" json:"service" yaml:"service"`
	// severity property.
	//
	// Specify an array of string values to match this event if the actual value of severity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Severity *[]*string `field:"optional" json:"severity" yaml:"severity"`
	// title property.
	//
	// Specify an array of string values to match this event if the actual value of title is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Title *[]*string `field:"optional" json:"title" yaml:"title"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
	// updatedAt property.
	//
	// Specify an array of string values to match this event if the actual value of updatedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UpdatedAt *[]*string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

