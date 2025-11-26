package previewawsguarddutyevents


// Type definition for InstanceDetailsItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var ipv6Addresses interface{}
//
//   instanceDetailsItem := &InstanceDetailsItem{
//   	Ipv6Addresses: []interface{}{
//   		ipv6Addresses,
//   	},
//   	NetworkInterfaceId: []*string{
//   		jsii.String("networkInterfaceId"),
//   	},
//   	PrivateDnsName: []*string{
//   		jsii.String("privateDnsName"),
//   	},
//   	PrivateIpAddress: []*string{
//   		jsii.String("privateIpAddress"),
//   	},
//   	PrivateIpAddresses: []InstanceDetailsItemItem{
//   		&InstanceDetailsItemItem{
//   			PrivateDnsName: []*string{
//   				jsii.String("privateDnsName"),
//   			},
//   			PrivateIpAddress: []*string{
//   				jsii.String("privateIpAddress"),
//   			},
//   		},
//   	},
//   	PublicDnsName: []*string{
//   		jsii.String("publicDnsName"),
//   	},
//   	PublicIp: []*string{
//   		jsii.String("publicIp"),
//   	},
//   	SecurityGroups: []InstanceDetailsItemItem1{
//   		&InstanceDetailsItemItem1{
//   			GroupId: []*string{
//   				jsii.String("groupId"),
//   			},
//   			GroupName: []*string{
//   				jsii.String("groupName"),
//   			},
//   		},
//   	},
//   	SubnetId: []*string{
//   		jsii.String("subnetId"),
//   	},
//   	VpcId: []*string{
//   		jsii.String("vpcId"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_InstanceDetailsItem struct {
	// ipv6Addresses property.
	//
	// Specify an array of string values to match this event if the actual value of ipv6Addresses is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Ipv6Addresses *[]interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// networkInterfaceId property.
	//
	// Specify an array of string values to match this event if the actual value of networkInterfaceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkInterfaceId *[]*string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// privateDnsName property.
	//
	// Specify an array of string values to match this event if the actual value of privateDnsName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrivateDnsName *[]*string `field:"optional" json:"privateDnsName" yaml:"privateDnsName"`
	// privateIpAddress property.
	//
	// Specify an array of string values to match this event if the actual value of privateIpAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrivateIpAddress *[]*string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// privateIpAddresses property.
	//
	// Specify an array of string values to match this event if the actual value of privateIpAddresses is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrivateIpAddresses *[]*DetectorEvents_GuardDutyFinding_InstanceDetailsItemItem `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// publicDnsName property.
	//
	// Specify an array of string values to match this event if the actual value of publicDnsName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PublicDnsName *[]*string `field:"optional" json:"publicDnsName" yaml:"publicDnsName"`
	// publicIp property.
	//
	// Specify an array of string values to match this event if the actual value of publicIp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PublicIp *[]*string `field:"optional" json:"publicIp" yaml:"publicIp"`
	// securityGroups property.
	//
	// Specify an array of string values to match this event if the actual value of securityGroups is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SecurityGroups *[]*DetectorEvents_GuardDutyFinding_InstanceDetailsItemItem1 `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// subnetId property.
	//
	// Specify an array of string values to match this event if the actual value of subnetId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SubnetId *[]*string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// vpcId property.
	//
	// Specify an array of string values to match this event if the actual value of vpcId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VpcId *[]*string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

