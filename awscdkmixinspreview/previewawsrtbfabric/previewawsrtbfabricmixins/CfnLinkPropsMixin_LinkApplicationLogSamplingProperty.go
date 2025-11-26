package previewawsrtbfabricmixins


// Describes a link application log sample.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-linkapplicationlogsampling.html
//
type CfnLinkPropsMixin_LinkApplicationLogSamplingProperty struct {
	// An error log entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-linkapplicationlogsampling.html#cfn-rtbfabric-link-linkapplicationlogsampling-errorlog
	//
	ErrorLog *float64 `field:"optional" json:"errorLog" yaml:"errorLog"`
	// A filter log entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-linkapplicationlogsampling.html#cfn-rtbfabric-link-linkapplicationlogsampling-filterlog
	//
	FilterLog *float64 `field:"optional" json:"filterLog" yaml:"filterLog"`
}

