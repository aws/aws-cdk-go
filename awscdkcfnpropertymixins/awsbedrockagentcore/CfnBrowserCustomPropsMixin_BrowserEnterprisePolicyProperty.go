package awsbedrockagentcore


// Browser enterprise policy configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   browserEnterprisePolicyProperty := &BrowserEnterprisePolicyProperty{
//   	Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-browserenterprisepolicy.html
//
type CfnBrowserCustomPropsMixin_BrowserEnterprisePolicyProperty struct {
	// S3 Location Configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-browserenterprisepolicy.html#cfn-bedrockagentcore-browsercustom-browserenterprisepolicy-location
	//
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// The type of browser enterprise policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-browserenterprisepolicy.html#cfn-bedrockagentcore-browsercustom-browserenterprisepolicy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

