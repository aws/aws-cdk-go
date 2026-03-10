package previewawscodebuildevents


// Type definition for Build.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   build := &Build{
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   	Artifacts: &Artifacts{
//   		EncryptionDisabled: []*string{
//   			jsii.String("encryptionDisabled"),
//   		},
//   		Location: []*string{
//   			jsii.String("location"),
//   		},
//   	},
//   	BuildComplete: []*string{
//   		jsii.String("buildComplete"),
//   	},
//   	BuildStatus: []*string{
//   		jsii.String("buildStatus"),
//   	},
//   	Cache: &Cache{
//   		Location: []*string{
//   			jsii.String("location"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	CurrentPhase: []*string{
//   		jsii.String("currentPhase"),
//   	},
//   	EncryptionKey: []*string{
//   		jsii.String("encryptionKey"),
//   	},
//   	EndTime: []*string{
//   		jsii.String("endTime"),
//   	},
//   	Environment: &Environment{
//   		Certificate: []*string{
//   			jsii.String("certificate"),
//   		},
//   		ComputeType: []*string{
//   			jsii.String("computeType"),
//   		},
//   		EnvironmentVariables: []EnvironmentItem{
//   			&EnvironmentItem{
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   		Image: []*string{
//   			jsii.String("image"),
//   		},
//   		ImagePullCredentialsType: []*string{
//   			jsii.String("imagePullCredentialsType"),
//   		},
//   		PrivilegedMode: []*string{
//   			jsii.String("privilegedMode"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	Id: []*string{
//   		jsii.String("id"),
//   	},
//   	Initiator: []*string{
//   		jsii.String("initiator"),
//   	},
//   	Logs: &Logs{
//   		DeepLink: []*string{
//   			jsii.String("deepLink"),
//   		},
//   		GroupName: []*string{
//   			jsii.String("groupName"),
//   		},
//   		StreamName: []*string{
//   			jsii.String("streamName"),
//   		},
//   	},
//   	Phases: []BuildPhase{
//   		&BuildPhase{
//   			DurationInSeconds: []*string{
//   				jsii.String("durationInSeconds"),
//   			},
//   			EndTime: []*string{
//   				jsii.String("endTime"),
//   			},
//   			PhaseContext: []*string{
//   				jsii.String("phaseContext"),
//   			},
//   			PhaseStatus: []*string{
//   				jsii.String("phaseStatus"),
//   			},
//   			PhaseType: []*string{
//   				jsii.String("phaseType"),
//   			},
//   			StartTime: []*string{
//   				jsii.String("startTime"),
//   			},
//   		},
//   	},
//   	ProjectName: []*string{
//   		jsii.String("projectName"),
//   	},
//   	QueuedTimeoutInMinutes: []*string{
//   		jsii.String("queuedTimeoutInMinutes"),
//   	},
//   	ResolvedSourceVersion: []*string{
//   		jsii.String("resolvedSourceVersion"),
//   	},
//   	ServiceRole: []*string{
//   		jsii.String("serviceRole"),
//   	},
//   	Source: &Source{
//   		Auth: &Auth{
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   		Buildspec: []*string{
//   			jsii.String("buildspec"),
//   		},
//   		InsecureSsl: []*string{
//   			jsii.String("insecureSsl"),
//   		},
//   		Location: []*string{
//   			jsii.String("location"),
//   		},
//   		ReportBuildStatus: []*string{
//   			jsii.String("reportBuildStatus"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	SourceVersion: []*string{
//   		jsii.String("sourceVersion"),
//   	},
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   	TimeoutInMinutes: []*string{
//   		jsii.String("timeoutInMinutes"),
//   	},
//   	VpcConfig: &VpcConfig{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		VpcId: []*string{
//   			jsii.String("vpcId"),
//   		},
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Build struct {
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
	// artifacts property.
	//
	// Specify an array of string values to match this event if the actual value of artifacts is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Artifacts *AWSAPICallViaCloudTrail_Artifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	// buildComplete property.
	//
	// Specify an array of string values to match this event if the actual value of buildComplete is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BuildComplete *[]*string `field:"optional" json:"buildComplete" yaml:"buildComplete"`
	// buildStatus property.
	//
	// Specify an array of string values to match this event if the actual value of buildStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BuildStatus *[]*string `field:"optional" json:"buildStatus" yaml:"buildStatus"`
	// cache property.
	//
	// Specify an array of string values to match this event if the actual value of cache is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Cache *AWSAPICallViaCloudTrail_Cache `field:"optional" json:"cache" yaml:"cache"`
	// currentPhase property.
	//
	// Specify an array of string values to match this event if the actual value of currentPhase is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CurrentPhase *[]*string `field:"optional" json:"currentPhase" yaml:"currentPhase"`
	// encryptionKey property.
	//
	// Specify an array of string values to match this event if the actual value of encryptionKey is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EncryptionKey *[]*string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// endTime property.
	//
	// Specify an array of string values to match this event if the actual value of endTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTime *[]*string `field:"optional" json:"endTime" yaml:"endTime"`
	// environment property.
	//
	// Specify an array of string values to match this event if the actual value of environment is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Environment *AWSAPICallViaCloudTrail_Environment `field:"optional" json:"environment" yaml:"environment"`
	// id property.
	//
	// Specify an array of string values to match this event if the actual value of id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
	// initiator property.
	//
	// Specify an array of string values to match this event if the actual value of initiator is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Initiator *[]*string `field:"optional" json:"initiator" yaml:"initiator"`
	// logs property.
	//
	// Specify an array of string values to match this event if the actual value of logs is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Logs *AWSAPICallViaCloudTrail_Logs `field:"optional" json:"logs" yaml:"logs"`
	// phases property.
	//
	// Specify an array of string values to match this event if the actual value of phases is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Phases *[]*AWSAPICallViaCloudTrail_BuildPhase `field:"optional" json:"phases" yaml:"phases"`
	// projectName property.
	//
	// Specify an array of string values to match this event if the actual value of projectName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProjectName *[]*string `field:"optional" json:"projectName" yaml:"projectName"`
	// queuedTimeoutInMinutes property.
	//
	// Specify an array of string values to match this event if the actual value of queuedTimeoutInMinutes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	QueuedTimeoutInMinutes *[]*string `field:"optional" json:"queuedTimeoutInMinutes" yaml:"queuedTimeoutInMinutes"`
	// resolvedSourceVersion property.
	//
	// Specify an array of string values to match this event if the actual value of resolvedSourceVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResolvedSourceVersion *[]*string `field:"optional" json:"resolvedSourceVersion" yaml:"resolvedSourceVersion"`
	// serviceRole property.
	//
	// Specify an array of string values to match this event if the actual value of serviceRole is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ServiceRole *[]*string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// source property.
	//
	// Specify an array of string values to match this event if the actual value of source is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Source *AWSAPICallViaCloudTrail_Source `field:"optional" json:"source" yaml:"source"`
	// sourceVersion property.
	//
	// Specify an array of string values to match this event if the actual value of sourceVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceVersion *[]*string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
	// startTime property.
	//
	// Specify an array of string values to match this event if the actual value of startTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
	// timeoutInMinutes property.
	//
	// Specify an array of string values to match this event if the actual value of timeoutInMinutes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TimeoutInMinutes *[]*string `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
	// vpcConfig property.
	//
	// Specify an array of string values to match this event if the actual value of vpcConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VpcConfig *AWSAPICallViaCloudTrail_VpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

