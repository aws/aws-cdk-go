package awsmacie


// Specifies, as a map, one or more property-based conditions that filter the results of a query for findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   findingCriteriaProperty := &findingCriteriaProperty{
//   	criterion: map[string]interface{}{
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
type CfnFindingsFilter_FindingCriteriaProperty struct {
	// Specifies a condition that defines the property, operator, and value to use to filter the results.
	Criterion interface{} `field:"optional" json:"criterion" yaml:"criterion"`
}

