package previewawsarcregionswitchmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsarcregionswitch/previewawsarcregionswitchmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Region switch plan.
//
// A plan defines the steps required to shift traffic from one AWS Region to another.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var stepProperty_ StepProperty
//
//   cfnPlanPropsMixin := awscdkmixinspreview.Mixins.NewCfnPlanPropsMixin(&CfnPlanMixinProps{
//   	AssociatedAlarms: map[string]interface{}{
//   		"associatedAlarmsKey": &AssociatedAlarmProperty{
//   			"alarmType": jsii.String("alarmType"),
//   			"crossAccountRole": jsii.String("crossAccountRole"),
//   			"externalId": jsii.String("externalId"),
//   			"resourceIdentifier": jsii.String("resourceIdentifier"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ExecutionRole: jsii.String("executionRole"),
//   	Name: jsii.String("name"),
//   	PrimaryRegion: jsii.String("primaryRegion"),
//   	RecoveryApproach: jsii.String("recoveryApproach"),
//   	RecoveryTimeObjectiveMinutes: jsii.Number(123),
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	ReportConfiguration: &ReportConfigurationProperty{
//   		ReportOutput: []interface{}{
//   			&ReportOutputConfigurationProperty{
//   				S3Configuration: &S3ReportOutputConfigurationProperty{
//   					BucketOwner: jsii.String("bucketOwner"),
//   					BucketPath: jsii.String("bucketPath"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Triggers: []interface{}{
//   		&TriggerProperty{
//   			Action: jsii.String("action"),
//   			Conditions: []interface{}{
//   				&TriggerConditionProperty{
//   					AssociatedAlarmName: jsii.String("associatedAlarmName"),
//   					Condition: jsii.String("condition"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			MinDelayMinutesBetweenExecutions: jsii.Number(123),
//   			TargetRegion: jsii.String("targetRegion"),
//   		},
//   	},
//   	Workflows: []interface{}{
//   		&WorkflowProperty{
//   			Steps: []interface{}{
//   				&StepProperty{
//   					Description: jsii.String("description"),
//   					ExecutionBlockConfiguration: &ExecutionBlockConfigurationProperty{
//   						ArcRoutingControlConfig: &ArcRoutingControlConfigurationProperty{
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
//   							RegionAndRoutingControls: map[string]interface{}{
//   								"regionAndRoutingControlsKey": []interface{}{
//   									&ArcRoutingControlStateProperty{
//   										"routingControlArn": jsii.String("routingControlArn"),
//   										"state": jsii.String("state"),
//   									},
//   								},
//   							},
//   							TimeoutMinutes: jsii.Number(123),
//   						},
//   						CustomActionLambdaConfig: &CustomActionLambdaConfigurationProperty{
//   							Lambdas: []interface{}{
//   								&LambdasProperty{
//   									Arn: jsii.String("arn"),
//   									CrossAccountRole: jsii.String("crossAccountRole"),
//   									ExternalId: jsii.String("externalId"),
//   								},
//   							},
//   							RegionToRun: jsii.String("regionToRun"),
//   							RetryIntervalMinutes: jsii.Number(123),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &LambdaUngracefulProperty{
//   								Behavior: jsii.String("behavior"),
//   							},
//   						},
//   						DocumentDbConfig: &DocumentDbConfigurationProperty{
//   							Behavior: jsii.String("behavior"),
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							DatabaseClusterArns: []*string{
//   								jsii.String("databaseClusterArns"),
//   							},
//   							ExternalId: jsii.String("externalId"),
//   							GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &DocumentDbUngracefulProperty{
//   								Ungraceful: jsii.String("ungraceful"),
//   							},
//   						},
//   						Ec2AsgCapacityIncreaseConfig: &Ec2AsgCapacityIncreaseConfigurationProperty{
//   							Asgs: []interface{}{
//   								&AsgProperty{
//   									Arn: jsii.String("arn"),
//   									CrossAccountRole: jsii.String("crossAccountRole"),
//   									ExternalId: jsii.String("externalId"),
//   								},
//   							},
//   							CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   							TargetPercent: jsii.Number(123),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &Ec2UngracefulProperty{
//   								MinimumSuccessPercentage: jsii.Number(123),
//   							},
//   						},
//   						EcsCapacityIncreaseConfig: &EcsCapacityIncreaseConfigurationProperty{
//   							CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   							Services: []interface{}{
//   								&ServiceProperty{
//   									ClusterArn: jsii.String("clusterArn"),
//   									CrossAccountRole: jsii.String("crossAccountRole"),
//   									ExternalId: jsii.String("externalId"),
//   									ServiceArn: jsii.String("serviceArn"),
//   								},
//   							},
//   							TargetPercent: jsii.Number(123),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &EcsUngracefulProperty{
//   								MinimumSuccessPercentage: jsii.Number(123),
//   							},
//   						},
//   						EksResourceScalingConfig: &EksResourceScalingConfigurationProperty{
//   							CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   							EksClusters: []interface{}{
//   								&EksClusterProperty{
//   									ClusterArn: jsii.String("clusterArn"),
//   									CrossAccountRole: jsii.String("crossAccountRole"),
//   									ExternalId: jsii.String("externalId"),
//   								},
//   							},
//   							KubernetesResourceType: &KubernetesResourceTypeProperty{
//   								ApiVersion: jsii.String("apiVersion"),
//   								Kind: jsii.String("kind"),
//   							},
//   							ScalingResources: []interface{}{
//   								map[string]interface{}{
//   									"scalingResourcesKey": map[string]interface{}{
//   										"scalingResourcesKey": &KubernetesScalingResourceProperty{
//   											"hpaName": jsii.String("hpaName"),
//   											"name": jsii.String("name"),
//   											"namespace": jsii.String("namespace"),
//   										},
//   									},
//   								},
//   							},
//   							TargetPercent: jsii.Number(123),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &EksResourceScalingUngracefulProperty{
//   								MinimumSuccessPercentage: jsii.Number(123),
//   							},
//   						},
//   						ExecutionApprovalConfig: &ExecutionApprovalConfigurationProperty{
//   							ApprovalRole: jsii.String("approvalRole"),
//   							TimeoutMinutes: jsii.Number(123),
//   						},
//   						GlobalAuroraConfig: &GlobalAuroraConfigurationProperty{
//   							Behavior: jsii.String("behavior"),
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							DatabaseClusterArns: []*string{
//   								jsii.String("databaseClusterArns"),
//   							},
//   							ExternalId: jsii.String("externalId"),
//   							GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   							TimeoutMinutes: jsii.Number(123),
//   							Ungraceful: &GlobalAuroraUngracefulProperty{
//   								Ungraceful: jsii.String("ungraceful"),
//   							},
//   						},
//   						ParallelConfig: &ParallelExecutionBlockConfigurationProperty{
//   							Steps: []interface{}{
//   								stepProperty_,
//   							},
//   						},
//   						RegionSwitchPlanConfig: &RegionSwitchPlanConfigurationProperty{
//   							Arn: jsii.String("arn"),
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
//   						},
//   						Route53HealthCheckConfig: &Route53HealthCheckConfigurationProperty{
//   							CrossAccountRole: jsii.String("crossAccountRole"),
//   							ExternalId: jsii.String("externalId"),
//   							HostedZoneId: jsii.String("hostedZoneId"),
//   							RecordName: jsii.String("recordName"),
//   							RecordSets: []interface{}{
//   								&Route53ResourceRecordSetProperty{
//   									RecordSetIdentifier: jsii.String("recordSetIdentifier"),
//   									Region: jsii.String("region"),
//   								},
//   							},
//   							TimeoutMinutes: jsii.Number(123),
//   						},
//   					},
//   					ExecutionBlockType: jsii.String("executionBlockType"),
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			WorkflowDescription: jsii.String("workflowDescription"),
//   			WorkflowTargetAction: jsii.String("workflowTargetAction"),
//   			WorkflowTargetRegion: jsii.String("workflowTargetRegion"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arcregionswitch-plan.html
//
type CfnPlanPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPlanMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPlanPropsMixin
type jsiiProxy_CfnPlanPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPlanPropsMixin) Props() *CfnPlanMixinProps {
	var returns *CfnPlanMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlanPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ARCRegionSwitch::Plan`.
func NewCfnPlanPropsMixin(props *CfnPlanMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPlanPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPlanPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPlanPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_arcregionswitch.mixins.CfnPlanPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ARCRegionSwitch::Plan`.
func NewCfnPlanPropsMixin_Override(c CfnPlanPropsMixin, props *CfnPlanMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_arcregionswitch.mixins.CfnPlanPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPlanPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPlanPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_arcregionswitch.mixins.CfnPlanPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPlanPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_arcregionswitch.mixins.CfnPlanPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPlanPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlanPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

