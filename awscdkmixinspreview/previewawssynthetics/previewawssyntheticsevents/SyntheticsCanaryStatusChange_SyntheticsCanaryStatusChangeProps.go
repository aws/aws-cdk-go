package previewawssyntheticsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.synthetics@SyntheticsCanaryStatusChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   syntheticsCanaryStatusChangeProps := &SyntheticsCanaryStatusChangeProps{
//   	AccountId: []*string{
//   		jsii.String("accountId"),
//   	},
//   	CanaryId: []*string{
//   		jsii.String("canaryId"),
//   	},
//   	CanaryName: []*string{
//   		jsii.String("canaryName"),
//   	},
//   	ChangedConfig: &ChangedConfig{
//   		ExecutionArn: &ExecutionArn{
//   			CurrentValue: []*string{
//   				jsii.String("currentValue"),
//   			},
//   			PreviousValue: []*string{
//   				jsii.String("previousValue"),
//   			},
//   		},
//   		TestCodeLayerVersionArn: &TestCodeLayerVersionArn{
//   			CurrentValue: []*string{
//   				jsii.String("currentValue"),
//   			},
//   			PreviousValue: []*string{
//   				jsii.String("previousValue"),
//   			},
//   		},
//   	},
//   	CurrentState: []*string{
//   		jsii.String("currentState"),
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
//   	PreviousState: []*string{
//   		jsii.String("previousState"),
//   	},
//   	SourceLocation: []*string{
//   		jsii.String("sourceLocation"),
//   	},
//   	UpdatedOn: []*string{
//   		jsii.String("updatedOn"),
//   	},
//   }
//
// Experimental.
type SyntheticsCanaryStatusChange_SyntheticsCanaryStatusChangeProps struct {
	// account-id property.
	//
	// Specify an array of string values to match this event if the actual value of account-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// canary-id property.
	//
	// Specify an array of string values to match this event if the actual value of canary-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CanaryId *[]*string `field:"optional" json:"canaryId" yaml:"canaryId"`
	// canary-name property.
	//
	// Specify an array of string values to match this event if the actual value of canary-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CanaryName *[]*string `field:"optional" json:"canaryName" yaml:"canaryName"`
	// changed-config property.
	//
	// Specify an array of string values to match this event if the actual value of changed-config is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChangedConfig *SyntheticsCanaryStatusChange_ChangedConfig `field:"optional" json:"changedConfig" yaml:"changedConfig"`
	// current-state property.
	//
	// Specify an array of string values to match this event if the actual value of current-state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CurrentState *[]*string `field:"optional" json:"currentState" yaml:"currentState"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// previous-state property.
	//
	// Specify an array of string values to match this event if the actual value of previous-state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousState *[]*string `field:"optional" json:"previousState" yaml:"previousState"`
	// source-location property.
	//
	// Specify an array of string values to match this event if the actual value of source-location is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceLocation *[]*string `field:"optional" json:"sourceLocation" yaml:"sourceLocation"`
	// updated-on property.
	//
	// Specify an array of string values to match this event if the actual value of updated-on is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UpdatedOn *[]*string `field:"optional" json:"updatedOn" yaml:"updatedOn"`
}

