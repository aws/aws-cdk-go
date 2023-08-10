package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Batch::ComputeEnvironment` resource defines your AWS Batch compute environment.
//
// You can define `MANAGED` or `UNMANAGED` compute environments. `MANAGED` compute environments can use Amazon EC2 or AWS Fargate resources. `UNMANAGED` compute environments can only use EC2 resources. For more information, see [Compute Environments](https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html) in the ** .
//
// In a managed compute environment, AWS Batch manages the capacity and instance types of the compute resources within the environment. This is based on the compute resource specification that you define or the [launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html) that you specify when you create the compute environment. You can choose either to use EC2 On-Demand Instances and EC2 Spot Instances, or to use Fargate and Fargate Spot capacity in your managed compute environment. You can optionally set a maximum price so that Spot Instances only launch when the Spot Instance price is below a specified percentage of the On-Demand price.
//
// > Multi-node parallel jobs are not supported on Spot Instances.
//
// In an unmanaged compute environment, you can manage your own EC2 compute resources and have a lot of flexibility with how you configure your compute resources. For example, you can use custom AMI. However, you need to verify that your AMI meets the Amazon ECS container instance AMI specification. For more information, see [container instance AMIs](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container_instance_AMIs.html) in the *Amazon Elastic Container Service Developer Guide* . After you have created your unmanaged compute environment, you can use the [DescribeComputeEnvironments](https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeComputeEnvironments.html) operation to find the Amazon ECS cluster that is associated with it. Then, manually launch your container instances into that Amazon ECS cluster. For more information, see [Launching an Amazon ECS container instance](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_container_instance.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// > To create a compute environment that uses EKS resources, the caller must have permissions to call `eks:DescribeCluster` . > AWS Batch doesn't upgrade the AMIs in a compute environment after it's created except under specific conditions. For example, it doesn't automatically update the AMIs when a newer version of the Amazon ECS optimized AMI is available. Therefore, you're responsible for the management of the guest operating system (including updates and security patches) and any additional application software or utilities that you install on the compute resources. There are two ways to use a new AMI for your AWS Batch jobs. The original method is to complete these steps:
// >
// > - Create a new compute environment with the new AMI.
// > - Add the compute environment to an existing job queue.
// > - Remove the earlier compute environment from your job queue.
// > - Delete the earlier compute environment.
// >
// > In April 2022, AWS Batch added enhanced support for updating compute environments. For example, the `UpdateComputeEnvironent` API lets you use the `ReplaceComputeEnvironment` property to dynamically update compute environment parameters such as the launch template or instance type without replacement. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
// >
// > To use the enhanced updating of compute environments to update AMIs, follow these rules:
// >
// > - Either do not set the [ServiceRole](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html#cfn-batch-computeenvironment-servicerole) property or set it to the *AWSServiceRoleForBatch* service-linked role.
// > - Set the [AllocationStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-allocationstrategy) property to `BEST_FIT_PROGRESSIVE` , `SPOT_CAPACITY_OPTIMIZED` or `SPOT_PRICE_CAPACITY_OPTIMIZED` .
// > - Set the [ReplaceComputeEnvironment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html#cfn-batch-computeenvironment-replacecomputeenvironment) property to `false` .
// >
// > > Set the `ReplaceComputeEnvironment` property to `false` if the compute environment uses the `BEST_FIT` allocation strategy. > If the `ReplaceComputeEnvironment` property is set to `false` , you might receive an error message when you update the CFN template for a compute environment. This issue occurs if the updated `desiredvcpus` value is less than the current `desiredvcpus` value. As a workaround, delete the `desiredvcpus` value from the updated template or use the `minvcpus` property to manage the number of vCPUs. For information, see [Error message when you update the `DesiredvCpus` setting](https://docs.aws.amazon.com/batch/latest/userguide/troubleshooting.html#error-desired-vcpus-update) .
// > - Set the [UpdateToLatestImageVersion](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-updatetolatestimageversion) property to `true` . This property is used when you update a compute environment. The [UpdateToLatestImageVersion](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-updatetolatestimageversion) property is ignored when you create a compute environment.
// > - Either do not specify an image ID in [ImageId](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-imageid) or [ImageIdOverride](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-ec2configurationobject.html#cfn-batch-computeenvironment-ec2configurationobject-imageidoverride) properties, or in the launch template identified by the [Launch Template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-launchtemplate) property. In that case AWS Batch will select the latest Amazon ECS optimized AMI supported by AWS Batch at the time the infrastructure update is initiated. Alternatively you can specify the AMI ID in the `ImageId` or `ImageIdOverride` properties, or the launch template identified by the `LaunchTemplate` properties. Changing any of these properties will trigger an infrastructure update.
// >
// > If these rules are followed, any update that triggers an infrastructure update will cause the AMI ID to be re-selected. If the [Version](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-launchtemplatespecification.html#cfn-batch-computeenvironment-launchtemplatespecification-version) property of the [LaunchTemplateSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-launchtemplatespecification.html) is set to `$Latest` or `$Default` , the latest or default version of the launch template will be evaluated up at the time of the infrastructure update, even if the `LaunchTemplateSpecification` was not updated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnComputeEnvironment := awscdk.Aws_batch.NewCfnComputeEnvironment(this, jsii.String("MyCfnComputeEnvironment"), &CfnComputeEnvironmentProps{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ComputeEnvironmentName: jsii.String("computeEnvironmentName"),
//   	ComputeResources: &ComputeResourcesProperty{
//   		MaxvCpus: jsii.Number(123),
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		BidPercentage: jsii.Number(123),
//   		DesiredvCpus: jsii.Number(123),
//   		Ec2Configuration: []interface{}{
//   			&Ec2ConfigurationObjectProperty{
//   				ImageType: jsii.String("imageType"),
//
//   				// the properties below are optional
//   				ImageIdOverride: jsii.String("imageIdOverride"),
//   				ImageKubernetesVersion: jsii.String("imageKubernetesVersion"),
//   			},
//   		},
//   		Ec2KeyPair: jsii.String("ec2KeyPair"),
//   		ImageId: jsii.String("imageId"),
//   		InstanceRole: jsii.String("instanceRole"),
//   		InstanceTypes: []*string{
//   			jsii.String("instanceTypes"),
//   		},
//   		LaunchTemplate: &LaunchTemplateSpecificationProperty{
//   			LaunchTemplateId: jsii.String("launchTemplateId"),
//   			LaunchTemplateName: jsii.String("launchTemplateName"),
//   			Version: jsii.String("version"),
//   		},
//   		MinvCpus: jsii.Number(123),
//   		PlacementGroup: jsii.String("placementGroup"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SpotIamFleetRole: jsii.String("spotIamFleetRole"),
//   		Tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   		UpdateToLatestImageVersion: jsii.Boolean(false),
//   	},
//   	EksConfiguration: &EksConfigurationProperty{
//   		EksClusterArn: jsii.String("eksClusterArn"),
//   		KubernetesNamespace: jsii.String("kubernetesNamespace"),
//   	},
//   	ReplaceComputeEnvironment: jsii.Boolean(false),
//   	ServiceRole: jsii.String("serviceRole"),
//   	State: jsii.String("state"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	UnmanagedvCpus: jsii.Number(123),
//   	UpdatePolicy: &UpdatePolicyProperty{
//   		JobExecutionTimeoutMinutes: jsii.Number(123),
//   		TerminateJobsOnUpdate: jsii.Boolean(false),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html
//
type CfnComputeEnvironment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// Returns the compute environment ARN, such as `batch: *us-east-1* : *111122223333* :compute-environment/ *ComputeEnvironmentName*` .
	AttrComputeEnvironmentArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The name for your compute environment.
	ComputeEnvironmentName() *string
	SetComputeEnvironmentName(val *string)
	// The ComputeResources property type specifies details of the compute resources managed by the compute environment.
	ComputeResources() interface{}
	SetComputeResources(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The details for the Amazon EKS cluster that supports the compute environment.
	EksConfiguration() interface{}
	SetEksConfiguration(val interface{})
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
	// Specifies whether the compute environment is replaced if an update is made that requires replacing the instances in the compute environment.
	ReplaceComputeEnvironment() interface{}
	SetReplaceComputeEnvironment(val interface{})
	// The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
	ServiceRole() *string
	SetServiceRole(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The state of the compute environment.
	State() *string
	SetState(val *string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The tags applied to the compute environment.
	TagsRaw() *map[string]*string
	SetTagsRaw(val *map[string]*string)
	// The type of the compute environment: `MANAGED` or `UNMANAGED` .
	Type() *string
	SetType(val *string)
	// The maximum number of vCPUs for an unmanaged compute environment.
	UnmanagedvCpus() *float64
	SetUnmanagedvCpus(val *float64)
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
	// Specifies the infrastructure update policy for the compute environment.
	UpdatePolicy() interface{}
	SetUpdatePolicy(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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

// The jsii proxy struct for CfnComputeEnvironment
type jsiiProxy_CfnComputeEnvironment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnComputeEnvironment) AttrComputeEnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrComputeEnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) ComputeEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) ComputeResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) EksConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) ReplaceComputeEnvironment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceComputeEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) TagsRaw() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) UnmanagedvCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unmanagedvCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputeEnvironment) UpdatePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}


func NewCfnComputeEnvironment(scope constructs.Construct, id *string, props *CfnComputeEnvironmentProps) CfnComputeEnvironment {
	_init_.Initialize()

	if err := validateNewCfnComputeEnvironmentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnComputeEnvironment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnComputeEnvironment_Override(c CfnComputeEnvironment, scope constructs.Construct, id *string, props *CfnComputeEnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment)SetComputeEnvironmentName(val *string) {
	_jsii_.Set(
		j,
		"computeEnvironmentName",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment)SetComputeResources(val interface{}) {
	if err := j.validateSetComputeResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeResources",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment)SetEksConfiguration(val interface{}) {
	if err := j.validateSetEksConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eksConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment)SetReplaceComputeEnvironment(val interface{}) {
	if err := j.validateSetReplaceComputeEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceComputeEnvironment",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment)SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment)SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment)SetTagsRaw(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment)SetUnmanagedvCpus(val *float64) {
	_jsii_.Set(
		j,
		"unmanagedvCpus",
		val,
	)
}

func (j *jsiiProxy_CfnComputeEnvironment)SetUpdatePolicy(val interface{}) {
	if err := j.validateSetUpdatePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatePolicy",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnComputeEnvironment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComputeEnvironment_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnComputeEnvironment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnComputeEnvironment_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
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
func CfnComputeEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComputeEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComputeEnvironment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnComputeEnvironment) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComputeEnvironment) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnComputeEnvironment) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComputeEnvironment) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComputeEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnComputeEnvironment) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnComputeEnvironment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComputeEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComputeEnvironment) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

