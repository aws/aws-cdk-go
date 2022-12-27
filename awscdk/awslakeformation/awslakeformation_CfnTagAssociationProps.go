package awslakeformation


// Properties for defining a `CfnTagAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   cfnTagAssociationProps := &cfnTagAssociationProps{
//   	lfTags: []interface{}{
//   		&lFTagPairProperty{
//   			catalogId: jsii.String("catalogId"),
//   			tagKey: jsii.String("tagKey"),
//   			tagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   	},
//   	resource: &resourceProperty{
//   		catalog: catalog,
//   		database: &databaseResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			name: jsii.String("name"),
//   		},
//   		table: &tableResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//
//   			// the properties below are optional
//   			name: jsii.String("name"),
//   			tableWildcard: tableWildcard,
//   		},
//   		tableWithColumns: &tableWithColumnsResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnTagAssociationProps struct {
	// `AWS::LakeFormation::TagAssociation.LFTags`.
	LfTags interface{} `field:"required" json:"lfTags" yaml:"lfTags"`
	// `AWS::LakeFormation::TagAssociation.Resource`.
	Resource interface{} `field:"required" json:"resource" yaml:"resource"`
}

