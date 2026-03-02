package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCRoute53ResolverQueryLogsLogGroupProps := &CfnVPCRoute53ResolverQueryLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnVPCRoute53ResolverQueryLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnVPCRoute53ResolverQueryLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnVPCRoute53ResolverQueryLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

