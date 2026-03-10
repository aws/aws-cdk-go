package previewawssyntheticsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.synthetics@SyntheticsCanaryTestRunFailure event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   syntheticsCanaryTestRunFailureProps := &SyntheticsCanaryTestRunFailureProps{
//   	AccountId: []*string{
//   		jsii.String("accountId"),
//   	},
//   	ArtifactLocation: []*string{
//   		jsii.String("artifactLocation"),
//   	},
//   	CanaryId: []*string{
//   		jsii.String("canaryId"),
//   	},
//   	CanaryName: []*string{
//   		jsii.String("canaryName"),
//   	},
//   	CanaryRunId: []*string{
//   		jsii.String("canaryRunId"),
//   	},
//   	CanaryRunTimeline: &CanaryRunTimeline{
//   		Completed: []*string{
//   			jsii.String("completed"),
//   		},
//   		Started: []*string{
//   			jsii.String("started"),
//   		},
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
//   	StateReason: []*string{
//   		jsii.String("stateReason"),
//   	},
//   	TestRunStatus: []*string{
//   		jsii.String("testRunStatus"),
//   	},
//   }
//
// Experimental.
type SyntheticsCanaryTestRunFailure_SyntheticsCanaryTestRunFailureProps struct {
	// account-id property.
	//
	// Specify an array of string values to match this event if the actual value of account-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// artifact-location property.
	//
	// Specify an array of string values to match this event if the actual value of artifact-location is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ArtifactLocation *[]*string `field:"optional" json:"artifactLocation" yaml:"artifactLocation"`
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
	// canary-run-id property.
	//
	// Specify an array of string values to match this event if the actual value of canary-run-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CanaryRunId *[]*string `field:"optional" json:"canaryRunId" yaml:"canaryRunId"`
	// canary-run-timeline property.
	//
	// Specify an array of string values to match this event if the actual value of canary-run-timeline is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CanaryRunTimeline *SyntheticsCanaryTestRunFailure_CanaryRunTimeline `field:"optional" json:"canaryRunTimeline" yaml:"canaryRunTimeline"`
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
	// state-reason property.
	//
	// Specify an array of string values to match this event if the actual value of state-reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StateReason *[]*string `field:"optional" json:"stateReason" yaml:"stateReason"`
	// test-run-status property.
	//
	// Specify an array of string values to match this event if the actual value of test-run-status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TestRunStatus *[]*string `field:"optional" json:"testRunStatus" yaml:"testRunStatus"`
}

