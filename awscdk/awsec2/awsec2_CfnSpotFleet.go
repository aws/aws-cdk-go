package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::EC2::SpotFleet`.
//
// Specifies a Spot Fleet request.
//
// The Spot Fleet request specifies the total target capacity and the On-Demand target capacity. Amazon EC2 calculates the difference between the total capacity and On-Demand capacity, and launches the difference as Spot capacity.
//
// You can submit a single request that includes multiple launch specifications that vary by instance type, AMI, Availability Zone, or subnet.
//
// By default, the Spot Fleet requests Spot Instances in the Spot Instance pool where the price per unit is the lowest. Each launch specification can include its own instance weighting that reflects the value of the instance type to your application workload.
//
// Alternatively, you can specify that the Spot Fleet distribute the target capacity across the Spot pools included in its launch specifications. By ensuring that the Spot Instances in your Spot Fleet are in different Spot pools, you can improve the availability of your fleet.
//
// You can specify tags for the Spot Fleet request and instances launched by the fleet. You cannot tag other resource types in a Spot Fleet request because only the `spot-fleet-request` and `instance` resource types are supported.
//
// For more information, see [Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet.html) in the *Amazon EC2 User Guide for Linux Instances* .
//
// > We strongly discourage using the RequestSpotFleet API because it is a legacy API with no planned investment. For options for requesting Spot Instances, see [Which is the best Spot request method to use?](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-best-practices.html#which-spot-request-method-to-use) in the *Amazon EC2 User Guide for Linux Instances* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSpotFleet := awscdk.Aws_ec2.NewCfnSpotFleet(this, jsii.String("MyCfnSpotFleet"), &CfnSpotFleetProps{
//   	SpotFleetRequestConfigData: &SpotFleetRequestConfigDataProperty{
//   		IamFleetRole: jsii.String("iamFleetRole"),
//   		TargetCapacity: jsii.Number(123),
//
//   		// the properties below are optional
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		Context: jsii.String("context"),
//   		ExcessCapacityTerminationPolicy: jsii.String("excessCapacityTerminationPolicy"),
//   		InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   		InstancePoolsToUseCount: jsii.Number(123),
//   		LaunchSpecifications: []interface{}{
//   			&SpotFleetLaunchSpecificationProperty{
//   				ImageId: jsii.String("imageId"),
//
//   				// the properties below are optional
//   				BlockDeviceMappings: []interface{}{
//   					&BlockDeviceMappingProperty{
//   						DeviceName: jsii.String("deviceName"),
//
//   						// the properties below are optional
//   						Ebs: &EbsBlockDeviceProperty{
//   							DeleteOnTermination: jsii.Boolean(false),
//   							Encrypted: jsii.Boolean(false),
//   							Iops: jsii.Number(123),
//   							SnapshotId: jsii.String("snapshotId"),
//   							VolumeSize: jsii.Number(123),
//   							VolumeType: jsii.String("volumeType"),
//   						},
//   						NoDevice: jsii.String("noDevice"),
//   						VirtualName: jsii.String("virtualName"),
//   					},
//   				},
//   				EbsOptimized: jsii.Boolean(false),
//   				IamInstanceProfile: &IamInstanceProfileSpecificationProperty{
//   					Arn: jsii.String("arn"),
//   				},
//   				InstanceRequirements: &InstanceRequirementsRequestProperty{
//   					AcceleratorCount: &AcceleratorCountRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					AcceleratorManufacturers: []*string{
//   						jsii.String("acceleratorManufacturers"),
//   					},
//   					AcceleratorNames: []*string{
//   						jsii.String("acceleratorNames"),
//   					},
//   					AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					AcceleratorTypes: []*string{
//   						jsii.String("acceleratorTypes"),
//   					},
//   					AllowedInstanceTypes: []*string{
//   						jsii.String("allowedInstanceTypes"),
//   					},
//   					BareMetal: jsii.String("bareMetal"),
//   					BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					BurstablePerformance: jsii.String("burstablePerformance"),
//   					CpuManufacturers: []*string{
//   						jsii.String("cpuManufacturers"),
//   					},
//   					ExcludedInstanceTypes: []*string{
//   						jsii.String("excludedInstanceTypes"),
//   					},
//   					InstanceGenerations: []*string{
//   						jsii.String("instanceGenerations"),
//   					},
//   					LocalStorage: jsii.String("localStorage"),
//   					LocalStorageTypes: []*string{
//   						jsii.String("localStorageTypes"),
//   					},
//   					MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					MemoryMiB: &MemoryMiBRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					RequireHibernateSupport: jsii.Boolean(false),
//   					SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					VCpuCount: &VCpuCountRangeRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   				},
//   				InstanceType: jsii.String("instanceType"),
//   				KernelId: jsii.String("kernelId"),
//   				KeyName: jsii.String("keyName"),
//   				Monitoring: &SpotFleetMonitoringProperty{
//   					Enabled: jsii.Boolean(false),
//   				},
//   				NetworkInterfaces: []interface{}{
//   					&InstanceNetworkInterfaceSpecificationProperty{
//   						AssociatePublicIpAddress: jsii.Boolean(false),
//   						DeleteOnTermination: jsii.Boolean(false),
//   						Description: jsii.String("description"),
//   						DeviceIndex: jsii.Number(123),
//   						Groups: []*string{
//   							jsii.String("groups"),
//   						},
//   						Ipv6AddressCount: jsii.Number(123),
//   						Ipv6Addresses: []interface{}{
//   							&InstanceIpv6AddressProperty{
//   								Ipv6Address: jsii.String("ipv6Address"),
//   							},
//   						},
//   						NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   						PrivateIpAddresses: []interface{}{
//   							&PrivateIpAddressSpecificationProperty{
//   								PrivateIpAddress: jsii.String("privateIpAddress"),
//
//   								// the properties below are optional
//   								Primary: jsii.Boolean(false),
//   							},
//   						},
//   						SecondaryPrivateIpAddressCount: jsii.Number(123),
//   						SubnetId: jsii.String("subnetId"),
//   					},
//   				},
//   				Placement: &SpotPlacementProperty{
//   					AvailabilityZone: jsii.String("availabilityZone"),
//   					GroupName: jsii.String("groupName"),
//   					Tenancy: jsii.String("tenancy"),
//   				},
//   				RamdiskId: jsii.String("ramdiskId"),
//   				SecurityGroups: []interface{}{
//   					&GroupIdentifierProperty{
//   						GroupId: jsii.String("groupId"),
//   					},
//   				},
//   				SpotPrice: jsii.String("spotPrice"),
//   				SubnetId: jsii.String("subnetId"),
//   				TagSpecifications: []interface{}{
//   					&SpotFleetTagSpecificationProperty{
//   						ResourceType: jsii.String("resourceType"),
//   						Tags: []cfnTag{
//   							&cfnTag{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				UserData: jsii.String("userData"),
//   				WeightedCapacity: jsii.Number(123),
//   			},
//   		},
//   		LaunchTemplateConfigs: []interface{}{
//   			&LaunchTemplateConfigProperty{
//   				LaunchTemplateSpecification: &FleetLaunchTemplateSpecificationProperty{
//   					Version: jsii.String("version"),
//
//   					// the properties below are optional
//   					LaunchTemplateId: jsii.String("launchTemplateId"),
//   					LaunchTemplateName: jsii.String("launchTemplateName"),
//   				},
//   				Overrides: []interface{}{
//   					&LaunchTemplateOverridesProperty{
//   						AvailabilityZone: jsii.String("availabilityZone"),
//   						InstanceRequirements: &InstanceRequirementsRequestProperty{
//   							AcceleratorCount: &AcceleratorCountRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							AcceleratorManufacturers: []*string{
//   								jsii.String("acceleratorManufacturers"),
//   							},
//   							AcceleratorNames: []*string{
//   								jsii.String("acceleratorNames"),
//   							},
//   							AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							AcceleratorTypes: []*string{
//   								jsii.String("acceleratorTypes"),
//   							},
//   							AllowedInstanceTypes: []*string{
//   								jsii.String("allowedInstanceTypes"),
//   							},
//   							BareMetal: jsii.String("bareMetal"),
//   							BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							BurstablePerformance: jsii.String("burstablePerformance"),
//   							CpuManufacturers: []*string{
//   								jsii.String("cpuManufacturers"),
//   							},
//   							ExcludedInstanceTypes: []*string{
//   								jsii.String("excludedInstanceTypes"),
//   							},
//   							InstanceGenerations: []*string{
//   								jsii.String("instanceGenerations"),
//   							},
//   							LocalStorage: jsii.String("localStorage"),
//   							LocalStorageTypes: []*string{
//   								jsii.String("localStorageTypes"),
//   							},
//   							MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							MemoryMiB: &MemoryMiBRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   							RequireHibernateSupport: jsii.Boolean(false),
//   							SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   							TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							VCpuCount: &VCpuCountRangeRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   						},
//   						InstanceType: jsii.String("instanceType"),
//   						Priority: jsii.Number(123),
//   						SpotPrice: jsii.String("spotPrice"),
//   						SubnetId: jsii.String("subnetId"),
//   						WeightedCapacity: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		LoadBalancersConfig: &LoadBalancersConfigProperty{
//   			ClassicLoadBalancersConfig: &ClassicLoadBalancersConfigProperty{
//   				ClassicLoadBalancers: []interface{}{
//   					&ClassicLoadBalancerProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			TargetGroupsConfig: &TargetGroupsConfigProperty{
//   				TargetGroups: []interface{}{
//   					&TargetGroupProperty{
//   						Arn: jsii.String("arn"),
//   					},
//   				},
//   			},
//   		},
//   		OnDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   		OnDemandMaxTotalPrice: jsii.String("onDemandMaxTotalPrice"),
//   		OnDemandTargetCapacity: jsii.Number(123),
//   		ReplaceUnhealthyInstances: jsii.Boolean(false),
//   		SpotMaintenanceStrategies: &SpotMaintenanceStrategiesProperty{
//   			CapacityRebalance: &SpotCapacityRebalanceProperty{
//   				ReplacementStrategy: jsii.String("replacementStrategy"),
//   				TerminationDelay: jsii.Number(123),
//   			},
//   		},
//   		SpotMaxTotalPrice: jsii.String("spotMaxTotalPrice"),
//   		SpotPrice: jsii.String("spotPrice"),
//   		TagSpecifications: []interface{}{
//   			&SpotFleetTagSpecificationProperty{
//   				ResourceType: jsii.String("resourceType"),
//   				Tags: []*cfnTag{
//   					&cfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		TargetCapacityUnitType: jsii.String("targetCapacityUnitType"),
//   		TerminateInstancesWithExpiration: jsii.Boolean(false),
//   		Type: jsii.String("type"),
//   		ValidFrom: jsii.String("validFrom"),
//   		ValidUntil: jsii.String("validUntil"),
//   	},
//   })
//
type CfnSpotFleet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID of the Spot Fleet.
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// Describes the configuration of a Spot Fleet request.
	SpotFleetRequestConfigData() interface{}
	SetSpotFleetRequestConfigData(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnSpotFleet
type jsiiProxy_CfnSpotFleet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSpotFleet) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpotFleet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpotFleet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpotFleet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpotFleet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpotFleet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpotFleet) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpotFleet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpotFleet) SpotFleetRequestConfigData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotFleetRequestConfigData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpotFleet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpotFleet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EC2::SpotFleet`.
func NewCfnSpotFleet(scope awscdk.Construct, id *string, props *CfnSpotFleetProps) CfnSpotFleet {
	_init_.Initialize()

	if err := validateNewCfnSpotFleetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSpotFleet{}

	_jsii_.Create(
		"monocdk.aws_ec2.CfnSpotFleet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EC2::SpotFleet`.
func NewCfnSpotFleet_Override(c CfnSpotFleet, scope awscdk.Construct, id *string, props *CfnSpotFleetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.CfnSpotFleet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSpotFleet)SetSpotFleetRequestConfigData(val interface{}) {
	if err := j.validateSetSpotFleetRequestConfigDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotFleetRequestConfigData",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnSpotFleet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSpotFleet_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.CfnSpotFleet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnSpotFleet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnSpotFleet_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.CfnSpotFleet",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnSpotFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSpotFleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.CfnSpotFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSpotFleet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ec2.CfnSpotFleet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSpotFleet) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSpotFleet) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSpotFleet) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSpotFleet) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSpotFleet) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSpotFleet) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSpotFleet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSpotFleet) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSpotFleet) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSpotFleet) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSpotFleet) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSpotFleet) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSpotFleet) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSpotFleet) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSpotFleet) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSpotFleet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSpotFleet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSpotFleet) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSpotFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSpotFleet) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSpotFleet) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

