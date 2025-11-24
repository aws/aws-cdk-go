package mixinsawsosis


// Container for the values required to configure logging for the pipeline.
//
// If you don't specify these values, OpenSearch Ingestion will not publish logs from your application to CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logPublishingOptionsProperty := &LogPublishingOptionsProperty{
//   	CloudWatchLogDestination: &CloudWatchLogDestinationProperty{
//   		LogGroup: jsii.String("logGroup"),
//   	},
//   	IsLoggingEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-logpublishingoptions.html
//
type CfnPipelinePropsMixin_LogPublishingOptionsProperty struct {
	// The destination for OpenSearch Ingestion logs sent to Amazon CloudWatch Logs.
	//
	// This parameter is required if `IsLoggingEnabled` is set to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-logpublishingoptions.html#cfn-osis-pipeline-logpublishingoptions-cloudwatchlogdestination
	//
	CloudWatchLogDestination interface{} `field:"optional" json:"cloudWatchLogDestination" yaml:"cloudWatchLogDestination"`
	// Whether logs should be published.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-logpublishingoptions.html#cfn-osis-pipeline-logpublishingoptions-isloggingenabled
	//
	IsLoggingEnabled interface{} `field:"optional" json:"isLoggingEnabled" yaml:"isLoggingEnabled"`
}

