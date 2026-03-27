package awselementalinference


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   clippingConfigProperty := &ClippingConfigProperty{
//   	CallbackMetadata: jsii.String("callbackMetadata"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-clippingconfig.html
//
type CfnFeedPropsMixin_ClippingConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-clippingconfig.html#cfn-elementalinference-feed-clippingconfig-callbackmetadata
	//
	CallbackMetadata *string `field:"optional" json:"callbackMetadata" yaml:"callbackMetadata"`
}

