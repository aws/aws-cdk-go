package previewawsguarddutyevents


// Type definition for InstanceDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var ipv6Addresses interface{}
//
//   instanceDetails := &InstanceDetails{
//   	AvailabilityZone: []*string{
//   		jsii.String("availabilityZone"),
//   	},
//   	IamInstanceProfile: &IamInstanceProfile{
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		Id: []*string{
//   			jsii.String("id"),
//   		},
//   	},
//   	ImageDescription: []*string{
//   		jsii.String("imageDescription"),
//   	},
//   	ImageId: []*string{
//   		jsii.String("imageId"),
//   	},
//   	InstanceId: []*string{
//   		jsii.String("instanceId"),
//   	},
//   	InstanceState: []*string{
//   		jsii.String("instanceState"),
//   	},
//   	InstanceType: []*string{
//   		jsii.String("instanceType"),
//   	},
//   	LaunchTime: []*string{
//   		jsii.String("launchTime"),
//   	},
//   	NetworkInterfaces: []InstanceDetailsItem{
//   		&InstanceDetailsItem{
//   			Ipv6Addresses: []interface{}{
//   				ipv6Addresses,
//   			},
//   			NetworkInterfaceId: []*string{
//   				jsii.String("networkInterfaceId"),
//   			},
//   			PrivateDnsName: []*string{
//   				jsii.String("privateDnsName"),
//   			},
//   			PrivateIpAddress: []*string{
//   				jsii.String("privateIpAddress"),
//   			},
//   			PrivateIpAddresses: []InstanceDetailsItemItem{
//   				&InstanceDetailsItemItem{
//   					PrivateDnsName: []*string{
//   						jsii.String("privateDnsName"),
//   					},
//   					PrivateIpAddress: []*string{
//   						jsii.String("privateIpAddress"),
//   					},
//   				},
//   			},
//   			PublicDnsName: []*string{
//   				jsii.String("publicDnsName"),
//   			},
//   			PublicIp: []*string{
//   				jsii.String("publicIp"),
//   			},
//   			SecurityGroups: []InstanceDetailsItemItem1{
//   				&InstanceDetailsItemItem1{
//   					GroupId: []*string{
//   						jsii.String("groupId"),
//   					},
//   					GroupName: []*string{
//   						jsii.String("groupName"),
//   					},
//   				},
//   			},
//   			SubnetId: []*string{
//   				jsii.String("subnetId"),
//   			},
//   			VpcId: []*string{
//   				jsii.String("vpcId"),
//   			},
//   		},
//   	},
//   	OutpostArn: []*string{
//   		jsii.String("outpostArn"),
//   	},
//   	Platform: []*string{
//   		jsii.String("platform"),
//   	},
//   	ProductCodes: []InstanceDetailsItem1{
//   		&InstanceDetailsItem1{
//   			ProductCodeId: []*string{
//   				jsii.String("productCodeId"),
//   			},
//   			ProductCodeType: []*string{
//   				jsii.String("productCodeType"),
//   			},
//   		},
//   	},
//   	Tags: []EcsClusterDetailsItem{
//   		&EcsClusterDetailsItem{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_InstanceDetails struct {
	// availabilityZone property.
	//
	// Specify an array of string values to match this event if the actual value of availabilityZone is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AvailabilityZone *[]*string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// iamInstanceProfile property.
	//
	// Specify an array of string values to match this event if the actual value of iamInstanceProfile is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IamInstanceProfile *DetectorEvents_GuardDutyFinding_IamInstanceProfile `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// imageDescription property.
	//
	// Specify an array of string values to match this event if the actual value of imageDescription is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageDescription *[]*string `field:"optional" json:"imageDescription" yaml:"imageDescription"`
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
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceId *[]*string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// instanceState property.
	//
	// Specify an array of string values to match this event if the actual value of instanceState is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceState *[]*string `field:"optional" json:"instanceState" yaml:"instanceState"`
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
	// networkInterfaces property.
	//
	// Specify an array of string values to match this event if the actual value of networkInterfaces is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkInterfaces *[]*DetectorEvents_GuardDutyFinding_InstanceDetailsItem `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// outpostArn property.
	//
	// Specify an array of string values to match this event if the actual value of outpostArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OutpostArn *[]*string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// platform property.
	//
	// Specify an array of string values to match this event if the actual value of platform is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Platform *[]*string `field:"optional" json:"platform" yaml:"platform"`
	// productCodes property.
	//
	// Specify an array of string values to match this event if the actual value of productCodes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProductCodes *[]*DetectorEvents_GuardDutyFinding_InstanceDetailsItem1 `field:"optional" json:"productCodes" yaml:"productCodes"`
	// tags property.
	//
	// Specify an array of string values to match this event if the actual value of tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*DetectorEvents_GuardDutyFinding_EcsClusterDetailsItem `field:"optional" json:"tags" yaml:"tags"`
}

