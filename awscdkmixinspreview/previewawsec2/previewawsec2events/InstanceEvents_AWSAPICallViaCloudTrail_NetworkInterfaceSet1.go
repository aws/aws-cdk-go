package previewawsec2events


// Type definition for NetworkInterfaceSet_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkInterfaceSet1 := &NetworkInterfaceSet1{
//   	Items: []NetworkInterfaceSet1Item{
//   		&NetworkInterfaceSet1Item{
//   			Attachment: &Attachment{
//   				AttachmentId: []*string{
//   					jsii.String("attachmentId"),
//   				},
//   				AttachTime: []*string{
//   					jsii.String("attachTime"),
//   				},
//   				DeleteOnTermination: []*string{
//   					jsii.String("deleteOnTermination"),
//   				},
//   				DeviceIndex: []*string{
//   					jsii.String("deviceIndex"),
//   				},
//   				Status: []*string{
//   					jsii.String("status"),
//   				},
//   			},
//   			GroupSet: &GroupSet3{
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
//   			PrivateIpAddressesSet: &PrivateIpAddressesSet2{
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
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_NetworkInterfaceSet1 struct {
	// items property.
	//
	// Specify an array of string values to match this event if the actual value of items is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Items *[]*InstanceEvents_AWSAPICallViaCloudTrail_NetworkInterfaceSet1Item `field:"optional" json:"items" yaml:"items"`
}

