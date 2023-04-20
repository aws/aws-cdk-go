package awspinpoint


// Specifies the settings for a campaign treatment.
//
// A *treatment* is a variation of a campaign that's used for A/B testing of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var customConfig interface{}
//   var metrics interface{}
//
//   writeTreatmentResourceProperty := &WriteTreatmentResourceProperty{
//   	CustomDeliveryConfiguration: &CustomDeliveryConfigurationProperty{
//   		DeliveryUri: jsii.String("deliveryUri"),
//   		EndpointTypes: []*string{
//   			jsii.String("endpointTypes"),
//   		},
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
//   	SizePercent: jsii.Number(123),
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
type CfnCampaign_WriteTreatmentResourceProperty struct {
	// `CfnCampaign.WriteTreatmentResourceProperty.CustomDeliveryConfiguration`.
	CustomDeliveryConfiguration interface{} `field:"optional" json:"customDeliveryConfiguration" yaml:"customDeliveryConfiguration"`
	// The message configuration settings for the treatment.
	MessageConfiguration interface{} `field:"optional" json:"messageConfiguration" yaml:"messageConfiguration"`
	// The schedule settings for the treatment.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// The allocated percentage of users (segment members) to send the treatment to.
	SizePercent *float64 `field:"optional" json:"sizePercent" yaml:"sizePercent"`
	// `CfnCampaign.WriteTreatmentResourceProperty.TemplateConfiguration`.
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// A custom description of the treatment.
	TreatmentDescription *string `field:"optional" json:"treatmentDescription" yaml:"treatmentDescription"`
	// A custom name for the treatment.
	TreatmentName *string `field:"optional" json:"treatmentName" yaml:"treatmentName"`
}

