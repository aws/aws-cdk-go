package awsrtbfabric


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linkApplicationLogSamplingProperty := &LinkApplicationLogSamplingProperty{
//   	ErrorLog: jsii.Number(123),
//   	FilterLog: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-outboundexternallink-linkapplicationlogsampling.html
//
type CfnOutboundExternalLink_LinkApplicationLogSamplingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-outboundexternallink-linkapplicationlogsampling.html#cfn-rtbfabric-outboundexternallink-linkapplicationlogsampling-errorlog
	//
	ErrorLog *float64 `field:"required" json:"errorLog" yaml:"errorLog"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-outboundexternallink-linkapplicationlogsampling.html#cfn-rtbfabric-outboundexternallink-linkapplicationlogsampling-filterlog
	//
	FilterLog *float64 `field:"required" json:"filterLog" yaml:"filterLog"`
}

