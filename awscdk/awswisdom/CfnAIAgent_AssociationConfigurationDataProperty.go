package awswisdom


// A typed union of the data of the configuration for an Amazon Q in Connect Assistant Association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   associationConfigurationDataProperty := &AssociationConfigurationDataProperty{
//   	KnowledgeBaseAssociationConfigurationData: &KnowledgeBaseAssociationConfigurationDataProperty{
//   		ContentTagFilter: &TagFilterProperty{
//   			AndConditions: []interface{}{
//   				&TagConditionProperty{
//   					Key: jsii.String("key"),
//
//   					// the properties below are optional
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			OrConditions: []interface{}{
//   				&OrConditionProperty{
//   					AndConditions: []interface{}{
//   						&TagConditionProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					TagCondition: &TagConditionProperty{
//   						Key: jsii.String("key"),
//
//   						// the properties below are optional
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			TagCondition: &TagConditionProperty{
//   				Key: jsii.String("key"),
//
//   				// the properties below are optional
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MaxResults: jsii.Number(123),
//   		OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-associationconfigurationdata.html
//
type CfnAIAgent_AssociationConfigurationDataProperty struct {
	// The data of the configuration for a `KNOWLEDGE_BASE` type Amazon Q in Connect Assistant Association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-associationconfigurationdata.html#cfn-wisdom-aiagent-associationconfigurationdata-knowledgebaseassociationconfigurationdata
	//
	KnowledgeBaseAssociationConfigurationData interface{} `field:"required" json:"knowledgeBaseAssociationConfigurationData" yaml:"knowledgeBaseAssociationConfigurationData"`
}

