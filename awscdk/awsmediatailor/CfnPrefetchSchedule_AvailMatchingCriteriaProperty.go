package awsmediatailor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availMatchingCriteriaProperty := &AvailMatchingCriteriaProperty{
//   	DynamicVariable: jsii.String("dynamicVariable"),
//   	Operator: jsii.String("operator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-availmatchingcriteria.html
//
type CfnPrefetchSchedule_AvailMatchingCriteriaProperty struct {
	// The dynamic variable(s) that MediaTailor should use as avail matching criteria.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-availmatchingcriteria.html#cfn-mediatailor-prefetchschedule-availmatchingcriteria-dynamicvariable
	//
	DynamicVariable *string `field:"required" json:"dynamicVariable" yaml:"dynamicVariable"`
	// For the DynamicVariable specified in AvailMatchingCriteria, the Operator that is used for the comparison.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-availmatchingcriteria.html#cfn-mediatailor-prefetchschedule-availmatchingcriteria-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
}

