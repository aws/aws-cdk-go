package previewawsiotanalyticsevents


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseElements := &ResponseElements{
//   	ChannelArn: []*string{
//   		jsii.String("channelArn"),
//   	},
//   	ChannelName: []*string{
//   		jsii.String("channelName"),
//   	},
//   	DatasetArn: []*string{
//   		jsii.String("datasetArn"),
//   	},
//   	DatasetName: []*string{
//   		jsii.String("datasetName"),
//   	},
//   	DatastoreArn: []*string{
//   		jsii.String("datastoreArn"),
//   	},
//   	DatastoreIndexArn: []*string{
//   		jsii.String("datastoreIndexArn"),
//   	},
//   	DatastoreIndexName: []*string{
//   		jsii.String("datastoreIndexName"),
//   	},
//   	DatastoreName: []*string{
//   		jsii.String("datastoreName"),
//   	},
//   	PipelineArn: []*string{
//   		jsii.String("pipelineArn"),
//   	},
//   	PipelineName: []*string{
//   		jsii.String("pipelineName"),
//   	},
//   	ReprocessingId: []*string{
//   		jsii.String("reprocessingId"),
//   	},
//   	RetentionPeriod: &RetentionPeriod1{
//   		NumberOfDays: []*string{
//   			jsii.String("numberOfDays"),
//   		},
//   		Unlimited: []*string{
//   			jsii.String("unlimited"),
//   		},
//   	},
//   	VersionId: []*string{
//   		jsii.String("versionId"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ResponseElements struct {
	// channelArn property.
	//
	// Specify an array of string values to match this event if the actual value of channelArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChannelArn *[]*string `field:"optional" json:"channelArn" yaml:"channelArn"`
	// channelName property.
	//
	// Specify an array of string values to match this event if the actual value of channelName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChannelName *[]*string `field:"optional" json:"channelName" yaml:"channelName"`
	// datasetArn property.
	//
	// Specify an array of string values to match this event if the actual value of datasetArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatasetArn *[]*string `field:"optional" json:"datasetArn" yaml:"datasetArn"`
	// datasetName property.
	//
	// Specify an array of string values to match this event if the actual value of datasetName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatasetName *[]*string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// datastoreArn property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatastoreArn *[]*string `field:"optional" json:"datastoreArn" yaml:"datastoreArn"`
	// datastoreIndexArn property.
	//
	// Specify an array of string values to match this event if the actual value of datastoreIndexArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DatastoreIndexArn *[]*string `field:"optional" json:"datastoreIndexArn" yaml:"datastoreIndexArn"`
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
	// pipelineArn property.
	//
	// Specify an array of string values to match this event if the actual value of pipelineArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PipelineArn *[]*string `field:"optional" json:"pipelineArn" yaml:"pipelineArn"`
	// pipelineName property.
	//
	// Specify an array of string values to match this event if the actual value of pipelineName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PipelineName *[]*string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// reprocessingId property.
	//
	// Specify an array of string values to match this event if the actual value of reprocessingId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReprocessingId *[]*string `field:"optional" json:"reprocessingId" yaml:"reprocessingId"`
	// retentionPeriod property.
	//
	// Specify an array of string values to match this event if the actual value of retentionPeriod is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RetentionPeriod *AWSAPICallViaCloudTrail_RetentionPeriod1 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// versionId property.
	//
	// Specify an array of string values to match this event if the actual value of versionId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VersionId *[]*string `field:"optional" json:"versionId" yaml:"versionId"`
}

