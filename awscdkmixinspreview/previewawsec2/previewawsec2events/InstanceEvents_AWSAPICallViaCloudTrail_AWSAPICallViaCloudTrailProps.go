package previewawsec2events

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Instance aws.ec2@AWSAPICallViaCloudTrail event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var overrides interface{}
//
//   aWSAPICallViaCloudTrailProps := &AWSAPICallViaCloudTrailProps{
//   	AwsRegion: []*string{
//   		jsii.String("awsRegion"),
//   	},
//   	ErrorCode: []*string{
//   		jsii.String("errorCode"),
//   	},
//   	ErrorMessage: []*string{
//   		jsii.String("errorMessage"),
//   	},
//   	EventId: []*string{
//   		jsii.String("eventId"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	EventName: []*string{
//   		jsii.String("eventName"),
//   	},
//   	EventSource: []*string{
//   		jsii.String("eventSource"),
//   	},
//   	EventTime: []*string{
//   		jsii.String("eventTime"),
//   	},
//   	EventType: []*string{
//   		jsii.String("eventType"),
//   	},
//   	EventVersion: []*string{
//   		jsii.String("eventVersion"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	RequestParameters: &RequestParameters{
//   		AvailabilityZone: []*string{
//   			jsii.String("availabilityZone"),
//   		},
//   		BlockDeviceMapping: []*string{
//   			jsii.String("blockDeviceMapping"),
//   		},
//   		ClientToken: []*string{
//   			jsii.String("clientToken"),
//   		},
//   		CreateFleetRequest: &CreateFleetRequest{
//   			ClientToken: []*string{
//   				jsii.String("clientToken"),
//   			},
//   			ExistingInstances: &ExistingInstances{
//   				AvailabilityZone: []*string{
//   					jsii.String("availabilityZone"),
//   				},
//   				Count: []*string{
//   					jsii.String("count"),
//   				},
//   				InstanceType: []*string{
//   					jsii.String("instanceType"),
//   				},
//   				MarketOption: []*string{
//   					jsii.String("marketOption"),
//   				},
//   				OperatingSystem: []*string{
//   					jsii.String("operatingSystem"),
//   				},
//   				Tag: []*string{
//   					jsii.String("tag"),
//   				},
//   			},
//   			LaunchTemplateConfigs: &LaunchTemplateConfigs{
//   				LaunchTemplateSpecification: &LaunchTemplateSpecification{
//   					LaunchTemplateId: []*string{
//   						jsii.String("launchTemplateId"),
//   					},
//   					Version: []*string{
//   						jsii.String("version"),
//   					},
//   				},
//   				Overrides: []interface{}{
//   					overrides,
//   				},
//   				Tag: []*string{
//   					jsii.String("tag"),
//   				},
//   			},
//   			OnDemandOptions: &OnDemandOptions{
//   				AllocationStrategy: []*string{
//   					jsii.String("allocationStrategy"),
//   				},
//   				InstancePoolConstraintFilterDisabled: []*string{
//   					jsii.String("instancePoolConstraintFilterDisabled"),
//   				},
//   				MaxInstanceCount: []*string{
//   					jsii.String("maxInstanceCount"),
//   				},
//   				MaxTargetCapacity: []*string{
//   					jsii.String("maxTargetCapacity"),
//   				},
//   			},
//   			SpotOptions: &SpotOptions{
//   				AllocationStrategy: []*string{
//   					jsii.String("allocationStrategy"),
//   				},
//   				InstancePoolConstraintFilterDisabled: []*string{
//   					jsii.String("instancePoolConstraintFilterDisabled"),
//   				},
//   				InstancePoolsToUseCount: []*string{
//   					jsii.String("instancePoolsToUseCount"),
//   				},
//   				MaxInstanceCount: []*string{
//   					jsii.String("maxInstanceCount"),
//   				},
//   				MaxTargetCapacity: []*string{
//   					jsii.String("maxTargetCapacity"),
//   				},
//   			},
//   			TagSpecification: &TagSpecification{
//   				ResourceType: []*string{
//   					jsii.String("resourceType"),
//   				},
//   				Tag: &Tag{
//   					Key: []*string{
//   						jsii.String("key"),
//   					},
//   					Tag: []*string{
//   						jsii.String("tag"),
//   					},
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   			TargetCapacitySpecification: &TargetCapacitySpecification{
//   				DefaultTargetCapacityType: []*string{
//   					jsii.String("defaultTargetCapacityType"),
//   				},
//   				OnDemandTargetCapacity: []*string{
//   					jsii.String("onDemandTargetCapacity"),
//   				},
//   				SpotTargetCapacity: []*string{
//   					jsii.String("spotTargetCapacity"),
//   				},
//   				TotalTargetCapacity: []*string{
//   					jsii.String("totalTargetCapacity"),
//   				},
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   		CreateLaunchTemplateRequest: &CreateLaunchTemplateRequest{
//   			LaunchTemplateData: &LaunchTemplateData{
//   				ImageId: []*string{
//   					jsii.String("imageId"),
//   				},
//   				InstanceMarketOptions: &InstanceMarketOptions1{
//   					MarketType: []*string{
//   						jsii.String("marketType"),
//   					},
//   					SpotOptions: &SpotOptions2{
//   						MaxPrice: []*string{
//   							jsii.String("maxPrice"),
//   						},
//   						SpotInstanceType: []*string{
//   							jsii.String("spotInstanceType"),
//   						},
//   					},
//   				},
//   				InstanceType: []*string{
//   					jsii.String("instanceType"),
//   				},
//   				NetworkInterface: &NetworkInterface1{
//   					DeviceIndex: []*string{
//   						jsii.String("deviceIndex"),
//   					},
//   					SecurityGroupId: &SecurityGroupId{
//   						Content: []*string{
//   							jsii.String("content"),
//   						},
//   						Tag: []*string{
//   							jsii.String("tag"),
//   						},
//   					},
//   					SubnetId: []*string{
//   						jsii.String("subnetId"),
//   					},
//   					Tag: []*string{
//   						jsii.String("tag"),
//   					},
//   				},
//   				UserData: []*string{
//   					jsii.String("userData"),
//   				},
//   			},
//   			LaunchTemplateName: []*string{
//   				jsii.String("launchTemplateName"),
//   			},
//   		},
//   		DeleteLaunchTemplateRequest: &DeleteLaunchTemplateRequest{
//   			LaunchTemplateName: []*string{
//   				jsii.String("launchTemplateName"),
//   			},
//   		},
//   		Description: []*string{
//   			jsii.String("description"),
//   		},
//   		DisableApiTermination: []*string{
//   			jsii.String("disableApiTermination"),
//   		},
//   		GroupDescription: []*string{
//   			jsii.String("groupDescription"),
//   		},
//   		GroupId: []*string{
//   			jsii.String("groupId"),
//   		},
//   		GroupName: []*string{
//   			jsii.String("groupName"),
//   		},
//   		GroupSet: &GroupSet1{
//   			Items: []GroupSet1Item{
//   				&GroupSet1Item{
//   					GroupId: []*string{
//   						jsii.String("groupId"),
//   					},
//   				},
//   			},
//   		},
//   		InstanceMarketOptions: &InstanceMarketOptions{
//   			MarketType: []*string{
//   				jsii.String("marketType"),
//   			},
//   			SpotOptions: &SpotOptions1{
//   				MaxPrice: []*string{
//   					jsii.String("maxPrice"),
//   				},
//   				SpotInstanceType: []*string{
//   					jsii.String("spotInstanceType"),
//   				},
//   			},
//   		},
//   		InstancesSet: &InstancesSet1{
//   			Items: []InstancesSet1Item{
//   				&InstancesSet1Item{
//   					ImageId: []*string{
//   						jsii.String("imageId"),
//   					},
//   					InstanceId: []*string{
//   						jsii.String("instanceId"),
//   					},
//   					MaxCount: []*string{
//   						jsii.String("maxCount"),
//   					},
//   					MinCount: []*string{
//   						jsii.String("minCount"),
//   					},
//   				},
//   			},
//   		},
//   		InstanceType: []*string{
//   			jsii.String("instanceType"),
//   		},
//   		Ipv6AddressCount: []*string{
//   			jsii.String("ipv6AddressCount"),
//   		},
//   		LaunchTemplate: &LaunchTemplate{
//   			LaunchTemplateId: []*string{
//   				jsii.String("launchTemplateId"),
//   			},
//   			Version: []*string{
//   				jsii.String("version"),
//   			},
//   		},
//   		Monitoring: &Monitoring{
//   			Enabled: []*string{
//   				jsii.String("enabled"),
//   			},
//   		},
//   		NetworkInterfaceId: []*string{
//   			jsii.String("networkInterfaceId"),
//   		},
//   		NetworkInterfaceSet: &NetworkInterfaceSet{
//   			Items: []NetworkInterfaceSetItem{
//   				&NetworkInterfaceSetItem{
//   					DeviceIndex: []*string{
//   						jsii.String("deviceIndex"),
//   					},
//   					SubnetId: []*string{
//   						jsii.String("subnetId"),
//   					},
//   				},
//   			},
//   		},
//   		PrivateIpAddressesSet: []*string{
//   			jsii.String("privateIpAddressesSet"),
//   		},
//   		SubnetId: []*string{
//   			jsii.String("subnetId"),
//   		},
//   		TagSpecificationSet: &TagSpecificationSet{
//   			Items: []TagSpecificationSetItem{
//   				&TagSpecificationSetItem{
//   					ResourceType: []*string{
//   						jsii.String("resourceType"),
//   					},
//   					Tags: []TagSpecificationSetItemItem{
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
//   			},
//   		},
//   		UserData: []*string{
//   			jsii.String("userData"),
//   		},
//   		VpcId: []*string{
//   			jsii.String("vpcId"),
//   		},
//   	},
//   	ResponseElements: &ResponseElements{
//   		CreateFleetResponse: &CreateFleetResponse{
//   			ErrorSet: []*string{
//   				jsii.String("errorSet"),
//   			},
//   			FleetId: []*string{
//   				jsii.String("fleetId"),
//   			},
//   			FleetInstanceSet: []*string{
//   				jsii.String("fleetInstanceSet"),
//   			},
//   			RequestId: []*string{
//   				jsii.String("requestId"),
//   			},
//   			Xmlns: []*string{
//   				jsii.String("xmlns"),
//   			},
//   		},
//   		CreateLaunchTemplateResponse: &DeleteLaunchTemplateResponse{
//   			LaunchTemplate: &LaunchTemplate1{
//   				CreatedBy: []*string{
//   					jsii.String("createdBy"),
//   				},
//   				CreateTime: []*string{
//   					jsii.String("createTime"),
//   				},
//   				DefaultVersionNumber: []*string{
//   					jsii.String("defaultVersionNumber"),
//   				},
//   				LatestVersionNumber: []*string{
//   					jsii.String("latestVersionNumber"),
//   				},
//   				LaunchTemplateId: []*string{
//   					jsii.String("launchTemplateId"),
//   				},
//   				LaunchTemplateName: []*string{
//   					jsii.String("launchTemplateName"),
//   				},
//   			},
//   			RequestId: []*string{
//   				jsii.String("requestId"),
//   			},
//   			Xmlns: []*string{
//   				jsii.String("xmlns"),
//   			},
//   		},
//   		DeleteLaunchTemplateResponse: &DeleteLaunchTemplateResponse{
//   			LaunchTemplate: &LaunchTemplate1{
//   				CreatedBy: []*string{
//   					jsii.String("createdBy"),
//   				},
//   				CreateTime: []*string{
//   					jsii.String("createTime"),
//   				},
//   				DefaultVersionNumber: []*string{
//   					jsii.String("defaultVersionNumber"),
//   				},
//   				LatestVersionNumber: []*string{
//   					jsii.String("latestVersionNumber"),
//   				},
//   				LaunchTemplateId: []*string{
//   					jsii.String("launchTemplateId"),
//   				},
//   				LaunchTemplateName: []*string{
//   					jsii.String("launchTemplateName"),
//   				},
//   			},
//   			RequestId: []*string{
//   				jsii.String("requestId"),
//   			},
//   			Xmlns: []*string{
//   				jsii.String("xmlns"),
//   			},
//   		},
//   		GroupId: []*string{
//   			jsii.String("groupId"),
//   		},
//   		GroupSet: []*string{
//   			jsii.String("groupSet"),
//   		},
//   		InstancesSet: &InstancesSet{
//   			Items: []InstancesSetItem{
//   				&InstancesSetItem{
//   					AmiLaunchIndex: []*string{
//   						jsii.String("amiLaunchIndex"),
//   					},
//   					Architecture: []*string{
//   						jsii.String("architecture"),
//   					},
//   					BlockDeviceMapping: []*string{
//   						jsii.String("blockDeviceMapping"),
//   					},
//   					CapacityReservationSpecification: &CapacityReservationSpecification{
//   						CapacityReservationPreference: []*string{
//   							jsii.String("capacityReservationPreference"),
//   						},
//   					},
//   					ClientToken: []*string{
//   						jsii.String("clientToken"),
//   					},
//   					CpuOptions: &CpuOptions{
//   						CoreCount: []*string{
//   							jsii.String("coreCount"),
//   						},
//   						ThreadsPerCore: []*string{
//   							jsii.String("threadsPerCore"),
//   						},
//   					},
//   					CurrentState: &InstanceState{
//   						Code: []*string{
//   							jsii.String("code"),
//   						},
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   					},
//   					EbsOptimized: []*string{
//   						jsii.String("ebsOptimized"),
//   					},
//   					EnclaveOptions: &EnclaveOptions{
//   						Enabled: []*string{
//   							jsii.String("enabled"),
//   						},
//   					},
//   					GroupSet: &GroupSet2{
//   						Items: []GroupSet2Item{
//   							&GroupSet2Item{
//   								GroupId: []*string{
//   									jsii.String("groupId"),
//   								},
//   								GroupName: []*string{
//   									jsii.String("groupName"),
//   								},
//   							},
//   						},
//   					},
//   					Hypervisor: []*string{
//   						jsii.String("hypervisor"),
//   					},
//   					ImageId: []*string{
//   						jsii.String("imageId"),
//   					},
//   					InstanceId: []*string{
//   						jsii.String("instanceId"),
//   					},
//   					InstanceLifecycle: []*string{
//   						jsii.String("instanceLifecycle"),
//   					},
//   					InstanceState: &InstanceState{
//   						Code: []*string{
//   							jsii.String("code"),
//   						},
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   					},
//   					InstanceType: []*string{
//   						jsii.String("instanceType"),
//   					},
//   					LaunchTime: []*string{
//   						jsii.String("launchTime"),
//   					},
//   					Monitoring: &Monitoring1{
//   						State: []*string{
//   							jsii.String("state"),
//   						},
//   					},
//   					NetworkInterfaceSet: &NetworkInterfaceSet1{
//   						Items: []NetworkInterfaceSet1Item{
//   							&NetworkInterfaceSet1Item{
//   								Attachment: &Attachment{
//   									AttachmentId: []*string{
//   										jsii.String("attachmentId"),
//   									},
//   									AttachTime: []*string{
//   										jsii.String("attachTime"),
//   									},
//   									DeleteOnTermination: []*string{
//   										jsii.String("deleteOnTermination"),
//   									},
//   									DeviceIndex: []*string{
//   										jsii.String("deviceIndex"),
//   									},
//   									Status: []*string{
//   										jsii.String("status"),
//   									},
//   								},
//   								GroupSet: &GroupSet3{
//   									Items: []GroupSet2Item{
//   										&GroupSet2Item{
//   											GroupId: []*string{
//   												jsii.String("groupId"),
//   											},
//   											GroupName: []*string{
//   												jsii.String("groupName"),
//   											},
//   										},
//   									},
//   								},
//   								InterfaceType: []*string{
//   									jsii.String("interfaceType"),
//   								},
//   								Ipv6AddressesSet: []*string{
//   									jsii.String("ipv6AddressesSet"),
//   								},
//   								MacAddress: []*string{
//   									jsii.String("macAddress"),
//   								},
//   								NetworkInterfaceId: []*string{
//   									jsii.String("networkInterfaceId"),
//   								},
//   								OwnerId: []*string{
//   									jsii.String("ownerId"),
//   								},
//   								PrivateIpAddress: []*string{
//   									jsii.String("privateIpAddress"),
//   								},
//   								PrivateIpAddressesSet: &PrivateIpAddressesSet2{
//   									Item: []PrivateIpAddressesSet1Item{
//   										&PrivateIpAddressesSet1Item{
//   											Primary: []*string{
//   												jsii.String("primary"),
//   											},
//   											PrivateIpAddress: []*string{
//   												jsii.String("privateIpAddress"),
//   											},
//   										},
//   									},
//   								},
//   								SourceDestCheck: []*string{
//   									jsii.String("sourceDestCheck"),
//   								},
//   								Status: []*string{
//   									jsii.String("status"),
//   								},
//   								SubnetId: []*string{
//   									jsii.String("subnetId"),
//   								},
//   								TagSet: []*string{
//   									jsii.String("tagSet"),
//   								},
//   								VpcId: []*string{
//   									jsii.String("vpcId"),
//   								},
//   							},
//   						},
//   					},
//   					Placement: &Placement{
//   						AvailabilityZone: []*string{
//   							jsii.String("availabilityZone"),
//   						},
//   						Tenancy: []*string{
//   							jsii.String("tenancy"),
//   						},
//   					},
//   					PreviousState: &InstanceState{
//   						Code: []*string{
//   							jsii.String("code"),
//   						},
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   					},
//   					PrivateIpAddress: []*string{
//   						jsii.String("privateIpAddress"),
//   					},
//   					ProductCodes: []*string{
//   						jsii.String("productCodes"),
//   					},
//   					RootDeviceName: []*string{
//   						jsii.String("rootDeviceName"),
//   					},
//   					RootDeviceType: []*string{
//   						jsii.String("rootDeviceType"),
//   					},
//   					SourceDestCheck: []*string{
//   						jsii.String("sourceDestCheck"),
//   					},
//   					SpotInstanceRequestId: []*string{
//   						jsii.String("spotInstanceRequestId"),
//   					},
//   					StateReason: &StateReason{
//   						Code: []*string{
//   							jsii.String("code"),
//   						},
//   						Message: []*string{
//   							jsii.String("message"),
//   						},
//   					},
//   					SubnetId: []*string{
//   						jsii.String("subnetId"),
//   					},
//   					TagSet: &TagSet{
//   						Items: []TagSpecificationSetItemItem{
//   							&TagSpecificationSetItemItem{
//   								Key: []*string{
//   									jsii.String("key"),
//   								},
//   								Value: []*string{
//   									jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					VirtualizationType: []*string{
//   						jsii.String("virtualizationType"),
//   					},
//   					VpcId: []*string{
//   						jsii.String("vpcId"),
//   					},
//   				},
//   			},
//   		},
//   		NetworkInterface: &NetworkInterface{
//   			AvailabilityZone: []*string{
//   				jsii.String("availabilityZone"),
//   			},
//   			Description: []*string{
//   				jsii.String("description"),
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
//   			InterfaceType: []*string{
//   				jsii.String("interfaceType"),
//   			},
//   			Ipv6AddressesSet: []*string{
//   				jsii.String("ipv6AddressesSet"),
//   			},
//   			MacAddress: []*string{
//   				jsii.String("macAddress"),
//   			},
//   			NetworkInterfaceId: []*string{
//   				jsii.String("networkInterfaceId"),
//   			},
//   			OwnerId: []*string{
//   				jsii.String("ownerId"),
//   			},
//   			PrivateIpAddress: []*string{
//   				jsii.String("privateIpAddress"),
//   			},
//   			PrivateIpAddressesSet: &PrivateIpAddressesSet1{
//   				Item: []PrivateIpAddressesSet1Item{
//   					&PrivateIpAddressesSet1Item{
//   						Primary: []*string{
//   							jsii.String("primary"),
//   						},
//   						PrivateIpAddress: []*string{
//   							jsii.String("privateIpAddress"),
//   						},
//   					},
//   				},
//   			},
//   			RequesterId: []*string{
//   				jsii.String("requesterId"),
//   			},
//   			RequesterManaged: []*string{
//   				jsii.String("requesterManaged"),
//   			},
//   			SourceDestCheck: []*string{
//   				jsii.String("sourceDestCheck"),
//   			},
//   			Status: []*string{
//   				jsii.String("status"),
//   			},
//   			SubnetId: []*string{
//   				jsii.String("subnetId"),
//   			},
//   			TagSet: []*string{
//   				jsii.String("tagSet"),
//   			},
//   			VpcId: []*string{
//   				jsii.String("vpcId"),
//   			},
//   		},
//   		OwnerId: []*string{
//   			jsii.String("ownerId"),
//   		},
//   		RequesterId: []*string{
//   			jsii.String("requesterId"),
//   		},
//   		RequestId: []*string{
//   			jsii.String("requestId"),
//   		},
//   		ReservationId: []*string{
//   			jsii.String("reservationId"),
//   		},
//   		Return: []*string{
//   			jsii.String("return"),
//   		},
//   	},
//   	SourceIpAddress: []*string{
//   		jsii.String("sourceIpAddress"),
//   	},
//   	UserAgent: []*string{
//   		jsii.String("userAgent"),
//   	},
//   	UserIdentity: &UserIdentity{
//   		AccessKeyId: []*string{
//   			jsii.String("accessKeyId"),
//   		},
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		InvokedBy: []*string{
//   			jsii.String("invokedBy"),
//   		},
//   		PrincipalId: []*string{
//   			jsii.String("principalId"),
//   		},
//   		SessionContext: &SessionContext{
//   			Attributes: &Attributes{
//   				CreationDate: []*string{
//   					jsii.String("creationDate"),
//   				},
//   				MfaAuthenticated: []*string{
//   					jsii.String("mfaAuthenticated"),
//   				},
//   			},
//   			SessionIssuer: &SessionIssuer{
//   				AccountId: []*string{
//   					jsii.String("accountId"),
//   				},
//   				Arn: []*string{
//   					jsii.String("arn"),
//   				},
//   				PrincipalId: []*string{
//   					jsii.String("principalId"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   				UserName: []*string{
//   					jsii.String("userName"),
//   				},
//   			},
//   			WebIdFederationData: []*string{
//   				jsii.String("webIdFederationData"),
//   			},
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps struct {
	// awsRegion property.
	//
	// Specify an array of string values to match this event if the actual value of awsRegion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsRegion *[]*string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// errorCode property.
	//
	// Specify an array of string values to match this event if the actual value of errorCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorCode *[]*string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// errorMessage property.
	//
	// Specify an array of string values to match this event if the actual value of errorMessage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorMessage *[]*string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// eventID property.
	//
	// Specify an array of string values to match this event if the actual value of eventID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventId *[]*string `field:"optional" json:"eventId" yaml:"eventId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// eventName property.
	//
	// Specify an array of string values to match this event if the actual value of eventName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventName *[]*string `field:"optional" json:"eventName" yaml:"eventName"`
	// eventSource property.
	//
	// Specify an array of string values to match this event if the actual value of eventSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventSource *[]*string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// eventTime property.
	//
	// Specify an array of string values to match this event if the actual value of eventTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventTime *[]*string `field:"optional" json:"eventTime" yaml:"eventTime"`
	// eventType property.
	//
	// Specify an array of string values to match this event if the actual value of eventType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// eventVersion property.
	//
	// Specify an array of string values to match this event if the actual value of eventVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventVersion *[]*string `field:"optional" json:"eventVersion" yaml:"eventVersion"`
	// requestID property.
	//
	// Specify an array of string values to match this event if the actual value of requestID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// requestParameters property.
	//
	// Specify an array of string values to match this event if the actual value of requestParameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestParameters *InstanceEvents_AWSAPICallViaCloudTrail_RequestParameters `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// responseElements property.
	//
	// Specify an array of string values to match this event if the actual value of responseElements is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResponseElements *InstanceEvents_AWSAPICallViaCloudTrail_ResponseElements `field:"optional" json:"responseElements" yaml:"responseElements"`
	// sourceIPAddress property.
	//
	// Specify an array of string values to match this event if the actual value of sourceIPAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIpAddress *[]*string `field:"optional" json:"sourceIpAddress" yaml:"sourceIpAddress"`
	// userAgent property.
	//
	// Specify an array of string values to match this event if the actual value of userAgent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserAgent *[]*string `field:"optional" json:"userAgent" yaml:"userAgent"`
	// userIdentity property.
	//
	// Specify an array of string values to match this event if the actual value of userIdentity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserIdentity *InstanceEvents_AWSAPICallViaCloudTrail_UserIdentity `field:"optional" json:"userIdentity" yaml:"userIdentity"`
}

