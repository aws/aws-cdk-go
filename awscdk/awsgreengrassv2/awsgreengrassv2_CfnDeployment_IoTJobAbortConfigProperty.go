package awsgreengrassv2


// Contains a list of criteria that define when and how to cancel a configuration deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ioTJobAbortConfigProperty := &ioTJobAbortConfigProperty{
//   	criteriaList: []interface{}{
//   		&ioTJobAbortCriteriaProperty{
//   			action: jsii.String("action"),
//   			failureType: jsii.String("failureType"),
//   			minNumberOfExecutedThings: jsii.Number(123),
//   			thresholdPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnDeployment_IoTJobAbortConfigProperty struct {
	// The list of criteria that define when and how to cancel the configuration deployment.
	CriteriaList interface{} `field:"required" json:"criteriaList" yaml:"criteriaList"`
}

