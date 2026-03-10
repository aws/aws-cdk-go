package previewawscodebuildevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.codebuild@AWSAPICallViaCloudTrail event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrailProps := &AWSAPICallViaCloudTrailProps{
//   	AwsRegion: []*string{
//   		jsii.String("awsRegion"),
//   	},
//   	ErrorCode: []*string{
//   		jsii.String("errorCode"),
//   	},
//   	ErrorMessage: []*string{
//   		jsii.String("errorMessage"),
//   	},
//   	EventId: []*string{
//   		jsii.String("eventId"),
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
//   	EventName: []*string{
//   		jsii.String("eventName"),
//   	},
//   	EventSource: []*string{
//   		jsii.String("eventSource"),
//   	},
//   	EventTime: []*string{
//   		jsii.String("eventTime"),
//   	},
//   	EventType: []*string{
//   		jsii.String("eventType"),
//   	},
//   	EventVersion: []*string{
//   		jsii.String("eventVersion"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	RequestParameters: &RequestParameters{
//   		AwsActId: []*string{
//   			jsii.String("awsActId"),
//   		},
//   		InvokedBy: []*string{
//   			jsii.String("invokedBy"),
//   		},
//   		PayerId: []*string{
//   			jsii.String("payerId"),
//   		},
//   		UserArn: []*string{
//   			jsii.String("userArn"),
//   		},
//   	},
//   	ResponseElements: &ResponseElements{
//   		BuildProperty: &Build{
//   			Arn: []*string{
//   				jsii.String("arn"),
//   			},
//   			Artifacts: &Artifacts{
//   				EncryptionDisabled: []*string{
//   					jsii.String("encryptionDisabled"),
//   				},
//   				Location: []*string{
//   					jsii.String("location"),
//   				},
//   			},
//   			BuildComplete: []*string{
//   				jsii.String("buildComplete"),
//   			},
//   			BuildStatus: []*string{
//   				jsii.String("buildStatus"),
//   			},
//   			Cache: &Cache{
//   				Location: []*string{
//   					jsii.String("location"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   			CurrentPhase: []*string{
//   				jsii.String("currentPhase"),
//   			},
//   			EncryptionKey: []*string{
//   				jsii.String("encryptionKey"),
//   			},
//   			EndTime: []*string{
//   				jsii.String("endTime"),
//   			},
//   			Environment: &Environment{
//   				Certificate: []*string{
//   					jsii.String("certificate"),
//   				},
//   				ComputeType: []*string{
//   					jsii.String("computeType"),
//   				},
//   				EnvironmentVariables: []EnvironmentItem{
//   					&EnvironmentItem{
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   						Type: []*string{
//   							jsii.String("type"),
//   						},
//   						Value: []*string{
//   							jsii.String("value"),
//   						},
//   					},
//   				},
//   				Image: []*string{
//   					jsii.String("image"),
//   				},
//   				ImagePullCredentialsType: []*string{
//   					jsii.String("imagePullCredentialsType"),
//   				},
//   				PrivilegedMode: []*string{
//   					jsii.String("privilegedMode"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   			Id: []*string{
//   				jsii.String("id"),
//   			},
//   			Initiator: []*string{
//   				jsii.String("initiator"),
//   			},
//   			Logs: &Logs{
//   				DeepLink: []*string{
//   					jsii.String("deepLink"),
//   				},
//   				GroupName: []*string{
//   					jsii.String("groupName"),
//   				},
//   				StreamName: []*string{
//   					jsii.String("streamName"),
//   				},
//   			},
//   			Phases: []BuildPhase{
//   				&BuildPhase{
//   					DurationInSeconds: []*string{
//   						jsii.String("durationInSeconds"),
//   					},
//   					EndTime: []*string{
//   						jsii.String("endTime"),
//   					},
//   					PhaseContext: []*string{
//   						jsii.String("phaseContext"),
//   					},
//   					PhaseStatus: []*string{
//   						jsii.String("phaseStatus"),
//   					},
//   					PhaseType: []*string{
//   						jsii.String("phaseType"),
//   					},
//   					StartTime: []*string{
//   						jsii.String("startTime"),
//   					},
//   				},
//   			},
//   			ProjectName: []*string{
//   				jsii.String("projectName"),
//   			},
//   			QueuedTimeoutInMinutes: []*string{
//   				jsii.String("queuedTimeoutInMinutes"),
//   			},
//   			ResolvedSourceVersion: []*string{
//   				jsii.String("resolvedSourceVersion"),
//   			},
//   			ServiceRole: []*string{
//   				jsii.String("serviceRole"),
//   			},
//   			Source: &Source{
//   				Auth: &Auth{
//   					Type: []*string{
//   						jsii.String("type"),
//   					},
//   				},
//   				Buildspec: []*string{
//   					jsii.String("buildspec"),
//   				},
//   				InsecureSsl: []*string{
//   					jsii.String("insecureSsl"),
//   				},
//   				Location: []*string{
//   					jsii.String("location"),
//   				},
//   				ReportBuildStatus: []*string{
//   					jsii.String("reportBuildStatus"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   			SourceVersion: []*string{
//   				jsii.String("sourceVersion"),
//   			},
//   			StartTime: []*string{
//   				jsii.String("startTime"),
//   			},
//   			TimeoutInMinutes: []*string{
//   				jsii.String("timeoutInMinutes"),
//   			},
//   			VpcConfig: &VpcConfig{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   				VpcId: []*string{
//   					jsii.String("vpcId"),
//   				},
//   			},
//   		},
//   		Message: []*string{
//   			jsii.String("message"),
//   		},
//   		Project: &Project{
//   			Arn: []*string{
//   				jsii.String("arn"),
//   			},
//   			Artifacts: &Artifacts1{
//   				EncryptionDisabled: []*string{
//   					jsii.String("encryptionDisabled"),
//   				},
//   				Location: []*string{
//   					jsii.String("location"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				NamespaceType: []*string{
//   					jsii.String("namespaceType"),
//   				},
//   				Packaging: []*string{
//   					jsii.String("packaging"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   			Badge: &Badge{
//   				BadgeEnabled: []*string{
//   					jsii.String("badgeEnabled"),
//   				},
//   				BadgeRequestUrl: []*string{
//   					jsii.String("badgeRequestUrl"),
//   				},
//   			},
//   			Cache: &Cache{
//   				Location: []*string{
//   					jsii.String("location"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   			Created: []*string{
//   				jsii.String("created"),
//   			},
//   			Description: []*string{
//   				jsii.String("description"),
//   			},
//   			EncryptionKey: []*string{
//   				jsii.String("encryptionKey"),
//   			},
//   			Environment: &Environment1{
//   				ComputeType: []*string{
//   					jsii.String("computeType"),
//   				},
//   				EnvironmentVariables: []EnvironmentItem{
//   					&EnvironmentItem{
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   						Type: []*string{
//   							jsii.String("type"),
//   						},
//   						Value: []*string{
//   							jsii.String("value"),
//   						},
//   					},
//   				},
//   				Image: []*string{
//   					jsii.String("image"),
//   				},
//   				ImagePullCredentialsType: []*string{
//   					jsii.String("imagePullCredentialsType"),
//   				},
//   				PrivilegedMode: []*string{
//   					jsii.String("privilegedMode"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   			LastModified: []*string{
//   				jsii.String("lastModified"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			QueuedTimeoutInMinutes: []*string{
//   				jsii.String("queuedTimeoutInMinutes"),
//   			},
//   			ServiceRole: []*string{
//   				jsii.String("serviceRole"),
//   			},
//   			Source: &Source1{
//   				Auth: &Auth{
//   					Type: []*string{
//   						jsii.String("type"),
//   					},
//   				},
//   				Buildspec: []*string{
//   					jsii.String("buildspec"),
//   				},
//   				GitCloneDepth: []*string{
//   					jsii.String("gitCloneDepth"),
//   				},
//   				GitSubmodulesConfig: &GitSubmodulesConfig{
//   					FetchSubmodules: []*string{
//   						jsii.String("fetchSubmodules"),
//   					},
//   				},
//   				InsecureSsl: []*string{
//   					jsii.String("insecureSsl"),
//   				},
//   				Location: []*string{
//   					jsii.String("location"),
//   				},
//   				ReportBuildStatus: []*string{
//   					jsii.String("reportBuildStatus"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   			SourceVersion: []*string{
//   				jsii.String("sourceVersion"),
//   			},
//   			Tags: []ProjectItem{
//   				&ProjectItem{
//   					Key: []*string{
//   						jsii.String("key"),
//   					},
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   			TimeoutInMinutes: []*string{
//   				jsii.String("timeoutInMinutes"),
//   			},
//   		},
//   		Webhook: &Webhook{
//   			BranchFilter: []*string{
//   				jsii.String("branchFilter"),
//   			},
//   			LastModifiedSecret: []*string{
//   				jsii.String("lastModifiedSecret"),
//   			},
//   			PayloadUrl: []*string{
//   				jsii.String("payloadUrl"),
//   			},
//   			Url: []*string{
//   				jsii.String("url"),
//   			},
//   		},
//   		WebhookDeletedStatus: []*string{
//   			jsii.String("webhookDeletedStatus"),
//   		},
//   	},
//   	SourceIpAddress: []*string{
//   		jsii.String("sourceIpAddress"),
//   	},
//   	UserAgent: []*string{
//   		jsii.String("userAgent"),
//   	},
//   	UserIdentity: &UserIdentity{
//   		AccessKeyId: []*string{
//   			jsii.String("accessKeyId"),
//   		},
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		InvokedBy: []*string{
//   			jsii.String("invokedBy"),
//   		},
//   		PrincipalId: []*string{
//   			jsii.String("principalId"),
//   		},
//   		SessionContext: &SessionContext{
//   			Attributes: &Attributes{
//   				CreationDate: []*string{
//   					jsii.String("creationDate"),
//   				},
//   				MfaAuthenticated: []*string{
//   					jsii.String("mfaAuthenticated"),
//   				},
//   			},
//   			SessionIssuer: &SessionIssuer{
//   				AccountId: []*string{
//   					jsii.String("accountId"),
//   				},
//   				Arn: []*string{
//   					jsii.String("arn"),
//   				},
//   				PrincipalId: []*string{
//   					jsii.String("principalId"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   				UserName: []*string{
//   					jsii.String("userName"),
//   				},
//   			},
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps struct {
	// awsRegion property.
	//
	// Specify an array of string values to match this event if the actual value of awsRegion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsRegion *[]*string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// errorCode property.
	//
	// Specify an array of string values to match this event if the actual value of errorCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorCode *[]*string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// errorMessage property.
	//
	// Specify an array of string values to match this event if the actual value of errorMessage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorMessage *[]*string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// eventID property.
	//
	// Specify an array of string values to match this event if the actual value of eventID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventId *[]*string `field:"optional" json:"eventId" yaml:"eventId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// eventName property.
	//
	// Specify an array of string values to match this event if the actual value of eventName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventName *[]*string `field:"optional" json:"eventName" yaml:"eventName"`
	// eventSource property.
	//
	// Specify an array of string values to match this event if the actual value of eventSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventSource *[]*string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// eventTime property.
	//
	// Specify an array of string values to match this event if the actual value of eventTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventTime *[]*string `field:"optional" json:"eventTime" yaml:"eventTime"`
	// eventType property.
	//
	// Specify an array of string values to match this event if the actual value of eventType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// eventVersion property.
	//
	// Specify an array of string values to match this event if the actual value of eventVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventVersion *[]*string `field:"optional" json:"eventVersion" yaml:"eventVersion"`
	// requestID property.
	//
	// Specify an array of string values to match this event if the actual value of requestID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// requestParameters property.
	//
	// Specify an array of string values to match this event if the actual value of requestParameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestParameters *AWSAPICallViaCloudTrail_RequestParameters `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// responseElements property.
	//
	// Specify an array of string values to match this event if the actual value of responseElements is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResponseElements *AWSAPICallViaCloudTrail_ResponseElements `field:"optional" json:"responseElements" yaml:"responseElements"`
	// sourceIPAddress property.
	//
	// Specify an array of string values to match this event if the actual value of sourceIPAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIpAddress *[]*string `field:"optional" json:"sourceIpAddress" yaml:"sourceIpAddress"`
	// userAgent property.
	//
	// Specify an array of string values to match this event if the actual value of userAgent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserAgent *[]*string `field:"optional" json:"userAgent" yaml:"userAgent"`
	// userIdentity property.
	//
	// Specify an array of string values to match this event if the actual value of userIdentity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserIdentity *AWSAPICallViaCloudTrail_UserIdentity `field:"optional" json:"userIdentity" yaml:"userIdentity"`
}

