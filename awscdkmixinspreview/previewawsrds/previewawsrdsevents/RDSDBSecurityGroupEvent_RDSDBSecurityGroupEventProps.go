package previewawsrdsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.rds@RDSDBSecurityGroupEvent event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rDSDBSecurityGroupEventProps := &RDSDBSecurityGroupEventProps{
//   	Date: []*string{
//   		jsii.String("date"),
//   	},
//   	EventCategories: []*string{
//   		jsii.String("eventCategories"),
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
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	SourceArn: []*string{
//   		jsii.String("sourceArn"),
//   	},
//   	SourceIdentifier: []*string{
//   		jsii.String("sourceIdentifier"),
//   	},
//   	SourceType: []*string{
//   		jsii.String("sourceType"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// Experimental.
type RDSDBSecurityGroupEvent_RDSDBSecurityGroupEventProps struct {
	// Date property.
	//
	// Specify an array of string values to match this event if the actual value of Date is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Date *[]*string `field:"optional" json:"date" yaml:"date"`
	// EventCategories property.
	//
	// Specify an array of string values to match this event if the actual value of EventCategories is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventCategories *[]*string `field:"optional" json:"eventCategories" yaml:"eventCategories"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// Message property.
	//
	// Specify an array of string values to match this event if the actual value of Message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// SourceArn property.
	//
	// Specify an array of string values to match this event if the actual value of SourceArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceArn *[]*string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
	// SourceIdentifier property.
	//
	// Specify an array of string values to match this event if the actual value of SourceIdentifier is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIdentifier *[]*string `field:"optional" json:"sourceIdentifier" yaml:"sourceIdentifier"`
	// SourceType property.
	//
	// Specify an array of string values to match this event if the actual value of SourceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceType *[]*string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// Tags property.
	//
	// Specify an array of string values to match this event if the actual value of Tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

