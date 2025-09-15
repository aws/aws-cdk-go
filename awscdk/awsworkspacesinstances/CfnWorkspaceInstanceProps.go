package awsworkspacesinstances

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWorkspaceInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspaceInstanceProps := &CfnWorkspaceInstanceProps{
//   	ManagedInstance: &ManagedInstanceProperty{
//   		ImageId: jsii.String("imageId"),
//   		InstanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		BlockDeviceMappings: []interface{}{
//   			&BlockDeviceMappingProperty{
//   				DeviceName: jsii.String("deviceName"),
//   				Ebs: &EbsBlockDeviceProperty{
//   					Encrypted: jsii.Boolean(false),
//   					Iops: jsii.Number(123),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					Throughput: jsii.Number(123),
//   					VolumeSize: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//   				},
//   				NoDevice: jsii.String("noDevice"),
//   				VirtualName: jsii.String("virtualName"),
//   			},
//   		},
//   		CpuOptions: &CpuOptionsRequestProperty{
//   			CoreCount: jsii.Number(123),
//   			ThreadsPerCore: jsii.Number(123),
//   		},
//   		CreditSpecification: &CreditSpecificationRequestProperty{
//   			CpuCredits: jsii.String("cpuCredits"),
//   		},
//   		DisableApiStop: jsii.Boolean(false),
//   		EbsOptimized: jsii.Boolean(false),
//   		EnclaveOptions: &EnclaveOptionsRequestProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		HibernationOptions: &HibernationOptionsRequestProperty{
//   			Configured: jsii.Boolean(false),
//   		},
//   		IamInstanceProfile: &IamInstanceProfileSpecificationProperty{
//   			Name: jsii.String("name"),
//   		},
//   		KeyName: jsii.String("keyName"),
//   		MaintenanceOptions: &InstanceMaintenanceOptionsRequestProperty{
//   			AutoRecovery: jsii.String("autoRecovery"),
//   		},
//   		MetadataOptions: &InstanceMetadataOptionsRequestProperty{
//   			HttpEndpoint: jsii.String("httpEndpoint"),
//   			HttpProtocolIpv6: jsii.String("httpProtocolIpv6"),
//   			HttpPutResponseHopLimit: jsii.Number(123),
//   			HttpTokens: jsii.String("httpTokens"),
//   			InstanceMetadataTags: jsii.String("instanceMetadataTags"),
//   		},
//   		Monitoring: &RunInstancesMonitoringEnabledProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		NetworkInterfaces: []interface{}{
//   			&InstanceNetworkInterfaceSpecificationProperty{
//   				Description: jsii.String("description"),
//   				DeviceIndex: jsii.Number(123),
//   				Groups: []*string{
//   					jsii.String("groups"),
//   				},
//   				SubnetId: jsii.String("subnetId"),
//   			},
//   		},
//   		NetworkPerformanceOptions: &InstanceNetworkPerformanceOptionsRequestProperty{
//   			BandwidthWeighting: jsii.String("bandwidthWeighting"),
//   		},
//   		Placement: &PlacementProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			GroupName: jsii.String("groupName"),
//   			Tenancy: jsii.String("tenancy"),
//   		},
//   		PrivateDnsNameOptions: &PrivateDnsNameOptionsRequestProperty{
//   			EnableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   			EnableResourceNameDnsARecord: jsii.Boolean(false),
//   			HostnameType: jsii.String("hostnameType"),
//   		},
//   		TagSpecifications: []interface{}{
//   			&TagSpecificationProperty{
//   				ResourceType: jsii.String("resourceType"),
//   				Tags: []cfnTag{
//   					&cfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		UserData: jsii.String("userData"),
//   	},
//   	Tags: []*cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-workspaceinstance.html
//
type CfnWorkspaceInstanceProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-workspaceinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance
	//
	ManagedInstance interface{} `field:"optional" json:"managedInstance" yaml:"managedInstance"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-workspaceinstance.html#cfn-workspacesinstances-workspaceinstance-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

