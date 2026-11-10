package awsmgn


// The aggregated migration status of the wave.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   waveAggregatedStatusProperty := &WaveAggregatedStatusProperty{
//   	HealthStatus: jsii.String("healthStatus"),
//   	ProgressStatus: jsii.String("progressStatus"),
//   	TotalApplications: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-wave-waveaggregatedstatus.html
//
type CfnWavePropsMixin_WaveAggregatedStatusProperty struct {
	// The health status of the wave.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-wave-waveaggregatedstatus.html#cfn-mgn-wave-waveaggregatedstatus-healthstatus
	//
	HealthStatus *string `field:"optional" json:"healthStatus" yaml:"healthStatus"`
	// The progress status of the wave.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-wave-waveaggregatedstatus.html#cfn-mgn-wave-waveaggregatedstatus-progressstatus
	//
	ProgressStatus *string `field:"optional" json:"progressStatus" yaml:"progressStatus"`
	// The total number of applications in the wave.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-wave-waveaggregatedstatus.html#cfn-mgn-wave-waveaggregatedstatus-totalapplications
	//
	TotalApplications *float64 `field:"optional" json:"totalApplications" yaml:"totalApplications"`
}

