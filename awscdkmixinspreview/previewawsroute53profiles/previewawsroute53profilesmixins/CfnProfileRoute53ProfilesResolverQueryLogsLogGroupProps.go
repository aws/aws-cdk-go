package previewawsroute53profilesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProfileRoute53ProfilesResolverQueryLogsLogGroupProps := &CfnProfileRoute53ProfilesResolverQueryLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnProfileRoute53ProfilesResolverQueryLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

