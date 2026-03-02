package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCRoute53ResolverQueryLogsFirehoseProps := &CfnVPCRoute53ResolverQueryLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnVPCRoute53ResolverQueryLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnVPCRoute53ResolverQueryLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnVPCRoute53ResolverQueryLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

