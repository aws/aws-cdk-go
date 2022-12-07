package awsmacie


// Specifies, as a map, one or more property-based conditions that filter the results of a query for findings.
//
// For more information, see [Filtering findings](https://docs.aws.amazon.com/macie/latest/user/findings-filter-overview.html) in the *Amazon Macie User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var criterion interface{}
//
//   findingCriteriaProperty := &findingCriteriaProperty{
//   	criterion: criterion,
//   }
//
type CfnFindingsFilter_FindingCriteriaProperty struct {
	// Specifies a condition that defines the property, operator, and one or more values to use to filter the results.
	Criterion interface{} `field:"optional" json:"criterion" yaml:"criterion"`
}

