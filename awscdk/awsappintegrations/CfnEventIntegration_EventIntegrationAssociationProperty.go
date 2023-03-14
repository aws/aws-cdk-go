package awsappintegrations


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventIntegrationAssociationProperty := &EventIntegrationAssociationProperty{
//   	ClientAssociationMetadata: []interface{}{
//   		&MetadataProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ClientId: jsii.String("clientId"),
//   	EventBridgeRuleName: jsii.String("eventBridgeRuleName"),
//   	EventIntegrationAssociationArn: jsii.String("eventIntegrationAssociationArn"),
//   	EventIntegrationAssociationId: jsii.String("eventIntegrationAssociationId"),
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

