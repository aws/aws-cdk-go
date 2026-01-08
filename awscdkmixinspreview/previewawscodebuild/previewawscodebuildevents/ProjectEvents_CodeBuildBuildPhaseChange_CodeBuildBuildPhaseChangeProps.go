package previewawscodebuildevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Project aws.codebuild@CodeBuildBuildPhaseChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeBuildBuildPhaseChangeProps := &CodeBuildBuildPhaseChangeProps{
//   	AdditionalInformation: &AdditionalInformation{
//   		Artifact: &Artifact{
//   			Location: []*string{
//   				jsii.String("location"),
//   			},
//   			Md5Sum: []*string{
//   				jsii.String("md5Sum"),
//   			},
//   			Sha256Sum: []*string{
//   				jsii.String("sha256Sum"),
//   			},
//   		},
//   		BuildComplete: []*string{
//   			jsii.String("buildComplete"),
//   		},
//   		BuildStartTime: []*string{
//   			jsii.String("buildStartTime"),
//   		},
//   		Cache: &Cache{
//   			Location: []*string{
//   				jsii.String("location"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   		Environment: &Environment{
//   			ComputeType: []*string{
//   				jsii.String("computeType"),
//   			},
//   			EnvironmentVariables: []EnvironmentItem{
//   				&EnvironmentItem{
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					Type: []*string{
//   						jsii.String("type"),
//   					},
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   			Image: []*string{
//   				jsii.String("image"),
//   			},
//   			ImagePullCredentialsType: []*string{
//   				jsii.String("imagePullCredentialsType"),
//   			},
//   			PrivilegedMode: []*string{
//   				jsii.String("privilegedMode"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   		Initiator: []*string{
//   			jsii.String("initiator"),
//   		},
//   		Logs: &Logs{
//   			DeepLink: []*string{
//   				jsii.String("deepLink"),
//   			},
//   			GroupName: []*string{
//   				jsii.String("groupName"),
//   			},
//   			StreamName: []*string{
//   				jsii.String("streamName"),
//   			},
//   		},
//   		NetworkInterface: &NetworkInterface{
//   			EniId: []*string{
//   				jsii.String("eniId"),
//   			},
//   			SubnetId: []*string{
//   				jsii.String("subnetId"),
//   			},
//   		},
//   		Phases: []AdditionalInformationItem{
//   			&AdditionalInformationItem{
//   				DurationInSeconds: []*string{
//   					jsii.String("durationInSeconds"),
//   				},
//   				EndTime: []*string{
//   					jsii.String("endTime"),
//   				},
//   				PhaseContext: []*string{
//   					jsii.String("phaseContext"),
//   				},
//   				PhaseStatus: []*string{
//   					jsii.String("phaseStatus"),
//   				},
//   				PhaseType: []*string{
//   					jsii.String("phaseType"),
//   				},
//   				StartTime: []*string{
//   					jsii.String("startTime"),
//   				},
//   			},
//   		},
//   		QueuedTimeoutInMinutes: []*string{
//   			jsii.String("queuedTimeoutInMinutes"),
//   		},
//   		Source: &Source{
//   			Auth: &Auth{
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   			Buildspec: []*string{
//   				jsii.String("buildspec"),
//   			},
//   			Location: []*string{
//   				jsii.String("location"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   		SourceVersion: []*string{
//   			jsii.String("sourceVersion"),
//   		},
//   		TimeoutInMinutes: []*string{
//   			jsii.String("timeoutInMinutes"),
//   		},
//   		VpcConfig: &VpcConfig{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			Subnets: []VpcConfigItem{
//   				&VpcConfigItem{
//   					BuildFleetAz: []*string{
//   						jsii.String("buildFleetAz"),
//   					},
//   					CustomerAz: []*string{
//   						jsii.String("customerAz"),
//   					},
//   					SubnetId: []*string{
//   						jsii.String("subnetId"),
//   					},
//   				},
//   			},
//   			VpcId: []*string{
//   				jsii.String("vpcId"),
//   			},
//   		},
//   	},
//   	BuildId: []*string{
//   		jsii.String("buildId"),
//   	},
//   	CompletedPhase: []*string{
//   		jsii.String("completedPhase"),
//   	},
//   	CompletedPhaseContext: []*string{
//   		jsii.String("completedPhaseContext"),
//   	},
//   	CompletedPhaseDurationSeconds: []*string{
//   		jsii.String("completedPhaseDurationSeconds"),
//   	},
//   	CompletedPhaseEnd: []*string{
//   		jsii.String("completedPhaseEnd"),
//   	},
//   	CompletedPhaseStart: []*string{
//   		jsii.String("completedPhaseStart"),
//   	},
//   	CompletedPhaseStatus: []*string{
//   		jsii.String("completedPhaseStatus"),
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
//   	ProjectName: []*string{
//   		jsii.String("projectName"),
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
// Experimental.
type ProjectEvents_CodeBuildBuildPhaseChange_CodeBuildBuildPhaseChangeProps struct {
	// additional-information property.
	//
	// Specify an array of string values to match this event if the actual value of additional-information is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AdditionalInformation *ProjectEvents_CodeBuildBuildPhaseChange_AdditionalInformation `field:"optional" json:"additionalInformation" yaml:"additionalInformation"`
	// build-id property.
	//
	// Specify an array of string values to match this event if the actual value of build-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BuildId *[]*string `field:"optional" json:"buildId" yaml:"buildId"`
	// completed-phase property.
	//
	// Specify an array of string values to match this event if the actual value of completed-phase is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CompletedPhase *[]*string `field:"optional" json:"completedPhase" yaml:"completedPhase"`
	// completed-phase-context property.
	//
	// Specify an array of string values to match this event if the actual value of completed-phase-context is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CompletedPhaseContext *[]*string `field:"optional" json:"completedPhaseContext" yaml:"completedPhaseContext"`
	// completed-phase-duration-seconds property.
	//
	// Specify an array of string values to match this event if the actual value of completed-phase-duration-seconds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CompletedPhaseDurationSeconds *[]*string `field:"optional" json:"completedPhaseDurationSeconds" yaml:"completedPhaseDurationSeconds"`
	// completed-phase-end property.
	//
	// Specify an array of string values to match this event if the actual value of completed-phase-end is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CompletedPhaseEnd *[]*string `field:"optional" json:"completedPhaseEnd" yaml:"completedPhaseEnd"`
	// completed-phase-start property.
	//
	// Specify an array of string values to match this event if the actual value of completed-phase-start is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CompletedPhaseStart *[]*string `field:"optional" json:"completedPhaseStart" yaml:"completedPhaseStart"`
	// completed-phase-status property.
	//
	// Specify an array of string values to match this event if the actual value of completed-phase-status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CompletedPhaseStatus *[]*string `field:"optional" json:"completedPhaseStatus" yaml:"completedPhaseStatus"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// project-name property.
	//
	// Specify an array of string values to match this event if the actual value of project-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Project reference.
	//
	// Experimental.
	ProjectName *[]*string `field:"optional" json:"projectName" yaml:"projectName"`
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

