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
//   cfnCampaignProps := &cfnCampaignProps{
//   	applicationId: jsii.String("applicationId"),
//   	name: jsii.String("name"),
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
//   	segmentId: jsii.String("segmentId"),
//
//   	// the properties below are optional
//   	additionalTreatments: []interface{}{
//   		&writeTreatmentResourceProperty{
//   			customDeliveryConfiguration: &customDeliveryConfigurationProperty{
//   				deliveryUri: jsii.String("deliveryUri"),
//   				endpointTypes: []*string{
//   					jsii.String("endpointTypes"),
//   				},
//   			},
//   			messageConfiguration: &messageConfigurationProperty{
//   				admMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				apnsMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				baiduMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				customMessage: &campaignCustomMessageProperty{
//   					data: jsii.String("data"),
//   				},
//   				defaultMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				emailMessage: &campaignEmailMessageProperty{
//   					body: jsii.String("body"),
//   					fromAddress: jsii.String("fromAddress"),
//   					htmlBody: jsii.String("htmlBody"),
//   					title: jsii.String("title"),
//   				},
//   				gcmMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				inAppMessage: &campaignInAppMessageProperty{
//   					content: []interface{}{
//   						&inAppMessageContentProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							bodyConfig: &inAppMessageBodyConfigProperty{
//   								alignment: jsii.String("alignment"),
//   								body: jsii.String("body"),
//   								textColor: jsii.String("textColor"),
//   							},
//   							headerConfig: &inAppMessageHeaderConfigProperty{
//   								alignment: jsii.String("alignment"),
//   								header: jsii.String("header"),
//   								textColor: jsii.String("textColor"),
//   							},
//   							imageUrl: jsii.String("imageUrl"),
//   							primaryBtn: &inAppMessageButtonProperty{
//   								android: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								defaultConfig: &defaultButtonConfigurationProperty{
//   									backgroundColor: jsii.String("backgroundColor"),
//   									borderRadius: jsii.Number(123),
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   									text: jsii.String("text"),
//   									textColor: jsii.String("textColor"),
//   								},
//   								ios: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								web: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   							},
//   							secondaryBtn: &inAppMessageButtonProperty{
//   								android: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								defaultConfig: &defaultButtonConfigurationProperty{
//   									backgroundColor: jsii.String("backgroundColor"),
//   									borderRadius: jsii.Number(123),
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   									text: jsii.String("text"),
//   									textColor: jsii.String("textColor"),
//   								},
//   								ios: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								web: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   							},
//   						},
//   					},
//   					customConfig: customConfig,
//   					layout: jsii.String("layout"),
//   				},
//   				smsMessage: &campaignSmsMessageProperty{
//   					body: jsii.String("body"),
//   					entityId: jsii.String("entityId"),
//   					messageType: jsii.String("messageType"),
//   					originationNumber: jsii.String("originationNumber"),
//   					senderId: jsii.String("senderId"),
//   					templateId: jsii.String("templateId"),
//   				},
//   			},
//   			schedule: &scheduleProperty{
//   				endTime: jsii.String("endTime"),
//   				eventFilter: &campaignEventFilterProperty{
//   					dimensions: &eventDimensionsProperty{
//   						attributes: attributes,
//   						eventType: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						metrics: metrics,
//   					},
//   					filterType: jsii.String("filterType"),
//   				},
//   				frequency: jsii.String("frequency"),
//   				isLocalTime: jsii.Boolean(false),
//   				quietTime: &quietTimeProperty{
//   					end: jsii.String("end"),
//   					start: jsii.String("start"),
//   				},
//   				startTime: jsii.String("startTime"),
//   				timeZone: jsii.String("timeZone"),
//   			},
//   			sizePercent: jsii.Number(123),
//   			templateConfiguration: &templateConfigurationProperty{
//   				emailTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   				pushTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   				smsTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   				voiceTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   			},
//   			treatmentDescription: jsii.String("treatmentDescription"),
//   			treatmentName: jsii.String("treatmentName"),
//   		},
//   	},
//   	campaignHook: &campaignHookProperty{
//   		lambdaFunctionName: jsii.String("lambdaFunctionName"),
//   		mode: jsii.String("mode"),
//   		webUrl: jsii.String("webUrl"),
//   	},
//   	customDeliveryConfiguration: &customDeliveryConfigurationProperty{
//   		deliveryUri: jsii.String("deliveryUri"),
//   		endpointTypes: []*string{
//   			jsii.String("endpointTypes"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	holdoutPercent: jsii.Number(123),
//   	isPaused: jsii.Boolean(false),
//   	limits: &limitsProperty{
//   		daily: jsii.Number(123),
//   		maximumDuration: jsii.Number(123),
//   		messagesPerSecond: jsii.Number(123),
//   		session: jsii.Number(123),
//   		total: jsii.Number(123),
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
//   	priority: jsii.Number(123),
//   	segmentVersion: jsii.Number(123),
//   	tags: tags,
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
type CfnCampaignProps struct {
	// The unique identifier for the Amazon Pinpoint application that the campaign is associated with.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The name of the campaign.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schedule settings for the campaign.
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// The unique identifier for the segment to associate with the campaign.
	SegmentId *string `field:"required" json:"segmentId" yaml:"segmentId"`
	// An array of requests that defines additional treatments for the campaign, in addition to the default treatment for the campaign.
	AdditionalTreatments interface{} `field:"optional" json:"additionalTreatments" yaml:"additionalTreatments"`
	// Specifies the Lambda function to use as a code hook for a campaign.
	CampaignHook interface{} `field:"optional" json:"campaignHook" yaml:"campaignHook"`
	// The delivery configuration settings for sending the treatment through a custom channel.
	//
	// This object is required if the `MessageConfiguration` object for the treatment specifies a `CustomMessage` object.
	CustomDeliveryConfiguration interface{} `field:"optional" json:"customDeliveryConfiguration" yaml:"customDeliveryConfiguration"`
	// A custom description of the campaign.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The allocated percentage of users (segment members) who shouldn't receive messages from the campaign.
	HoldoutPercent *float64 `field:"optional" json:"holdoutPercent" yaml:"holdoutPercent"`
	// Specifies whether to pause the campaign.
	//
	// A paused campaign doesn't run unless you resume it by changing this value to `false` . If you restart a campaign, the campaign restarts from the beginning and not at the point you paused it.
	IsPaused interface{} `field:"optional" json:"isPaused" yaml:"isPaused"`
	// The messaging limits for the campaign.
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
	// The message configuration settings for the campaign.
	MessageConfiguration interface{} `field:"optional" json:"messageConfiguration" yaml:"messageConfiguration"`
	// An integer between 1 and 5, inclusive, that represents the priority of the in-app message campaign, where 1 is the highest priority and 5 is the lowest.
	//
	// If there are multiple messages scheduled to be displayed at the same time, the priority determines the order in which those messages are displayed.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The version of the segment to associate with the campaign.
	SegmentVersion *float64 `field:"optional" json:"segmentVersion" yaml:"segmentVersion"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The message template to use for the treatment.
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// A custom description of the default treatment for the campaign.
	TreatmentDescription *string `field:"optional" json:"treatmentDescription" yaml:"treatmentDescription"`
	// A custom name of the default treatment for the campaign, if the campaign has multiple treatments.
	//
	// A *treatment* is a variation of a campaign that's used for A/B testing.
	TreatmentName *string `field:"optional" json:"treatmentName" yaml:"treatmentName"`
}

