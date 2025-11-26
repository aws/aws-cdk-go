package previewawskendrarankingmixins


// Sets additional capacity units configured for your rescore execution plan.
//
// A rescore execution plan is an Amazon Kendra Intelligent Ranking resource used for provisioning the `Rescore` API. You can add and remove capacity units to fit your usage requirements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   capacityUnitsConfigurationProperty := &CapacityUnitsConfigurationProperty{
//   	RescoreCapacityUnits: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendraranking-executionplan-capacityunitsconfiguration.html
//
type CfnExecutionPlanPropsMixin_CapacityUnitsConfigurationProperty struct {
	// The amount of extra capacity for your rescore execution plan.
	//
	// A single extra capacity unit for a rescore execution plan provides 0.01 rescore requests per second. You can add up to 1000 extra capacity units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendraranking-executionplan-capacityunitsconfiguration.html#cfn-kendraranking-executionplan-capacityunitsconfiguration-rescorecapacityunits
	//
	RescoreCapacityUnits *float64 `field:"optional" json:"rescoreCapacityUnits" yaml:"rescoreCapacityUnits"`
}

