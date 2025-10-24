package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceProps := &CfnInstanceProps{
//   	BlueprintId: jsii.String("blueprintId"),
//   	BundleId: jsii.String("bundleId"),
//   	InstanceName: jsii.String("instanceName"),
//
//   	// the properties below are optional
//   	AddOns: []interface{}{
//   		&AddOnProperty{
//   			AddOnType: jsii.String("addOnType"),
//
//   			// the properties below are optional
//   			AutoSnapshotAddOnRequest: &AutoSnapshotAddOnProperty{
//   				SnapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   			},
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	Hardware: &HardwareProperty{
//   		CpuCount: jsii.Number(123),
//   		Disks: []interface{}{
//   			&DiskProperty{
//   				DiskName: jsii.String("diskName"),
//   				Path: jsii.String("path"),
//
//   				// the properties below are optional
//   				AttachedTo: jsii.String("attachedTo"),
//   				AttachmentState: jsii.String("attachmentState"),
//   				Iops: jsii.Number(123),
//   				IsSystemDisk: jsii.Boolean(false),
//   				SizeInGb: jsii.String("sizeInGb"),
//   			},
//   		},
//   		RamSizeInGb: jsii.Number(123),
//   	},
//   	KeyPairName: jsii.String("keyPairName"),
//   	Location: &LocationProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		RegionName: jsii.String("regionName"),
//   	},
//   	Networking: &NetworkingProperty{
//   		Ports: []interface{}{
//   			&PortProperty{
//   				AccessDirection: jsii.String("accessDirection"),
//   				AccessFrom: jsii.String("accessFrom"),
//   				AccessType: jsii.String("accessType"),
//   				CidrListAliases: []*string{
//   					jsii.String("cidrListAliases"),
//   				},
//   				Cidrs: []*string{
//   					jsii.String("cidrs"),
//   				},
//   				CommonName: jsii.String("commonName"),
//   				FromPort: jsii.Number(123),
//   				Ipv6Cidrs: []*string{
//   					jsii.String("ipv6Cidrs"),
//   				},
//   				Protocol: jsii.String("protocol"),
//   				ToPort: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		MonthlyTransfer: &MonthlyTransferProperty{
//   			GbPerMonthAllocated: jsii.String("gbPerMonthAllocated"),
//   		},
//   	},
//   	State: &StateProperty{
//   		Code: jsii.Number(123),
//   		Name: jsii.String("name"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserData: jsii.String("userData"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html
//
type CfnInstanceProps struct {
	// The blueprint ID for the instance (for example, `os_amlinux_2016_03` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html#cfn-lightsail-instance-blueprintid
	//
	BlueprintId *string `field:"required" json:"blueprintId" yaml:"blueprintId"`
	// The bundle ID for the instance (for example, `micro_1_0` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html#cfn-lightsail-instance-bundleid
	//
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// The name of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html#cfn-lightsail-instance-instancename
	//
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// An array of add-ons for the instance.
	//
	// > If the instance has an add-on enabled when performing a delete instance request, the add-on is automatically disabled before the instance is deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html#cfn-lightsail-instance-addons
	//
	AddOns interface{} `field:"optional" json:"addOns" yaml:"addOns"`
	// The Availability Zone for the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html#cfn-lightsail-instance-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The hardware properties for the instance, such as the vCPU count, attached disks, and amount of RAM.
	//
	// > The instance restarts when performing an attach disk or detach disk request. This resets the public IP address of your instance if a static IP isn't attached to it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html#cfn-lightsail-instance-hardware
	//
	Hardware interface{} `field:"optional" json:"hardware" yaml:"hardware"`
	// The name of the key pair to use for the instance.
	//
	// If no key pair name is specified, the Regional Lightsail default key pair is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html#cfn-lightsail-instance-keypairname
	//
	KeyPairName *string `field:"optional" json:"keyPairName" yaml:"keyPairName"`
	// The location for the instance, such as the AWS Region and Availability Zone.
	//
	// > The `Location` property is read-only and should not be specified in a create instance or update instance request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html#cfn-lightsail-instance-location
	//
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// The public ports and the monthly amount of data transfer allocated for the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html#cfn-lightsail-instance-networking
	//
	Networking interface{} `field:"optional" json:"networking" yaml:"networking"`
	// The status code and the state (for example, `running` ) of the instance.
	//
	// > The `State` property is read-only and should not be specified in a create instance or update instance request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html#cfn-lightsail-instance-state
	//
	State interface{} `field:"optional" json:"state" yaml:"state"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html#cfn-lightsail-instance-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The optional launch script for the instance.
	//
	// Specify a launch script to configure an instance with additional user data. For example, you might want to specify `apt-get -y update` as a launch script.
	//
	// > Depending on the blueprint of your instance, the command to get software on your instance varies. Amazon Linux and CentOS use `yum` , Debian and Ubuntu use `apt-get` , and FreeBSD uses `pkg` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html#cfn-lightsail-instance-userdata
	//
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

