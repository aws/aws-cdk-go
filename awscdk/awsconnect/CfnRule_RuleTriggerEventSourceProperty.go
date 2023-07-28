package awsconnect


// The name of the event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleTriggerEventSourceProperty := &RuleTriggerEventSourceProperty{
//   	EventSourceName: jsii.String("eventSourceName"),
//
//   	// the properties below are optional
//   	IntegrationAssociationArn: jsii.String("integrationAssociationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-ruletriggereventsource.html
//
type CfnRule_RuleTriggerEventSourceProperty struct {
	// The name of the event source.
	//
	// *Allowed values* : `OnPostCallAnalysisAvailable` | `OnRealTimeCallAnalysisAvailable` | `OnPostChatAnalysisAvailable` | `OnZendeskTicketCreate` | `OnZendeskTicketStatusUpdate` | `OnSalesforceCaseCreate`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-ruletriggereventsource.html#cfn-connect-rule-ruletriggereventsource-eventsourcename
	//
	EventSourceName *string `field:"required" json:"eventSourceName" yaml:"eventSourceName"`
	// The Amazon Resource Name (ARN) of the integration association.
	//
	// `IntegrationAssociationArn` is required if `TriggerEventSource` is one of the following values: `OnZendeskTicketCreate` | `OnZendeskTicketStatusUpdate` | `OnSalesforceCaseCreate`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-ruletriggereventsource.html#cfn-connect-rule-ruletriggereventsource-integrationassociationarn
	//
	IntegrationAssociationArn *string `field:"optional" json:"integrationAssociationArn" yaml:"integrationAssociationArn"`
}

