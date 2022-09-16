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
//   cfnInstanceProps := &cfnInstanceProps{
//   	blueprintId: jsii.String("blueprintId"),
//   	bundleId: jsii.String("bundleId"),
//   	instanceName: jsii.String("instanceName"),
//
//   	// the properties below are optional
//   	addOns: []interface{}{
//   		&addOnProperty{
//   			addOnType: jsii.String("addOnType"),
//
//   			// the properties below are optional
//   			autoSnapshotAddOnRequest: &autoSnapshotAddOnProperty{
//   				snapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   			},
//   			status: jsii.String("status"),
//   		},
//   	},
//   	availabilityZone: jsii.String("availabilityZone"),
//   	hardware: &hardwareProperty{
//   		cpuCount: jsii.Number(123),
//   		disks: []interface{}{
//   			&diskProperty{
//   				diskName: jsii.String("diskName"),
//   				path: jsii.String("path"),
//
//   				// the properties below are optional
//   				attachedTo: jsii.String("attachedTo"),
//   				attachmentState: jsii.String("attachmentState"),
//   				iops: jsii.Number(123),
//   				isSystemDisk: jsii.Boolean(false),
//   				sizeInGb: jsii.String("sizeInGb"),
//   			},
//   		},
//   		ramSizeInGb: jsii.Number(123),
//   	},
//   	keyPairName: jsii.String("keyPairName"),
//   	networking: &networkingProperty{
//   		ports: []interface{}{
//   			&portProperty{
//   				accessDirection: jsii.String("accessDirection"),
//   				accessFrom: jsii.String("accessFrom"),
//   				accessType: jsii.String("accessType"),
//   				cidrListAliases: []*string{
//   					jsii.String("cidrListAliases"),
//   				},
//   				cidrs: []*string{
//   					jsii.String("cidrs"),
//   				},
//   				commonName: jsii.String("commonName"),
//   				fromPort: jsii.Number(123),
//   				ipv6Cidrs: []*string{
//   					jsii.String("ipv6Cidrs"),
//   				},
//   				protocol: jsii.String("protocol"),
//   				toPort: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		monthlyTransfer: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userData: jsii.String("userData"),
//   }
//
type CfnInstanceProps struct {
	// The blueprint ID for the instance (for example, `os_amlinux_2016_03` ).
	BlueprintId *string `field:"required" json:"blueprintId" yaml:"blueprintId"`
	// The bundle ID for the instance (for example, `micro_1_0` ).
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// The name of the instance.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// An array of add-ons for the instance.
	//
	// > If the instance has an add-on enabled when performing a delete instance request, the add-on is automatically disabled before the instance is deleted.
	AddOns interface{} `field:"optional" json:"addOns" yaml:"addOns"`
	// The Availability Zone for the instance.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The hardware properties for the instance, such as the vCPU count, attached disks, and amount of RAM.
	//
	// > The instance restarts when performing an attach disk or detach disk request. This resets the public IP address of your instance if a static IP isn't attached to it.
	Hardware interface{} `field:"optional" json:"hardware" yaml:"hardware"`
	// The name of the key pair to use for the instance.
	//
	// If no key pair name is specified, the Regional Lightsail default key pair is used.
	KeyPairName *string `field:"optional" json:"keyPairName" yaml:"keyPairName"`
	// The public ports and the monthly amount of data transfer allocated for the instance.
	Networking interface{} `field:"optional" json:"networking" yaml:"networking"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The optional launch script for the instance.
	//
	// Specify a launch script to configure an instance with additional user data. For example, you might want to specify `apt-get -y update` as a launch script.
	//
	// > Depending on the blueprint of your instance, the command to get software on your instance varies. Amazon Linux and CentOS use `yum` , Debian and Ubuntu use `apt-get` , and FreeBSD uses `pkg` .
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

