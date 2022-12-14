package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleTriggerEventSourceProperty := &ruleTriggerEventSourceProperty{
//   	eventSourceName: jsii.String("eventSourceName"),
//
//   	// the properties below are optional
//   	integrationAssociationArn: jsii.String("integrationAssociationArn"),
//   }
//
type CfnRule_RuleTriggerEventSourceProperty struct {
	// `CfnRule.RuleTriggerEventSourceProperty.EventSourceName`.
	EventSourceName *string `field:"required" json:"eventSourceName" yaml:"eventSourceName"`
	// `CfnRule.RuleTriggerEventSourceProperty.IntegrationAssociationArn`.
	IntegrationAssociationArn *string `field:"optional" json:"integrationAssociationArn" yaml:"integrationAssociationArn"`
}

