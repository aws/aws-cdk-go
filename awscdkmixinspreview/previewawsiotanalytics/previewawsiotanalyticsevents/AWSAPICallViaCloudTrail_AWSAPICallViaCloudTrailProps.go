package previewawsiotanalyticsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.iotanalytics@AWSAPICallViaCloudTrail event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrailProps := &AWSAPICallViaCloudTrailProps{
//   	AwsRegion: []*string{
//   		jsii.String("awsRegion"),
//   	},
//   	ErrorCode: []*string{
//   		jsii.String("errorCode"),
//   	},
//   	ErrorMessage: []*string{
//   		jsii.String("errorMessage"),
//   	},
//   	EventId: []*string{
//   		jsii.String("eventId"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	EventName: []*string{
//   		jsii.String("eventName"),
//   	},
//   	EventSource: []*string{
//   		jsii.String("eventSource"),
//   	},
//   	EventTime: []*string{
//   		jsii.String("eventTime"),
//   	},
//   	EventType: []*string{
//   		jsii.String("eventType"),
//   	},
//   	EventVersion: []*string{
//   		jsii.String("eventVersion"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	RequestParameters: &RequestParameters{
//   		ChannelName: []*string{
//   			jsii.String("channelName"),
//   		},
//   		ChannelStorage: &ChannelStorage{
//   			CustomerManagedS3: &CustomerManagedS3{
//   				Bucket: []*string{
//   					jsii.String("bucket"),
//   				},
//   				KeyPrefix: []*string{
//   					jsii.String("keyPrefix"),
//   				},
//   				RoleArn: []*string{
//   					jsii.String("roleArn"),
//   				},
//   			},
//   		},
//   		DatasetName: []*string{
//   			jsii.String("datasetName"),
//   		},
//   		DatastoreIndexName: []*string{
//   			jsii.String("datastoreIndexName"),
//   		},
//   		DatastoreName: []*string{
//   			jsii.String("datastoreName"),
//   		},
//   		DatastoreStorage: &ChannelStorage{
//   			CustomerManagedS3: &CustomerManagedS3{
//   				Bucket: []*string{
//   					jsii.String("bucket"),
//   				},
//   				KeyPrefix: []*string{
//   					jsii.String("keyPrefix"),
//   				},
//   				RoleArn: []*string{
//   					jsii.String("roleArn"),
//   				},
//   			},
//   		},
//   		EndTime: []*string{
//   			jsii.String("endTime"),
//   		},
//   		LoggingOptions: &LoggingOptions{
//   			Enabled: []*string{
//   				jsii.String("enabled"),
//   			},
//   		},
//   		MaxMessages: []*string{
//   			jsii.String("maxMessages"),
//   		},
//   		Payloads: []RequestParametersItem{
//   			&RequestParametersItem{
//   				Address: []*string{
//   					jsii.String("address"),
//   				},
//   				BigEndian: []*string{
//   					jsii.String("bigEndian"),
//   				},
//   				Capacity: []*string{
//   					jsii.String("capacity"),
//   				},
//   				Hb: []*f64{
//   					jsii.Number(123),
//   				},
//   				IsReadOnly: []*string{
//   					jsii.String("isReadOnly"),
//   				},
//   				Limit: []*string{
//   					jsii.String("limit"),
//   				},
//   				Mark: []*string{
//   					jsii.String("mark"),
//   				},
//   				NativeByteOrder: []*string{
//   					jsii.String("nativeByteOrder"),
//   				},
//   				Offset: []*string{
//   					jsii.String("offset"),
//   				},
//   				Position: []*string{
//   					jsii.String("position"),
//   				},
//   			},
//   		},
//   		PipelineActivity: &PipelineActivity{
//   			Math: &MathType{
//   				Attribute: []*string{
//   					jsii.String("attribute"),
//   				},
//   				Math: []*string{
//   					jsii.String("math"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   			},
//   		},
//   		PipelineName: []*string{
//   			jsii.String("pipelineName"),
//   		},
//   		QueryAction: &QueryAction{
//   			SqlQuery: []*string{
//   				jsii.String("sqlQuery"),
//   			},
//   		},
//   		ReprocessingId: []*string{
//   			jsii.String("reprocessingId"),
//   		},
//   		ResourceArn: []*string{
//   			jsii.String("resourceArn"),
//   		},
//   		RetentionPeriod: &RetentionPeriod{
//   			NumberOfDays: []*string{
//   				jsii.String("numberOfDays"),
//   			},
//   			Unlimited: []*string{
//   				jsii.String("unlimited"),
//   			},
//   		},
//   		Settings: &Settings{
//   			Timeseries: &Timeseries{
//   				DimensionAttributes: []TimeseriesItem{
//   					&TimeseriesItem{
//   						Attribute: []*string{
//   							jsii.String("attribute"),
//   						},
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   					},
//   				},
//   				MeasureAttributes: []TimeseriesItem{
//   					&TimeseriesItem{
//   						Attribute: []*string{
//   							jsii.String("attribute"),
//   						},
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   					},
//   				},
//   				TimestampAttribute: []*string{
//   					jsii.String("timestampAttribute"),
//   				},
//   			},
//   		},
//   		StartTime: []*string{
//   			jsii.String("startTime"),
//   		},
//   		TagKeys: []*string{
//   			jsii.String("tagKeys"),
//   		},
//   		Tags: []RequestParametersItem1{
//   			&RequestParametersItem1{
//   				Key: []*string{
//   					jsii.String("key"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   		VersionId: []*string{
//   			jsii.String("versionId"),
//   		},
//   	},
//   	ResponseElements: &ResponseElements{
//   		ChannelArn: []*string{
//   			jsii.String("channelArn"),
//   		},
//   		ChannelName: []*string{
//   			jsii.String("channelName"),
//   		},
//   		DatasetArn: []*string{
//   			jsii.String("datasetArn"),
//   		},
//   		DatasetName: []*string{
//   			jsii.String("datasetName"),
//   		},
//   		DatastoreArn: []*string{
//   			jsii.String("datastoreArn"),
//   		},
//   		DatastoreIndexArn: []*string{
//   			jsii.String("datastoreIndexArn"),
//   		},
//   		DatastoreIndexName: []*string{
//   			jsii.String("datastoreIndexName"),
//   		},
//   		DatastoreName: []*string{
//   			jsii.String("datastoreName"),
//   		},
//   		PipelineArn: []*string{
//   			jsii.String("pipelineArn"),
//   		},
//   		PipelineName: []*string{
//   			jsii.String("pipelineName"),
//   		},
//   		ReprocessingId: []*string{
//   			jsii.String("reprocessingId"),
//   		},
//   		RetentionPeriod: &RetentionPeriod1{
//   			NumberOfDays: []*string{
//   				jsii.String("numberOfDays"),
//   			},
//   			Unlimited: []*string{
//   				jsii.String("unlimited"),
//   			},
//   		},
//   		VersionId: []*string{
//   			jsii.String("versionId"),
//   		},
//   	},
//   	SourceIpAddress: []*string{
//   		jsii.String("sourceIpAddress"),
//   	},
//   	UserAgent: []*string{
//   		jsii.String("userAgent"),
//   	},
//   	UserIdentity: &UserIdentity{
//   		AccessKeyId: []*string{
//   			jsii.String("accessKeyId"),
//   		},
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		PrincipalId: []*string{
//   			jsii.String("principalId"),
//   		},
//   		SessionContext: &SessionContext{
//   			Attributes: &Attributes{
//   				CreationDate: []*string{
//   					jsii.String("creationDate"),
//   				},
//   				MfaAuthenticated: []*string{
//   					jsii.String("mfaAuthenticated"),
//   				},
//   			},
//   			SessionIssuer: &SessionIssuer{
//   				AccountId: []*string{
//   					jsii.String("accountId"),
//   				},
//   				Arn: []*string{
//   					jsii.String("arn"),
//   				},
//   				PrincipalId: []*string{
//   					jsii.String("principalId"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   				UserName: []*string{
//   					jsii.String("userName"),
//   				},
//   			},
//   			WebIdFederationData: &WebIdFederationData{
//   				Attributes: map[string]*string{
//   					"attributesKey": jsii.String("attributes"),
//   				},
//   				FederatedProvider: []*string{
//   					jsii.String("federatedProvider"),
//   				},
//   			},
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps struct {
	// awsRegion property.
	//
	// Specify an array of string values to match this event if the actual value of awsRegion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsRegion *[]*string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// errorCode property.
	//
	// Specify an array of string values to match this event if the actual value of errorCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorCode *[]*string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// errorMessage property.
	//
	// Specify an array of string values to match this event if the actual value of errorMessage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorMessage *[]*string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// eventID property.
	//
	// Specify an array of string values to match this event if the actual value of eventID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventId *[]*string `field:"optional" json:"eventId" yaml:"eventId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// eventName property.
	//
	// Specify an array of string values to match this event if the actual value of eventName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventName *[]*string `field:"optional" json:"eventName" yaml:"eventName"`
	// eventSource property.
	//
	// Specify an array of string values to match this event if the actual value of eventSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventSource *[]*string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// eventTime property.
	//
	// Specify an array of string values to match this event if the actual value of eventTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventTime *[]*string `field:"optional" json:"eventTime" yaml:"eventTime"`
	// eventType property.
	//
	// Specify an array of string values to match this event if the actual value of eventType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// eventVersion property.
	//
	// Specify an array of string values to match this event if the actual value of eventVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventVersion *[]*string `field:"optional" json:"eventVersion" yaml:"eventVersion"`
	// requestID property.
	//
	// Specify an array of string values to match this event if the actual value of requestID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// requestParameters property.
	//
	// Specify an array of string values to match this event if the actual value of requestParameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestParameters *AWSAPICallViaCloudTrail_RequestParameters `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// responseElements property.
	//
	// Specify an array of string values to match this event if the actual value of responseElements is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResponseElements *AWSAPICallViaCloudTrail_ResponseElements `field:"optional" json:"responseElements" yaml:"responseElements"`
	// sourceIPAddress property.
	//
	// Specify an array of string values to match this event if the actual value of sourceIPAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIpAddress *[]*string `field:"optional" json:"sourceIpAddress" yaml:"sourceIpAddress"`
	// userAgent property.
	//
	// Specify an array of string values to match this event if the actual value of userAgent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserAgent *[]*string `field:"optional" json:"userAgent" yaml:"userAgent"`
	// userIdentity property.
	//
	// Specify an array of string values to match this event if the actual value of userIdentity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserIdentity *AWSAPICallViaCloudTrail_UserIdentity `field:"optional" json:"userIdentity" yaml:"userIdentity"`
}

