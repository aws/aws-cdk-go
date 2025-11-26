package previewawsauditmanagermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsauditmanager/previewawsauditmanagermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AuditManager::Assessment` resource is an Audit Manager resource type that defines the scope of audit evidence collected by Audit Manager .
//
// An Audit Manager assessment is an implementation of an Audit Manager framework.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAssessmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnAssessmentPropsMixin(&CfnAssessmentMixinProps{
//   	AssessmentReportsDestination: &AssessmentReportsDestinationProperty{
//   		Destination: jsii.String("destination"),
//   		DestinationType: jsii.String("destinationType"),
//   	},
//   	AwsAccount: &AWSAccountProperty{
//   		EmailAddress: jsii.String("emailAddress"),
//   		Id: jsii.String("id"),
//   		Name: jsii.String("name"),
//   	},
//   	Delegations: []interface{}{
//   		&DelegationProperty{
//   			AssessmentId: jsii.String("assessmentId"),
//   			AssessmentName: jsii.String("assessmentName"),
//   			Comment: jsii.String("comment"),
//   			ControlSetId: jsii.String("controlSetId"),
//   			CreatedBy: jsii.String("createdBy"),
//   			CreationTime: jsii.Number(123),
//   			Id: jsii.String("id"),
//   			LastUpdated: jsii.Number(123),
//   			RoleArn: jsii.String("roleArn"),
//   			RoleType: jsii.String("roleType"),
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	FrameworkId: jsii.String("frameworkId"),
//   	Name: jsii.String("name"),
//   	Roles: []interface{}{
//   		&RoleProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			RoleType: jsii.String("roleType"),
//   		},
//   	},
//   	Scope: &ScopeProperty{
//   		AwsAccounts: []interface{}{
//   			&AWSAccountProperty{
//   				EmailAddress: jsii.String("emailAddress"),
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		AwsServices: []interface{}{
//   			&AWSServiceProperty{
//   				ServiceName: jsii.String("serviceName"),
//   			},
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html
//
type CfnAssessmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAssessmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAssessmentPropsMixin
type jsiiProxy_CfnAssessmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAssessmentPropsMixin) Props() *CfnAssessmentMixinProps {
	var returns *CfnAssessmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssessmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AuditManager::Assessment`.
func NewCfnAssessmentPropsMixin(props *CfnAssessmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAssessmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAssessmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAssessmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_auditmanager.mixins.CfnAssessmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AuditManager::Assessment`.
func NewCfnAssessmentPropsMixin_Override(c CfnAssessmentPropsMixin, props *CfnAssessmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_auditmanager.mixins.CfnAssessmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAssessmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssessmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_auditmanager.mixins.CfnAssessmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssessmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_auditmanager.mixins.CfnAssessmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssessmentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAssessmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

