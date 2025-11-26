package previewawsworkspacesinstancesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   spotMarketOptionsProperty := &SpotMarketOptionsProperty{
//   	InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   	MaxPrice: jsii.String("maxPrice"),
//   	SpotInstanceType: jsii.String("spotInstanceType"),
//   	ValidUntilUtc: jsii.String("validUntilUtc"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-spotmarketoptions.html
//
type CfnWorkspaceInstancePropsMixin_SpotMarketOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-spotmarketoptions.html#cfn-workspacesinstances-workspaceinstance-spotmarketoptions-instanceinterruptionbehavior
	//
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-spotmarketoptions.html#cfn-workspacesinstances-workspaceinstance-spotmarketoptions-maxprice
	//
	MaxPrice *string `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-spotmarketoptions.html#cfn-workspacesinstances-workspaceinstance-spotmarketoptions-spotinstancetype
	//
	SpotInstanceType *string `field:"optional" json:"spotInstanceType" yaml:"spotInstanceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-spotmarketoptions.html#cfn-workspacesinstances-workspaceinstance-spotmarketoptions-validuntilutc
	//
	ValidUntilUtc *string `field:"optional" json:"validUntilUtc" yaml:"validUntilUtc"`
}

