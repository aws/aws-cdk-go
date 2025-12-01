package previewawsrtbfabricmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   applicationLogsProperty := &ApplicationLogsProperty{
//   	LinkApplicationLogSampling: &LinkApplicationLogSamplingProperty{
//   		ErrorLog: jsii.Number(123),
//   		FilterLog: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-outboundexternallink-applicationlogs.html
//
type CfnOutboundExternalLinkPropsMixin_ApplicationLogsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-outboundexternallink-applicationlogs.html#cfn-rtbfabric-outboundexternallink-applicationlogs-linkapplicationlogsampling
	//
	LinkApplicationLogSampling interface{} `field:"optional" json:"linkApplicationLogSampling" yaml:"linkApplicationLogSampling"`
}

