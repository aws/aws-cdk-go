package previewawscodebuildevents


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseElements := &ResponseElements{
//   	BuildProperty: &Build{
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		Artifacts: &Artifacts{
//   			EncryptionDisabled: []*string{
//   				jsii.String("encryptionDisabled"),
//   			},
//   			Location: []*string{
//   				jsii.String("location"),
//   			},
//   		},
//   		BuildComplete: []*string{
//   			jsii.String("buildComplete"),
//   		},
//   		BuildStatus: []*string{
//   			jsii.String("buildStatus"),
//   		},
//   		Cache: &Cache{
//   			Location: []*string{
//   				jsii.String("location"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   		CurrentPhase: []*string{
//   			jsii.String("currentPhase"),
//   		},
//   		EncryptionKey: []*string{
//   			jsii.String("encryptionKey"),
//   		},
//   		EndTime: []*string{
//   			jsii.String("endTime"),
//   		},
//   		Environment: &Environment{
//   			Certificate: []*string{
//   				jsii.String("certificate"),
//   			},
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
//   		Id: []*string{
//   			jsii.String("id"),
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
//   		Phases: []BuildPhase{
//   			&BuildPhase{
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
//   		ProjectName: []*string{
//   			jsii.String("projectName"),
//   		},
//   		QueuedTimeoutInMinutes: []*string{
//   			jsii.String("queuedTimeoutInMinutes"),
//   		},
//   		ResolvedSourceVersion: []*string{
//   			jsii.String("resolvedSourceVersion"),
//   		},
//   		ServiceRole: []*string{
//   			jsii.String("serviceRole"),
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
//   			InsecureSsl: []*string{
//   				jsii.String("insecureSsl"),
//   			},
//   			Location: []*string{
//   				jsii.String("location"),
//   			},
//   			ReportBuildStatus: []*string{
//   				jsii.String("reportBuildStatus"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   		SourceVersion: []*string{
//   			jsii.String("sourceVersion"),
//   		},
//   		StartTime: []*string{
//   			jsii.String("startTime"),
//   		},
//   		TimeoutInMinutes: []*string{
//   			jsii.String("timeoutInMinutes"),
//   		},
//   		VpcConfig: &VpcConfig{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   			VpcId: []*string{
//   				jsii.String("vpcId"),
//   			},
//   		},
//   	},
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	Project: &Project{
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		Artifacts: &Artifacts1{
//   			EncryptionDisabled: []*string{
//   				jsii.String("encryptionDisabled"),
//   			},
//   			Location: []*string{
//   				jsii.String("location"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			NamespaceType: []*string{
//   				jsii.String("namespaceType"),
//   			},
//   			Packaging: []*string{
//   				jsii.String("packaging"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   		Badge: &Badge{
//   			BadgeEnabled: []*string{
//   				jsii.String("badgeEnabled"),
//   			},
//   			BadgeRequestUrl: []*string{
//   				jsii.String("badgeRequestUrl"),
//   			},
//   		},
//   		Cache: &Cache{
//   			Location: []*string{
//   				jsii.String("location"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   		Created: []*string{
//   			jsii.String("created"),
//   		},
//   		Description: []*string{
//   			jsii.String("description"),
//   		},
//   		EncryptionKey: []*string{
//   			jsii.String("encryptionKey"),
//   		},
//   		Environment: &Environment1{
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
//   		LastModified: []*string{
//   			jsii.String("lastModified"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		QueuedTimeoutInMinutes: []*string{
//   			jsii.String("queuedTimeoutInMinutes"),
//   		},
//   		ServiceRole: []*string{
//   			jsii.String("serviceRole"),
//   		},
//   		Source: &Source1{
//   			Auth: &Auth{
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   			Buildspec: []*string{
//   				jsii.String("buildspec"),
//   			},
//   			GitCloneDepth: []*string{
//   				jsii.String("gitCloneDepth"),
//   			},
//   			GitSubmodulesConfig: &GitSubmodulesConfig{
//   				FetchSubmodules: []*string{
//   					jsii.String("fetchSubmodules"),
//   				},
//   			},
//   			InsecureSsl: []*string{
//   				jsii.String("insecureSsl"),
//   			},
//   			Location: []*string{
//   				jsii.String("location"),
//   			},
//   			ReportBuildStatus: []*string{
//   				jsii.String("reportBuildStatus"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   		SourceVersion: []*string{
//   			jsii.String("sourceVersion"),
//   		},
//   		Tags: []ProjectItem{
//   			&ProjectItem{
//   				Key: []*string{
//   					jsii.String("key"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   		TimeoutInMinutes: []*string{
//   			jsii.String("timeoutInMinutes"),
//   		},
//   	},
//   	Webhook: &Webhook{
//   		BranchFilter: []*string{
//   			jsii.String("branchFilter"),
//   		},
//   		LastModifiedSecret: []*string{
//   			jsii.String("lastModifiedSecret"),
//   		},
//   		PayloadUrl: []*string{
//   			jsii.String("payloadUrl"),
//   		},
//   		Url: []*string{
//   			jsii.String("url"),
//   		},
//   	},
//   	WebhookDeletedStatus: []*string{
//   		jsii.String("webhookDeletedStatus"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ResponseElements struct {
	// build property.
	//
	// Specify an array of string values to match this event if the actual value of build is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BuildProperty *AWSAPICallViaCloudTrail_Build `field:"optional" json:"buildProperty" yaml:"buildProperty"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// project property.
	//
	// Specify an array of string values to match this event if the actual value of project is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Project *AWSAPICallViaCloudTrail_Project `field:"optional" json:"project" yaml:"project"`
	// webhook property.
	//
	// Specify an array of string values to match this event if the actual value of webhook is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Webhook *AWSAPICallViaCloudTrail_Webhook `field:"optional" json:"webhook" yaml:"webhook"`
	// webhookDeletedStatus property.
	//
	// Specify an array of string values to match this event if the actual value of webhookDeletedStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	WebhookDeletedStatus *[]*string `field:"optional" json:"webhookDeletedStatus" yaml:"webhookDeletedStatus"`
}

