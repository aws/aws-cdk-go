package previewawsguarddutyevents


// Type definition for Resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var ipv6Addresses interface{}
//
//   resource := &Resource{
//   	AccessKeyDetails: &AccessKeyDetails{
//   		AccessKeyId: []*string{
//   			jsii.String("accessKeyId"),
//   		},
//   		PrincipalId: []*string{
//   			jsii.String("principalId"),
//   		},
//   		UserName: []*string{
//   			jsii.String("userName"),
//   		},
//   		UserType: []*string{
//   			jsii.String("userType"),
//   		},
//   	},
//   	ContainerDetails: &ContainerDetails{
//   		Id: []*string{
//   			jsii.String("id"),
//   		},
//   		Image: []*string{
//   			jsii.String("image"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   	},
//   	EbsVolumeDetails: &EbsVolumeDetails{
//   		ScannedVolumeDetails: []EbsVolumeDetailsItem{
//   			&EbsVolumeDetailsItem{
//   				DeviceName: []*string{
//   					jsii.String("deviceName"),
//   				},
//   				EncryptionType: []*string{
//   					jsii.String("encryptionType"),
//   				},
//   				KmsKeyArn: []*string{
//   					jsii.String("kmsKeyArn"),
//   				},
//   				SnapshotArn: []*string{
//   					jsii.String("snapshotArn"),
//   				},
//   				VolumeArn: []*string{
//   					jsii.String("volumeArn"),
//   				},
//   				VolumeSizeInGb: []*string{
//   					jsii.String("volumeSizeInGb"),
//   				},
//   				VolumeType: []*string{
//   					jsii.String("volumeType"),
//   				},
//   			},
//   		},
//   		SkippedVolumeDetails: []*string{
//   			jsii.String("skippedVolumeDetails"),
//   		},
//   	},
//   	EcsClusterDetails: &EcsClusterDetails{
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		Status: []*string{
//   			jsii.String("status"),
//   		},
//   		Tags: []EcsClusterDetailsItem{
//   			&EcsClusterDetailsItem{
//   				Key: []*string{
//   					jsii.String("key"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   		TaskDetails: &TaskDetails{
//   			Arn: []*string{
//   				jsii.String("arn"),
//   			},
//   			Containers: []TaskDetailsItem{
//   				&TaskDetailsItem{
//   					Image: []*string{
//   						jsii.String("image"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   				},
//   			},
//   			CreatedAt: []*string{
//   				jsii.String("createdAt"),
//   			},
//   			DefinitionArn: []*string{
//   				jsii.String("definitionArn"),
//   			},
//   			StartedAt: []*string{
//   				jsii.String("startedAt"),
//   			},
//   			StartedBy: []*string{
//   				jsii.String("startedBy"),
//   			},
//   			Version: []*string{
//   				jsii.String("version"),
//   			},
//   		},
//   	},
//   	EksClusterDetails: &EksClusterDetails{
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		CreatedAt: []*string{
//   			jsii.String("createdAt"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		Status: []*string{
//   			jsii.String("status"),
//   		},
//   		Tags: []EcsClusterDetailsItem{
//   			&EcsClusterDetailsItem{
//   				Key: []*string{
//   					jsii.String("key"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   		VpcId: []*string{
//   			jsii.String("vpcId"),
//   		},
//   	},
//   	InstanceDetails: &InstanceDetails{
//   		AvailabilityZone: []*string{
//   			jsii.String("availabilityZone"),
//   		},
//   		IamInstanceProfile: &IamInstanceProfile{
//   			Arn: []*string{
//   				jsii.String("arn"),
//   			},
//   			Id: []*string{
//   				jsii.String("id"),
//   			},
//   		},
//   		ImageDescription: []*string{
//   			jsii.String("imageDescription"),
//   		},
//   		ImageId: []*string{
//   			jsii.String("imageId"),
//   		},
//   		InstanceId: []*string{
//   			jsii.String("instanceId"),
//   		},
//   		InstanceState: []*string{
//   			jsii.String("instanceState"),
//   		},
//   		InstanceType: []*string{
//   			jsii.String("instanceType"),
//   		},
//   		LaunchTime: []*string{
//   			jsii.String("launchTime"),
//   		},
//   		NetworkInterfaces: []InstanceDetailsItem{
//   			&InstanceDetailsItem{
//   				Ipv6Addresses: []interface{}{
//   					ipv6Addresses,
//   				},
//   				NetworkInterfaceId: []*string{
//   					jsii.String("networkInterfaceId"),
//   				},
//   				PrivateDnsName: []*string{
//   					jsii.String("privateDnsName"),
//   				},
//   				PrivateIpAddress: []*string{
//   					jsii.String("privateIpAddress"),
//   				},
//   				PrivateIpAddresses: []InstanceDetailsItemItem{
//   					&InstanceDetailsItemItem{
//   						PrivateDnsName: []*string{
//   							jsii.String("privateDnsName"),
//   						},
//   						PrivateIpAddress: []*string{
//   							jsii.String("privateIpAddress"),
//   						},
//   					},
//   				},
//   				PublicDnsName: []*string{
//   					jsii.String("publicDnsName"),
//   				},
//   				PublicIp: []*string{
//   					jsii.String("publicIp"),
//   				},
//   				SecurityGroups: []InstanceDetailsItemItem1{
//   					&InstanceDetailsItemItem1{
//   						GroupId: []*string{
//   							jsii.String("groupId"),
//   						},
//   						GroupName: []*string{
//   							jsii.String("groupName"),
//   						},
//   					},
//   				},
//   				SubnetId: []*string{
//   					jsii.String("subnetId"),
//   				},
//   				VpcId: []*string{
//   					jsii.String("vpcId"),
//   				},
//   			},
//   		},
//   		OutpostArn: []*string{
//   			jsii.String("outpostArn"),
//   		},
//   		Platform: []*string{
//   			jsii.String("platform"),
//   		},
//   		ProductCodes: []InstanceDetailsItem1{
//   			&InstanceDetailsItem1{
//   				ProductCodeId: []*string{
//   					jsii.String("productCodeId"),
//   				},
//   				ProductCodeType: []*string{
//   					jsii.String("productCodeType"),
//   				},
//   			},
//   		},
//   		Tags: []EcsClusterDetailsItem{
//   			&EcsClusterDetailsItem{
//   				Key: []*string{
//   					jsii.String("key"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	KubernetesDetails: &KubernetesDetails{
//   		KubernetesUserDetails: &KubernetesUserDetails{
//   			Groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			Uid: []*string{
//   				jsii.String("uid"),
//   			},
//   			Username: []*string{
//   				jsii.String("username"),
//   			},
//   		},
//   		KubernetesWorkloadDetails: &KubernetesWorkloadDetails{
//   			Containers: []KubernetesWorkloadDetailsItem{
//   				&KubernetesWorkloadDetailsItem{
//   					Image: []*string{
//   						jsii.String("image"),
//   					},
//   					ImagePrefix: []*string{
//   						jsii.String("imagePrefix"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					SecurityContext: &SecurityContext{
//   						Privileged: []*string{
//   							jsii.String("privileged"),
//   						},
//   					},
//   				},
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Namespace: []*string{
//   				jsii.String("namespace"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   			Uid: []*string{
//   				jsii.String("uid"),
//   			},
//   		},
//   	},
//   	ResourceType: []*string{
//   		jsii.String("resourceType"),
//   	},
//   	S3BucketDetails: []ResourceItem{
//   		&ResourceItem{
//   			Arn: []*string{
//   				jsii.String("arn"),
//   			},
//   			CreatedAt: []*string{
//   				jsii.String("createdAt"),
//   			},
//   			DefaultServerSideEncryption: &DefaultServerSideEncryption{
//   				EncryptionType: []*string{
//   					jsii.String("encryptionType"),
//   				},
//   				KmsMasterKeyArn: []*string{
//   					jsii.String("kmsMasterKeyArn"),
//   				},
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Owner: &Owner{
//   				Id: []*string{
//   					jsii.String("id"),
//   				},
//   			},
//   			PublicAccess: &PublicAccess{
//   				EffectivePermission: []*string{
//   					jsii.String("effectivePermission"),
//   				},
//   				PermissionConfiguration: &PermissionConfiguration{
//   					AccountLevelPermissions: &AccountLevelPermissions{
//   						BlockPublicAccess: &BlockPublicAccess{
//   							BlockPublicAcls: []*string{
//   								jsii.String("blockPublicAcls"),
//   							},
//   							BlockPublicPolicy: []*string{
//   								jsii.String("blockPublicPolicy"),
//   							},
//   							IgnorePublicAcls: []*string{
//   								jsii.String("ignorePublicAcls"),
//   							},
//   							RestrictPublicBuckets: []*string{
//   								jsii.String("restrictPublicBuckets"),
//   							},
//   						},
//   					},
//   					BucketLevelPermissions: &BucketLevelPermissions{
//   						AccessControlList: &AccessControlList{
//   							AllowsPublicReadAccess: []*string{
//   								jsii.String("allowsPublicReadAccess"),
//   							},
//   							AllowsPublicWriteAccess: []*string{
//   								jsii.String("allowsPublicWriteAccess"),
//   							},
//   						},
//   						BlockPublicAccess: &BlockPublicAccess{
//   							BlockPublicAcls: []*string{
//   								jsii.String("blockPublicAcls"),
//   							},
//   							BlockPublicPolicy: []*string{
//   								jsii.String("blockPublicPolicy"),
//   							},
//   							IgnorePublicAcls: []*string{
//   								jsii.String("ignorePublicAcls"),
//   							},
//   							RestrictPublicBuckets: []*string{
//   								jsii.String("restrictPublicBuckets"),
//   							},
//   						},
//   						BucketPolicy: &AccessControlList{
//   							AllowsPublicReadAccess: []*string{
//   								jsii.String("allowsPublicReadAccess"),
//   							},
//   							AllowsPublicWriteAccess: []*string{
//   								jsii.String("allowsPublicWriteAccess"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			Tags: []EcsClusterDetailsItem{
//   				&EcsClusterDetailsItem{
//   					Key: []*string{
//   						jsii.String("key"),
//   					},
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_Resource struct {
	// accessKeyDetails property.
	//
	// Specify an array of string values to match this event if the actual value of accessKeyDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccessKeyDetails *DetectorEvents_GuardDutyFinding_AccessKeyDetails `field:"optional" json:"accessKeyDetails" yaml:"accessKeyDetails"`
	// containerDetails property.
	//
	// Specify an array of string values to match this event if the actual value of containerDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerDetails *DetectorEvents_GuardDutyFinding_ContainerDetails `field:"optional" json:"containerDetails" yaml:"containerDetails"`
	// ebsVolumeDetails property.
	//
	// Specify an array of string values to match this event if the actual value of ebsVolumeDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EbsVolumeDetails *DetectorEvents_GuardDutyFinding_EbsVolumeDetails `field:"optional" json:"ebsVolumeDetails" yaml:"ebsVolumeDetails"`
	// ecsClusterDetails property.
	//
	// Specify an array of string values to match this event if the actual value of ecsClusterDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EcsClusterDetails *DetectorEvents_GuardDutyFinding_EcsClusterDetails `field:"optional" json:"ecsClusterDetails" yaml:"ecsClusterDetails"`
	// eksClusterDetails property.
	//
	// Specify an array of string values to match this event if the actual value of eksClusterDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EksClusterDetails *DetectorEvents_GuardDutyFinding_EksClusterDetails `field:"optional" json:"eksClusterDetails" yaml:"eksClusterDetails"`
	// instanceDetails property.
	//
	// Specify an array of string values to match this event if the actual value of instanceDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceDetails *DetectorEvents_GuardDutyFinding_InstanceDetails `field:"optional" json:"instanceDetails" yaml:"instanceDetails"`
	// kubernetesDetails property.
	//
	// Specify an array of string values to match this event if the actual value of kubernetesDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	KubernetesDetails *DetectorEvents_GuardDutyFinding_KubernetesDetails `field:"optional" json:"kubernetesDetails" yaml:"kubernetesDetails"`
	// resourceType property.
	//
	// Specify an array of string values to match this event if the actual value of resourceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceType *[]*string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// s3BucketDetails property.
	//
	// Specify an array of string values to match this event if the actual value of s3BucketDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3BucketDetails *[]*DetectorEvents_GuardDutyFinding_ResourceItem `field:"optional" json:"s3BucketDetails" yaml:"s3BucketDetails"`
}

