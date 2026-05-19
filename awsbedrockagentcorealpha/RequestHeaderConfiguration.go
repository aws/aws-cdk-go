package awsbedrockagentcorealpha


// Configuration for HTTP request headers that will be passed through to the runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   requestHeaderConfiguration := &RequestHeaderConfiguration{
//   	AllowlistedHeaders: []*string{
//   		jsii.String("allowlistedHeaders"),
//   	},
//   }
//
// Experimental.
type RequestHeaderConfiguration struct {
	// A list of HTTP request headers that are allowed to be passed through to the runtime.
	// Default: - No request headers allowed.
	//
	// Experimental.
	AllowlistedHeaders *[]*string `field:"optional" json:"allowlistedHeaders" yaml:"allowlistedHeaders"`
}

