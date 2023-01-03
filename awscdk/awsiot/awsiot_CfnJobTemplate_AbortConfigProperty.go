package awsiot


// The criteria that determine when and how a job abort takes place.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   abortConfigProperty := &abortConfigProperty{
//   	criteriaList: []interface{}{
//   		&abortCriteriaProperty{
//   			action: jsii.String("action"),
//   			failureType: jsii.String("failureType"),
//   			minNumberOfExecutedThings: jsii.Number(123),
//   			thresholdPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnJobTemplate_AbortConfigProperty struct {
	// The list of criteria that determine when and how to abort the job.
	CriteriaList interface{} `field:"required" json:"criteriaList" yaml:"criteriaList"`
}

