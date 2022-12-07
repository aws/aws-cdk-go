package awsbatch


// Properties for defining a `CfnJobDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var limits interface{}
//   var options interface{}
//   var parameters interface{}
//   var requests interface{}
//   var tags interface{}
//
//   cfnJobDefinitionProps := &cfnJobDefinitionProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	containerProperties: &containerPropertiesProperty{
//   		image: jsii.String("image"),
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		environment: []interface{}{
//   			&environmentProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		executionRoleArn: jsii.String("executionRoleArn"),
//   		fargatePlatformConfiguration: &fargatePlatformConfigurationProperty{
//   			platformVersion: jsii.String("platformVersion"),
//   		},
//   		instanceType: jsii.String("instanceType"),
//   		jobRoleArn: jsii.String("jobRoleArn"),
//   		linuxParameters: &linuxParametersProperty{
//   			devices: []interface{}{
//   				&deviceProperty{
//   					containerPath: jsii.String("containerPath"),
//   					hostPath: jsii.String("hostPath"),
//   					permissions: []*string{
//   						jsii.String("permissions"),
//   					},
//   				},
//   			},
//   			initProcessEnabled: jsii.Boolean(false),
//   			maxSwap: jsii.Number(123),
//   			sharedMemorySize: jsii.Number(123),
//   			swappiness: jsii.Number(123),
//   			tmpfs: []interface{}{
//   				&tmpfsProperty{
//   					containerPath: jsii.String("containerPath"),
//   					size: jsii.Number(123),
//
//   					// the properties below are optional
//   					mountOptions: []*string{
//   						jsii.String("mountOptions"),
//   					},
//   				},
//   			},
//   		},
//   		logConfiguration: &logConfigurationProperty{
//   			logDriver: jsii.String("logDriver"),
//
//   			// the properties below are optional
//   			options: options,
//   			secretOptions: []interface{}{
//   				&secretProperty{
//   					name: jsii.String("name"),
//   					valueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   		},
//   		memory: jsii.Number(123),
//   		mountPoints: []interface{}{
//   			&mountPointsProperty{
//   				containerPath: jsii.String("containerPath"),
//   				readOnly: jsii.Boolean(false),
//   				sourceVolume: jsii.String("sourceVolume"),
//   			},
//   		},
//   		networkConfiguration: &networkConfigurationProperty{
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   		},
//   		privileged: jsii.Boolean(false),
//   		readonlyRootFilesystem: jsii.Boolean(false),
//   		resourceRequirements: []interface{}{
//   			&resourceRequirementProperty{
//   				type: jsii.String("type"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		secrets: []interface{}{
//   			&secretProperty{
//   				name: jsii.String("name"),
//   				valueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   		ulimits: []interface{}{
//   			&ulimitProperty{
//   				hardLimit: jsii.Number(123),
//   				name: jsii.String("name"),
//   				softLimit: jsii.Number(123),
//   			},
//   		},
//   		user: jsii.String("user"),
//   		vcpus: jsii.Number(123),
//   		volumes: []interface{}{
//   			&volumesProperty{
//   				efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   					fileSystemId: jsii.String("fileSystemId"),
//
//   					// the properties below are optional
//   					authorizationConfig: &authorizationConfigProperty{
//   						accessPointId: jsii.String("accessPointId"),
//   						iam: jsii.String("iam"),
//   					},
//   					rootDirectory: jsii.String("rootDirectory"),
//   					transitEncryption: jsii.String("transitEncryption"),
//   					transitEncryptionPort: jsii.Number(123),
//   				},
//   				host: &volumesHostProperty{
//   					sourcePath: jsii.String("sourcePath"),
//   				},
//   				name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	eksProperties: &eksPropertiesProperty{
//   		podProperties: &podPropertiesProperty{
//   			containers: []interface{}{
//   				&eksContainerProperty{
//   					image: jsii.String("image"),
//
//   					// the properties below are optional
//   					args: []*string{
//   						jsii.String("args"),
//   					},
//   					command: []*string{
//   						jsii.String("command"),
//   					},
//   					env: []interface{}{
//   						&eksContainerEnvironmentVariableProperty{
//   							name: jsii.String("name"),
//
//   							// the properties below are optional
//   							value: jsii.String("value"),
//   						},
//   					},
//   					imagePullPolicy: jsii.String("imagePullPolicy"),
//   					name: jsii.String("name"),
//   					resources: &resourcesProperty{
//   						limits: limits,
//   						requests: requests,
//   					},
//   					securityContext: &securityContextProperty{
//   						privileged: jsii.Boolean(false),
//   						readOnlyRootFilesystem: jsii.Boolean(false),
//   						runAsGroup: jsii.Number(123),
//   						runAsNonRoot: jsii.Boolean(false),
//   						runAsUser: jsii.Number(123),
//   					},
//   					volumeMounts: []interface{}{
//   						&eksContainerVolumeMountProperty{
//   							mountPath: jsii.String("mountPath"),
//   							name: jsii.String("name"),
//   							readOnly: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   			},
//   			dnsPolicy: jsii.String("dnsPolicy"),
//   			hostNetwork: jsii.Boolean(false),
//   			serviceAccountName: jsii.String("serviceAccountName"),
//   			volumes: []interface{}{
//   				&eksVolumeProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					emptyDir: &emptyDirProperty{
//   						medium: jsii.String("medium"),
//   						sizeLimit: jsii.String("sizeLimit"),
//   					},
//   					hostPath: &hostPathProperty{
//   						path: jsii.String("path"),
//   					},
//   					secret: &secretProperty{
//   						name: jsii.String("name"),
//   						valueFrom: jsii.String("valueFrom"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	nodeProperties: &nodePropertiesProperty{
//   		mainNode: jsii.Number(123),
//   		nodeRangeProperties: []interface{}{
//   			&nodeRangePropertyProperty{
//   				targetNodes: jsii.String("targetNodes"),
//
//   				// the properties below are optional
//   				container: &containerPropertiesProperty{
//   					image: jsii.String("image"),
//
//   					// the properties below are optional
//   					command: []*string{
//   						jsii.String("command"),
//   					},
//   					environment: []interface{}{
//   						&environmentProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					executionRoleArn: jsii.String("executionRoleArn"),
//   					fargatePlatformConfiguration: &fargatePlatformConfigurationProperty{
//   						platformVersion: jsii.String("platformVersion"),
//   					},
//   					instanceType: jsii.String("instanceType"),
//   					jobRoleArn: jsii.String("jobRoleArn"),
//   					linuxParameters: &linuxParametersProperty{
//   						devices: []interface{}{
//   							&deviceProperty{
//   								containerPath: jsii.String("containerPath"),
//   								hostPath: jsii.String("hostPath"),
//   								permissions: []*string{
//   									jsii.String("permissions"),
//   								},
//   							},
//   						},
//   						initProcessEnabled: jsii.Boolean(false),
//   						maxSwap: jsii.Number(123),
//   						sharedMemorySize: jsii.Number(123),
//   						swappiness: jsii.Number(123),
//   						tmpfs: []interface{}{
//   							&tmpfsProperty{
//   								containerPath: jsii.String("containerPath"),
//   								size: jsii.Number(123),
//
//   								// the properties below are optional
//   								mountOptions: []*string{
//   									jsii.String("mountOptions"),
//   								},
//   							},
//   						},
//   					},
//   					logConfiguration: &logConfigurationProperty{
//   						logDriver: jsii.String("logDriver"),
//
//   						// the properties below are optional
//   						options: options,
//   						secretOptions: []interface{}{
//   							&secretProperty{
//   								name: jsii.String("name"),
//   								valueFrom: jsii.String("valueFrom"),
//   							},
//   						},
//   					},
//   					memory: jsii.Number(123),
//   					mountPoints: []interface{}{
//   						&mountPointsProperty{
//   							containerPath: jsii.String("containerPath"),
//   							readOnly: jsii.Boolean(false),
//   							sourceVolume: jsii.String("sourceVolume"),
//   						},
//   					},
//   					networkConfiguration: &networkConfigurationProperty{
//   						assignPublicIp: jsii.String("assignPublicIp"),
//   					},
//   					privileged: jsii.Boolean(false),
//   					readonlyRootFilesystem: jsii.Boolean(false),
//   					resourceRequirements: []interface{}{
//   						&resourceRequirementProperty{
//   							type: jsii.String("type"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					secrets: []interface{}{
//   						&secretProperty{
//   							name: jsii.String("name"),
//   							valueFrom: jsii.String("valueFrom"),
//   						},
//   					},
//   					ulimits: []interface{}{
//   						&ulimitProperty{
//   							hardLimit: jsii.Number(123),
//   							name: jsii.String("name"),
//   							softLimit: jsii.Number(123),
//   						},
//   					},
//   					user: jsii.String("user"),
//   					vcpus: jsii.Number(123),
//   					volumes: []interface{}{
//   						&volumesProperty{
//   							efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   								fileSystemId: jsii.String("fileSystemId"),
//
//   								// the properties below are optional
//   								authorizationConfig: &authorizationConfigProperty{
//   									accessPointId: jsii.String("accessPointId"),
//   									iam: jsii.String("iam"),
//   								},
//   								rootDirectory: jsii.String("rootDirectory"),
//   								transitEncryption: jsii.String("transitEncryption"),
//   								transitEncryptionPort: jsii.Number(123),
//   							},
//   							host: &volumesHostProperty{
//   								sourcePath: jsii.String("sourcePath"),
//   							},
//   							name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		numNodes: jsii.Number(123),
//   	},
//   	parameters: parameters,
//   	platformCapabilities: []*string{
//   		jsii.String("platformCapabilities"),
//   	},
//   	propagateTags: jsii.Boolean(false),
//   	retryStrategy: &retryStrategyProperty{
//   		attempts: jsii.Number(123),
//   		evaluateOnExit: []interface{}{
//   			&evaluateOnExitProperty{
//   				action: jsii.String("action"),
//
//   				// the properties below are optional
//   				onExitCode: jsii.String("onExitCode"),
//   				onReason: jsii.String("onReason"),
//   				onStatusReason: jsii.String("onStatusReason"),
//   			},
//   		},
//   	},
//   	schedulingPriority: jsii.Number(123),
//   	tags: tags,
//   	timeout: &timeoutProperty{
//   		attemptDurationSeconds: jsii.Number(123),
//   	},
//   }
//
type CfnJobDefinitionProps struct {
	// The type of job definition.
	//
	// For more information about multi-node parallel jobs, see [Creating a multi-node parallel job definition](https://docs.aws.amazon.com/batch/latest/userguide/multi-node-job-def.html) in the *AWS Batch User Guide* .
	//
	// > If the job is run on Fargate resources, then `multinode` isn't supported.
	Type *string `field:"required" json:"type" yaml:"type"`
	// An object with various properties specific to container-based jobs.
	ContainerProperties interface{} `field:"optional" json:"containerProperties" yaml:"containerProperties"`
	// `AWS::Batch::JobDefinition.EksProperties`.
	EksProperties interface{} `field:"optional" json:"eksProperties" yaml:"eksProperties"`
	// The name of the job definition.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// An object with various properties specific to multi-node parallel jobs.
	//
	// > If the job runs on Fargate resources, then you must not specify `nodeProperties` ; use `containerProperties` instead.
	NodeProperties interface{} `field:"optional" json:"nodeProperties" yaml:"nodeProperties"`
	// Default parameters or parameter substitution placeholders that are set in the job definition.
	//
	// Parameters are specified as a key-value pair mapping. Parameters in a `SubmitJob` request override any corresponding parameter defaults from the job definition. For more information about specifying parameters, see [Job definition parameters](https://docs.aws.amazon.com/batch/latest/userguide/job_definition_parameters.html) in the *AWS Batch User Guide* .
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The platform capabilities required by the job definition.
	//
	// If no value is specified, it defaults to `EC2` . Jobs run on Fargate resources specify `FARGATE` .
	PlatformCapabilities *[]*string `field:"optional" json:"platformCapabilities" yaml:"platformCapabilities"`
	// Specifies whether to propagate the tags from the job or job definition to the corresponding Amazon ECS task.
	//
	// If no value is specified, the tags aren't propagated. Tags can only be propagated to the tasks during task creation. For tags with the same name, job tags are given priority over job definitions tags. If the total number of combined tags from the job and job definition is over 50, the job is moved to the `FAILED` state.
	PropagateTags interface{} `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The retry strategy to use for failed jobs that are submitted with this job definition.
	RetryStrategy interface{} `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	// The scheduling priority of the job definition.
	//
	// This only affects jobs in job queues with a fair share policy. Jobs with a higher scheduling priority are scheduled before jobs with a lower scheduling priority.
	SchedulingPriority *float64 `field:"optional" json:"schedulingPriority" yaml:"schedulingPriority"`
	// The tags applied to the job definition.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The timeout configuration for jobs that are submitted with this job definition.
	//
	// You can specify a timeout duration after which AWS Batch terminates your jobs if they haven't finished.
	Timeout interface{} `field:"optional" json:"timeout" yaml:"timeout"`
}

