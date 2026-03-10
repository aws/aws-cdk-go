package previewawsiotanalyticsevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	ChannelName: []*string{
//   		jsii.String("channelName"),
//   	},
//   	ChannelStorage: &ChannelStorage{
//   		CustomerManagedS3: &CustomerManagedS3{
//   			Bucket: []*string{
//   				jsii.String("bucket"),
//   			},
//   			KeyPrefix: []*string{
//   				jsii.String("keyPrefix"),
//   			},
//   			RoleArn: []*string{
//   				jsii.String("roleArn"),
//   			},
//   		},
//   	},
//   	DatasetName: []*string{
//   		jsii.String("datasetName"),
//   	},
//   	DatastoreIndexName: []*string{
//   		jsii.String("datastoreIndexName"),
//   	},
//   	DatastoreName: []*string{
//   		jsii.String("datastoreName"),
//   	},
//   	DatastoreStorage: &ChannelStorage{
//   		CustomerManagedS3: &CustomerManagedS3{
//   			Bucket: []*string{
//   				jsii.String("bucket"),
//   			},
//   			KeyPrefix: []*string{
//   				jsii.String("keyPrefix"),
//   			},
//   			RoleArn: []*string{
//   				jsii.String("roleArn"),
//   			},
//   		},
//   	},
//   	EndTime: []*string{
//   		jsii.String("endTime"),
//   	},
//   	LoggingOptions: &LoggingOptions{
//   		Enabled: []*string{
//   			jsii.String("enabled"),
//   		},
//   	},
//   	MaxMessages: []*string{
//   		jsii.String("maxMessages"),
//   	},
//   	Payloads: []RequestParametersItem{
//   		&RequestParametersItem{
//   			Address: []*string{
//   				jsii.String("address"),
//   			},
//   			BigEndian: []*string{
//   				jsii.String("bigEndian"),
//   			},
//   			Capacity: []*string{
//   				jsii.String("capacity"),
//   			},
//   			Hb: []*f64{
//   				jsii.Number(123),
//   			},
//   			IsReadOnly: []*string{
//   				jsii.String("isReadOnly"),
//   			},
//   			Limit: []*string{
//   				jsii.String("limit"),
//   			},
//   			Mark: []*string{
//   				jsii.String("mark"),
//   			},
//   			NativeByteOrder: []*string{
//   				jsii.String("nativeByteOrder"),
//   			},
//   			Offset: []*string{
//   				jsii.String("offset"),
//   			},
//   			Position: []*string{
//   				jsii.String("position"),
//   			},
//   		},
//   	},
//   	PipelineActivity: &PipelineActivity{
//   		Math: &MathType{
//   			Attribute: []*string{
//   				jsii.String("attribute"),
//   			},
//   			Math: []*string{
//   				jsii.String("math"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   		},
//   	},
//   	PipelineName: []*string{
//   		jsii.String("pipelineName"),
//   	},
//   	QueryAction: &QueryAction{
//   		SqlQuery: []*string{
//   			jsii.String("sqlQuery"),
//   		},
//   	},
//   	ReprocessingId: []*string{
//   		jsii.String("reprocessingId"),
//   	},
//   	ResourceArn: []*string{
//   		jsii.String("resourceArn"),
//   	},
//   	RetentionPeriod: &RetentionPeriod{
//   		NumberOfDays: []*string{
//   			jsii.String("numberOfDays"),
//   		},
//   		Unlimited: []*string{
//   			jsii.String("unlimited"),
//   		},
//   	},
//   	Settings: &Settings{
//   		Timeseries: &Timeseries{
//   			DimensionAttributes: []TimeseriesItem{
//   				&TimeseriesItem{
//   					Attribute: []*string{
//   						jsii.String("attribute"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   				},
//   			},
//   			MeasureAttributes: []TimeseriesItem{
//   				&TimeseriesItem{
//   					Attribute: []*string{
//   						jsii.String("attribute"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   				},
//   			},
//   			TimestampAttribute: []*string{
//   				jsii.String("timestampAttribute"),
//   			},
//   		},
//   	},
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   	TagKeys: []*string{
//   		jsii.String("tagKeys"),
//   	},
//   	Tags: []RequestParametersItem1{
//   		&RequestParametersItem1{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	VersionId: []*string{
//   		jsii.String("versionId"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RequestParameters struct {
	// channelName property.
	//
	// Specify an array of string values to match this event if the actual value of channelName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChannelName *[]*string `field:"optional" json:"channelName" yaml:"channelName"`
	// channelStorage property.
	//
	// Specify an array of string values to match this event if the actual value of channelStorage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChannelStorage *AWSAPICallViaCloudTrail_ChannelStorage `field:"optional" json:"channelStorage" yaml:"channelStorage"`
	// datasetName property.
	//
	// Specify an array of string values to match this event if the actual value of datasetName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatasetName *[]*string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// datastoreIndexName property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreIndexName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatastoreIndexName *[]*string `field:"optional" json:"datastoreIndexName" yaml:"datastoreIndexName"`
	// datastoreName property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatastoreName *[]*string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// datastoreStorage property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreStorage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatastoreStorage *AWSAPICallViaCloudTrail_ChannelStorage `field:"optional" json:"datastoreStorage" yaml:"datastoreStorage"`
	// endTime property.
	//
	// Specify an array of string values to match this event if the actual value of endTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTime *[]*string `field:"optional" json:"endTime" yaml:"endTime"`
	// loggingOptions property.
	//
	// Specify an array of string values to match this event if the actual value of loggingOptions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LoggingOptions *AWSAPICallViaCloudTrail_LoggingOptions `field:"optional" json:"loggingOptions" yaml:"loggingOptions"`
	// maxMessages property.
	//
	// Specify an array of string values to match this event if the actual value of maxMessages is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MaxMessages *[]*string `field:"optional" json:"maxMessages" yaml:"maxMessages"`
	// payloads property.
	//
	// Specify an array of string values to match this event if the actual value of payloads is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Payloads *[]*AWSAPICallViaCloudTrail_RequestParametersItem `field:"optional" json:"payloads" yaml:"payloads"`
	// pipelineActivity property.
	//
	// Specify an array of string values to match this event if the actual value of pipelineActivity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PipelineActivity *AWSAPICallViaCloudTrail_PipelineActivity `field:"optional" json:"pipelineActivity" yaml:"pipelineActivity"`
	// pipelineName property.
	//
	// Specify an array of string values to match this event if the actual value of pipelineName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PipelineName *[]*string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// queryAction property.
	//
	// Specify an array of string values to match this event if the actual value of queryAction is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	QueryAction *AWSAPICallViaCloudTrail_QueryAction `field:"optional" json:"queryAction" yaml:"queryAction"`
	// reprocessingId property.
	//
	// Specify an array of string values to match this event if the actual value of reprocessingId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReprocessingId *[]*string `field:"optional" json:"reprocessingId" yaml:"reprocessingId"`
	// resourceArn property.
	//
	// Specify an array of string values to match this event if the actual value of resourceArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceArn *[]*string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// retentionPeriod property.
	//
	// Specify an array of string values to match this event if the actual value of retentionPeriod is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RetentionPeriod *AWSAPICallViaCloudTrail_RetentionPeriod `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// settings property.
	//
	// Specify an array of string values to match this event if the actual value of settings is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Settings *AWSAPICallViaCloudTrail_Settings `field:"optional" json:"settings" yaml:"settings"`
	// startTime property.
	//
	// Specify an array of string values to match this event if the actual value of startTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
	// tagKeys property.
	//
	// Specify an array of string values to match this event if the actual value of tagKeys is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TagKeys *[]*string `field:"optional" json:"tagKeys" yaml:"tagKeys"`
	// tags property.
	//
	// Specify an array of string values to match this event if the actual value of tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*AWSAPICallViaCloudTrail_RequestParametersItem1 `field:"optional" json:"tags" yaml:"tags"`
	// versionId property.
	//
	// Specify an array of string values to match this event if the actual value of versionId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VersionId *[]*string `field:"optional" json:"versionId" yaml:"versionId"`
}

