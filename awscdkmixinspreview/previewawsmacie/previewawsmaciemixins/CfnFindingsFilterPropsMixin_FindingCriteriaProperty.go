package previewawsmaciemixins


// Specifies, as a map, one or more property-based conditions for a findings filter.
//
// A *findings filter* , also referred to as a *filter rule* , is a set of custom criteria that specifies which findings to include or exclude from the results of a query for findings. You can also configure a findings filter to suppress (automatically archive) findings that match the filter's criteria. For more information, see [Filtering Macie findings](https://docs.aws.amazon.com/macie/latest/user/findings-filter-overview.html) in the *Amazon Macie User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   findingCriteriaProperty := &FindingCriteriaProperty{
//   	Criterion: map[string]interface{}{
//   		"criterionKey": &CriterionAdditionalPropertiesProperty{
//   			"eq": []*string{
//   				jsii.String("eq"),
//   			},
//   			"gt": jsii.Number(123),
//   			"gte": jsii.Number(123),
//   			"lt": jsii.Number(123),
//   			"lte": jsii.Number(123),
//   			"neq": []*string{
//   				jsii.String("neq"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-findingcriteria.html
//
type CfnFindingsFilterPropsMixin_FindingCriteriaProperty struct {
	// Specifies a condition that defines the property, operator, and one or more values to use to filter the results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-findingcriteria.html#cfn-macie-findingsfilter-findingcriteria-criterion
	//
	Criterion interface{} `field:"optional" json:"criterion" yaml:"criterion"`
}

