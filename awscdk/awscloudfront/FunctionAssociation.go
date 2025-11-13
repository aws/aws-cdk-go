package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront"
)

// Represents a CloudFront function and event type when using CF Functions.
//
// The type of the `AddBehaviorOptions.functionAssociations` property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var functionRef IFunctionRef
//
//   functionAssociation := &FunctionAssociation{
//   	EventType: awscdk.Aws_cloudfront.FunctionEventType_VIEWER_REQUEST,
//   	Function: functionRef,
//   }
//
type FunctionAssociation struct {
	// The type of event which should invoke the function.
	EventType FunctionEventType `field:"required" json:"eventType" yaml:"eventType"`
	// The CloudFront function that will be invoked.
	Function interfacesawscloudfront.IFunctionRef `field:"required" json:"function" yaml:"function"`
}

