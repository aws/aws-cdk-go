package awsworkspacesinstances


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedInstanceProperty := &ManagedInstanceProperty{
//   	ImageId: jsii.String("imageId"),
//   	InstanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	BlockDeviceMappings: []interface{}{
//   		&BlockDeviceMappingProperty{
//   			DeviceName: jsii.String("deviceName"),
//   			Ebs: &EbsBlockDeviceProperty{
//   				Encrypted: jsii.Boolean(false),
//   				Iops: jsii.Number(123),
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   				Throughput: jsii.Number(123),
//   				VolumeSize: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//   			},
//   			NoDevice: jsii.String("noDevice"),
//   			VirtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	CpuOptions: &CpuOptionsRequestProperty{
//   		CoreCount: jsii.Number(123),
//   		ThreadsPerCore: jsii.Number(123),
//   	},
//   	CreditSpecification: &CreditSpecificationRequestProperty{
//   		CpuCredits: jsii.String("cpuCredits"),
//   	},
//   	DisableApiStop: jsii.Boolean(false),
//   	EbsOptimized: jsii.Boolean(false),
//   	EnclaveOptions: &EnclaveOptionsRequestProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	HibernationOptions: &HibernationOptionsRequestProperty{
//   		Configured: jsii.Boolean(false),
//   	},
//   	IamInstanceProfile: &IamInstanceProfileSpecificationProperty{
//   		Name: jsii.String("name"),
//   	},
//   	KeyName: jsii.String("keyName"),
//   	MaintenanceOptions: &InstanceMaintenanceOptionsRequestProperty{
//   		AutoRecovery: jsii.String("autoRecovery"),
//   	},
//   	MetadataOptions: &InstanceMetadataOptionsRequestProperty{
//   		HttpEndpoint: jsii.String("httpEndpoint"),
//   		HttpProtocolIpv6: jsii.String("httpProtocolIpv6"),
//   		HttpPutResponseHopLimit: jsii.Number(123),
//   		HttpTokens: jsii.String("httpTokens"),
//   		InstanceMetadataTags: jsii.String("instanceMetadataTags"),
//   	},
//   	Monitoring: &RunInstancesMonitoringEnabledProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	NetworkInterfaces: []interface{}{
//   		&InstanceNetworkInterfaceSpecificationProperty{
//   			Description: jsii.String("description"),
//   			DeviceIndex: jsii.Number(123),
//   			Groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	NetworkPerformanceOptions: &InstanceNetworkPerformanceOptionsRequestProperty{
//   		BandwidthWeighting: jsii.String("bandwidthWeighting"),
//   	},
//   	Placement: &PlacementProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		GroupName: jsii.String("groupName"),
//   		Tenancy: jsii.String("tenancy"),
//   	},
//   	PrivateDnsNameOptions: &PrivateDnsNameOptionsRequestProperty{
//   		EnableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   		EnableResourceNameDnsARecord: jsii.Boolean(false),
//   		HostnameType: jsii.String("hostnameType"),
//   	},
//   	TagSpecifications: []interface{}{
//   		&TagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []cfnTag{
//   				&cfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	UserData: jsii.String("userData"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html
//
type CfnWorkspaceInstance_ManagedInstanceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-imageid
	//
	ImageId *string `field:"required" json:"imageId" yaml:"imageId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-blockdevicemappings
	//
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-cpuoptions
	//
	CpuOptions interface{} `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-creditspecification
	//
	CreditSpecification interface{} `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-disableapistop
	//
	DisableApiStop interface{} `field:"optional" json:"disableApiStop" yaml:"disableApiStop"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-ebsoptimized
	//
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-enclaveoptions
	//
	EnclaveOptions interface{} `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-hibernationoptions
	//
	HibernationOptions interface{} `field:"optional" json:"hibernationOptions" yaml:"hibernationOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-iaminstanceprofile
	//
	IamInstanceProfile interface{} `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-keyname
	//
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-maintenanceoptions
	//
	MaintenanceOptions interface{} `field:"optional" json:"maintenanceOptions" yaml:"maintenanceOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-metadataoptions
	//
	MetadataOptions interface{} `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-monitoring
	//
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-networkinterfaces
	//
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-networkperformanceoptions
	//
	NetworkPerformanceOptions interface{} `field:"optional" json:"networkPerformanceOptions" yaml:"networkPerformanceOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-placement
	//
	Placement interface{} `field:"optional" json:"placement" yaml:"placement"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-privatednsnameoptions
	//
	PrivateDnsNameOptions interface{} `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-tagspecifications
	//
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-managedinstance.html#cfn-workspacesinstances-workspaceinstance-managedinstance-userdata
	//
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

