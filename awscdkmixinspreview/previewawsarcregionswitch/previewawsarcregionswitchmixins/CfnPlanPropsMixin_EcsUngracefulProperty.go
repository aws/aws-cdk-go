package previewawsarcregionswitchmixins


// The settings for ungraceful execution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ecsUngracefulProperty := &EcsUngracefulProperty{
//   	MinimumSuccessPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ecsungraceful.html
//
type CfnPlanPropsMixin_EcsUngracefulProperty struct {
	// The minimum success percentage specified for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ecsungraceful.html#cfn-arcregionswitch-plan-ecsungraceful-minimumsuccesspercentage
	//
	MinimumSuccessPercentage *float64 `field:"optional" json:"minimumSuccessPercentage" yaml:"minimumSuccessPercentage"`
}

