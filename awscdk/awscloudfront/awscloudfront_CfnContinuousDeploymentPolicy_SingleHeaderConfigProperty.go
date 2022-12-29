package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singleHeaderConfigProperty := &singleHeaderConfigProperty{
//   	header: jsii.String("header"),
//   	value: jsii.String("value"),
//   }
//
type CfnContinuousDeploymentPolicy_SingleHeaderConfigProperty struct {
	// `CfnContinuousDeploymentPolicy.SingleHeaderConfigProperty.Header`.
	Header *string `field:"required" json:"header" yaml:"header"`
	// `CfnContinuousDeploymentPolicy.SingleHeaderConfigProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

