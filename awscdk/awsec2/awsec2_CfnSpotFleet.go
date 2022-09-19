package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::EC2::SpotFleet`.
//
// Specifies a Spot Fleet request. A Spot Fleet request contains the configuration information to launch a fleet, or group, of instances.
//
// The Spot Fleet request specifies the total target capacity and the On-Demand target capacity for the fleet. Amazon EC2 calculates the difference between the total capacity and On-Demand capacity, and launches the difference as Spot capacity.
//
// The Spot Fleet request can include multiple launch specifications that vary by instance type, AMI, Availability Zone, or subnet.
//
// By default, the Spot Fleet requests Spot Instances in the Spot pool where the price per unit is the lowest. Each launch specification can include its own instance weighting that reflects the value of the instance type to your application workload.
//
// Alternatively, you can specify that the Spot Fleet distribute the target capacity across the Spot pools included in its launch specifications. By ensuring that the Spot Instances in your Spot Fleet are in different Spot pools, you can improve the availability of your fleet.
//
// You can specify tags for the Spot Instances. You cannot tag other resource types in a Spot Fleet request because only the `instance` resource type is supported.
//
// For more information, see [Spot Fleet Requests](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-requests.html) in the *Amazon EC2 User Guide for Linux Instances* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSpotFleet := awscdk.Aws_ec2.NewCfnSpotFleet(this, jsii.String("MyCfnSpotFleet"), &cfnSpotFleetProps{
//   	spotFleetRequestConfigData: &spotFleetRequestConfigDataProperty{
//   		iamFleetRole: jsii.String("iamFleetRole"),
//   		targetCapacity: jsii.Number(123),
//
//   		// the properties below are optional
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   		context: jsii.String("context"),
//   		excessCapacityTerminationPolicy: jsii.String("excessCapacityTerminationPolicy"),
//   		instanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   		instancePoolsToUseCount: jsii.Number(123),
//   		launchSpecifications: []interface{}{
//   			&spotFleetLaunchSpecificationProperty{
//   				imageId: jsii.String("imageId"),
//
//   				// the properties below are optional
//   				blockDeviceMappings: []interface{}{
//   					&blockDeviceMappingProperty{
//   						deviceName: jsii.String("deviceName"),
//
//   						// the properties below are optional
//   						ebs: &ebsBlockDeviceProperty{
//   							deleteOnTermination: jsii.Boolean(false),
//   							encrypted: jsii.Boolean(false),
//   							iops: jsii.Number(123),
//   							snapshotId: jsii.String("snapshotId"),
//   							volumeSize: jsii.Number(123),
//   							volumeType: jsii.String("volumeType"),
//   						},
//   						noDevice: jsii.String("noDevice"),
//   						virtualName: jsii.String("virtualName"),
//   					},
//   				},
//   				ebsOptimized: jsii.Boolean(false),
//   				iamInstanceProfile: &iamInstanceProfileSpecificationProperty{
//   					arn: jsii.String("arn"),
//   				},
//   				instanceRequirements: &instanceRequirementsRequestProperty{
//   					acceleratorCount: &acceleratorCountRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					acceleratorManufacturers: []*string{
//   						jsii.String("acceleratorManufacturers"),
//   					},
//   					acceleratorNames: []*string{
//   						jsii.String("acceleratorNames"),
//   					},
//   					acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					acceleratorTypes: []*string{
//   						jsii.String("acceleratorTypes"),
//   					},
//   					bareMetal: jsii.String("bareMetal"),
//   					baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					burstablePerformance: jsii.String("burstablePerformance"),
//   					cpuManufacturers: []*string{
//   						jsii.String("cpuManufacturers"),
//   					},
//   					excludedInstanceTypes: []*string{
//   						jsii.String("excludedInstanceTypes"),
//   					},
//   					instanceGenerations: []*string{
//   						jsii.String("instanceGenerations"),
//   					},
//   					localStorage: jsii.String("localStorage"),
//   					localStorageTypes: []*string{
//   						jsii.String("localStorageTypes"),
//   					},
//   					memoryGiBPerVCpu: &memoryGiBPerVCpuRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					memoryMiB: &memoryMiBRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					networkInterfaceCount: &networkInterfaceCountRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					requireHibernateSupport: jsii.Boolean(false),
//   					spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					totalLocalStorageGb: &totalLocalStorageGBRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					vCpuCount: &vCpuCountRangeRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   				},
//   				instanceType: jsii.String("instanceType"),
//   				kernelId: jsii.String("kernelId"),
//   				keyName: jsii.String("keyName"),
//   				monitoring: &spotFleetMonitoringProperty{
//   					enabled: jsii.Boolean(false),
//   				},
//   				networkInterfaces: []interface{}{
//   					&instanceNetworkInterfaceSpecificationProperty{
//   						associatePublicIpAddress: jsii.Boolean(false),
//   						deleteOnTermination: jsii.Boolean(false),
//   						description: jsii.String("description"),
//   						deviceIndex: jsii.Number(123),
//   						groups: []*string{
//   							jsii.String("groups"),
//   						},
//   						ipv6AddressCount: jsii.Number(123),
//   						ipv6Addresses: []interface{}{
//   							&instanceIpv6AddressProperty{
//   								ipv6Address: jsii.String("ipv6Address"),
//   							},
//   						},
//   						networkInterfaceId: jsii.String("networkInterfaceId"),
//   						privateIpAddresses: []interface{}{
//   							&privateIpAddressSpecificationProperty{
//   								privateIpAddress: jsii.String("privateIpAddress"),
//
//   								// the properties below are optional
//   								primary: jsii.Boolean(false),
//   							},
//   						},
//   						secondaryPrivateIpAddressCount: jsii.Number(123),
//   						subnetId: jsii.String("subnetId"),
//   					},
//   				},
//   				placement: &spotPlacementProperty{
//   					availabilityZone: jsii.String("availabilityZone"),
//   					groupName: jsii.String("groupName"),
//   					tenancy: jsii.String("tenancy"),
//   				},
//   				ramdiskId: jsii.String("ramdiskId"),
//   				securityGroups: []interface{}{
//   					&groupIdentifierProperty{
//   						groupId: jsii.String("groupId"),
//   					},
//   				},
//   				spotPrice: jsii.String("spotPrice"),
//   				subnetId: jsii.String("subnetId"),
//   				tagSpecifications: []interface{}{
//   					&spotFleetTagSpecificationProperty{
//   						resourceType: jsii.String("resourceType"),
//   						tags: []cfnTag{
//   							&cfnTag{
//   								key: jsii.String("key"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				userData: jsii.String("userData"),
//   				weightedCapacity: jsii.Number(123),
//   			},
//   		},
//   		launchTemplateConfigs: []interface{}{
//   			&launchTemplateConfigProperty{
//   				launchTemplateSpecification: &fleetLaunchTemplateSpecificationProperty{
//   					version: jsii.String("version"),
//
//   					// the properties below are optional
//   					launchTemplateId: jsii.String("launchTemplateId"),
//   					launchTemplateName: jsii.String("launchTemplateName"),
//   				},
//   				overrides: []interface{}{
//   					&launchTemplateOverridesProperty{
//   						availabilityZone: jsii.String("availabilityZone"),
//   						instanceRequirements: &instanceRequirementsRequestProperty{
//   							acceleratorCount: &acceleratorCountRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							acceleratorManufacturers: []*string{
//   								jsii.String("acceleratorManufacturers"),
//   							},
//   							acceleratorNames: []*string{
//   								jsii.String("acceleratorNames"),
//   							},
//   							acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							acceleratorTypes: []*string{
//   								jsii.String("acceleratorTypes"),
//   							},
//   							bareMetal: jsii.String("bareMetal"),
//   							baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							burstablePerformance: jsii.String("burstablePerformance"),
//   							cpuManufacturers: []*string{
//   								jsii.String("cpuManufacturers"),
//   							},
//   							excludedInstanceTypes: []*string{
//   								jsii.String("excludedInstanceTypes"),
//   							},
//   							instanceGenerations: []*string{
//   								jsii.String("instanceGenerations"),
//   							},
//   							localStorage: jsii.String("localStorage"),
//   							localStorageTypes: []*string{
//   								jsii.String("localStorageTypes"),
//   							},
//   							memoryGiBPerVCpu: &memoryGiBPerVCpuRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							memoryMiB: &memoryMiBRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							networkInterfaceCount: &networkInterfaceCountRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   							requireHibernateSupport: jsii.Boolean(false),
//   							spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   							totalLocalStorageGb: &totalLocalStorageGBRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							vCpuCount: &vCpuCountRangeRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   						},
//   						instanceType: jsii.String("instanceType"),
//   						priority: jsii.Number(123),
//   						spotPrice: jsii.String("spotPrice"),
//   						subnetId: jsii.String("subnetId"),
//   						weightedCapacity: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		loadBalancersConfig: &loadBalancersConfigProperty{
//   			classicLoadBalancersConfig: &classicLoadBalancersConfigProperty{
//   				classicLoadBalancers: []interface{}{
//   					&classicLoadBalancerProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			targetGroupsConfig: &targetGroupsConfigProperty{
//   				targetGroups: []interface{}{
//   					&targetGroupProperty{
//   						arn: jsii.String("arn"),
//   					},
//   				},
//   			},
//   		},
//   		onDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   		onDemandMaxTotalPrice: jsii.String("onDemandMaxTotalPrice"),
//   		onDemandTargetCapacity: jsii.Number(123),
//   		replaceUnhealthyInstances: jsii.Boolean(false),
//   		spotMaintenanceStrategies: &spotMaintenanceStrategiesProperty{
//   			capacityRebalance: &spotCapacityRebalanceProperty{
//   				replacementStrategy: jsii.String("replacementStrategy"),
//   				terminationDelay: jsii.Number(123),
//   			},
//   		},
//   		spotMaxTotalPrice: jsii.String("spotMaxTotalPrice"),
//   		spotPrice: jsii.String("spotPrice"),
//   		targetCapacityUnitType: jsii.String("targetCapacityUnitType"),
//   		terminateInstancesWithExpiration: jsii.Boolean(false),
//   		type: jsii.String("type"),
//   		validFrom: jsii.String("validFrom"),
//   		validUntil: jsii.String("validUntil"),
//   	},
//   })
//
type CfnSpotFleet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID of the Spot Fleet.
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Describes the configuration of a Spot Fleet request.
	SpotFleetRequestConfigData() interface{}
	SetSpotFleetRequestConfigData(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
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

func (j *jsiiProxy_CfnSpotFleet) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnSpotFleet) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::EC2::SpotFleet`.
func NewCfnSpotFleet(scope constructs.Construct, id *string, props *CfnSpotFleetProps) CfnSpotFleet {
	_init_.Initialize()

	if err := validateNewCfnSpotFleetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSpotFleet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EC2::SpotFleet`.
func NewCfnSpotFleet_Override(c CfnSpotFleet, scope constructs.Construct, id *string, props *CfnSpotFleetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet",
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
func CfnSpotFleet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSpotFleet_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSpotFleet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnSpotFleet_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnSpotFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSpotFleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet",
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
		"aws-cdk-lib.aws_ec2.CfnSpotFleet",
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

