package previewawsec2events


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseElements := &ResponseElements{
//   	CreateFleetResponse: &CreateFleetResponse{
//   		ErrorSet: []*string{
//   			jsii.String("errorSet"),
//   		},
//   		FleetId: []*string{
//   			jsii.String("fleetId"),
//   		},
//   		FleetInstanceSet: []*string{
//   			jsii.String("fleetInstanceSet"),
//   		},
//   		RequestId: []*string{
//   			jsii.String("requestId"),
//   		},
//   		Xmlns: []*string{
//   			jsii.String("xmlns"),
//   		},
//   	},
//   	CreateLaunchTemplateResponse: &DeleteLaunchTemplateResponse{
//   		LaunchTemplate: &LaunchTemplate1{
//   			CreatedBy: []*string{
//   				jsii.String("createdBy"),
//   			},
//   			CreateTime: []*string{
//   				jsii.String("createTime"),
//   			},
//   			DefaultVersionNumber: []*string{
//   				jsii.String("defaultVersionNumber"),
//   			},
//   			LatestVersionNumber: []*string{
//   				jsii.String("latestVersionNumber"),
//   			},
//   			LaunchTemplateId: []*string{
//   				jsii.String("launchTemplateId"),
//   			},
//   			LaunchTemplateName: []*string{
//   				jsii.String("launchTemplateName"),
//   			},
//   		},
//   		RequestId: []*string{
//   			jsii.String("requestId"),
//   		},
//   		Xmlns: []*string{
//   			jsii.String("xmlns"),
//   		},
//   	},
//   	DeleteLaunchTemplateResponse: &DeleteLaunchTemplateResponse{
//   		LaunchTemplate: &LaunchTemplate1{
//   			CreatedBy: []*string{
//   				jsii.String("createdBy"),
//   			},
//   			CreateTime: []*string{
//   				jsii.String("createTime"),
//   			},
//   			DefaultVersionNumber: []*string{
//   				jsii.String("defaultVersionNumber"),
//   			},
//   			LatestVersionNumber: []*string{
//   				jsii.String("latestVersionNumber"),
//   			},
//   			LaunchTemplateId: []*string{
//   				jsii.String("launchTemplateId"),
//   			},
//   			LaunchTemplateName: []*string{
//   				jsii.String("launchTemplateName"),
//   			},
//   		},
//   		RequestId: []*string{
//   			jsii.String("requestId"),
//   		},
//   		Xmlns: []*string{
//   			jsii.String("xmlns"),
//   		},
//   	},
//   	GroupId: []*string{
//   		jsii.String("groupId"),
//   	},
//   	GroupSet: []*string{
//   		jsii.String("groupSet"),
//   	},
//   	InstancesSet: &InstancesSet{
//   		Items: []InstancesSetItem{
//   			&InstancesSetItem{
//   				AmiLaunchIndex: []*string{
//   					jsii.String("amiLaunchIndex"),
//   				},
//   				Architecture: []*string{
//   					jsii.String("architecture"),
//   				},
//   				BlockDeviceMapping: []*string{
//   					jsii.String("blockDeviceMapping"),
//   				},
//   				CapacityReservationSpecification: &CapacityReservationSpecification{
//   					CapacityReservationPreference: []*string{
//   						jsii.String("capacityReservationPreference"),
//   					},
//   				},
//   				ClientToken: []*string{
//   					jsii.String("clientToken"),
//   				},
//   				CpuOptions: &CpuOptions{
//   					CoreCount: []*string{
//   						jsii.String("coreCount"),
//   					},
//   					ThreadsPerCore: []*string{
//   						jsii.String("threadsPerCore"),
//   					},
//   				},
//   				CurrentState: &InstanceState{
//   					Code: []*string{
//   						jsii.String("code"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   				},
//   				EbsOptimized: []*string{
//   					jsii.String("ebsOptimized"),
//   				},
//   				EnclaveOptions: &EnclaveOptions{
//   					Enabled: []*string{
//   						jsii.String("enabled"),
//   					},
//   				},
//   				GroupSet: &GroupSet2{
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
//   				Hypervisor: []*string{
//   					jsii.String("hypervisor"),
//   				},
//   				ImageId: []*string{
//   					jsii.String("imageId"),
//   				},
//   				InstanceId: []*string{
//   					jsii.String("instanceId"),
//   				},
//   				InstanceLifecycle: []*string{
//   					jsii.String("instanceLifecycle"),
//   				},
//   				InstanceState: &InstanceState{
//   					Code: []*string{
//   						jsii.String("code"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   				},
//   				InstanceType: []*string{
//   					jsii.String("instanceType"),
//   				},
//   				LaunchTime: []*string{
//   					jsii.String("launchTime"),
//   				},
//   				Monitoring: &Monitoring1{
//   					State: []*string{
//   						jsii.String("state"),
//   					},
//   				},
//   				NetworkInterfaceSet: &NetworkInterfaceSet1{
//   					Items: []NetworkInterfaceSet1Item{
//   						&NetworkInterfaceSet1Item{
//   							Attachment: &Attachment{
//   								AttachmentId: []*string{
//   									jsii.String("attachmentId"),
//   								},
//   								AttachTime: []*string{
//   									jsii.String("attachTime"),
//   								},
//   								DeleteOnTermination: []*string{
//   									jsii.String("deleteOnTermination"),
//   								},
//   								DeviceIndex: []*string{
//   									jsii.String("deviceIndex"),
//   								},
//   								Status: []*string{
//   									jsii.String("status"),
//   								},
//   							},
//   							GroupSet: &GroupSet3{
//   								Items: []GroupSet2Item{
//   									&GroupSet2Item{
//   										GroupId: []*string{
//   											jsii.String("groupId"),
//   										},
//   										GroupName: []*string{
//   											jsii.String("groupName"),
//   										},
//   									},
//   								},
//   							},
//   							InterfaceType: []*string{
//   								jsii.String("interfaceType"),
//   							},
//   							Ipv6AddressesSet: []*string{
//   								jsii.String("ipv6AddressesSet"),
//   							},
//   							MacAddress: []*string{
//   								jsii.String("macAddress"),
//   							},
//   							NetworkInterfaceId: []*string{
//   								jsii.String("networkInterfaceId"),
//   							},
//   							OwnerId: []*string{
//   								jsii.String("ownerId"),
//   							},
//   							PrivateIpAddress: []*string{
//   								jsii.String("privateIpAddress"),
//   							},
//   							PrivateIpAddressesSet: &PrivateIpAddressesSet2{
//   								Item: []PrivateIpAddressesSet1Item{
//   									&PrivateIpAddressesSet1Item{
//   										Primary: []*string{
//   											jsii.String("primary"),
//   										},
//   										PrivateIpAddress: []*string{
//   											jsii.String("privateIpAddress"),
//   										},
//   									},
//   								},
//   							},
//   							SourceDestCheck: []*string{
//   								jsii.String("sourceDestCheck"),
//   							},
//   							Status: []*string{
//   								jsii.String("status"),
//   							},
//   							SubnetId: []*string{
//   								jsii.String("subnetId"),
//   							},
//   							TagSet: []*string{
//   								jsii.String("tagSet"),
//   							},
//   							VpcId: []*string{
//   								jsii.String("vpcId"),
//   							},
//   						},
//   					},
//   				},
//   				Placement: &Placement{
//   					AvailabilityZone: []*string{
//   						jsii.String("availabilityZone"),
//   					},
//   					Tenancy: []*string{
//   						jsii.String("tenancy"),
//   					},
//   				},
//   				PreviousState: &InstanceState{
//   					Code: []*string{
//   						jsii.String("code"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   				},
//   				PrivateIpAddress: []*string{
//   					jsii.String("privateIpAddress"),
//   				},
//   				ProductCodes: []*string{
//   					jsii.String("productCodes"),
//   				},
//   				RootDeviceName: []*string{
//   					jsii.String("rootDeviceName"),
//   				},
//   				RootDeviceType: []*string{
//   					jsii.String("rootDeviceType"),
//   				},
//   				SourceDestCheck: []*string{
//   					jsii.String("sourceDestCheck"),
//   				},
//   				SpotInstanceRequestId: []*string{
//   					jsii.String("spotInstanceRequestId"),
//   				},
//   				StateReason: &StateReason{
//   					Code: []*string{
//   						jsii.String("code"),
//   					},
//   					Message: []*string{
//   						jsii.String("message"),
//   					},
//   				},
//   				SubnetId: []*string{
//   					jsii.String("subnetId"),
//   				},
//   				TagSet: &TagSet{
//   					Items: []TagSpecificationSetItemItem{
//   						&TagSpecificationSetItemItem{
//   							Key: []*string{
//   								jsii.String("key"),
//   							},
//   							Value: []*string{
//   								jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				VirtualizationType: []*string{
//   					jsii.String("virtualizationType"),
//   				},
//   				VpcId: []*string{
//   					jsii.String("vpcId"),
//   				},
//   			},
//   		},
//   	},
//   	NetworkInterface: &NetworkInterface{
//   		AvailabilityZone: []*string{
//   			jsii.String("availabilityZone"),
//   		},
//   		Description: []*string{
//   			jsii.String("description"),
//   		},
//   		GroupSet: &GroupSet2{
//   			Items: []GroupSet2Item{
//   				&GroupSet2Item{
//   					GroupId: []*string{
//   						jsii.String("groupId"),
//   					},
//   					GroupName: []*string{
//   						jsii.String("groupName"),
//   					},
//   				},
//   			},
//   		},
//   		InterfaceType: []*string{
//   			jsii.String("interfaceType"),
//   		},
//   		Ipv6AddressesSet: []*string{
//   			jsii.String("ipv6AddressesSet"),
//   		},
//   		MacAddress: []*string{
//   			jsii.String("macAddress"),
//   		},
//   		NetworkInterfaceId: []*string{
//   			jsii.String("networkInterfaceId"),
//   		},
//   		OwnerId: []*string{
//   			jsii.String("ownerId"),
//   		},
//   		PrivateIpAddress: []*string{
//   			jsii.String("privateIpAddress"),
//   		},
//   		PrivateIpAddressesSet: &PrivateIpAddressesSet1{
//   			Item: []PrivateIpAddressesSet1Item{
//   				&PrivateIpAddressesSet1Item{
//   					Primary: []*string{
//   						jsii.String("primary"),
//   					},
//   					PrivateIpAddress: []*string{
//   						jsii.String("privateIpAddress"),
//   					},
//   				},
//   			},
//   		},
//   		RequesterId: []*string{
//   			jsii.String("requesterId"),
//   		},
//   		RequesterManaged: []*string{
//   			jsii.String("requesterManaged"),
//   		},
//   		SourceDestCheck: []*string{
//   			jsii.String("sourceDestCheck"),
//   		},
//   		Status: []*string{
//   			jsii.String("status"),
//   		},
//   		SubnetId: []*string{
//   			jsii.String("subnetId"),
//   		},
//   		TagSet: []*string{
//   			jsii.String("tagSet"),
//   		},
//   		VpcId: []*string{
//   			jsii.String("vpcId"),
//   		},
//   	},
//   	OwnerId: []*string{
//   		jsii.String("ownerId"),
//   	},
//   	RequesterId: []*string{
//   		jsii.String("requesterId"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	ReservationId: []*string{
//   		jsii.String("reservationId"),
//   	},
//   	Return: []*string{
//   		jsii.String("return"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_ResponseElements struct {
	// CreateFleetResponse property.
	//
	// Specify an array of string values to match this event if the actual value of CreateFleetResponse is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreateFleetResponse *InstanceEvents_AWSAPICallViaCloudTrail_CreateFleetResponse `field:"optional" json:"createFleetResponse" yaml:"createFleetResponse"`
	// CreateLaunchTemplateResponse property.
	//
	// Specify an array of string values to match this event if the actual value of CreateLaunchTemplateResponse is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreateLaunchTemplateResponse *InstanceEvents_AWSAPICallViaCloudTrail_DeleteLaunchTemplateResponse `field:"optional" json:"createLaunchTemplateResponse" yaml:"createLaunchTemplateResponse"`
	// DeleteLaunchTemplateResponse property.
	//
	// Specify an array of string values to match this event if the actual value of DeleteLaunchTemplateResponse is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeleteLaunchTemplateResponse *InstanceEvents_AWSAPICallViaCloudTrail_DeleteLaunchTemplateResponse `field:"optional" json:"deleteLaunchTemplateResponse" yaml:"deleteLaunchTemplateResponse"`
	// groupId property.
	//
	// Specify an array of string values to match this event if the actual value of groupId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupId *[]*string `field:"optional" json:"groupId" yaml:"groupId"`
	// groupSet property.
	//
	// Specify an array of string values to match this event if the actual value of groupSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupSet *[]*string `field:"optional" json:"groupSet" yaml:"groupSet"`
	// instancesSet property.
	//
	// Specify an array of string values to match this event if the actual value of instancesSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstancesSet *InstanceEvents_AWSAPICallViaCloudTrail_InstancesSet `field:"optional" json:"instancesSet" yaml:"instancesSet"`
	// networkInterface property.
	//
	// Specify an array of string values to match this event if the actual value of networkInterface is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkInterface *InstanceEvents_AWSAPICallViaCloudTrail_NetworkInterface `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// ownerId property.
	//
	// Specify an array of string values to match this event if the actual value of ownerId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OwnerId *[]*string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// requesterId property.
	//
	// Specify an array of string values to match this event if the actual value of requesterId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequesterId *[]*string `field:"optional" json:"requesterId" yaml:"requesterId"`
	// requestId property.
	//
	// Specify an array of string values to match this event if the actual value of requestId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// reservationId property.
	//
	// Specify an array of string values to match this event if the actual value of reservationId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReservationId *[]*string `field:"optional" json:"reservationId" yaml:"reservationId"`
	// _return property.
	//
	// Specify an array of string values to match this event if the actual value of _return is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Return *[]*string `field:"optional" json:"return" yaml:"return"`
}

