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
//   writeTreatmentResourceProperty := &writeTreatmentResourceProperty{
//   	customDeliveryConfiguration: &customDeliveryConfigurationProperty{
//   		deliveryUri: jsii.String("deliveryUri"),
//   		endpointTypes: []*string{
//   			jsii.String("endpointTypes"),
//   		},
//   	},
//   	messageConfiguration: &messageConfigurationProperty{
//   		admMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		apnsMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		baiduMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		customMessage: &campaignCustomMessageProperty{
//   			data: jsii.String("data"),
//   		},
//   		defaultMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		emailMessage: &campaignEmailMessageProperty{
//   			body: jsii.String("body"),
//   			fromAddress: jsii.String("fromAddress"),
//   			htmlBody: jsii.String("htmlBody"),
//   			title: jsii.String("title"),
//   		},
//   		gcmMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		inAppMessage: &campaignInAppMessageProperty{
//   			content: []interface{}{
//   				&inAppMessageContentProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					bodyConfig: &inAppMessageBodyConfigProperty{
//   						alignment: jsii.String("alignment"),
//   						body: jsii.String("body"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					headerConfig: &inAppMessageHeaderConfigProperty{
//   						alignment: jsii.String("alignment"),
//   						header: jsii.String("header"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					imageUrl: jsii.String("imageUrl"),
//   					primaryBtn: &inAppMessageButtonProperty{
//   						android: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						defaultConfig: &defaultButtonConfigurationProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							borderRadius: jsii.Number(123),
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   							text: jsii.String("text"),
//   							textColor: jsii.String("textColor"),
//   						},
//   						ios: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						web: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   					},
//   					secondaryBtn: &inAppMessageButtonProperty{
//   						android: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						defaultConfig: &defaultButtonConfigurationProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							borderRadius: jsii.Number(123),
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   							text: jsii.String("text"),
//   							textColor: jsii.String("textColor"),
//   						},
//   						ios: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						web: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   					},
//   				},
//   			},
//   			customConfig: customConfig,
//   			layout: jsii.String("layout"),
//   		},
//   		smsMessage: &campaignSmsMessageProperty{
//   			body: jsii.String("body"),
//   			entityId: jsii.String("entityId"),
//   			messageType: jsii.String("messageType"),
//   			originationNumber: jsii.String("originationNumber"),
//   			senderId: jsii.String("senderId"),
//   			templateId: jsii.String("templateId"),
//   		},
//   	},
//   	schedule: &scheduleProperty{
//   		endTime: jsii.String("endTime"),
//   		eventFilter: &campaignEventFilterProperty{
//   			dimensions: &eventDimensionsProperty{
//   				attributes: attributes,
//   				eventType: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				metrics: metrics,
//   			},
//   			filterType: jsii.String("filterType"),
//   		},
//   		frequency: jsii.String("frequency"),
//   		isLocalTime: jsii.Boolean(false),
//   		quietTime: &quietTimeProperty{
//   			end: jsii.String("end"),
//   			start: jsii.String("start"),
//   		},
//   		startTime: jsii.String("startTime"),
//   		timeZone: jsii.String("timeZone"),
//   	},
//   	sizePercent: jsii.Number(123),
//   	templateConfiguration: &templateConfigurationProperty{
//   		emailTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		pushTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		smsTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		voiceTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   	},
//   	treatmentDescription: jsii.String("treatmentDescription"),
//   	treatmentName: jsii.String("treatmentName"),
//   }
//
type CfnCampaign_WriteTreatmentResourceProperty struct {
	// The delivery configuration settings for sending the treatment through a custom channel.
	//
	// This object is required if the `MessageConfiguration` object for the treatment specifies a `CustomMessage` object.
	CustomDeliveryConfiguration interface{} `field:"optional" json:"customDeliveryConfiguration" yaml:"customDeliveryConfiguration"`
	// The message configuration settings for the treatment.
	MessageConfiguration interface{} `field:"optional" json:"messageConfiguration" yaml:"messageConfiguration"`
	// The schedule settings for the treatment.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// The allocated percentage of users (segment members) to send the treatment to.
	SizePercent *float64 `field:"optional" json:"sizePercent" yaml:"sizePercent"`
	// The message template to use for the treatment.
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// A custom description of the treatment.
	TreatmentDescription *string `field:"optional" json:"treatmentDescription" yaml:"treatmentDescription"`
	// A custom name for the treatment.
	TreatmentName *string `field:"optional" json:"treatmentName" yaml:"treatmentName"`
}

