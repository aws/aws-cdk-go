package mixinsawsbedrockagentcore


// Configuration for HTTP request headers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestHeaderConfigurationProperty := &RequestHeaderConfigurationProperty{
//   	RequestHeaderAllowlist: []*string{
//   		jsii.String("requestHeaderAllowlist"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-requestheaderconfiguration.html
//
type CfnRuntimePropsMixin_RequestHeaderConfigurationProperty struct {
	// List of allowed HTTP headers for agent runtime requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-requestheaderconfiguration.html#cfn-bedrockagentcore-runtime-requestheaderconfiguration-requestheaderallowlist
	//
	RequestHeaderAllowlist *[]*string `field:"optional" json:"requestHeaderAllowlist" yaml:"requestHeaderAllowlist"`
}

