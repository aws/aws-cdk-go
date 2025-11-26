package previewawsworkspacesinstancesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instanceMarketOptionsRequestProperty := &InstanceMarketOptionsRequestProperty{
//   	MarketType: jsii.String("marketType"),
//   	SpotOptions: &SpotMarketOptionsProperty{
//   		InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   		MaxPrice: jsii.String("maxPrice"),
//   		SpotInstanceType: jsii.String("spotInstanceType"),
//   		ValidUntilUtc: jsii.String("validUntilUtc"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancemarketoptionsrequest.html
//
type CfnWorkspaceInstancePropsMixin_InstanceMarketOptionsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancemarketoptionsrequest.html#cfn-workspacesinstances-workspaceinstance-instancemarketoptionsrequest-markettype
	//
	MarketType *string `field:"optional" json:"marketType" yaml:"marketType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancemarketoptionsrequest.html#cfn-workspacesinstances-workspaceinstance-instancemarketoptionsrequest-spotoptions
	//
	SpotOptions interface{} `field:"optional" json:"spotOptions" yaml:"spotOptions"`
}

