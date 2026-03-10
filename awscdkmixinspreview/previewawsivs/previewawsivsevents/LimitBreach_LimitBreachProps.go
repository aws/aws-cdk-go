package previewawsivsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.ivs@LimitBreach event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   limitBreachProps := &LimitBreachProps{
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
//   	ExceededBy: []*string{
//   		jsii.String("exceededBy"),
//   	},
//   	Limit: []*string{
//   		jsii.String("limit"),
//   	},
//   	LimitUnit: []*string{
//   		jsii.String("limitUnit"),
//   	},
//   	LimitValue: []*string{
//   		jsii.String("limitValue"),
//   	},
//   }
//
// Experimental.
type LimitBreach_LimitBreachProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// exceeded_by property.
	//
	// Specify an array of string values to match this event if the actual value of exceeded_by is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExceededBy *[]*string `field:"optional" json:"exceededBy" yaml:"exceededBy"`
	// limit property.
	//
	// Specify an array of string values to match this event if the actual value of limit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Limit *[]*string `field:"optional" json:"limit" yaml:"limit"`
	// limit_unit property.
	//
	// Specify an array of string values to match this event if the actual value of limit_unit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LimitUnit *[]*string `field:"optional" json:"limitUnit" yaml:"limitUnit"`
	// limit_value property.
	//
	// Specify an array of string values to match this event if the actual value of limit_value is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LimitValue *[]*string `field:"optional" json:"limitValue" yaml:"limitValue"`
}

