package mixinsawsapplicationsignals


// The start and end time of the time exclusion window.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   windowProperty := &WindowProperty{
//   	Duration: jsii.Number(123),
//   	DurationUnit: jsii.String("durationUnit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-window.html
//
type CfnServiceLevelObjectivePropsMixin_WindowProperty struct {
	// The start and end time of the time exclusion window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-window.html#cfn-applicationsignals-servicelevelobjective-window-duration
	//
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
	// The unit of measurement to use during the time window exclusion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-window.html#cfn-applicationsignals-servicelevelobjective-window-durationunit
	//
	DurationUnit *string `field:"optional" json:"durationUnit" yaml:"durationUnit"`
}

