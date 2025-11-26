package previewawsec2events


// Type definition for InstancesSet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instancesSet := &InstancesSet{
//   	Items: []InstancesSetItem{
//   		&InstancesSetItem{
//   			AmiLaunchIndex: []*string{
//   				jsii.String("amiLaunchIndex"),
//   			},
//   			Architecture: []*string{
//   				jsii.String("architecture"),
//   			},
//   			BlockDeviceMapping: []*string{
//   				jsii.String("blockDeviceMapping"),
//   			},
//   			CapacityReservationSpecification: &CapacityReservationSpecification{
//   				CapacityReservationPreference: []*string{
//   					jsii.String("capacityReservationPreference"),
//   				},
//   			},
//   			ClientToken: []*string{
//   				jsii.String("clientToken"),
//   			},
//   			CpuOptions: &CpuOptions{
//   				CoreCount: []*string{
//   					jsii.String("coreCount"),
//   				},
//   				ThreadsPerCore: []*string{
//   					jsii.String("threadsPerCore"),
//   				},
//   			},
//   			CurrentState: &InstanceState{
//   				Code: []*string{
//   					jsii.String("code"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   			},
//   			EbsOptimized: []*string{
//   				jsii.String("ebsOptimized"),
//   			},
//   			EnclaveOptions: &EnclaveOptions{
//   				Enabled: []*string{
//   					jsii.String("enabled"),
//   				},
//   			},
//   			GroupSet: &GroupSet2{
//   				Items: []GroupSet2Item{
//   					&GroupSet2Item{
//   						GroupId: []*string{
//   							jsii.String("groupId"),
//   						},
//   						GroupName: []*string{
//   							jsii.String("groupName"),
//   						},
//   					},
//   				},
//   			},
//   			Hypervisor: []*string{
//   				jsii.String("hypervisor"),
//   			},
//   			ImageId: []*string{
//   				jsii.String("imageId"),
//   			},
//   			InstanceId: []*string{
//   				jsii.String("instanceId"),
//   			},
//   			InstanceLifecycle: []*string{
//   				jsii.String("instanceLifecycle"),
//   			},
//   			InstanceState: &InstanceState{
//   				Code: []*string{
//   					jsii.String("code"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   			},
//   			InstanceType: []*string{
//   				jsii.String("instanceType"),
//   			},
//   			LaunchTime: []*string{
//   				jsii.String("launchTime"),
//   			},
//   			Monitoring: &Monitoring1{
//   				State: []*string{
//   					jsii.String("state"),
//   				},
//   			},
//   			NetworkInterfaceSet: &NetworkInterfaceSet1{
//   				Items: []NetworkInterfaceSet1Item{
//   					&NetworkInterfaceSet1Item{
//   						Attachment: &Attachment{
//   							AttachmentId: []*string{
//   								jsii.String("attachmentId"),
//   							},
//   							AttachTime: []*string{
//   								jsii.String("attachTime"),
//   							},
//   							DeleteOnTermination: []*string{
//   								jsii.String("deleteOnTermination"),
//   							},
//   							DeviceIndex: []*string{
//   								jsii.String("deviceIndex"),
//   							},
//   							Status: []*string{
//   								jsii.String("status"),
//   							},
//   						},
//   						GroupSet: &GroupSet3{
//   							Items: []GroupSet2Item{
//   								&GroupSet2Item{
//   									GroupId: []*string{
//   										jsii.String("groupId"),
//   									},
//   									GroupName: []*string{
//   										jsii.String("groupName"),
//   									},
//   								},
//   							},
//   						},
//   						InterfaceType: []*string{
//   							jsii.String("interfaceType"),
//   						},
//   						Ipv6AddressesSet: []*string{
//   							jsii.String("ipv6AddressesSet"),
//   						},
//   						MacAddress: []*string{
//   							jsii.String("macAddress"),
//   						},
//   						NetworkInterfaceId: []*string{
//   							jsii.String("networkInterfaceId"),
//   						},
//   						OwnerId: []*string{
//   							jsii.String("ownerId"),
//   						},
//   						PrivateIpAddress: []*string{
//   							jsii.String("privateIpAddress"),
//   						},
//   						PrivateIpAddressesSet: &PrivateIpAddressesSet2{
//   							Item: []PrivateIpAddressesSet1Item{
//   								&PrivateIpAddressesSet1Item{
//   									Primary: []*string{
//   										jsii.String("primary"),
//   									},
//   									PrivateIpAddress: []*string{
//   										jsii.String("privateIpAddress"),
//   									},
//   								},
//   							},
//   						},
//   						SourceDestCheck: []*string{
//   							jsii.String("sourceDestCheck"),
//   						},
//   						Status: []*string{
//   							jsii.String("status"),
//   						},
//   						SubnetId: []*string{
//   							jsii.String("subnetId"),
//   						},
//   						TagSet: []*string{
//   							jsii.String("tagSet"),
//   						},
//   						VpcId: []*string{
//   							jsii.String("vpcId"),
//   						},
//   					},
//   				},
//   			},
//   			Placement: &Placement{
//   				AvailabilityZone: []*string{
//   					jsii.String("availabilityZone"),
//   				},
//   				Tenancy: []*string{
//   					jsii.String("tenancy"),
//   				},
//   			},
//   			PreviousState: &InstanceState{
//   				Code: []*string{
//   					jsii.String("code"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   			},
//   			PrivateIpAddress: []*string{
//   				jsii.String("privateIpAddress"),
//   			},
//   			ProductCodes: []*string{
//   				jsii.String("productCodes"),
//   			},
//   			RootDeviceName: []*string{
//   				jsii.String("rootDeviceName"),
//   			},
//   			RootDeviceType: []*string{
//   				jsii.String("rootDeviceType"),
//   			},
//   			SourceDestCheck: []*string{
//   				jsii.String("sourceDestCheck"),
//   			},
//   			SpotInstanceRequestId: []*string{
//   				jsii.String("spotInstanceRequestId"),
//   			},
//   			StateReason: &StateReason{
//   				Code: []*string{
//   					jsii.String("code"),
//   				},
//   				Message: []*string{
//   					jsii.String("message"),
//   				},
//   			},
//   			SubnetId: []*string{
//   				jsii.String("subnetId"),
//   			},
//   			TagSet: &TagSet{
//   				Items: []TagSpecificationSetItemItem{
//   					&TagSpecificationSetItemItem{
//   						Key: []*string{
//   							jsii.String("key"),
//   						},
//   						Value: []*string{
//   							jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			VirtualizationType: []*string{
//   				jsii.String("virtualizationType"),
//   			},
//   			VpcId: []*string{
//   				jsii.String("vpcId"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_InstancesSet struct {
	// items property.
	//
	// Specify an array of string values to match this event if the actual value of items is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Items *[]*InstanceEvents_AWSAPICallViaCloudTrail_InstancesSetItem `field:"optional" json:"items" yaml:"items"`
}

