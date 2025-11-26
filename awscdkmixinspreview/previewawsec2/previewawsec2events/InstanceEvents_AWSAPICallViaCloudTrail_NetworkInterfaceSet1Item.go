package previewawsec2events


// Type definition for NetworkInterfaceSet_1Item.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkInterfaceSet1Item := &NetworkInterfaceSet1Item{
//   	Attachment: &Attachment{
//   		AttachmentId: []*string{
//   			jsii.String("attachmentId"),
//   		},
//   		AttachTime: []*string{
//   			jsii.String("attachTime"),
//   		},
//   		DeleteOnTermination: []*string{
//   			jsii.String("deleteOnTermination"),
//   		},
//   		DeviceIndex: []*string{
//   			jsii.String("deviceIndex"),
//   		},
//   		Status: []*string{
//   			jsii.String("status"),
//   		},
//   	},
//   	GroupSet: &GroupSet3{
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
//   	InterfaceType: []*string{
//   		jsii.String("interfaceType"),
//   	},
//   	Ipv6AddressesSet: []*string{
//   		jsii.String("ipv6AddressesSet"),
//   	},
//   	MacAddress: []*string{
//   		jsii.String("macAddress"),
//   	},
//   	NetworkInterfaceId: []*string{
//   		jsii.String("networkInterfaceId"),
//   	},
//   	OwnerId: []*string{
//   		jsii.String("ownerId"),
//   	},
//   	PrivateIpAddress: []*string{
//   		jsii.String("privateIpAddress"),
//   	},
//   	PrivateIpAddressesSet: &PrivateIpAddressesSet2{
//   		Item: []PrivateIpAddressesSet1Item{
//   			&PrivateIpAddressesSet1Item{
//   				Primary: []*string{
//   					jsii.String("primary"),
//   				},
//   				PrivateIpAddress: []*string{
//   					jsii.String("privateIpAddress"),
//   				},
//   			},
//   		},
//   	},
//   	SourceDestCheck: []*string{
//   		jsii.String("sourceDestCheck"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	SubnetId: []*string{
//   		jsii.String("subnetId"),
//   	},
//   	TagSet: []*string{
//   		jsii.String("tagSet"),
//   	},
//   	VpcId: []*string{
//   		jsii.String("vpcId"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_NetworkInterfaceSet1Item struct {
	// attachment property.
	//
	// Specify an array of string values to match this event if the actual value of attachment is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Attachment *InstanceEvents_AWSAPICallViaCloudTrail_Attachment `field:"optional" json:"attachment" yaml:"attachment"`
	// groupSet property.
	//
	// Specify an array of string values to match this event if the actual value of groupSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupSet *InstanceEvents_AWSAPICallViaCloudTrail_GroupSet3 `field:"optional" json:"groupSet" yaml:"groupSet"`
	// interfaceType property.
	//
	// Specify an array of string values to match this event if the actual value of interfaceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InterfaceType *[]*string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// ipv6AddressesSet property.
	//
	// Specify an array of string values to match this event if the actual value of ipv6AddressesSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Ipv6AddressesSet *[]*string `field:"optional" json:"ipv6AddressesSet" yaml:"ipv6AddressesSet"`
	// macAddress property.
	//
	// Specify an array of string values to match this event if the actual value of macAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MacAddress *[]*string `field:"optional" json:"macAddress" yaml:"macAddress"`
	// networkInterfaceId property.
	//
	// Specify an array of string values to match this event if the actual value of networkInterfaceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkInterfaceId *[]*string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// ownerId property.
	//
	// Specify an array of string values to match this event if the actual value of ownerId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OwnerId *[]*string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// privateIpAddress property.
	//
	// Specify an array of string values to match this event if the actual value of privateIpAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrivateIpAddress *[]*string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// privateIpAddressesSet property.
	//
	// Specify an array of string values to match this event if the actual value of privateIpAddressesSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrivateIpAddressesSet *InstanceEvents_AWSAPICallViaCloudTrail_PrivateIpAddressesSet2 `field:"optional" json:"privateIpAddressesSet" yaml:"privateIpAddressesSet"`
	// sourceDestCheck property.
	//
	// Specify an array of string values to match this event if the actual value of sourceDestCheck is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceDestCheck *[]*string `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
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
	TagSet *[]*string `field:"optional" json:"tagSet" yaml:"tagSet"`
	// vpcId property.
	//
	// Specify an array of string values to match this event if the actual value of vpcId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VpcId *[]*string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

