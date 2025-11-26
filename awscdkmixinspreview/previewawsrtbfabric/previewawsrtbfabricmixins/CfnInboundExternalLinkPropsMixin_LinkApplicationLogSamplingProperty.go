package previewawsrtbfabricmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   linkApplicationLogSamplingProperty := &LinkApplicationLogSamplingProperty{
//   	ErrorLog: jsii.Number(123),
//   	FilterLog: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-linkapplicationlogsampling.html
//
type CfnInboundExternalLinkPropsMixin_LinkApplicationLogSamplingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-linkapplicationlogsampling.html#cfn-rtbfabric-inboundexternallink-linkapplicationlogsampling-errorlog
	//
	ErrorLog *float64 `field:"optional" json:"errorLog" yaml:"errorLog"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-linkapplicationlogsampling.html#cfn-rtbfabric-inboundexternallink-linkapplicationlogsampling-filterlog
	//
	FilterLog *float64 `field:"optional" json:"filterLog" yaml:"filterLog"`
}

