package previewawspinpointmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspinpoint/previewawspinpointmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies the settings for a campaign.
//
// A *campaign* is a messaging initiative that engages a specific segment of users for an Amazon Pinpoint application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var attributes interface{}
//   var customConfig interface{}
//   var metrics interface{}
//   var tags interface{}
//
//   cfnCampaignPropsMixin := awscdkmixinspreview.Mixins.NewCfnCampaignPropsMixin(&CfnCampaignMixinProps{
//   	AdditionalTreatments: []interface{}{
//   		&WriteTreatmentResourceProperty{
//   			CustomDeliveryConfiguration: &CustomDeliveryConfigurationProperty{
//   				DeliveryUri: jsii.String("deliveryUri"),
//   				EndpointTypes: []*string{
//   					jsii.String("endpointTypes"),
//   				},
//   			},
//   			MessageConfiguration: &MessageConfigurationProperty{
//   				AdmMessage: &MessageProperty{
//   					Action: jsii.String("action"),
//   					Body: jsii.String("body"),
//   					ImageIconUrl: jsii.String("imageIconUrl"),
//   					ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					ImageUrl: jsii.String("imageUrl"),
//   					JsonBody: jsii.String("jsonBody"),
//   					MediaUrl: jsii.String("mediaUrl"),
//   					RawContent: jsii.String("rawContent"),
//   					SilentPush: jsii.Boolean(false),
//   					TimeToLive: jsii.Number(123),
//   					Title: jsii.String("title"),
//   					Url: jsii.String("url"),
//   				},
//   				ApnsMessage: &MessageProperty{
//   					Action: jsii.String("action"),
//   					Body: jsii.String("body"),
//   					ImageIconUrl: jsii.String("imageIconUrl"),
//   					ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					ImageUrl: jsii.String("imageUrl"),
//   					JsonBody: jsii.String("jsonBody"),
//   					MediaUrl: jsii.String("mediaUrl"),
//   					RawContent: jsii.String("rawContent"),
//   					SilentPush: jsii.Boolean(false),
//   					TimeToLive: jsii.Number(123),
//   					Title: jsii.String("title"),
//   					Url: jsii.String("url"),
//   				},
//   				BaiduMessage: &MessageProperty{
//   					Action: jsii.String("action"),
//   					Body: jsii.String("body"),
//   					ImageIconUrl: jsii.String("imageIconUrl"),
//   					ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					ImageUrl: jsii.String("imageUrl"),
//   					JsonBody: jsii.String("jsonBody"),
//   					MediaUrl: jsii.String("mediaUrl"),
//   					RawContent: jsii.String("rawContent"),
//   					SilentPush: jsii.Boolean(false),
//   					TimeToLive: jsii.Number(123),
//   					Title: jsii.String("title"),
//   					Url: jsii.String("url"),
//   				},
//   				CustomMessage: &CampaignCustomMessageProperty{
//   					Data: jsii.String("data"),
//   				},
//   				DefaultMessage: &MessageProperty{
//   					Action: jsii.String("action"),
//   					Body: jsii.String("body"),
//   					ImageIconUrl: jsii.String("imageIconUrl"),
//   					ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					ImageUrl: jsii.String("imageUrl"),
//   					JsonBody: jsii.String("jsonBody"),
//   					MediaUrl: jsii.String("mediaUrl"),
//   					RawContent: jsii.String("rawContent"),
//   					SilentPush: jsii.Boolean(false),
//   					TimeToLive: jsii.Number(123),
//   					Title: jsii.String("title"),
//   					Url: jsii.String("url"),
//   				},
//   				EmailMessage: &CampaignEmailMessageProperty{
//   					Body: jsii.String("body"),
//   					FromAddress: jsii.String("fromAddress"),
//   					HtmlBody: jsii.String("htmlBody"),
//   					Title: jsii.String("title"),
//   				},
//   				GcmMessage: &MessageProperty{
//   					Action: jsii.String("action"),
//   					Body: jsii.String("body"),
//   					ImageIconUrl: jsii.String("imageIconUrl"),
//   					ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					ImageUrl: jsii.String("imageUrl"),
//   					JsonBody: jsii.String("jsonBody"),
//   					MediaUrl: jsii.String("mediaUrl"),
//   					RawContent: jsii.String("rawContent"),
//   					SilentPush: jsii.Boolean(false),
//   					TimeToLive: jsii.Number(123),
//   					Title: jsii.String("title"),
//   					Url: jsii.String("url"),
//   				},
//   				InAppMessage: &CampaignInAppMessageProperty{
//   					Content: []interface{}{
//   						&InAppMessageContentProperty{
//   							BackgroundColor: jsii.String("backgroundColor"),
//   							BodyConfig: &InAppMessageBodyConfigProperty{
//   								Alignment: jsii.String("alignment"),
//   								Body: jsii.String("body"),
//   								TextColor: jsii.String("textColor"),
//   							},
//   							HeaderConfig: &InAppMessageHeaderConfigProperty{
//   								Alignment: jsii.String("alignment"),
//   								Header: jsii.String("header"),
//   								TextColor: jsii.String("textColor"),
//   							},
//   							ImageUrl: jsii.String("imageUrl"),
//   							PrimaryBtn: &InAppMessageButtonProperty{
//   								Android: &OverrideButtonConfigurationProperty{
//   									ButtonAction: jsii.String("buttonAction"),
//   									Link: jsii.String("link"),
//   								},
//   								DefaultConfig: &DefaultButtonConfigurationProperty{
//   									BackgroundColor: jsii.String("backgroundColor"),
//   									BorderRadius: jsii.Number(123),
//   									ButtonAction: jsii.String("buttonAction"),
//   									Link: jsii.String("link"),
//   									Text: jsii.String("text"),
//   									TextColor: jsii.String("textColor"),
//   								},
//   								Ios: &OverrideButtonConfigurationProperty{
//   									ButtonAction: jsii.String("buttonAction"),
//   									Link: jsii.String("link"),
//   								},
//   								Web: &OverrideButtonConfigurationProperty{
//   									ButtonAction: jsii.String("buttonAction"),
//   									Link: jsii.String("link"),
//   								},
//   							},
//   							SecondaryBtn: &InAppMessageButtonProperty{
//   								Android: &OverrideButtonConfigurationProperty{
//   									ButtonAction: jsii.String("buttonAction"),
//   									Link: jsii.String("link"),
//   								},
//   								DefaultConfig: &DefaultButtonConfigurationProperty{
//   									BackgroundColor: jsii.String("backgroundColor"),
//   									BorderRadius: jsii.Number(123),
//   									ButtonAction: jsii.String("buttonAction"),
//   									Link: jsii.String("link"),
//   									Text: jsii.String("text"),
//   									TextColor: jsii.String("textColor"),
//   								},
//   								Ios: &OverrideButtonConfigurationProperty{
//   									ButtonAction: jsii.String("buttonAction"),
//   									Link: jsii.String("link"),
//   								},
//   								Web: &OverrideButtonConfigurationProperty{
//   									ButtonAction: jsii.String("buttonAction"),
//   									Link: jsii.String("link"),
//   								},
//   							},
//   						},
//   					},
//   					CustomConfig: customConfig,
//   					Layout: jsii.String("layout"),
//   				},
//   				SmsMessage: &CampaignSmsMessageProperty{
//   					Body: jsii.String("body"),
//   					EntityId: jsii.String("entityId"),
//   					MessageType: jsii.String("messageType"),
//   					OriginationNumber: jsii.String("originationNumber"),
//   					SenderId: jsii.String("senderId"),
//   					TemplateId: jsii.String("templateId"),
//   				},
//   			},
//   			Schedule: &ScheduleProperty{
//   				EndTime: jsii.String("endTime"),
//   				EventFilter: &CampaignEventFilterProperty{
//   					Dimensions: &EventDimensionsProperty{
//   						Attributes: attributes,
//   						EventType: &SetDimensionProperty{
//   							DimensionType: jsii.String("dimensionType"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						Metrics: metrics,
//   					},
//   					FilterType: jsii.String("filterType"),
//   				},
//   				Frequency: jsii.String("frequency"),
//   				IsLocalTime: jsii.Boolean(false),
//   				QuietTime: &QuietTimeProperty{
//   					End: jsii.String("end"),
//   					Start: jsii.String("start"),
//   				},
//   				StartTime: jsii.String("startTime"),
//   				TimeZone: jsii.String("timeZone"),
//   			},
//   			SizePercent: jsii.Number(123),
//   			TemplateConfiguration: &TemplateConfigurationProperty{
//   				EmailTemplate: &TemplateProperty{
//   					Name: jsii.String("name"),
//   					Version: jsii.String("version"),
//   				},
//   				PushTemplate: &TemplateProperty{
//   					Name: jsii.String("name"),
//   					Version: jsii.String("version"),
//   				},
//   				SmsTemplate: &TemplateProperty{
//   					Name: jsii.String("name"),
//   					Version: jsii.String("version"),
//   				},
//   				VoiceTemplate: &TemplateProperty{
//   					Name: jsii.String("name"),
//   					Version: jsii.String("version"),
//   				},
//   			},
//   			TreatmentDescription: jsii.String("treatmentDescription"),
//   			TreatmentName: jsii.String("treatmentName"),
//   		},
//   	},
//   	ApplicationId: jsii.String("applicationId"),
//   	CampaignHook: &CampaignHookProperty{
//   		LambdaFunctionName: jsii.String("lambdaFunctionName"),
//   		Mode: jsii.String("mode"),
//   		WebUrl: jsii.String("webUrl"),
//   	},
//   	CustomDeliveryConfiguration: &CustomDeliveryConfigurationProperty{
//   		DeliveryUri: jsii.String("deliveryUri"),
//   		EndpointTypes: []*string{
//   			jsii.String("endpointTypes"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	HoldoutPercent: jsii.Number(123),
//   	IsPaused: jsii.Boolean(false),
//   	Limits: &LimitsProperty{
//   		Daily: jsii.Number(123),
//   		MaximumDuration: jsii.Number(123),
//   		MessagesPerSecond: jsii.Number(123),
//   		Session: jsii.Number(123),
//   		Total: jsii.Number(123),
//   	},
//   	MessageConfiguration: &MessageConfigurationProperty{
//   		AdmMessage: &MessageProperty{
//   			Action: jsii.String("action"),
//   			Body: jsii.String("body"),
//   			ImageIconUrl: jsii.String("imageIconUrl"),
//   			ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			ImageUrl: jsii.String("imageUrl"),
//   			JsonBody: jsii.String("jsonBody"),
//   			MediaUrl: jsii.String("mediaUrl"),
//   			RawContent: jsii.String("rawContent"),
//   			SilentPush: jsii.Boolean(false),
//   			TimeToLive: jsii.Number(123),
//   			Title: jsii.String("title"),
//   			Url: jsii.String("url"),
//   		},
//   		ApnsMessage: &MessageProperty{
//   			Action: jsii.String("action"),
//   			Body: jsii.String("body"),
//   			ImageIconUrl: jsii.String("imageIconUrl"),
//   			ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			ImageUrl: jsii.String("imageUrl"),
//   			JsonBody: jsii.String("jsonBody"),
//   			MediaUrl: jsii.String("mediaUrl"),
//   			RawContent: jsii.String("rawContent"),
//   			SilentPush: jsii.Boolean(false),
//   			TimeToLive: jsii.Number(123),
//   			Title: jsii.String("title"),
//   			Url: jsii.String("url"),
//   		},
//   		BaiduMessage: &MessageProperty{
//   			Action: jsii.String("action"),
//   			Body: jsii.String("body"),
//   			ImageIconUrl: jsii.String("imageIconUrl"),
//   			ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			ImageUrl: jsii.String("imageUrl"),
//   			JsonBody: jsii.String("jsonBody"),
//   			MediaUrl: jsii.String("mediaUrl"),
//   			RawContent: jsii.String("rawContent"),
//   			SilentPush: jsii.Boolean(false),
//   			TimeToLive: jsii.Number(123),
//   			Title: jsii.String("title"),
//   			Url: jsii.String("url"),
//   		},
//   		CustomMessage: &CampaignCustomMessageProperty{
//   			Data: jsii.String("data"),
//   		},
//   		DefaultMessage: &MessageProperty{
//   			Action: jsii.String("action"),
//   			Body: jsii.String("body"),
//   			ImageIconUrl: jsii.String("imageIconUrl"),
//   			ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			ImageUrl: jsii.String("imageUrl"),
//   			JsonBody: jsii.String("jsonBody"),
//   			MediaUrl: jsii.String("mediaUrl"),
//   			RawContent: jsii.String("rawContent"),
//   			SilentPush: jsii.Boolean(false),
//   			TimeToLive: jsii.Number(123),
//   			Title: jsii.String("title"),
//   			Url: jsii.String("url"),
//   		},
//   		EmailMessage: &CampaignEmailMessageProperty{
//   			Body: jsii.String("body"),
//   			FromAddress: jsii.String("fromAddress"),
//   			HtmlBody: jsii.String("htmlBody"),
//   			Title: jsii.String("title"),
//   		},
//   		GcmMessage: &MessageProperty{
//   			Action: jsii.String("action"),
//   			Body: jsii.String("body"),
//   			ImageIconUrl: jsii.String("imageIconUrl"),
//   			ImageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			ImageUrl: jsii.String("imageUrl"),
//   			JsonBody: jsii.String("jsonBody"),
//   			MediaUrl: jsii.String("mediaUrl"),
//   			RawContent: jsii.String("rawContent"),
//   			SilentPush: jsii.Boolean(false),
//   			TimeToLive: jsii.Number(123),
//   			Title: jsii.String("title"),
//   			Url: jsii.String("url"),
//   		},
//   		InAppMessage: &CampaignInAppMessageProperty{
//   			Content: []interface{}{
//   				&InAppMessageContentProperty{
//   					BackgroundColor: jsii.String("backgroundColor"),
//   					BodyConfig: &InAppMessageBodyConfigProperty{
//   						Alignment: jsii.String("alignment"),
//   						Body: jsii.String("body"),
//   						TextColor: jsii.String("textColor"),
//   					},
//   					HeaderConfig: &InAppMessageHeaderConfigProperty{
//   						Alignment: jsii.String("alignment"),
//   						Header: jsii.String("header"),
//   						TextColor: jsii.String("textColor"),
//   					},
//   					ImageUrl: jsii.String("imageUrl"),
//   					PrimaryBtn: &InAppMessageButtonProperty{
//   						Android: &OverrideButtonConfigurationProperty{
//   							ButtonAction: jsii.String("buttonAction"),
//   							Link: jsii.String("link"),
//   						},
//   						DefaultConfig: &DefaultButtonConfigurationProperty{
//   							BackgroundColor: jsii.String("backgroundColor"),
//   							BorderRadius: jsii.Number(123),
//   							ButtonAction: jsii.String("buttonAction"),
//   							Link: jsii.String("link"),
//   							Text: jsii.String("text"),
//   							TextColor: jsii.String("textColor"),
//   						},
//   						Ios: &OverrideButtonConfigurationProperty{
//   							ButtonAction: jsii.String("buttonAction"),
//   							Link: jsii.String("link"),
//   						},
//   						Web: &OverrideButtonConfigurationProperty{
//   							ButtonAction: jsii.String("buttonAction"),
//   							Link: jsii.String("link"),
//   						},
//   					},
//   					SecondaryBtn: &InAppMessageButtonProperty{
//   						Android: &OverrideButtonConfigurationProperty{
//   							ButtonAction: jsii.String("buttonAction"),
//   							Link: jsii.String("link"),
//   						},
//   						DefaultConfig: &DefaultButtonConfigurationProperty{
//   							BackgroundColor: jsii.String("backgroundColor"),
//   							BorderRadius: jsii.Number(123),
//   							ButtonAction: jsii.String("buttonAction"),
//   							Link: jsii.String("link"),
//   							Text: jsii.String("text"),
//   							TextColor: jsii.String("textColor"),
//   						},
//   						Ios: &OverrideButtonConfigurationProperty{
//   							ButtonAction: jsii.String("buttonAction"),
//   							Link: jsii.String("link"),
//   						},
//   						Web: &OverrideButtonConfigurationProperty{
//   							ButtonAction: jsii.String("buttonAction"),
//   							Link: jsii.String("link"),
//   						},
//   					},
//   				},
//   			},
//   			CustomConfig: customConfig,
//   			Layout: jsii.String("layout"),
//   		},
//   		SmsMessage: &CampaignSmsMessageProperty{
//   			Body: jsii.String("body"),
//   			EntityId: jsii.String("entityId"),
//   			MessageType: jsii.String("messageType"),
//   			OriginationNumber: jsii.String("originationNumber"),
//   			SenderId: jsii.String("senderId"),
//   			TemplateId: jsii.String("templateId"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   	Schedule: &ScheduleProperty{
//   		EndTime: jsii.String("endTime"),
//   		EventFilter: &CampaignEventFilterProperty{
//   			Dimensions: &EventDimensionsProperty{
//   				Attributes: attributes,
//   				EventType: &SetDimensionProperty{
//   					DimensionType: jsii.String("dimensionType"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				Metrics: metrics,
//   			},
//   			FilterType: jsii.String("filterType"),
//   		},
//   		Frequency: jsii.String("frequency"),
//   		IsLocalTime: jsii.Boolean(false),
//   		QuietTime: &QuietTimeProperty{
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   		StartTime: jsii.String("startTime"),
//   		TimeZone: jsii.String("timeZone"),
//   	},
//   	SegmentId: jsii.String("segmentId"),
//   	SegmentVersion: jsii.Number(123),
//   	Tags: tags,
//   	TemplateConfiguration: &TemplateConfigurationProperty{
//   		EmailTemplate: &TemplateProperty{
//   			Name: jsii.String("name"),
//   			Version: jsii.String("version"),
//   		},
//   		PushTemplate: &TemplateProperty{
//   			Name: jsii.String("name"),
//   			Version: jsii.String("version"),
//   		},
//   		SmsTemplate: &TemplateProperty{
//   			Name: jsii.String("name"),
//   			Version: jsii.String("version"),
//   		},
//   		VoiceTemplate: &TemplateProperty{
//   			Name: jsii.String("name"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	TreatmentDescription: jsii.String("treatmentDescription"),
//   	TreatmentName: jsii.String("treatmentName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html
//
type CfnCampaignPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCampaignMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCampaignPropsMixin
type jsiiProxy_CfnCampaignPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCampaignPropsMixin) Props() *CfnCampaignMixinProps {
	var returns *CfnCampaignMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaignPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Pinpoint::Campaign`.
func NewCfnCampaignPropsMixin(props *CfnCampaignMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCampaignPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCampaignPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCampaignPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnCampaignPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Pinpoint::Campaign`.
func NewCfnCampaignPropsMixin_Override(c CfnCampaignPropsMixin, props *CfnCampaignMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnCampaignPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCampaignPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaignPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnCampaignPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCampaignPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnCampaignPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCampaignPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCampaignPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

