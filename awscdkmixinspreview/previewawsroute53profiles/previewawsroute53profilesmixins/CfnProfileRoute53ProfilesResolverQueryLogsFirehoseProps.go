package previewawsroute53profilesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProfileRoute53ProfilesResolverQueryLogsFirehoseProps := &CfnProfileRoute53ProfilesResolverQueryLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnProfileRoute53ProfilesResolverQueryLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

