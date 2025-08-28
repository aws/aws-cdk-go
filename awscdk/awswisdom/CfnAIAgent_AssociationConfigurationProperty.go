package awswisdom


// The configuration for an Amazon Q in Connect Assistant Association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   associationConfigurationProperty := &AssociationConfigurationProperty{
//   	AssociationConfigurationData: &AssociationConfigurationDataProperty{
//   		KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   			ContentTagFilter: &TagFilterProperty{
//   				AndConditions: []interface{}{
//   					&TagConditionProperty{
//   						Key: jsii.String("key"),
//
//   						// the properties below are optional
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				OrConditions: []interface{}{
//   					&OrConditionProperty{
//   						AndConditions: []interface{}{
//   							&TagConditionProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						TagCondition: &TagConditionProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				TagCondition: &TagConditionProperty{
//   					Key: jsii.String("key"),
//
//   					// the properties below are optional
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MaxResults: jsii.Number(123),
//   			OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   		},
//   	},
//   	AssociationId: jsii.String("associationId"),
//   	AssociationType: jsii.String("associationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-associationconfiguration.html
//
type CfnAIAgent_AssociationConfigurationProperty struct {
	// A typed union of the data of the configuration for an Amazon Q in Connect Assistant Association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-associationconfiguration.html#cfn-wisdom-aiagent-associationconfiguration-associationconfigurationdata
	//
	AssociationConfigurationData interface{} `field:"optional" json:"associationConfigurationData" yaml:"associationConfigurationData"`
	// The identifier of the association for this Association Configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-associationconfiguration.html#cfn-wisdom-aiagent-associationconfiguration-associationid
	//
	AssociationId *string `field:"optional" json:"associationId" yaml:"associationId"`
	// The type of the association for this Association Configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-associationconfiguration.html#cfn-wisdom-aiagent-associationconfiguration-associationtype
	//
	AssociationType *string `field:"optional" json:"associationType" yaml:"associationType"`
}

