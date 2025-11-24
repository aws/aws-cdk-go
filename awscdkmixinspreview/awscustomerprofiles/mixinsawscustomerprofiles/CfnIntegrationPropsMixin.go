package mixinsawscustomerprofiles

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awscustomerprofiles/mixinsawscustomerprofiles/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an Amazon Connect Customer Profiles Integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIntegrationPropsMixin := awscdkmixinspreview.Mixins.NewCfnIntegrationPropsMixin(&CfnIntegrationMixinProps{
//   	DomainName: jsii.String("domainName"),
//   	EventTriggerNames: []*string{
//   		jsii.String("eventTriggerNames"),
//   	},
//   	FlowDefinition: &FlowDefinitionProperty{
//   		Description: jsii.String("description"),
//   		FlowName: jsii.String("flowName"),
//   		KmsArn: jsii.String("kmsArn"),
//   		SourceFlowConfig: &SourceFlowConfigProperty{
//   			ConnectorProfileName: jsii.String("connectorProfileName"),
//   			ConnectorType: jsii.String("connectorType"),
//   			IncrementalPullConfig: &IncrementalPullConfigProperty{
//   				DatetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   			},
//   			SourceConnectorProperties: &SourceConnectorPropertiesProperty{
//   				Marketo: &MarketoSourcePropertiesProperty{
//   					Object: jsii.String("object"),
//   				},
//   				S3: &S3SourcePropertiesProperty{
//   					BucketName: jsii.String("bucketName"),
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   				},
//   				Salesforce: &SalesforceSourcePropertiesProperty{
//   					EnableDynamicFieldUpdate: jsii.Boolean(false),
//   					IncludeDeletedRecords: jsii.Boolean(false),
//   					Object: jsii.String("object"),
//   				},
//   				ServiceNow: &ServiceNowSourcePropertiesProperty{
//   					Object: jsii.String("object"),
//   				},
//   				Zendesk: &ZendeskSourcePropertiesProperty{
//   					Object: jsii.String("object"),
//   				},
//   			},
//   		},
//   		Tasks: []interface{}{
//   			&TaskProperty{
//   				ConnectorOperator: &ConnectorOperatorProperty{
//   					Marketo: jsii.String("marketo"),
//   					S3: jsii.String("s3"),
//   					Salesforce: jsii.String("salesforce"),
//   					ServiceNow: jsii.String("serviceNow"),
//   					Zendesk: jsii.String("zendesk"),
//   				},
//   				DestinationField: jsii.String("destinationField"),
//   				SourceFields: []*string{
//   					jsii.String("sourceFields"),
//   				},
//   				TaskProperties: []interface{}{
//   					&TaskPropertiesMapProperty{
//   						OperatorPropertyKey: jsii.String("operatorPropertyKey"),
//   						Property: jsii.String("property"),
//   					},
//   				},
//   				TaskType: jsii.String("taskType"),
//   			},
//   		},
//   		TriggerConfig: &TriggerConfigProperty{
//   			TriggerProperties: &TriggerPropertiesProperty{
//   				Scheduled: &ScheduledTriggerPropertiesProperty{
//   					DataPullMode: jsii.String("dataPullMode"),
//   					FirstExecutionFrom: jsii.Number(123),
//   					ScheduleEndTime: jsii.Number(123),
//   					ScheduleExpression: jsii.String("scheduleExpression"),
//   					ScheduleOffset: jsii.Number(123),
//   					ScheduleStartTime: jsii.Number(123),
//   					Timezone: jsii.String("timezone"),
//   				},
//   			},
//   			TriggerType: jsii.String("triggerType"),
//   		},
//   	},
//   	ObjectTypeName: jsii.String("objectTypeName"),
//   	ObjectTypeNames: []interface{}{
//   		&ObjectTypeMappingProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Uri: jsii.String("uri"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-integration.html
//
type CfnIntegrationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIntegrationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIntegrationPropsMixin
type jsiiProxy_CfnIntegrationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIntegrationPropsMixin) Props() *CfnIntegrationMixinProps {
	var returns *CfnIntegrationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CustomerProfiles::Integration`.
func NewCfnIntegrationPropsMixin(props *CfnIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIntegrationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIntegrationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIntegrationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnIntegrationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CustomerProfiles::Integration`.
func NewCfnIntegrationPropsMixin_Override(c CfnIntegrationPropsMixin, props *CfnIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnIntegrationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIntegrationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIntegrationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnIntegrationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIntegrationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnIntegrationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIntegrationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnIntegrationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

