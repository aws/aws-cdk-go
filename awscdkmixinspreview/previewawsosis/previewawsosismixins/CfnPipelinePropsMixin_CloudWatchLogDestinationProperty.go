package previewawsosismixins


// The destination for OpenSearch Ingestion logs sent to Amazon CloudWatch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchLogDestinationProperty := &CloudWatchLogDestinationProperty{
//   	LogGroup: jsii.String("logGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-cloudwatchlogdestination.html
//
type CfnPipelinePropsMixin_CloudWatchLogDestinationProperty struct {
	// The name of the CloudWatch Logs group to send pipeline logs to.
	//
	// You can specify an existing log group or create a new one. For example, `/aws/vendedlogs/OpenSearchService/pipelines` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-cloudwatchlogdestination.html#cfn-osis-pipeline-cloudwatchlogdestination-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

