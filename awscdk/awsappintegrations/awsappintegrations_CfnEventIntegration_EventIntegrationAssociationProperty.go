package awsappintegrations


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventIntegrationAssociationProperty := &eventIntegrationAssociationProperty{
//   	clientAssociationMetadata: []interface{}{
//   		&metadataProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	clientId: jsii.String("clientId"),
//   	eventBridgeRuleName: jsii.String("eventBridgeRuleName"),
//   	eventIntegrationAssociationArn: jsii.String("eventIntegrationAssociationArn"),
//   	eventIntegrationAssociationId: jsii.String("eventIntegrationAssociationId"),
//   }
//
type CfnEventIntegration_EventIntegrationAssociationProperty struct {
	// `CfnEventIntegration.EventIntegrationAssociationProperty.ClientAssociationMetadata`.
	ClientAssociationMetadata interface{} `field:"optional" json:"clientAssociationMetadata" yaml:"clientAssociationMetadata"`
	// `CfnEventIntegration.EventIntegrationAssociationProperty.ClientId`.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// `CfnEventIntegration.EventIntegrationAssociationProperty.EventBridgeRuleName`.
	EventBridgeRuleName *string `field:"optional" json:"eventBridgeRuleName" yaml:"eventBridgeRuleName"`
	// `CfnEventIntegration.EventIntegrationAssociationProperty.EventIntegrationAssociationArn`.
	EventIntegrationAssociationArn *string `field:"optional" json:"eventIntegrationAssociationArn" yaml:"eventIntegrationAssociationArn"`
	// `CfnEventIntegration.EventIntegrationAssociationProperty.EventIntegrationAssociationId`.
	EventIntegrationAssociationId *string `field:"optional" json:"eventIntegrationAssociationId" yaml:"eventIntegrationAssociationId"`
}

