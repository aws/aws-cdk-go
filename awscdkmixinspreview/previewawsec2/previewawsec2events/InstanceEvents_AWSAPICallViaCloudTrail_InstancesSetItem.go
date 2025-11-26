package previewawsec2events


// Type definition for InstancesSetItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instancesSetItem := &InstancesSetItem{
//   	AmiLaunchIndex: []*string{
//   		jsii.String("amiLaunchIndex"),
//   	},
//   	Architecture: []*string{
//   		jsii.String("architecture"),
//   	},
//   	BlockDeviceMapping: []*string{
//   		jsii.String("blockDeviceMapping"),
//   	},
//   	CapacityReservationSpecification: &CapacityReservationSpecification{
//   		CapacityReservationPreference: []*string{
//   			jsii.String("capacityReservationPreference"),
//   		},
//   	},
//   	ClientToken: []*string{
//   		jsii.String("clientToken"),
//   	},
//   	CpuOptions: &CpuOptions{
//   		CoreCount: []*string{
//   			jsii.String("coreCount"),
//   		},
//   		ThreadsPerCore: []*string{
//   			jsii.String("threadsPerCore"),
//   		},
//   	},
//   	CurrentState: &InstanceState{
//   		Code: []*string{
//   			jsii.String("code"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   	},
//   	EbsOptimized: []*string{
//   		jsii.String("ebsOptimized"),
//   	},
//   	EnclaveOptions: &EnclaveOptions{
//   		Enabled: []*string{
//   			jsii.String("enabled"),
//   		},
//   	},
//   	GroupSet: &GroupSet2{
//   		Items: []GroupSet2Item{
//   			&GroupSet2Item{
//   				GroupId: []*string{
//   					jsii.String("groupId"),
//   				},
//   				GroupName: []*string{
//   					jsii.String("groupName"),
//   				},
//   			},
//   		},
//   	},
//   	Hypervisor: []*string{
//   		jsii.String("hypervisor"),
//   	},
//   	ImageId: []*string{
//   		jsii.String("imageId"),
//   	},
//   	InstanceId: []*string{
//   		jsii.String("instanceId"),
//   	},
//   	InstanceLifecycle: []*string{
//   		jsii.String("instanceLifecycle"),
//   	},
//   	InstanceState: &InstanceState{
//   		Code: []*string{
//   			jsii.String("code"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   	},
//   	InstanceType: []*string{
//   		jsii.String("instanceType"),
//   	},
//   	LaunchTime: []*string{
//   		jsii.String("launchTime"),
//   	},
//   	Monitoring: &Monitoring1{
//   		State: []*string{
//   			jsii.String("state"),
//   		},
//   	},
//   	NetworkInterfaceSet: &NetworkInterfaceSet1{
//   		Items: []NetworkInterfaceSet1Item{
//   			&NetworkInterfaceSet1Item{
//   				Attachment: &Attachment{
//   					AttachmentId: []*string{
//   						jsii.String("attachmentId"),
//   					},
//   					AttachTime: []*string{
//   						jsii.String("attachTime"),
//   					},
//   					DeleteOnTermination: []*string{
//   						jsii.String("deleteOnTermination"),
//   					},
//   					DeviceIndex: []*string{
//   						jsii.String("deviceIndex"),
//   					},
//   					Status: []*string{
//   						jsii.String("status"),
//   					},
//   				},
//   				GroupSet: &GroupSet3{
//   					Items: []GroupSet2Item{
//   						&GroupSet2Item{
//   							GroupId: []*string{
//   								jsii.String("groupId"),
//   							},
//   							GroupName: []*string{
//   								jsii.String("groupName"),
//   							},
//   						},
//   					},
//   				},
//   				InterfaceType: []*string{
//   					jsii.String("interfaceType"),
//   				},
//   				Ipv6AddressesSet: []*string{
//   					jsii.String("ipv6AddressesSet"),
//   				},
//   				MacAddress: []*string{
//   					jsii.String("macAddress"),
//   				},
//   				NetworkInterfaceId: []*string{
//   					jsii.String("networkInterfaceId"),
//   				},
//   				OwnerId: []*string{
//   					jsii.String("ownerId"),
//   				},
//   				PrivateIpAddress: []*string{
//   					jsii.String("privateIpAddress"),
//   				},
//   				PrivateIpAddressesSet: &PrivateIpAddressesSet2{
//   					Item: []PrivateIpAddressesSet1Item{
//   						&PrivateIpAddressesSet1Item{
//   							Primary: []*string{
//   								jsii.String("primary"),
//   							},
//   							PrivateIpAddress: []*string{
//   								jsii.String("privateIpAddress"),
//   							},
//   						},
//   					},
//   				},
//   				SourceDestCheck: []*string{
//   					jsii.String("sourceDestCheck"),
//   				},
//   				Status: []*string{
//   					jsii.String("status"),
//   				},
//   				SubnetId: []*string{
//   					jsii.String("subnetId"),
//   				},
//   				TagSet: []*string{
//   					jsii.String("tagSet"),
//   				},
//   				VpcId: []*string{
//   					jsii.String("vpcId"),
//   				},
//   			},
//   		},
//   	},
//   	Placement: &Placement{
//   		AvailabilityZone: []*string{
//   			jsii.String("availabilityZone"),
//   		},
//   		Tenancy: []*string{
//   			jsii.String("tenancy"),
//   		},
//   	},
//   	PreviousState: &InstanceState{
//   		Code: []*string{
//   			jsii.String("code"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   	},
//   	PrivateIpAddress: []*string{
//   		jsii.String("privateIpAddress"),
//   	},
//   	ProductCodes: []*string{
//   		jsii.String("productCodes"),
//   	},
//   	RootDeviceName: []*string{
//   		jsii.String("rootDeviceName"),
//   	},
//   	RootDeviceType: []*string{
//   		jsii.String("rootDeviceType"),
//   	},
//   	SourceDestCheck: []*string{
//   		jsii.String("sourceDestCheck"),
//   	},
//   	SpotInstanceRequestId: []*string{
//   		jsii.String("spotInstanceRequestId"),
//   	},
//   	StateReason: &StateReason{
//   		Code: []*string{
//   			jsii.String("code"),
//   		},
//   		Message: []*string{
//   			jsii.String("message"),
//   		},
//   	},
//   	SubnetId: []*string{
//   		jsii.String("subnetId"),
//   	},
//   	TagSet: &TagSet{
//   		Items: []TagSpecificationSetItemItem{
//   			&TagSpecificationSetItemItem{
//   				Key: []*string{
//   					jsii.String("key"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	VirtualizationType: []*string{
//   		jsii.String("virtualizationType"),
//   	},
//   	VpcId: []*string{
//   		jsii.String("vpcId"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_InstancesSetItem struct {
	// amiLaunchIndex property.
	//
	// Specify an array of string values to match this event if the actual value of amiLaunchIndex is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AmiLaunchIndex *[]*string `field:"optional" json:"amiLaunchIndex" yaml:"amiLaunchIndex"`
	// architecture property.
	//
	// Specify an array of string values to match this event if the actual value of architecture is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Architecture *[]*string `field:"optional" json:"architecture" yaml:"architecture"`
	// blockDeviceMapping property.
	//
	// Specify an array of string values to match this event if the actual value of blockDeviceMapping is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BlockDeviceMapping *[]*string `field:"optional" json:"blockDeviceMapping" yaml:"blockDeviceMapping"`
	// capacityReservationSpecification property.
	//
	// Specify an array of string values to match this event if the actual value of capacityReservationSpecification is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CapacityReservationSpecification *InstanceEvents_AWSAPICallViaCloudTrail_CapacityReservationSpecification `field:"optional" json:"capacityReservationSpecification" yaml:"capacityReservationSpecification"`
	// clientToken property.
	//
	// Specify an array of string values to match this event if the actual value of clientToken is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientToken *[]*string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// cpuOptions property.
	//
	// Specify an array of string values to match this event if the actual value of cpuOptions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CpuOptions *InstanceEvents_AWSAPICallViaCloudTrail_CpuOptions `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// currentState property.
	//
	// Specify an array of string values to match this event if the actual value of currentState is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CurrentState *InstanceEvents_AWSAPICallViaCloudTrail_InstanceState `field:"optional" json:"currentState" yaml:"currentState"`
	// ebsOptimized property.
	//
	// Specify an array of string values to match this event if the actual value of ebsOptimized is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EbsOptimized *[]*string `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// enclaveOptions property.
	//
	// Specify an array of string values to match this event if the actual value of enclaveOptions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EnclaveOptions *InstanceEvents_AWSAPICallViaCloudTrail_EnclaveOptions `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// groupSet property.
	//
	// Specify an array of string values to match this event if the actual value of groupSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupSet *InstanceEvents_AWSAPICallViaCloudTrail_GroupSet2 `field:"optional" json:"groupSet" yaml:"groupSet"`
	// hypervisor property.
	//
	// Specify an array of string values to match this event if the actual value of hypervisor is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Hypervisor *[]*string `field:"optional" json:"hypervisor" yaml:"hypervisor"`
	// imageId property.
	//
	// Specify an array of string values to match this event if the actual value of imageId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageId *[]*string `field:"optional" json:"imageId" yaml:"imageId"`
	// instanceId property.
	//
	// Specify an array of string values to match this event if the actual value of instanceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Instance reference.
	//
	// Experimental.
	InstanceId *[]*string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// instanceLifecycle property.
	//
	// Specify an array of string values to match this event if the actual value of instanceLifecycle is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceLifecycle *[]*string `field:"optional" json:"instanceLifecycle" yaml:"instanceLifecycle"`
	// instanceState property.
	//
	// Specify an array of string values to match this event if the actual value of instanceState is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceState *InstanceEvents_AWSAPICallViaCloudTrail_InstanceState `field:"optional" json:"instanceState" yaml:"instanceState"`
	// instanceType property.
	//
	// Specify an array of string values to match this event if the actual value of instanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// launchTime property.
	//
	// Specify an array of string values to match this event if the actual value of launchTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTime *[]*string `field:"optional" json:"launchTime" yaml:"launchTime"`
	// monitoring property.
	//
	// Specify an array of string values to match this event if the actual value of monitoring is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Monitoring *InstanceEvents_AWSAPICallViaCloudTrail_Monitoring1 `field:"optional" json:"monitoring" yaml:"monitoring"`
	// networkInterfaceSet property.
	//
	// Specify an array of string values to match this event if the actual value of networkInterfaceSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkInterfaceSet *InstanceEvents_AWSAPICallViaCloudTrail_NetworkInterfaceSet1 `field:"optional" json:"networkInterfaceSet" yaml:"networkInterfaceSet"`
	// placement property.
	//
	// Specify an array of string values to match this event if the actual value of placement is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Placement *InstanceEvents_AWSAPICallViaCloudTrail_Placement `field:"optional" json:"placement" yaml:"placement"`
	// previousState property.
	//
	// Specify an array of string values to match this event if the actual value of previousState is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousState *InstanceEvents_AWSAPICallViaCloudTrail_InstanceState `field:"optional" json:"previousState" yaml:"previousState"`
	// privateIpAddress property.
	//
	// Specify an array of string values to match this event if the actual value of privateIpAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrivateIpAddress *[]*string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// productCodes property.
	//
	// Specify an array of string values to match this event if the actual value of productCodes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProductCodes *[]*string `field:"optional" json:"productCodes" yaml:"productCodes"`
	// rootDeviceName property.
	//
	// Specify an array of string values to match this event if the actual value of rootDeviceName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RootDeviceName *[]*string `field:"optional" json:"rootDeviceName" yaml:"rootDeviceName"`
	// rootDeviceType property.
	//
	// Specify an array of string values to match this event if the actual value of rootDeviceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RootDeviceType *[]*string `field:"optional" json:"rootDeviceType" yaml:"rootDeviceType"`
	// sourceDestCheck property.
	//
	// Specify an array of string values to match this event if the actual value of sourceDestCheck is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceDestCheck *[]*string `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// spotInstanceRequestId property.
	//
	// Specify an array of string values to match this event if the actual value of spotInstanceRequestId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SpotInstanceRequestId *[]*string `field:"optional" json:"spotInstanceRequestId" yaml:"spotInstanceRequestId"`
	// stateReason property.
	//
	// Specify an array of string values to match this event if the actual value of stateReason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StateReason *InstanceEvents_AWSAPICallViaCloudTrail_StateReason `field:"optional" json:"stateReason" yaml:"stateReason"`
	// subnetId property.
	//
	// Specify an array of string values to match this event if the actual value of subnetId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SubnetId *[]*string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// tagSet property.
	//
	// Specify an array of string values to match this event if the actual value of tagSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TagSet *InstanceEvents_AWSAPICallViaCloudTrail_TagSet `field:"optional" json:"tagSet" yaml:"tagSet"`
	// virtualizationType property.
	//
	// Specify an array of string values to match this event if the actual value of virtualizationType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VirtualizationType *[]*string `field:"optional" json:"virtualizationType" yaml:"virtualizationType"`
	// vpcId property.
	//
	// Specify an array of string values to match this event if the actual value of vpcId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VpcId *[]*string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

