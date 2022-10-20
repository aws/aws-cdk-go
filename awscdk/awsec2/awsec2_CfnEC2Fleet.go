package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::EC2::EC2Fleet`.
//
// Specifies the configuration information to launch a fleet--or group--of instances. An EC2 Fleet can launch multiple instance types across multiple Availability Zones, using the On-Demand Instance, Reserved Instance, and Spot Instance purchasing models together. Using EC2 Fleet, you can define separate On-Demand and Spot capacity targets, specify the instance types that work best for your applications, and specify how Amazon EC2 should distribute your fleet capacity within each purchasing model. For more information, see [Launching an EC2 Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet.html) in the *Amazon EC2 User Guide for Linux Instances* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEC2Fleet := awscdk.Aws_ec2.NewCfnEC2Fleet(this, jsii.String("MyCfnEC2Fleet"), &cfnEC2FleetProps{
//   	launchTemplateConfigs: []interface{}{
//   		&fleetLaunchTemplateConfigRequestProperty{
//   			launchTemplateSpecification: &fleetLaunchTemplateSpecificationRequestProperty{
//   				launchTemplateId: jsii.String("launchTemplateId"),
//   				launchTemplateName: jsii.String("launchTemplateName"),
//   				version: jsii.String("version"),
//   			},
//   			overrides: []interface{}{
//   				&fleetLaunchTemplateOverridesRequestProperty{
//   					availabilityZone: jsii.String("availabilityZone"),
//   					instanceRequirements: &instanceRequirementsRequestProperty{
//   						acceleratorCount: &acceleratorCountRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						acceleratorManufacturers: []*string{
//   							jsii.String("acceleratorManufacturers"),
//   						},
//   						acceleratorNames: []*string{
//   							jsii.String("acceleratorNames"),
//   						},
//   						acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						acceleratorTypes: []*string{
//   							jsii.String("acceleratorTypes"),
//   						},
//   						bareMetal: jsii.String("bareMetal"),
//   						baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						burstablePerformance: jsii.String("burstablePerformance"),
//   						cpuManufacturers: []*string{
//   							jsii.String("cpuManufacturers"),
//   						},
//   						excludedInstanceTypes: []*string{
//   							jsii.String("excludedInstanceTypes"),
//   						},
//   						instanceGenerations: []*string{
//   							jsii.String("instanceGenerations"),
//   						},
//   						localStorage: jsii.String("localStorage"),
//   						localStorageTypes: []*string{
//   							jsii.String("localStorageTypes"),
//   						},
//   						memoryGiBPerVCpu: &memoryGiBPerVCpuRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						memoryMiB: &memoryMiBRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						networkInterfaceCount: &networkInterfaceCountRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						requireHibernateSupport: jsii.Boolean(false),
//   						spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						totalLocalStorageGb: &totalLocalStorageGBRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						vCpuCount: &vCpuCountRangeRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   					},
//   					instanceType: jsii.String("instanceType"),
//   					maxPrice: jsii.String("maxPrice"),
//   					placement: &placementProperty{
//   						affinity: jsii.String("affinity"),
//   						availabilityZone: jsii.String("availabilityZone"),
//   						groupName: jsii.String("groupName"),
//   						hostId: jsii.String("hostId"),
//   						hostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   						partitionNumber: jsii.Number(123),
//   						spreadDomain: jsii.String("spreadDomain"),
//   						tenancy: jsii.String("tenancy"),
//   					},
//   					priority: jsii.Number(123),
//   					subnetId: jsii.String("subnetId"),
//   					weightedCapacity: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	targetCapacitySpecification: &targetCapacitySpecificationRequestProperty{
//   		totalTargetCapacity: jsii.Number(123),
//
//   		// the properties below are optional
//   		defaultTargetCapacityType: jsii.String("defaultTargetCapacityType"),
//   		onDemandTargetCapacity: jsii.Number(123),
//   		spotTargetCapacity: jsii.Number(123),
//   		targetCapacityUnitType: jsii.String("targetCapacityUnitType"),
//   	},
//
//   	// the properties below are optional
//   	context: jsii.String("context"),
//   	excessCapacityTerminationPolicy: jsii.String("excessCapacityTerminationPolicy"),
//   	onDemandOptions: &onDemandOptionsRequestProperty{
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   		capacityReservationOptions: &capacityReservationOptionsRequestProperty{
//   			usageStrategy: jsii.String("usageStrategy"),
//   		},
//   		maxTotalPrice: jsii.String("maxTotalPrice"),
//   		minTargetCapacity: jsii.Number(123),
//   		singleAvailabilityZone: jsii.Boolean(false),
//   		singleInstanceType: jsii.Boolean(false),
//   	},
//   	replaceUnhealthyInstances: jsii.Boolean(false),
//   	spotOptions: &spotOptionsRequestProperty{
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   		instanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   		instancePoolsToUseCount: jsii.Number(123),
//   		maintenanceStrategies: &maintenanceStrategiesProperty{
//   			capacityRebalance: &capacityRebalanceProperty{
//   				replacementStrategy: jsii.String("replacementStrategy"),
//   				terminationDelay: jsii.Number(123),
//   			},
//   		},
//   		maxTotalPrice: jsii.String("maxTotalPrice"),
//   		minTargetCapacity: jsii.Number(123),
//   		singleAvailabilityZone: jsii.Boolean(false),
//   		singleInstanceType: jsii.Boolean(false),
//   	},
//   	tagSpecifications: []interface{}{
//   		&tagSpecificationProperty{
//   			resourceType: jsii.String("resourceType"),
//   			tags: []cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	terminateInstancesWithExpiration: jsii.Boolean(false),
//   	type: jsii.String("type"),
//   	validFrom: jsii.String("validFrom"),
//   	validUntil: jsii.String("validUntil"),
//   })
//
type CfnEC2Fleet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID of the EC2 Fleet.
	AttrFleetId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Reserved.
	Context() *string
	SetContext(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Indicates whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2 Fleet.
	ExcessCapacityTerminationPolicy() *string
	SetExcessCapacityTerminationPolicy(val *string)
	// The configuration for the EC2 Fleet.
	LaunchTemplateConfigs() interface{}
	SetLaunchTemplateConfigs(val interface{})
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
	// Describes the configuration of On-Demand Instances in an EC2 Fleet.
	OnDemandOptions() interface{}
	SetOnDemandOptions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Indicates whether EC2 Fleet should replace unhealthy Spot Instances.
	//
	// Supported only for fleets of type `maintain` . For more information, see [EC2 Fleet health checks](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/manage-ec2-fleet.html#ec2-fleet-health-checks) in the *Amazon EC2 User Guide* .
	ReplaceUnhealthyInstances() interface{}
	SetReplaceUnhealthyInstances(val interface{})
	// Describes the configuration of Spot Instances in an EC2 Fleet.
	SpotOptions() interface{}
	SetSpotOptions(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The key-value pair for tagging the EC2 Fleet request on creation. For more information, see [Tagging your resources](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#tag-resources) .
	//
	// If the fleet type is `instant` , specify a resource type of `fleet` to tag the fleet or `instance` to tag the instances at launch.
	//
	// If the fleet type is `maintain` or `request` , specify a resource type of `fleet` to tag the fleet. You cannot specify a resource type of `instance` . To tag instances at launch, specify the tags in a [launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#create-launch-template) .
	TagSpecifications() interface{}
	SetTagSpecifications(val interface{})
	// The number of units to request.
	TargetCapacitySpecification() interface{}
	SetTargetCapacitySpecification(val interface{})
	// Indicates whether running instances should be terminated when the EC2 Fleet expires.
	TerminateInstancesWithExpiration() interface{}
	SetTerminateInstancesWithExpiration(val interface{})
	// The fleet type. The default value is `maintain` .
	//
	// - `maintain` - The EC2 Fleet places an asynchronous request for your desired capacity, and continues to maintain your desired Spot capacity by replenishing interrupted Spot Instances.
	// - `request` - The EC2 Fleet places an asynchronous one-time request for your desired capacity, but does submit Spot requests in alternative capacity pools if Spot capacity is unavailable, and does not maintain Spot capacity if Spot Instances are interrupted.
	// - `instant` - The EC2 Fleet places a synchronous one-time request for your desired capacity, and returns errors for any instances that could not be launched.
	//
	// For more information, see [EC2 Fleet request types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-request-type.html) in the *Amazon EC2 User Guide* .
	Type() *string
	SetType(val *string)
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
	// The start date and time of the request, in UTC format (for example, *YYYY* - *MM* - *DD* T *HH* : *MM* : *SS* Z).
	//
	// The default is to start fulfilling the request immediately.
	ValidFrom() *string
	SetValidFrom(val *string)
	// The end date and time of the request, in UTC format (for example, *YYYY* - *MM* - *DD* T *HH* : *MM* : *SS* Z).
	//
	// At this point, no new EC2 Fleet requests are placed or able to fulfill the request. If no value is specified, the request remains until you cancel it.
	ValidUntil() *string
	SetValidUntil(val *string)
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

// The jsii proxy struct for CfnEC2Fleet
type jsiiProxy_CfnEC2Fleet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEC2Fleet) AttrFleetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFleetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) ExcessCapacityTerminationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) LaunchTemplateConfigs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) OnDemandOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDemandOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) ReplaceUnhealthyInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) SpotOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) TagSpecifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) TargetCapacitySpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetCapacitySpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) TerminateInstancesWithExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) ValidFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEC2Fleet) ValidUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntil",
		&returns,
	)
	return returns
}


// Create a new `AWS::EC2::EC2Fleet`.
func NewCfnEC2Fleet(scope constructs.Construct, id *string, props *CfnEC2FleetProps) CfnEC2Fleet {
	_init_.Initialize()

	if err := validateNewCfnEC2FleetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEC2Fleet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EC2::EC2Fleet`.
func NewCfnEC2Fleet_Override(c CfnEC2Fleet, scope constructs.Construct, id *string, props *CfnEC2FleetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEC2Fleet)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_CfnEC2Fleet)SetExcessCapacityTerminationPolicy(val *string) {
	_jsii_.Set(
		j,
		"excessCapacityTerminationPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnEC2Fleet)SetLaunchTemplateConfigs(val interface{}) {
	if err := j.validateSetLaunchTemplateConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchTemplateConfigs",
		val,
	)
}

func (j *jsiiProxy_CfnEC2Fleet)SetOnDemandOptions(val interface{}) {
	if err := j.validateSetOnDemandOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandOptions",
		val,
	)
}

func (j *jsiiProxy_CfnEC2Fleet)SetReplaceUnhealthyInstances(val interface{}) {
	if err := j.validateSetReplaceUnhealthyInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceUnhealthyInstances",
		val,
	)
}

func (j *jsiiProxy_CfnEC2Fleet)SetSpotOptions(val interface{}) {
	if err := j.validateSetSpotOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotOptions",
		val,
	)
}

func (j *jsiiProxy_CfnEC2Fleet)SetTagSpecifications(val interface{}) {
	if err := j.validateSetTagSpecificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagSpecifications",
		val,
	)
}

func (j *jsiiProxy_CfnEC2Fleet)SetTargetCapacitySpecification(val interface{}) {
	if err := j.validateSetTargetCapacitySpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCapacitySpecification",
		val,
	)
}

func (j *jsiiProxy_CfnEC2Fleet)SetTerminateInstancesWithExpiration(val interface{}) {
	if err := j.validateSetTerminateInstancesWithExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateInstancesWithExpiration",
		val,
	)
}

func (j *jsiiProxy_CfnEC2Fleet)SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnEC2Fleet)SetValidFrom(val *string) {
	_jsii_.Set(
		j,
		"validFrom",
		val,
	)
}

func (j *jsiiProxy_CfnEC2Fleet)SetValidUntil(val *string) {
	_jsii_.Set(
		j,
		"validUntil",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEC2Fleet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEC2Fleet_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEC2Fleet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnEC2Fleet_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet",
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
func CfnEC2Fleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEC2Fleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEC2Fleet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEC2Fleet) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEC2Fleet) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEC2Fleet) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEC2Fleet) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEC2Fleet) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEC2Fleet) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEC2Fleet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEC2Fleet) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnEC2Fleet) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnEC2Fleet) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEC2Fleet) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEC2Fleet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnEC2Fleet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEC2Fleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEC2Fleet) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

