package awslookoutmetrics


// Details about an Amazon AppFlow flow datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appFlowConfigProperty := &AppFlowConfigProperty{
//   	FlowName: jsii.String("flowName"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
type CfnAnomalyDetector_AppFlowConfigProperty struct {
	// name of the flow.
	FlowName *string `field:"required" json:"flowName" yaml:"flowName"`
	// An IAM role that gives Amazon Lookout for Metrics permission to access the flow.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

