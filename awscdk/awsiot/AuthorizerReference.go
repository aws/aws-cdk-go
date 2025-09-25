package awsiot


// A reference to a Authorizer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizerReference := &AuthorizerReference{
//   	AuthorizerArn: jsii.String("authorizerArn"),
//   	AuthorizerName: jsii.String("authorizerName"),
//   }
//
type AuthorizerReference struct {
	// The ARN of the Authorizer resource.
	AuthorizerArn *string `field:"required" json:"authorizerArn" yaml:"authorizerArn"`
	// The AuthorizerName of the Authorizer resource.
	AuthorizerName *string `field:"required" json:"authorizerName" yaml:"authorizerName"`
}

