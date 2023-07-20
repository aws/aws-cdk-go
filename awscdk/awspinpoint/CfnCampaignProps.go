package awspinpoint


// Properties for defining a `CfnCampaign`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var customConfig interface{}
//   var metrics interface{}
//   var tags interface{}
//
//   cfnCampaignProps := &CfnCampaignProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	Name: jsii.String("name"),
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
//
//   	// the properties below are optional
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
//   	Priority: jsii.Number(123),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html
//
type CfnCampaignProps struct {
	// The unique identifier for the Amazon Pinpoint application that the campaign is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The name of the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schedule settings for the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-schedule
	//
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// The unique identifier for the segment to associate with the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-segmentid
	//
	SegmentId *string `field:"required" json:"segmentId" yaml:"segmentId"`
	// An array of requests that defines additional treatments for the campaign, in addition to the default treatment for the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-additionaltreatments
	//
	AdditionalTreatments interface{} `field:"optional" json:"additionalTreatments" yaml:"additionalTreatments"`
	// Specifies the Lambda function to use as a code hook for a campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-campaignhook
	//
	CampaignHook interface{} `field:"optional" json:"campaignHook" yaml:"campaignHook"`
	// The delivery configuration settings for sending the treatment through a custom channel.
	//
	// This object is required if the `MessageConfiguration` object for the treatment specifies a `CustomMessage` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-customdeliveryconfiguration
	//
	CustomDeliveryConfiguration interface{} `field:"optional" json:"customDeliveryConfiguration" yaml:"customDeliveryConfiguration"`
	// A custom description of the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The allocated percentage of users (segment members) who shouldn't receive messages from the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-holdoutpercent
	//
	HoldoutPercent *float64 `field:"optional" json:"holdoutPercent" yaml:"holdoutPercent"`
	// Specifies whether to pause the campaign.
	//
	// A paused campaign doesn't run unless you resume it by changing this value to `false` . If you restart a campaign, the campaign restarts from the beginning and not at the point you paused it. If a campaign is running it will complete and then pause. Pause only pauses or skips the next run for a recurring future scheduled campaign. A campaign scheduled for immediate can't be paused.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-ispaused
	//
	IsPaused interface{} `field:"optional" json:"isPaused" yaml:"isPaused"`
	// The messaging limits for the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-limits
	//
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
	// The message configuration settings for the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-messageconfiguration
	//
	MessageConfiguration interface{} `field:"optional" json:"messageConfiguration" yaml:"messageConfiguration"`
	// An integer between 1 and 5, inclusive, that represents the priority of the in-app message campaign, where 1 is the highest priority and 5 is the lowest.
	//
	// If there are multiple messages scheduled to be displayed at the same time, the priority determines the order in which those messages are displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The version of the segment to associate with the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-segmentversion
	//
	SegmentVersion *float64 `field:"optional" json:"segmentVersion" yaml:"segmentVersion"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The message template to use for the treatment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-templateconfiguration
	//
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// A custom description of the default treatment for the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-treatmentdescription
	//
	TreatmentDescription *string `field:"optional" json:"treatmentDescription" yaml:"treatmentDescription"`
	// A custom name of the default treatment for the campaign, if the campaign has multiple treatments.
	//
	// A *treatment* is a variation of a campaign that's used for A/B testing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html#cfn-pinpoint-campaign-treatmentname
	//
	TreatmentName *string `field:"optional" json:"treatmentName" yaml:"treatmentName"`
}

