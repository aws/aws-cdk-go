package previewawsroute53recoveryreadinessevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for ReadinessCheck aws.route53recoveryreadiness@Route53ApplicationRecoveryControllerReadinessCheckStatusChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   route53ApplicationRecoveryControllerReadinessCheckStatusChangeProps := &Route53ApplicationRecoveryControllerReadinessCheckStatusChangeProps{
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
//   	NewState: &State{
//   		ReadinessStatus: []*string{
//   			jsii.String("readinessStatus"),
//   		},
//   	},
//   	PreviousState: &State{
//   		ReadinessStatus: []*string{
//   			jsii.String("readinessStatus"),
//   		},
//   	},
//   	ReadinessCheckName: []*string{
//   		jsii.String("readinessCheckName"),
//   	},
//   }
//
// Experimental.
type ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange_Route53ApplicationRecoveryControllerReadinessCheckStatusChangeProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// new-state property.
	//
	// Specify an array of string values to match this event if the actual value of new-state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NewState *ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange_State `field:"optional" json:"newState" yaml:"newState"`
	// previous-state property.
	//
	// Specify an array of string values to match this event if the actual value of previous-state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousState *ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange_State `field:"optional" json:"previousState" yaml:"previousState"`
	// readiness-check-name property.
	//
	// Specify an array of string values to match this event if the actual value of readiness-check-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the ReadinessCheck reference.
	//
	// Experimental.
	ReadinessCheckName *[]*string `field:"optional" json:"readinessCheckName" yaml:"readinessCheckName"`
}

