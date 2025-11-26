package previewawsrtbfabricmixins


// Describes the parameters of a no bid module.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   noBidModuleParametersProperty := &NoBidModuleParametersProperty{
//   	PassThroughPercentage: jsii.Number(123),
//   	Reason: jsii.String("reason"),
//   	ReasonCode: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-nobidmoduleparameters.html
//
type CfnLinkPropsMixin_NoBidModuleParametersProperty struct {
	// The pass through percentage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-nobidmoduleparameters.html#cfn-rtbfabric-link-nobidmoduleparameters-passthroughpercentage
	//
	PassThroughPercentage *float64 `field:"optional" json:"passThroughPercentage" yaml:"passThroughPercentage"`
	// The reason description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-nobidmoduleparameters.html#cfn-rtbfabric-link-nobidmoduleparameters-reason
	//
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// The reason code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-nobidmoduleparameters.html#cfn-rtbfabric-link-nobidmoduleparameters-reasoncode
	//
	ReasonCode *float64 `field:"optional" json:"reasonCode" yaml:"reasonCode"`
}

