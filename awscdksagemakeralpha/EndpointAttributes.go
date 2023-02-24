// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha


// Represents an Endpoint resource defined outside this stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import sagemaker_alpha "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   endpointAttributes := &EndpointAttributes{
//   	EndpointArn: jsii.String("endpointArn"),
//   }
//
// Experimental.
type EndpointAttributes struct {
	// The ARN of this endpoint.
	// Experimental.
	EndpointArn *string `field:"required" json:"endpointArn" yaml:"endpointArn"`
}

