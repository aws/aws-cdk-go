package interfacesawsevents


// A reference to a EventBusPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBusPolicyReference := &EventBusPolicyReference{
//   	EventBusName: jsii.String("eventBusName"),
//   	StatementId: jsii.String("statementId"),
//   }
//
type EventBusPolicyReference struct {
	// The EventBusName of the EventBusPolicy resource.
	EventBusName *string `field:"required" json:"eventBusName" yaml:"eventBusName"`
	// The StatementId of the EventBusPolicy resource.
	StatementId *string `field:"required" json:"statementId" yaml:"statementId"`
}

