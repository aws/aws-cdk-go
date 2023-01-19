package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobExecutionsRetryConfigProperty := &jobExecutionsRetryConfigProperty{
//   	retryCriteriaList: []interface{}{
//   		&retryCriteriaProperty{
//   			failureType: jsii.String("failureType"),
//   			numberOfRetries: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnJobTemplate_JobExecutionsRetryConfigProperty struct {
	// `CfnJobTemplate.JobExecutionsRetryConfigProperty.RetryCriteriaList`.
	RetryCriteriaList interface{} `field:"optional" json:"retryCriteriaList" yaml:"retryCriteriaList"`
}

