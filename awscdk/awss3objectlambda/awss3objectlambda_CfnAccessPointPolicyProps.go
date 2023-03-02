package awss3objectlambda


// Properties for defining a `CfnAccessPointPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnAccessPointPolicyProps := &cfnAccessPointPolicyProps{
//   	objectLambdaAccessPoint: jsii.String("objectLambdaAccessPoint"),
//   	policyDocument: policyDocument,
//   }
//
type CfnAccessPointPolicyProps struct {
	// An access point with an attached AWS Lambda function used to access transformed data from an Amazon S3 bucket.
	ObjectLambdaAccessPoint *string `field:"required" json:"objectLambdaAccessPoint" yaml:"objectLambdaAccessPoint"`
	// Object Lambda Access Point resource policy document.
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
}

