package previewawsec2events


// Type definition for NetworkInterface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkInterface := &NetworkInterface{
//   	AvailabilityZone: []*string{
//   		jsii.String("availabilityZone"),
//   	},
//   	Description: []*string{
//   		jsii.String("description"),
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
//   	PrivateIpAddressesSet: &PrivateIpAddressesSet1{
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
//   	RequesterId: []*string{
//   		jsii.String("requesterId"),
//   	},
//   	RequesterManaged: []*string{
//   		jsii.String("requesterManaged"),
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
type InstanceEvents_AWSAPICallViaCloudTrail_NetworkInterface struct {
	// availabilityZone property.
	//
	// Specify an array of string values to match this event if the actual value of availabilityZone is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AvailabilityZone *[]*string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// groupSet property.
	//
	// Specify an array of string values to match this event if the actual value of groupSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupSet *InstanceEvents_AWSAPICallViaCloudTrail_GroupSet2 `field:"optional" json:"groupSet" yaml:"groupSet"`
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
	PrivateIpAddressesSet *InstanceEvents_AWSAPICallViaCloudTrail_PrivateIpAddressesSet1 `field:"optional" json:"privateIpAddressesSet" yaml:"privateIpAddressesSet"`
	// requesterId property.
	//
	// Specify an array of string values to match this event if the actual value of requesterId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequesterId *[]*string `field:"optional" json:"requesterId" yaml:"requesterId"`
	// requesterManaged property.
	//
	// Specify an array of string values to match this event if the actual value of requesterManaged is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequesterManaged *[]*string `field:"optional" json:"requesterManaged" yaml:"requesterManaged"`
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

