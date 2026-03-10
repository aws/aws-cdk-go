package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.sagemaker@SageMakerGroundTruthLabelingJobTaskStatusChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerGroundTruthLabelingJobTaskStatusChangeProps := &SageMakerGroundTruthLabelingJobTaskStatusChangeProps{
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
//   	LabelingJobArn: []*string{
//   		jsii.String("labelingJobArn"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   	WorkteamTeamArn: []*string{
//   		jsii.String("workteamTeamArn"),
//   	},
//   }
//
// Experimental.
type SageMakerGroundTruthLabelingJobTaskStatusChange_SageMakerGroundTruthLabelingJobTaskStatusChangeProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// labeling-job-arn property.
	//
	// Specify an array of string values to match this event if the actual value of labeling-job-arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LabelingJobArn *[]*string `field:"optional" json:"labelingJobArn" yaml:"labelingJobArn"`
	// State property.
	//
	// Specify an array of string values to match this event if the actual value of State is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
	// workteam-team-arn property.
	//
	// Specify an array of string values to match this event if the actual value of workteam-team-arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	WorkteamTeamArn *[]*string `field:"optional" json:"workteamTeamArn" yaml:"workteamTeamArn"`
}

