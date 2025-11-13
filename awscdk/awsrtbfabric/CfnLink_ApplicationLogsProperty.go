package awsrtbfabric


// Describes the configuration of a link application log.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationLogsProperty := &ApplicationLogsProperty{
//   	LinkApplicationLogSampling: &LinkApplicationLogSamplingProperty{
//   		ErrorLog: jsii.Number(123),
//   		FilterLog: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-applicationlogs.html
//
type CfnLink_ApplicationLogsProperty struct {
	// Describes a link application log sample.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-applicationlogs.html#cfn-rtbfabric-link-applicationlogs-linkapplicationlogsampling
	//
	LinkApplicationLogSampling interface{} `field:"required" json:"linkApplicationLogSampling" yaml:"linkApplicationLogSampling"`
}

