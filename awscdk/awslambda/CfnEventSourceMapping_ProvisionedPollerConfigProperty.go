package awslambda


// ProvisionedPollerConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedPollerConfigProperty := &ProvisionedPollerConfigProperty{
//   	MaximumPollers: jsii.Number(123),
//   	MinimumPollers: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-provisionedpollerconfig.html
//
type CfnEventSourceMapping_ProvisionedPollerConfigProperty struct {
	// Maximum poller count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-provisionedpollerconfig.html#cfn-lambda-eventsourcemapping-provisionedpollerconfig-maximumpollers
	//
	MaximumPollers *float64 `field:"optional" json:"maximumPollers" yaml:"maximumPollers"`
	// Minimum poller count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-provisionedpollerconfig.html#cfn-lambda-eventsourcemapping-provisionedpollerconfig-minimumpollers
	//
	MinimumPollers *float64 `field:"optional" json:"minimumPollers" yaml:"minimumPollers"`
}

