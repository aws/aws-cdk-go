package awsbatch


// Properties for defining a `CfnJobDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var labels interface{}
//   var limits interface{}
//   var options interface{}
//   var parameters interface{}
//   var requests interface{}
//   var tags interface{}
//
//   cfnJobDefinitionProps := &CfnJobDefinitionProps{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ContainerProperties: &ContainerPropertiesProperty{
//   		Image: jsii.String("image"),
//
//   		// the properties below are optional
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		Environment: []interface{}{
//   			&EnvironmentProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		EphemeralStorage: &EphemeralStorageProperty{
//   			SizeInGiB: jsii.Number(123),
//   		},
//   		ExecutionRoleArn: jsii.String("executionRoleArn"),
//   		FargatePlatformConfiguration: &FargatePlatformConfigurationProperty{
//   			PlatformVersion: jsii.String("platformVersion"),
//   		},
//   		InstanceType: jsii.String("instanceType"),
//   		JobRoleArn: jsii.String("jobRoleArn"),
//   		LinuxParameters: &LinuxParametersProperty{
//   			Devices: []interface{}{
//   				&DeviceProperty{
//   					ContainerPath: jsii.String("containerPath"),
//   					HostPath: jsii.String("hostPath"),
//   					Permissions: []*string{
//   						jsii.String("permissions"),
//   					},
//   				},
//   			},
//   			InitProcessEnabled: jsii.Boolean(false),
//   			MaxSwap: jsii.Number(123),
//   			SharedMemorySize: jsii.Number(123),
//   			Swappiness: jsii.Number(123),
//   			Tmpfs: []interface{}{
//   				&TmpfsProperty{
//   					ContainerPath: jsii.String("containerPath"),
//   					Size: jsii.Number(123),
//
//   					// the properties below are optional
//   					MountOptions: []*string{
//   						jsii.String("mountOptions"),
//   					},
//   				},
//   			},
//   		},
//   		LogConfiguration: &LogConfigurationProperty{
//   			LogDriver: jsii.String("logDriver"),
//
//   			// the properties below are optional
//   			Options: options,
//   			SecretOptions: []interface{}{
//   				&SecretProperty{
//   					Name: jsii.String("name"),
//   					ValueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   		},
//   		Memory: jsii.Number(123),
//   		MountPoints: []interface{}{
//   			&MountPointsProperty{
//   				ContainerPath: jsii.String("containerPath"),
//   				ReadOnly: jsii.Boolean(false),
//   				SourceVolume: jsii.String("sourceVolume"),
//   			},
//   		},
//   		NetworkConfiguration: &NetworkConfigurationProperty{
//   			AssignPublicIp: jsii.String("assignPublicIp"),
//   		},
//   		Privileged: jsii.Boolean(false),
//   		ReadonlyRootFilesystem: jsii.Boolean(false),
//   		ResourceRequirements: []interface{}{
//   			&ResourceRequirementProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Secrets: []interface{}{
//   			&SecretProperty{
//   				Name: jsii.String("name"),
//   				ValueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   		Ulimits: []interface{}{
//   			&UlimitProperty{
//   				HardLimit: jsii.Number(123),
//   				Name: jsii.String("name"),
//   				SoftLimit: jsii.Number(123),
//   			},
//   		},
//   		User: jsii.String("user"),
//   		Vcpus: jsii.Number(123),
//   		Volumes: []interface{}{
//   			&VolumesProperty{
//   				EfsVolumeConfiguration: &EfsVolumeConfigurationProperty{
//   					FileSystemId: jsii.String("fileSystemId"),
//
//   					// the properties below are optional
//   					AuthorizationConfig: &AuthorizationConfigProperty{
//   						AccessPointId: jsii.String("accessPointId"),
//   						Iam: jsii.String("iam"),
//   					},
//   					RootDirectory: jsii.String("rootDirectory"),
//   					TransitEncryption: jsii.String("transitEncryption"),
//   					TransitEncryptionPort: jsii.Number(123),
//   				},
//   				Host: &VolumesHostProperty{
//   					SourcePath: jsii.String("sourcePath"),
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	EksProperties: &EksPropertiesProperty{
//   		PodProperties: &PodPropertiesProperty{
//   			Containers: []interface{}{
//   				&EksContainerProperty{
//   					Image: jsii.String("image"),
//
//   					// the properties below are optional
//   					Args: []*string{
//   						jsii.String("args"),
//   					},
//   					Command: []*string{
//   						jsii.String("command"),
//   					},
//   					Env: []interface{}{
//   						&EksContainerEnvironmentVariableProperty{
//   							Name: jsii.String("name"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					ImagePullPolicy: jsii.String("imagePullPolicy"),
//   					Name: jsii.String("name"),
//   					Resources: &ResourcesProperty{
//   						Limits: limits,
//   						Requests: requests,
//   					},
//   					SecurityContext: &SecurityContextProperty{
//   						Privileged: jsii.Boolean(false),
//   						ReadOnlyRootFilesystem: jsii.Boolean(false),
//   						RunAsGroup: jsii.Number(123),
//   						RunAsNonRoot: jsii.Boolean(false),
//   						RunAsUser: jsii.Number(123),
//   					},
//   					VolumeMounts: []interface{}{
//   						&EksContainerVolumeMountProperty{
//   							MountPath: jsii.String("mountPath"),
//   							Name: jsii.String("name"),
//   							ReadOnly: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   			},
//   			DnsPolicy: jsii.String("dnsPolicy"),
//   			HostNetwork: jsii.Boolean(false),
//   			Metadata: &MetadataProperty{
//   				Labels: labels,
//   			},
//   			ServiceAccountName: jsii.String("serviceAccountName"),
//   			Volumes: []interface{}{
//   				&EksVolumeProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					EmptyDir: &EmptyDirProperty{
//   						Medium: jsii.String("medium"),
//   						SizeLimit: jsii.String("sizeLimit"),
//   					},
//   					HostPath: &HostPathProperty{
//   						Path: jsii.String("path"),
//   					},
//   					Secret: &EksSecretProperty{
//   						SecretName: jsii.String("secretName"),
//
//   						// the properties below are optional
//   						Optional: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	JobDefinitionName: jsii.String("jobDefinitionName"),
//   	NodeProperties: &NodePropertiesProperty{
//   		MainNode: jsii.Number(123),
//   		NodeRangeProperties: []interface{}{
//   			&NodeRangePropertyProperty{
//   				TargetNodes: jsii.String("targetNodes"),
//
//   				// the properties below are optional
//   				Container: &ContainerPropertiesProperty{
//   					Image: jsii.String("image"),
//
//   					// the properties below are optional
//   					Command: []*string{
//   						jsii.String("command"),
//   					},
//   					Environment: []interface{}{
//   						&EnvironmentProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					EphemeralStorage: &EphemeralStorageProperty{
//   						SizeInGiB: jsii.Number(123),
//   					},
//   					ExecutionRoleArn: jsii.String("executionRoleArn"),
//   					FargatePlatformConfiguration: &FargatePlatformConfigurationProperty{
//   						PlatformVersion: jsii.String("platformVersion"),
//   					},
//   					InstanceType: jsii.String("instanceType"),
//   					JobRoleArn: jsii.String("jobRoleArn"),
//   					LinuxParameters: &LinuxParametersProperty{
//   						Devices: []interface{}{
//   							&DeviceProperty{
//   								ContainerPath: jsii.String("containerPath"),
//   								HostPath: jsii.String("hostPath"),
//   								Permissions: []*string{
//   									jsii.String("permissions"),
//   								},
//   							},
//   						},
//   						InitProcessEnabled: jsii.Boolean(false),
//   						MaxSwap: jsii.Number(123),
//   						SharedMemorySize: jsii.Number(123),
//   						Swappiness: jsii.Number(123),
//   						Tmpfs: []interface{}{
//   							&TmpfsProperty{
//   								ContainerPath: jsii.String("containerPath"),
//   								Size: jsii.Number(123),
//
//   								// the properties below are optional
//   								MountOptions: []*string{
//   									jsii.String("mountOptions"),
//   								},
//   							},
//   						},
//   					},
//   					LogConfiguration: &LogConfigurationProperty{
//   						LogDriver: jsii.String("logDriver"),
//
//   						// the properties below are optional
//   						Options: options,
//   						SecretOptions: []interface{}{
//   							&SecretProperty{
//   								Name: jsii.String("name"),
//   								ValueFrom: jsii.String("valueFrom"),
//   							},
//   						},
//   					},
//   					Memory: jsii.Number(123),
//   					MountPoints: []interface{}{
//   						&MountPointsProperty{
//   							ContainerPath: jsii.String("containerPath"),
//   							ReadOnly: jsii.Boolean(false),
//   							SourceVolume: jsii.String("sourceVolume"),
//   						},
//   					},
//   					NetworkConfiguration: &NetworkConfigurationProperty{
//   						AssignPublicIp: jsii.String("assignPublicIp"),
//   					},
//   					Privileged: jsii.Boolean(false),
//   					ReadonlyRootFilesystem: jsii.Boolean(false),
//   					ResourceRequirements: []interface{}{
//   						&ResourceRequirementProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Secrets: []interface{}{
//   						&SecretProperty{
//   							Name: jsii.String("name"),
//   							ValueFrom: jsii.String("valueFrom"),
//   						},
//   					},
//   					Ulimits: []interface{}{
//   						&UlimitProperty{
//   							HardLimit: jsii.Number(123),
//   							Name: jsii.String("name"),
//   							SoftLimit: jsii.Number(123),
//   						},
//   					},
//   					User: jsii.String("user"),
//   					Vcpus: jsii.Number(123),
//   					Volumes: []interface{}{
//   						&VolumesProperty{
//   							EfsVolumeConfiguration: &EfsVolumeConfigurationProperty{
//   								FileSystemId: jsii.String("fileSystemId"),
//
//   								// the properties below are optional
//   								AuthorizationConfig: &AuthorizationConfigProperty{
//   									AccessPointId: jsii.String("accessPointId"),
//   									Iam: jsii.String("iam"),
//   								},
//   								RootDirectory: jsii.String("rootDirectory"),
//   								TransitEncryption: jsii.String("transitEncryption"),
//   								TransitEncryptionPort: jsii.Number(123),
//   							},
//   							Host: &VolumesHostProperty{
//   								SourcePath: jsii.String("sourcePath"),
//   							},
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		NumNodes: jsii.Number(123),
//   	},
//   	Parameters: parameters,
//   	PlatformCapabilities: []*string{
//   		jsii.String("platformCapabilities"),
//   	},
//   	PropagateTags: jsii.Boolean(false),
//   	RetryStrategy: &RetryStrategyProperty{
//   		Attempts: jsii.Number(123),
//   		EvaluateOnExit: []interface{}{
//   			&EvaluateOnExitProperty{
//   				Action: jsii.String("action"),
//
//   				// the properties below are optional
//   				OnExitCode: jsii.String("onExitCode"),
//   				OnReason: jsii.String("onReason"),
//   				OnStatusReason: jsii.String("onStatusReason"),
//   			},
//   		},
//   	},
//   	SchedulingPriority: jsii.Number(123),
//   	Tags: tags,
//   	Timeout: &TimeoutProperty{
//   		AttemptDurationSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html
//
type CfnJobDefinitionProps struct {
	// The type of job definition.
	//
	// For more information about multi-node parallel jobs, see [Creating a multi-node parallel job definition](https://docs.aws.amazon.com/batch/latest/userguide/multi-node-job-def.html) in the *AWS Batch User Guide* .
	//
	// > If the job is run on Fargate resources, then `multinode` isn't supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// An object with various properties specific to Amazon ECS based jobs.
	//
	// Valid values are `containerProperties` , `eksProperties` , and `nodeProperties` . Only one can be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-containerproperties
	//
	ContainerProperties interface{} `field:"optional" json:"containerProperties" yaml:"containerProperties"`
	// An object with various properties that are specific to Amazon EKS based jobs.
	//
	// Valid values are `containerProperties` , `eksProperties` , and `nodeProperties` . Only one can be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-eksproperties
	//
	EksProperties interface{} `field:"optional" json:"eksProperties" yaml:"eksProperties"`
	// The name of the job definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-jobdefinitionname
	//
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// An object with various properties that are specific to multi-node parallel jobs.
	//
	// Valid values are `containerProperties` , `eksProperties` , and `nodeProperties` . Only one can be specified.
	//
	// > If the job runs on Fargate resources, don't specify `nodeProperties` . Use `containerProperties` instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-nodeproperties
	//
	NodeProperties interface{} `field:"optional" json:"nodeProperties" yaml:"nodeProperties"`
	// Default parameters or parameter substitution placeholders that are set in the job definition.
	//
	// Parameters are specified as a key-value pair mapping. Parameters in a `SubmitJob` request override any corresponding parameter defaults from the job definition. For more information about specifying parameters, see [Job definition parameters](https://docs.aws.amazon.com/batch/latest/userguide/job_definition_parameters.html) in the *AWS Batch User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The platform capabilities required by the job definition.
	//
	// If no value is specified, it defaults to `EC2` . Jobs run on Fargate resources specify `FARGATE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-platformcapabilities
	//
	PlatformCapabilities *[]*string `field:"optional" json:"platformCapabilities" yaml:"platformCapabilities"`
	// Specifies whether to propagate the tags from the job or job definition to the corresponding Amazon ECS task.
	//
	// If no value is specified, the tags aren't propagated. Tags can only be propagated to the tasks when the tasks are created. For tags with the same name, job tags are given priority over job definitions tags. If the total number of combined tags from the job and job definition is over 50, the job is moved to the `FAILED` state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-propagatetags
	//
	PropagateTags interface{} `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The retry strategy to use for failed jobs that are submitted with this job definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-retrystrategy
	//
	RetryStrategy interface{} `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	// The scheduling priority of the job definition.
	//
	// This only affects jobs in job queues with a fair share policy. Jobs with a higher scheduling priority are scheduled before jobs with a lower scheduling priority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-schedulingpriority
	//
	SchedulingPriority *float64 `field:"optional" json:"schedulingPriority" yaml:"schedulingPriority"`
	// The tags that are applied to the job definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The timeout time for jobs that are submitted with this job definition.
	//
	// After the amount of time you specify passes, AWS Batch terminates your jobs if they aren't finished.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-timeout
	//
	Timeout interface{} `field:"optional" json:"timeout" yaml:"timeout"`
}

