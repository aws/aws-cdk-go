package awsauditmanager


// The `AWSService` property type specifies an AWS service such as Amazon S3 , AWS CloudTrail , and so on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aWSServiceProperty := &aWSServiceProperty{
//   	serviceName: jsii.String("serviceName"),
//   }
//
type CfnAssessment_AWSServiceProperty struct {
	// The name of the AWS service .
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

