package mixinsawsbatch


// This is an object that represents the properties of the node range for a multi-node parallel job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var labels interface{}
//   var limits interface{}
//   var options interface{}
//   var requests interface{}
//
//   nodeRangePropertyProperty := &NodeRangePropertyProperty{
//   	ConsumableResourceProperties: &ConsumableResourcePropertiesProperty{
//   		ConsumableResourceList: []interface{}{
//   			&ConsumableResourceRequirementProperty{
//   				ConsumableResource: jsii.String("consumableResource"),
//   				Quantity: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Container: &ContainerPropertiesProperty{
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		EnableExecuteCommand: jsii.Boolean(false),
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
//   		Image: jsii.String("image"),
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
//   					MountOptions: []*string{
//   						jsii.String("mountOptions"),
//   					},
//   					Size: jsii.Number(123),
//   				},
//   			},
//   		},
//   		LogConfiguration: &LogConfigurationProperty{
//   			LogDriver: jsii.String("logDriver"),
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
//   		RepositoryCredentials: &RepositoryCredentialsProperty{
//   			CredentialsParameter: jsii.String("credentialsParameter"),
//   		},
//   		ResourceRequirements: []interface{}{
//   			&ResourceRequirementProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RuntimePlatform: &RuntimePlatformProperty{
//   			CpuArchitecture: jsii.String("cpuArchitecture"),
//   			OperatingSystemFamily: jsii.String("operatingSystemFamily"),
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
//   					AuthorizationConfig: &AuthorizationConfigProperty{
//   						AccessPointId: jsii.String("accessPointId"),
//   						Iam: jsii.String("iam"),
//   					},
//   					FileSystemId: jsii.String("fileSystemId"),
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
//   	EcsProperties: &MultiNodeEcsPropertiesProperty{
//   		TaskProperties: []interface{}{
//   			&MultiNodeEcsTaskPropertiesProperty{
//   				Containers: []interface{}{
//   					&TaskContainerPropertiesProperty{
//   						Command: []*string{
//   							jsii.String("command"),
//   						},
//   						DependsOn: []interface{}{
//   							&TaskContainerDependencyProperty{
//   								Condition: jsii.String("condition"),
//   								ContainerName: jsii.String("containerName"),
//   							},
//   						},
//   						Environment: []interface{}{
//   							&EnvironmentProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						Essential: jsii.Boolean(false),
//   						FirelensConfiguration: &FirelensConfigurationProperty{
//   							Options: map[string]*string{
//   								"optionsKey": jsii.String("options"),
//   							},
//   							Type: jsii.String("type"),
//   						},
//   						Image: jsii.String("image"),
//   						LinuxParameters: &LinuxParametersProperty{
//   							Devices: []interface{}{
//   								&DeviceProperty{
//   									ContainerPath: jsii.String("containerPath"),
//   									HostPath: jsii.String("hostPath"),
//   									Permissions: []*string{
//   										jsii.String("permissions"),
//   									},
//   								},
//   							},
//   							InitProcessEnabled: jsii.Boolean(false),
//   							MaxSwap: jsii.Number(123),
//   							SharedMemorySize: jsii.Number(123),
//   							Swappiness: jsii.Number(123),
//   							Tmpfs: []interface{}{
//   								&TmpfsProperty{
//   									ContainerPath: jsii.String("containerPath"),
//   									MountOptions: []*string{
//   										jsii.String("mountOptions"),
//   									},
//   									Size: jsii.Number(123),
//   								},
//   							},
//   						},
//   						LogConfiguration: &LogConfigurationProperty{
//   							LogDriver: jsii.String("logDriver"),
//   							Options: options,
//   							SecretOptions: []interface{}{
//   								&SecretProperty{
//   									Name: jsii.String("name"),
//   									ValueFrom: jsii.String("valueFrom"),
//   								},
//   							},
//   						},
//   						MountPoints: []interface{}{
//   							&MountPointProperty{
//   								ContainerPath: jsii.String("containerPath"),
//   								ReadOnly: jsii.Boolean(false),
//   								SourceVolume: jsii.String("sourceVolume"),
//   							},
//   						},
//   						Name: jsii.String("name"),
//   						Privileged: jsii.Boolean(false),
//   						ReadonlyRootFilesystem: jsii.Boolean(false),
//   						RepositoryCredentials: &RepositoryCredentialsProperty{
//   							CredentialsParameter: jsii.String("credentialsParameter"),
//   						},
//   						ResourceRequirements: []interface{}{
//   							&ResourceRequirementProperty{
//   								Type: jsii.String("type"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						Secrets: []interface{}{
//   							&SecretProperty{
//   								Name: jsii.String("name"),
//   								ValueFrom: jsii.String("valueFrom"),
//   							},
//   						},
//   						Ulimits: []interface{}{
//   							&UlimitProperty{
//   								HardLimit: jsii.Number(123),
//   								Name: jsii.String("name"),
//   								SoftLimit: jsii.Number(123),
//   							},
//   						},
//   						User: jsii.String("user"),
//   					},
//   				},
//   				EnableExecuteCommand: jsii.Boolean(false),
//   				ExecutionRoleArn: jsii.String("executionRoleArn"),
//   				IpcMode: jsii.String("ipcMode"),
//   				PidMode: jsii.String("pidMode"),
//   				TaskRoleArn: jsii.String("taskRoleArn"),
//   				Volumes: []interface{}{
//   					&VolumesProperty{
//   						EfsVolumeConfiguration: &EfsVolumeConfigurationProperty{
//   							AuthorizationConfig: &AuthorizationConfigProperty{
//   								AccessPointId: jsii.String("accessPointId"),
//   								Iam: jsii.String("iam"),
//   							},
//   							FileSystemId: jsii.String("fileSystemId"),
//   							RootDirectory: jsii.String("rootDirectory"),
//   							TransitEncryption: jsii.String("transitEncryption"),
//   							TransitEncryptionPort: jsii.Number(123),
//   						},
//   						Host: &VolumesHostProperty{
//   							SourcePath: jsii.String("sourcePath"),
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	EksProperties: &EksPropertiesProperty{
//   		PodProperties: &PodPropertiesProperty{
//   			Containers: []interface{}{
//   				&EksContainerProperty{
//   					Args: []*string{
//   						jsii.String("args"),
//   					},
//   					Command: []*string{
//   						jsii.String("command"),
//   					},
//   					Env: []interface{}{
//   						&EksContainerEnvironmentVariableProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Image: jsii.String("image"),
//   					ImagePullPolicy: jsii.String("imagePullPolicy"),
//   					Name: jsii.String("name"),
//   					Resources: &ResourcesProperty{
//   						Limits: limits,
//   						Requests: requests,
//   					},
//   					SecurityContext: &SecurityContextProperty{
//   						AllowPrivilegeEscalation: jsii.Boolean(false),
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
//   							SubPath: jsii.String("subPath"),
//   						},
//   					},
//   				},
//   			},
//   			DnsPolicy: jsii.String("dnsPolicy"),
//   			HostNetwork: jsii.Boolean(false),
//   			ImagePullSecrets: []interface{}{
//   				&ImagePullSecretProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			InitContainers: []interface{}{
//   				&EksContainerProperty{
//   					Args: []*string{
//   						jsii.String("args"),
//   					},
//   					Command: []*string{
//   						jsii.String("command"),
//   					},
//   					Env: []interface{}{
//   						&EksContainerEnvironmentVariableProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Image: jsii.String("image"),
//   					ImagePullPolicy: jsii.String("imagePullPolicy"),
//   					Name: jsii.String("name"),
//   					Resources: &ResourcesProperty{
//   						Limits: limits,
//   						Requests: requests,
//   					},
//   					SecurityContext: &SecurityContextProperty{
//   						AllowPrivilegeEscalation: jsii.Boolean(false),
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
//   							SubPath: jsii.String("subPath"),
//   						},
//   					},
//   				},
//   			},
//   			Metadata: &MetadataProperty{
//   				Labels: labels,
//   			},
//   			ServiceAccountName: jsii.String("serviceAccountName"),
//   			ShareProcessNamespace: jsii.Boolean(false),
//   			Volumes: []interface{}{
//   				&EksVolumeProperty{
//   					EmptyDir: &EmptyDirProperty{
//   						Medium: jsii.String("medium"),
//   						SizeLimit: jsii.String("sizeLimit"),
//   					},
//   					HostPath: &HostPathProperty{
//   						Path: jsii.String("path"),
//   					},
//   					Name: jsii.String("name"),
//   					PersistentVolumeClaim: &EksPersistentVolumeClaimProperty{
//   						ClaimName: jsii.String("claimName"),
//   						ReadOnly: jsii.Boolean(false),
//   					},
//   					Secret: &EksSecretProperty{
//   						Optional: jsii.Boolean(false),
//   						SecretName: jsii.String("secretName"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	InstanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	TargetNodes: jsii.String("targetNodes"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html
//
type CfnJobDefinitionPropsMixin_NodeRangePropertyProperty struct {
	// Contains a list of consumable resources required by a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html#cfn-batch-jobdefinition-noderangeproperty-consumableresourceproperties
	//
	ConsumableResourceProperties interface{} `field:"optional" json:"consumableResourceProperties" yaml:"consumableResourceProperties"`
	// The container details for the node range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html#cfn-batch-jobdefinition-noderangeproperty-container
	//
	Container interface{} `field:"optional" json:"container" yaml:"container"`
	// This is an object that represents the properties of the node range for a multi-node parallel job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html#cfn-batch-jobdefinition-noderangeproperty-ecsproperties
	//
	EcsProperties interface{} `field:"optional" json:"ecsProperties" yaml:"ecsProperties"`
	// This is an object that represents the properties of the node range for a multi-node parallel job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html#cfn-batch-jobdefinition-noderangeproperty-eksproperties
	//
	EksProperties interface{} `field:"optional" json:"eksProperties" yaml:"eksProperties"`
	// The instance types of the underlying host infrastructure of a multi-node parallel job.
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources.
	// >
	// > In addition, this list object is currently limited to one element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html#cfn-batch-jobdefinition-noderangeproperty-instancetypes
	//
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The range of nodes, using node index values.
	//
	// A range of `0:3` indicates nodes with index values of `0` through `3` . If the starting range value is omitted ( `:n` ), then `0` is used to start the range. If the ending range value is omitted ( `n:` ), then the highest possible node index is used to end the range. Your accumulative node ranges must account for all nodes ( `0:n` ). You can nest node ranges (for example, `0:10` and `4:5` ). In this case, the `4:5` range properties override the `0:10` properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html#cfn-batch-jobdefinition-noderangeproperty-targetnodes
	//
	TargetNodes *string `field:"optional" json:"targetNodes" yaml:"targetNodes"`
}

