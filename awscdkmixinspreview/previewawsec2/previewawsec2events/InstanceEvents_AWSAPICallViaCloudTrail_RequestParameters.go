package previewawsec2events


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var overrides interface{}
//
//   requestParameters := &RequestParameters{
//   	AvailabilityZone: []*string{
//   		jsii.String("availabilityZone"),
//   	},
//   	BlockDeviceMapping: []*string{
//   		jsii.String("blockDeviceMapping"),
//   	},
//   	ClientToken: []*string{
//   		jsii.String("clientToken"),
//   	},
//   	CreateFleetRequest: &CreateFleetRequest{
//   		ClientToken: []*string{
//   			jsii.String("clientToken"),
//   		},
//   		ExistingInstances: &ExistingInstances{
//   			AvailabilityZone: []*string{
//   				jsii.String("availabilityZone"),
//   			},
//   			Count: []*string{
//   				jsii.String("count"),
//   			},
//   			InstanceType: []*string{
//   				jsii.String("instanceType"),
//   			},
//   			MarketOption: []*string{
//   				jsii.String("marketOption"),
//   			},
//   			OperatingSystem: []*string{
//   				jsii.String("operatingSystem"),
//   			},
//   			Tag: []*string{
//   				jsii.String("tag"),
//   			},
//   		},
//   		LaunchTemplateConfigs: &LaunchTemplateConfigs{
//   			LaunchTemplateSpecification: &LaunchTemplateSpecification{
//   				LaunchTemplateId: []*string{
//   					jsii.String("launchTemplateId"),
//   				},
//   				Version: []*string{
//   					jsii.String("version"),
//   				},
//   			},
//   			Overrides: []interface{}{
//   				overrides,
//   			},
//   			Tag: []*string{
//   				jsii.String("tag"),
//   			},
//   		},
//   		OnDemandOptions: &OnDemandOptions{
//   			AllocationStrategy: []*string{
//   				jsii.String("allocationStrategy"),
//   			},
//   			InstancePoolConstraintFilterDisabled: []*string{
//   				jsii.String("instancePoolConstraintFilterDisabled"),
//   			},
//   			MaxInstanceCount: []*string{
//   				jsii.String("maxInstanceCount"),
//   			},
//   			MaxTargetCapacity: []*string{
//   				jsii.String("maxTargetCapacity"),
//   			},
//   		},
//   		SpotOptions: &SpotOptions{
//   			AllocationStrategy: []*string{
//   				jsii.String("allocationStrategy"),
//   			},
//   			InstancePoolConstraintFilterDisabled: []*string{
//   				jsii.String("instancePoolConstraintFilterDisabled"),
//   			},
//   			InstancePoolsToUseCount: []*string{
//   				jsii.String("instancePoolsToUseCount"),
//   			},
//   			MaxInstanceCount: []*string{
//   				jsii.String("maxInstanceCount"),
//   			},
//   			MaxTargetCapacity: []*string{
//   				jsii.String("maxTargetCapacity"),
//   			},
//   		},
//   		TagSpecification: &TagSpecification{
//   			ResourceType: []*string{
//   				jsii.String("resourceType"),
//   			},
//   			Tag: &Tag{
//   				Key: []*string{
//   					jsii.String("key"),
//   				},
//   				Tag: []*string{
//   					jsii.String("tag"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   		TargetCapacitySpecification: &TargetCapacitySpecification{
//   			DefaultTargetCapacityType: []*string{
//   				jsii.String("defaultTargetCapacityType"),
//   			},
//   			OnDemandTargetCapacity: []*string{
//   				jsii.String("onDemandTargetCapacity"),
//   			},
//   			SpotTargetCapacity: []*string{
//   				jsii.String("spotTargetCapacity"),
//   			},
//   			TotalTargetCapacity: []*string{
//   				jsii.String("totalTargetCapacity"),
//   			},
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	CreateLaunchTemplateRequest: &CreateLaunchTemplateRequest{
//   		LaunchTemplateData: &LaunchTemplateData{
//   			ImageId: []*string{
//   				jsii.String("imageId"),
//   			},
//   			InstanceMarketOptions: &InstanceMarketOptions1{
//   				MarketType: []*string{
//   					jsii.String("marketType"),
//   				},
//   				SpotOptions: &SpotOptions2{
//   					MaxPrice: []*string{
//   						jsii.String("maxPrice"),
//   					},
//   					SpotInstanceType: []*string{
//   						jsii.String("spotInstanceType"),
//   					},
//   				},
//   			},
//   			InstanceType: []*string{
//   				jsii.String("instanceType"),
//   			},
//   			NetworkInterface: &NetworkInterface1{
//   				DeviceIndex: []*string{
//   					jsii.String("deviceIndex"),
//   				},
//   				SecurityGroupId: &SecurityGroupId{
//   					Content: []*string{
//   						jsii.String("content"),
//   					},
//   					Tag: []*string{
//   						jsii.String("tag"),
//   					},
//   				},
//   				SubnetId: []*string{
//   					jsii.String("subnetId"),
//   				},
//   				Tag: []*string{
//   					jsii.String("tag"),
//   				},
//   			},
//   			UserData: []*string{
//   				jsii.String("userData"),
//   			},
//   		},
//   		LaunchTemplateName: []*string{
//   			jsii.String("launchTemplateName"),
//   		},
//   	},
//   	DeleteLaunchTemplateRequest: &DeleteLaunchTemplateRequest{
//   		LaunchTemplateName: []*string{
//   			jsii.String("launchTemplateName"),
//   		},
//   	},
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	DisableApiTermination: []*string{
//   		jsii.String("disableApiTermination"),
//   	},
//   	GroupDescription: []*string{
//   		jsii.String("groupDescription"),
//   	},
//   	GroupId: []*string{
//   		jsii.String("groupId"),
//   	},
//   	GroupName: []*string{
//   		jsii.String("groupName"),
//   	},
//   	GroupSet: &GroupSet1{
//   		Items: []GroupSet1Item{
//   			&GroupSet1Item{
//   				GroupId: []*string{
//   					jsii.String("groupId"),
//   				},
//   			},
//   		},
//   	},
//   	InstanceMarketOptions: &InstanceMarketOptions{
//   		MarketType: []*string{
//   			jsii.String("marketType"),
//   		},
//   		SpotOptions: &SpotOptions1{
//   			MaxPrice: []*string{
//   				jsii.String("maxPrice"),
//   			},
//   			SpotInstanceType: []*string{
//   				jsii.String("spotInstanceType"),
//   			},
//   		},
//   	},
//   	InstancesSet: &InstancesSet1{
//   		Items: []InstancesSet1Item{
//   			&InstancesSet1Item{
//   				ImageId: []*string{
//   					jsii.String("imageId"),
//   				},
//   				InstanceId: []*string{
//   					jsii.String("instanceId"),
//   				},
//   				MaxCount: []*string{
//   					jsii.String("maxCount"),
//   				},
//   				MinCount: []*string{
//   					jsii.String("minCount"),
//   				},
//   			},
//   		},
//   	},
//   	InstanceType: []*string{
//   		jsii.String("instanceType"),
//   	},
//   	Ipv6AddressCount: []*string{
//   		jsii.String("ipv6AddressCount"),
//   	},
//   	LaunchTemplate: &LaunchTemplate{
//   		LaunchTemplateId: []*string{
//   			jsii.String("launchTemplateId"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	Monitoring: &Monitoring{
//   		Enabled: []*string{
//   			jsii.String("enabled"),
//   		},
//   	},
//   	NetworkInterfaceId: []*string{
//   		jsii.String("networkInterfaceId"),
//   	},
//   	NetworkInterfaceSet: &NetworkInterfaceSet{
//   		Items: []NetworkInterfaceSetItem{
//   			&NetworkInterfaceSetItem{
//   				DeviceIndex: []*string{
//   					jsii.String("deviceIndex"),
//   				},
//   				SubnetId: []*string{
//   					jsii.String("subnetId"),
//   				},
//   			},
//   		},
//   	},
//   	PrivateIpAddressesSet: []*string{
//   		jsii.String("privateIpAddressesSet"),
//   	},
//   	SubnetId: []*string{
//   		jsii.String("subnetId"),
//   	},
//   	TagSpecificationSet: &TagSpecificationSet{
//   		Items: []TagSpecificationSetItem{
//   			&TagSpecificationSetItem{
//   				ResourceType: []*string{
//   					jsii.String("resourceType"),
//   				},
//   				Tags: []TagSpecificationSetItemItem{
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
//   		},
//   	},
//   	UserData: []*string{
//   		jsii.String("userData"),
//   	},
//   	VpcId: []*string{
//   		jsii.String("vpcId"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_RequestParameters struct {
	// availabilityZone property.
	//
	// Specify an array of string values to match this event if the actual value of availabilityZone is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AvailabilityZone *[]*string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// blockDeviceMapping property.
	//
	// Specify an array of string values to match this event if the actual value of blockDeviceMapping is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BlockDeviceMapping *[]*string `field:"optional" json:"blockDeviceMapping" yaml:"blockDeviceMapping"`
	// clientToken property.
	//
	// Specify an array of string values to match this event if the actual value of clientToken is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientToken *[]*string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// CreateFleetRequest property.
	//
	// Specify an array of string values to match this event if the actual value of CreateFleetRequest is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreateFleetRequest *InstanceEvents_AWSAPICallViaCloudTrail_CreateFleetRequest `field:"optional" json:"createFleetRequest" yaml:"createFleetRequest"`
	// CreateLaunchTemplateRequest property.
	//
	// Specify an array of string values to match this event if the actual value of CreateLaunchTemplateRequest is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreateLaunchTemplateRequest *InstanceEvents_AWSAPICallViaCloudTrail_CreateLaunchTemplateRequest `field:"optional" json:"createLaunchTemplateRequest" yaml:"createLaunchTemplateRequest"`
	// DeleteLaunchTemplateRequest property.
	//
	// Specify an array of string values to match this event if the actual value of DeleteLaunchTemplateRequest is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeleteLaunchTemplateRequest *InstanceEvents_AWSAPICallViaCloudTrail_DeleteLaunchTemplateRequest `field:"optional" json:"deleteLaunchTemplateRequest" yaml:"deleteLaunchTemplateRequest"`
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// disableApiTermination property.
	//
	// Specify an array of string values to match this event if the actual value of disableApiTermination is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DisableApiTermination *[]*string `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// groupDescription property.
	//
	// Specify an array of string values to match this event if the actual value of groupDescription is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupDescription *[]*string `field:"optional" json:"groupDescription" yaml:"groupDescription"`
	// groupId property.
	//
	// Specify an array of string values to match this event if the actual value of groupId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupId *[]*string `field:"optional" json:"groupId" yaml:"groupId"`
	// groupName property.
	//
	// Specify an array of string values to match this event if the actual value of groupName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupName *[]*string `field:"optional" json:"groupName" yaml:"groupName"`
	// groupSet property.
	//
	// Specify an array of string values to match this event if the actual value of groupSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupSet *InstanceEvents_AWSAPICallViaCloudTrail_GroupSet1 `field:"optional" json:"groupSet" yaml:"groupSet"`
	// instanceMarketOptions property.
	//
	// Specify an array of string values to match this event if the actual value of instanceMarketOptions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceMarketOptions *InstanceEvents_AWSAPICallViaCloudTrail_InstanceMarketOptions `field:"optional" json:"instanceMarketOptions" yaml:"instanceMarketOptions"`
	// instancesSet property.
	//
	// Specify an array of string values to match this event if the actual value of instancesSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstancesSet *InstanceEvents_AWSAPICallViaCloudTrail_InstancesSet1 `field:"optional" json:"instancesSet" yaml:"instancesSet"`
	// instanceType property.
	//
	// Specify an array of string values to match this event if the actual value of instanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// ipv6AddressCount property.
	//
	// Specify an array of string values to match this event if the actual value of ipv6AddressCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Ipv6AddressCount *[]*string `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// launchTemplate property.
	//
	// Specify an array of string values to match this event if the actual value of launchTemplate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplate *InstanceEvents_AWSAPICallViaCloudTrail_LaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// monitoring property.
	//
	// Specify an array of string values to match this event if the actual value of monitoring is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Monitoring *InstanceEvents_AWSAPICallViaCloudTrail_Monitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// networkInterfaceId property.
	//
	// Specify an array of string values to match this event if the actual value of networkInterfaceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkInterfaceId *[]*string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// networkInterfaceSet property.
	//
	// Specify an array of string values to match this event if the actual value of networkInterfaceSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkInterfaceSet *InstanceEvents_AWSAPICallViaCloudTrail_NetworkInterfaceSet `field:"optional" json:"networkInterfaceSet" yaml:"networkInterfaceSet"`
	// privateIpAddressesSet property.
	//
	// Specify an array of string values to match this event if the actual value of privateIpAddressesSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrivateIpAddressesSet *[]*string `field:"optional" json:"privateIpAddressesSet" yaml:"privateIpAddressesSet"`
	// subnetId property.
	//
	// Specify an array of string values to match this event if the actual value of subnetId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SubnetId *[]*string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// tagSpecificationSet property.
	//
	// Specify an array of string values to match this event if the actual value of tagSpecificationSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TagSpecificationSet *InstanceEvents_AWSAPICallViaCloudTrail_TagSpecificationSet `field:"optional" json:"tagSpecificationSet" yaml:"tagSpecificationSet"`
	// userData property.
	//
	// Specify an array of string values to match this event if the actual value of userData is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserData *[]*string `field:"optional" json:"userData" yaml:"userData"`
	// vpcId property.
	//
	// Specify an array of string values to match this event if the actual value of vpcId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VpcId *[]*string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

