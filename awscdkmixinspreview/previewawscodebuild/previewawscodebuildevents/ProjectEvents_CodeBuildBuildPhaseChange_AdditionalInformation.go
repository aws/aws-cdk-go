package previewawscodebuildevents


// Type definition for Additional-information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   additionalInformation := &AdditionalInformation{
//   	Artifact: &Artifact{
//   		Location: []*string{
//   			jsii.String("location"),
//   		},
//   		Md5Sum: []*string{
//   			jsii.String("md5Sum"),
//   		},
//   		Sha256Sum: []*string{
//   			jsii.String("sha256Sum"),
//   		},
//   	},
//   	BuildComplete: []*string{
//   		jsii.String("buildComplete"),
//   	},
//   	BuildStartTime: []*string{
//   		jsii.String("buildStartTime"),
//   	},
//   	Cache: &Cache{
//   		Location: []*string{
//   			jsii.String("location"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	Environment: &Environment{
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
//   	NetworkInterface: &NetworkInterface{
//   		EniId: []*string{
//   			jsii.String("eniId"),
//   		},
//   		SubnetId: []*string{
//   			jsii.String("subnetId"),
//   		},
//   	},
//   	Phases: []AdditionalInformationItem{
//   		&AdditionalInformationItem{
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
//   	QueuedTimeoutInMinutes: []*string{
//   		jsii.String("queuedTimeoutInMinutes"),
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
//   		Location: []*string{
//   			jsii.String("location"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	SourceVersion: []*string{
//   		jsii.String("sourceVersion"),
//   	},
//   	TimeoutInMinutes: []*string{
//   		jsii.String("timeoutInMinutes"),
//   	},
//   	VpcConfig: &VpcConfig{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []VpcConfigItem{
//   			&VpcConfigItem{
//   				BuildFleetAz: []*string{
//   					jsii.String("buildFleetAz"),
//   				},
//   				CustomerAz: []*string{
//   					jsii.String("customerAz"),
//   				},
//   				SubnetId: []*string{
//   					jsii.String("subnetId"),
//   				},
//   			},
//   		},
//   		VpcId: []*string{
//   			jsii.String("vpcId"),
//   		},
//   	},
//   }
//
// Experimental.
type ProjectEvents_CodeBuildBuildPhaseChange_AdditionalInformation struct {
	// artifact property.
	//
	// Specify an array of string values to match this event if the actual value of artifact is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Artifact *ProjectEvents_CodeBuildBuildPhaseChange_Artifact `field:"optional" json:"artifact" yaml:"artifact"`
	// build-complete property.
	//
	// Specify an array of string values to match this event if the actual value of build-complete is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BuildComplete *[]*string `field:"optional" json:"buildComplete" yaml:"buildComplete"`
	// build-start-time property.
	//
	// Specify an array of string values to match this event if the actual value of build-start-time is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BuildStartTime *[]*string `field:"optional" json:"buildStartTime" yaml:"buildStartTime"`
	// cache property.
	//
	// Specify an array of string values to match this event if the actual value of cache is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Cache *ProjectEvents_CodeBuildBuildPhaseChange_Cache `field:"optional" json:"cache" yaml:"cache"`
	// environment property.
	//
	// Specify an array of string values to match this event if the actual value of environment is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Environment *ProjectEvents_CodeBuildBuildPhaseChange_Environment `field:"optional" json:"environment" yaml:"environment"`
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
	Logs *ProjectEvents_CodeBuildBuildPhaseChange_Logs `field:"optional" json:"logs" yaml:"logs"`
	// network-interface property.
	//
	// Specify an array of string values to match this event if the actual value of network-interface is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkInterface *ProjectEvents_CodeBuildBuildPhaseChange_NetworkInterface `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// phases property.
	//
	// Specify an array of string values to match this event if the actual value of phases is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Phases *[]*ProjectEvents_CodeBuildBuildPhaseChange_AdditionalInformationItem `field:"optional" json:"phases" yaml:"phases"`
	// queued-timeout-in-minutes property.
	//
	// Specify an array of string values to match this event if the actual value of queued-timeout-in-minutes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	QueuedTimeoutInMinutes *[]*string `field:"optional" json:"queuedTimeoutInMinutes" yaml:"queuedTimeoutInMinutes"`
	// source property.
	//
	// Specify an array of string values to match this event if the actual value of source is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Source *ProjectEvents_CodeBuildBuildPhaseChange_Source `field:"optional" json:"source" yaml:"source"`
	// source-version property.
	//
	// Specify an array of string values to match this event if the actual value of source-version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceVersion *[]*string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
	// timeout-in-minutes property.
	//
	// Specify an array of string values to match this event if the actual value of timeout-in-minutes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TimeoutInMinutes *[]*string `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
	// vpc-config property.
	//
	// Specify an array of string values to match this event if the actual value of vpc-config is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VpcConfig *ProjectEvents_CodeBuildBuildPhaseChange_VpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

