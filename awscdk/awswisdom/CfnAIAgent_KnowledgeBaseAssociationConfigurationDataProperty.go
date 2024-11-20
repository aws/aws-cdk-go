package awswisdom


// The data of the configuration for a `KNOWLEDGE_BASE` type Amazon Q in Connect Assistant Association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   knowledgeBaseAssociationConfigurationDataProperty := &KnowledgeBaseAssociationConfigurationDataProperty{
//   	ContentTagFilter: &TagFilterProperty{
//   		AndConditions: []interface{}{
//   			&TagConditionProperty{
//   				Key: jsii.String("key"),
//
//   				// the properties below are optional
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		OrConditions: []interface{}{
//   			&OrConditionProperty{
//   				AndConditions: []interface{}{
//   					&TagConditionProperty{
//   						Key: jsii.String("key"),
//
//   						// the properties below are optional
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				TagCondition: &TagConditionProperty{
//   					Key: jsii.String("key"),
//
//   					// the properties below are optional
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		TagCondition: &TagConditionProperty{
//   			Key: jsii.String("key"),
//
//   			// the properties below are optional
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MaxResults: jsii.Number(123),
//   	OverrideKnowledgeBaseSearchType: jsii.String("overrideKnowledgeBaseSearchType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-knowledgebaseassociationconfigurationdata.html
//
type CfnAIAgent_KnowledgeBaseAssociationConfigurationDataProperty struct {
	// An object that can be used to specify Tag conditions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-knowledgebaseassociationconfigurationdata.html#cfn-wisdom-aiagent-knowledgebaseassociationconfigurationdata-contenttagfilter
	//
	ContentTagFilter interface{} `field:"optional" json:"contentTagFilter" yaml:"contentTagFilter"`
	// The maximum number of results to return per page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-knowledgebaseassociationconfigurationdata.html#cfn-wisdom-aiagent-knowledgebaseassociationconfigurationdata-maxresults
	//
	MaxResults *float64 `field:"optional" json:"maxResults" yaml:"maxResults"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-knowledgebaseassociationconfigurationdata.html#cfn-wisdom-aiagent-knowledgebaseassociationconfigurationdata-overrideknowledgebasesearchtype
	//
	OverrideKnowledgeBaseSearchType *string `field:"optional" json:"overrideKnowledgeBaseSearchType" yaml:"overrideKnowledgeBaseSearchType"`
}

