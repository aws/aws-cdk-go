package mixinsawskendra


// Provides the configuration information for the knowledge article types that Amazon Kendra indexes.
//
// Amazon Kendra indexes standard knowledge articles and the standard fields of knowledge articles, or the custom fields of custom knowledge articles, but not both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   salesforceKnowledgeArticleConfigurationProperty := &SalesforceKnowledgeArticleConfigurationProperty{
//   	CustomKnowledgeArticleTypeConfigurations: []interface{}{
//   		&SalesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   			DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	IncludedStates: []*string{
//   		jsii.String("includedStates"),
//   	},
//   	StandardKnowledgeArticleTypeConfiguration: &SalesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   		DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   		DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceknowledgearticleconfiguration.html
//
type CfnDataSourcePropsMixin_SalesforceKnowledgeArticleConfigurationProperty struct {
	// Configuration information for custom Salesforce knowledge articles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceknowledgearticleconfiguration.html#cfn-kendra-datasource-salesforceknowledgearticleconfiguration-customknowledgearticletypeconfigurations
	//
	CustomKnowledgeArticleTypeConfigurations interface{} `field:"optional" json:"customKnowledgeArticleTypeConfigurations" yaml:"customKnowledgeArticleTypeConfigurations"`
	// Specifies the document states that should be included when Amazon Kendra indexes knowledge articles.
	//
	// You must specify at least one state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceknowledgearticleconfiguration.html#cfn-kendra-datasource-salesforceknowledgearticleconfiguration-includedstates
	//
	IncludedStates *[]*string `field:"optional" json:"includedStates" yaml:"includedStates"`
	// Configuration information for standard Salesforce knowledge articles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceknowledgearticleconfiguration.html#cfn-kendra-datasource-salesforceknowledgearticleconfiguration-standardknowledgearticletypeconfiguration
	//
	StandardKnowledgeArticleTypeConfiguration interface{} `field:"optional" json:"standardKnowledgeArticleTypeConfiguration" yaml:"standardKnowledgeArticleTypeConfiguration"`
}

