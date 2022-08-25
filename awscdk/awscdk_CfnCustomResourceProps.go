// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Properties for defining a `CfnCustomResource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomResourceProps := &cfnCustomResourceProps{
//   	serviceToken: jsii.String("serviceToken"),
//   }
//
type CfnCustomResourceProps struct {
	// > Only one property is defined by AWS for a custom resource: `ServiceToken` .
	//
	// All other properties are defined by the service provider.
	//
	// The service token that was given to the template developer by the service provider to access the service, such as an Amazon SNS topic ARN or Lambda function ARN. The service token must be from the same Region in which you are creating the stack.
	//
	// Updates aren't supported.
	ServiceToken *string `field:"required" json:"serviceToken" yaml:"serviceToken"`
}

