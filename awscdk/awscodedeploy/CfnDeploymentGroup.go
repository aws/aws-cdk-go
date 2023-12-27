package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CodeDeploy::DeploymentGroup` resource creates an AWS CodeDeploy deployment group that specifies which instances your application revisions are deployed to, along with other deployment options.
//
// For more information, see [CreateDeploymentGroup](https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_CreateDeploymentGroup.html) in the *CodeDeploy API Reference* .
//
// > Amazon ECS blue/green deployments through CodeDeploy do not use the `AWS::CodeDeploy::DeploymentGroup` resource. To perform Amazon ECS blue/green deployments, use the `AWS::CodeDeploy::BlueGreen` hook. See [Perform Amazon ECS blue/green deployments through CodeDeploy using AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green.html) for more information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentGroup := awscdk.Aws_codedeploy.NewCfnDeploymentGroup(this, jsii.String("MyCfnDeploymentGroup"), &CfnDeploymentGroupProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	ServiceRoleArn: jsii.String("serviceRoleArn"),
//
//   	// the properties below are optional
//   	AlarmConfiguration: &AlarmConfigurationProperty{
//   		Alarms: []interface{}{
//   			&AlarmProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Enabled: jsii.Boolean(false),
//   		IgnorePollAlarmFailure: jsii.Boolean(false),
//   	},
//   	AutoRollbackConfiguration: &AutoRollbackConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		Events: []*string{
//   			jsii.String("events"),
//   		},
//   	},
//   	AutoScalingGroups: []*string{
//   		jsii.String("autoScalingGroups"),
//   	},
//   	BlueGreenDeploymentConfiguration: &BlueGreenDeploymentConfigurationProperty{
//   		DeploymentReadyOption: &DeploymentReadyOptionProperty{
//   			ActionOnTimeout: jsii.String("actionOnTimeout"),
//   			WaitTimeInMinutes: jsii.Number(123),
//   		},
//   		GreenFleetProvisioningOption: &GreenFleetProvisioningOptionProperty{
//   			Action: jsii.String("action"),
//   		},
//   		TerminateBlueInstancesOnDeploymentSuccess: &BlueInstanceTerminationOptionProperty{
//   			Action: jsii.String("action"),
//   			TerminationWaitTimeInMinutes: jsii.Number(123),
//   		},
//   	},
//   	Deployment: &DeploymentProperty{
//   		Revision: &RevisionLocationProperty{
//   			GitHubLocation: &GitHubLocationProperty{
//   				CommitId: jsii.String("commitId"),
//   				Repository: jsii.String("repository"),
//   			},
//   			RevisionType: jsii.String("revisionType"),
//   			S3Location: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//
//   				// the properties below are optional
//   				BundleType: jsii.String("bundleType"),
//   				ETag: jsii.String("eTag"),
//   				Version: jsii.String("version"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   		IgnoreApplicationStopFailures: jsii.Boolean(false),
//   	},
//   	DeploymentConfigName: jsii.String("deploymentConfigName"),
//   	DeploymentGroupName: jsii.String("deploymentGroupName"),
//   	DeploymentStyle: &DeploymentStyleProperty{
//   		DeploymentOption: jsii.String("deploymentOption"),
//   		DeploymentType: jsii.String("deploymentType"),
//   	},
//   	Ec2TagFilters: []interface{}{
//   		&EC2TagFilterProperty{
//   			Key: jsii.String("key"),
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Ec2TagSet: &EC2TagSetProperty{
//   		Ec2TagSetList: []interface{}{
//   			&EC2TagSetListObjectProperty{
//   				Ec2TagGroup: []interface{}{
//   					&EC2TagFilterProperty{
//   						Key: jsii.String("key"),
//   						Type: jsii.String("type"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	EcsServices: []interface{}{
//   		&ECSServiceProperty{
//   			ClusterName: jsii.String("clusterName"),
//   			ServiceName: jsii.String("serviceName"),
//   		},
//   	},
//   	LoadBalancerInfo: &LoadBalancerInfoProperty{
//   		ElbInfoList: []interface{}{
//   			&ELBInfoProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		TargetGroupInfoList: []interface{}{
//   			&TargetGroupInfoProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		TargetGroupPairInfoList: []interface{}{
//   			&TargetGroupPairInfoProperty{
//   				ProdTrafficRoute: &TrafficRouteProperty{
//   					ListenerArns: []*string{
//   						jsii.String("listenerArns"),
//   					},
//   				},
//   				TargetGroups: []interface{}{
//   					&TargetGroupInfoProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				TestTrafficRoute: &TrafficRouteProperty{
//   					ListenerArns: []*string{
//   						jsii.String("listenerArns"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	OnPremisesInstanceTagFilters: []interface{}{
//   		&TagFilterProperty{
//   			Key: jsii.String("key"),
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	OnPremisesTagSet: &OnPremisesTagSetProperty{
//   		OnPremisesTagSetList: []interface{}{
//   			&OnPremisesTagSetListObjectProperty{
//   				OnPremisesTagGroup: []interface{}{
//   					&TagFilterProperty{
//   						Key: jsii.String("key"),
//   						Type: jsii.String("type"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	OutdatedInstancesStrategy: jsii.String("outdatedInstancesStrategy"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TerminationHookEnabled: jsii.Boolean(false),
//   	TriggerConfigurations: []interface{}{
//   		&TriggerConfigProperty{
//   			TriggerEvents: []*string{
//   				jsii.String("triggerEvents"),
//   			},
//   			TriggerName: jsii.String("triggerName"),
//   			TriggerTargetArn: jsii.String("triggerTargetArn"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html
//
type CfnDeploymentGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// Information about the Amazon CloudWatch alarms that are associated with the deployment group.
	AlarmConfiguration() interface{}
	SetAlarmConfiguration(val interface{})
	// The name of an existing CodeDeploy application to associate this deployment group with.
	ApplicationName() *string
	SetApplicationName(val *string)
	AttrId() *string
	// Information about the automatic rollback configuration that is associated with the deployment group.
	AutoRollbackConfiguration() interface{}
	SetAutoRollbackConfiguration(val interface{})
	// A list of associated Auto Scaling groups that CodeDeploy automatically deploys revisions to when new instances are created.
	AutoScalingGroups() *[]*string
	SetAutoScalingGroups(val *[]*string)
	// Information about blue/green deployment options for a deployment group.
	BlueGreenDeploymentConfiguration() interface{}
	SetBlueGreenDeploymentConfiguration(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The application revision to deploy to this deployment group.
	Deployment() interface{}
	SetDeployment(val interface{})
	// A deployment configuration name or a predefined configuration name.
	DeploymentConfigName() *string
	SetDeploymentConfigName(val *string)
	// A name for the deployment group.
	DeploymentGroupName() *string
	SetDeploymentGroupName(val *string)
	// Attributes that determine the type of deployment to run and whether to route deployment traffic behind a load balancer.
	DeploymentStyle() interface{}
	SetDeploymentStyle(val interface{})
	// The Amazon EC2 tags that are already applied to Amazon EC2 instances that you want to include in the deployment group.
	Ec2TagFilters() interface{}
	SetEc2TagFilters(val interface{})
	// Information about groups of tags applied to Amazon EC2 instances.
	Ec2TagSet() interface{}
	SetEc2TagSet(val interface{})
	// The target Amazon ECS services in the deployment group.
	EcsServices() interface{}
	SetEcsServices(val interface{})
	// Information about the load balancer to use in a deployment.
	LoadBalancerInfo() interface{}
	SetLoadBalancerInfo(val interface{})
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
	// The on-premises instance tags already applied to on-premises instances that you want to include in the deployment group.
	OnPremisesInstanceTagFilters() interface{}
	SetOnPremisesInstanceTagFilters(val interface{})
	// Information about groups of tags applied to on-premises instances.
	OnPremisesTagSet() interface{}
	SetOnPremisesTagSet(val interface{})
	// Indicates what happens when new Amazon EC2 instances are launched mid-deployment and do not receive the deployed application revision.
	OutdatedInstancesStrategy() *string
	SetOutdatedInstancesStrategy(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A service role Amazon Resource Name (ARN) that grants CodeDeploy permission to make calls to AWS services on your behalf.
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The metadata that you apply to CodeDeploy deployment groups to help you organize and categorize them.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// Indicates whether the deployment group was configured to have CodeDeploy install a termination hook into an Auto Scaling group.
	TerminationHookEnabled() interface{}
	SetTerminationHookEnabled(val interface{})
	// Information about triggers associated with the deployment group.
	TriggerConfigurations() interface{}
	SetTriggerConfigurations(val interface{})
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

// The jsii proxy struct for CfnDeploymentGroup
type jsiiProxy_CfnDeploymentGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnDeploymentGroup) AlarmConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) AutoRollbackConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRollbackConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) AutoScalingGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoScalingGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) BlueGreenDeploymentConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blueGreenDeploymentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Deployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) DeploymentStyle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Ec2TagFilters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2TagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Ec2TagSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2TagSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) EcsServices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) LoadBalancerInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) OnPremisesInstanceTagFilters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onPremisesInstanceTagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) OnPremisesTagSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onPremisesTagSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) OutdatedInstancesStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdatedInstancesStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) TerminationHookEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationHookEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) TriggerConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroup) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnDeploymentGroup(scope constructs.Construct, id *string, props *CfnDeploymentGroupProps) CfnDeploymentGroup {
	_init_.Initialize()

	if err := validateNewCfnDeploymentGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDeploymentGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnDeploymentGroup_Override(c CfnDeploymentGroup, scope constructs.Construct, id *string, props *CfnDeploymentGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetAlarmConfiguration(val interface{}) {
	if err := j.validateSetAlarmConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetApplicationName(val *string) {
	if err := j.validateSetApplicationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetAutoRollbackConfiguration(val interface{}) {
	if err := j.validateSetAutoRollbackConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRollbackConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetAutoScalingGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"autoScalingGroups",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetBlueGreenDeploymentConfiguration(val interface{}) {
	if err := j.validateSetBlueGreenDeploymentConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blueGreenDeploymentConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetDeployment(val interface{}) {
	if err := j.validateSetDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deployment",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetDeploymentConfigName(val *string) {
	_jsii_.Set(
		j,
		"deploymentConfigName",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetDeploymentGroupName(val *string) {
	_jsii_.Set(
		j,
		"deploymentGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetDeploymentStyle(val interface{}) {
	if err := j.validateSetDeploymentStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentStyle",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetEc2TagFilters(val interface{}) {
	if err := j.validateSetEc2TagFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2TagFilters",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetEc2TagSet(val interface{}) {
	if err := j.validateSetEc2TagSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2TagSet",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetEcsServices(val interface{}) {
	if err := j.validateSetEcsServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecsServices",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetLoadBalancerInfo(val interface{}) {
	if err := j.validateSetLoadBalancerInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerInfo",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetOnPremisesInstanceTagFilters(val interface{}) {
	if err := j.validateSetOnPremisesInstanceTagFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onPremisesInstanceTagFilters",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetOnPremisesTagSet(val interface{}) {
	if err := j.validateSetOnPremisesTagSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onPremisesTagSet",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetOutdatedInstancesStrategy(val *string) {
	_jsii_.Set(
		j,
		"outdatedInstancesStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetTerminationHookEnabled(val interface{}) {
	if err := j.validateSetTerminationHookEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationHookEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnDeploymentGroup)SetTriggerConfigurations(val interface{}) {
	if err := j.validateSetTriggerConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerConfigurations",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDeploymentGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeploymentGroup_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDeploymentGroup_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeploymentGroup_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
		"isCfnResource",
		[]interface{}{x},
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
func CfnDeploymentGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeploymentGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeploymentGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeploymentGroup) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnDeploymentGroup) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnDeploymentGroup) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeploymentGroup) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeploymentGroup) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnDeploymentGroup) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnDeploymentGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeploymentGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeploymentGroup) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

