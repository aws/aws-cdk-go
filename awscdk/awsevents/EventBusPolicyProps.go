package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties to associate Event Buses with a policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventBus eventBus
//   var policyStatement policyStatement
//
//   eventBusPolicyProps := &EventBusPolicyProps{
//   	EventBus: eventBus,
//   	Statement: policyStatement,
//   	StatementId: jsii.String("statementId"),
//   }
//
type EventBusPolicyProps struct {
	// The event bus to which the policy applies.
	EventBus IEventBus `field:"required" json:"eventBus" yaml:"eventBus"`
	// An IAM Policy Statement to apply to the Event Bus.
	Statement awsiam.PolicyStatement `field:"required" json:"statement" yaml:"statement"`
	// An identifier string for the external account that you are granting permissions to.
	StatementId *string `field:"required" json:"statementId" yaml:"statementId"`
}

