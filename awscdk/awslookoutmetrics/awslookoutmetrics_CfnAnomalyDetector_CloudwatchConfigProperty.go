package awslookoutmetrics


// Details about an Amazon CloudWatch datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudwatchConfigProperty := &cloudwatchConfigProperty{
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnAnomalyDetector_CloudwatchConfigProperty struct {
	// An IAM role that gives Amazon Lookout for Metrics permission to access data in Amazon CloudWatch.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

