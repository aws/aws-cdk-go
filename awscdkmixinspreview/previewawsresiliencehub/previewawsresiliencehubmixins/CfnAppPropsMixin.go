package previewawsresiliencehubmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsresiliencehub/previewawsresiliencehubmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an AWS Resilience Hub application.
//
// An AWS Resilience Hub application is a collection of AWS resources structured to prevent and recover AWS application disruptions. To describe a AWS Resilience Hub application, you provide an application name, resources from one or more AWS CloudFormation stacks, Resource Groups , Terraform state files, AppRegistry applications, and an appropriate resiliency policy. In addition, you can also add resources that are located on Amazon Elastic Kubernetes Service (Amazon EKS) clusters as optional resources. For more information about the number of resources supported per application, see [Service quotas](https://docs.aws.amazon.com/general/latest/gr/resiliencehub.html#limits_resiliencehub) .
//
// After you create an AWS Resilience Hub application, you publish it so that you can run a resiliency assessment on it. You can then use recommendations from the assessment to improve resiliency by running another assessment, comparing results, and then iterating the process until you achieve your goals for recovery time objective (RTO) and recovery point objective (RPO).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppPropsMixin := awscdkmixinspreview.Mixins.NewCfnAppPropsMixin(&CfnAppMixinProps{
//   	AppAssessmentSchedule: jsii.String("appAssessmentSchedule"),
//   	AppTemplateBody: jsii.String("appTemplateBody"),
//   	Description: jsii.String("description"),
//   	EventSubscriptions: []interface{}{
//   		&EventSubscriptionProperty{
//   			EventType: jsii.String("eventType"),
//   			Name: jsii.String("name"),
//   			SnsTopicArn: jsii.String("snsTopicArn"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	PermissionModel: &PermissionModelProperty{
//   		CrossAccountRoleArns: []*string{
//   			jsii.String("crossAccountRoleArns"),
//   		},
//   		InvokerRoleName: jsii.String("invokerRoleName"),
//   		Type: jsii.String("type"),
//   	},
//   	ResiliencyPolicyArn: jsii.String("resiliencyPolicyArn"),
//   	ResourceMappings: []interface{}{
//   		&ResourceMappingProperty{
//   			EksSourceName: jsii.String("eksSourceName"),
//   			LogicalStackName: jsii.String("logicalStackName"),
//   			MappingType: jsii.String("mappingType"),
//   			PhysicalResourceId: &PhysicalResourceIdProperty{
//   				AwsAccountId: jsii.String("awsAccountId"),
//   				AwsRegion: jsii.String("awsRegion"),
//   				Identifier: jsii.String("identifier"),
//   				Type: jsii.String("type"),
//   			},
//   			ResourceName: jsii.String("resourceName"),
//   			TerraformSourceName: jsii.String("terraformSourceName"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehub-app.html
//
type CfnAppPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAppMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAppPropsMixin
type jsiiProxy_CfnAppPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAppPropsMixin) Props() *CfnAppMixinProps {
	var returns *CfnAppMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ResilienceHub::App`.
func NewCfnAppPropsMixin(props *CfnAppMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAppPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAppPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAppPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_resiliencehub.mixins.CfnAppPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ResilienceHub::App`.
func NewCfnAppPropsMixin_Override(c CfnAppPropsMixin, props *CfnAppMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_resiliencehub.mixins.CfnAppPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAppPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAppPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_resiliencehub.mixins.CfnAppPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAppPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_resiliencehub.mixins.CfnAppPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAppPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAppPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

