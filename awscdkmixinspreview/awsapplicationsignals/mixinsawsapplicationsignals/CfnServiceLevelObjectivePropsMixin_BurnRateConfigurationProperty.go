package mixinsawsapplicationsignals


// This object defines the length of the look-back window used to calculate one burn rate metric for this SLO.
//
// The burn rate measures how fast the service is consuming the error budget, relative to the attainment goal of the SLO. A burn rate of exactly 1 indicates that the SLO goal will be met exactly.
//
// For example, if you specify 60 as the number of minutes in the look-back window, the burn rate is calculated as the following:
//
// *burn rate = error rate over the look-back window / (100% - attainment goal percentage)*
//
// For more information about burn rates, see [Calculate burn rates](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-ServiceLevelObjectives.html#CloudWatch-ServiceLevelObjectives-burn) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   burnRateConfigurationProperty := &BurnRateConfigurationProperty{
//   	LookBackWindowMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-burnrateconfiguration.html
//
type CfnServiceLevelObjectivePropsMixin_BurnRateConfigurationProperty struct {
	// The number of minutes to use as the look-back window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-burnrateconfiguration.html#cfn-applicationsignals-servicelevelobjective-burnrateconfiguration-lookbackwindowminutes
	//
	LookBackWindowMinutes *float64 `field:"optional" json:"lookBackWindowMinutes" yaml:"lookBackWindowMinutes"`
}

