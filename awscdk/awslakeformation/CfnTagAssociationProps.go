package awslakeformation


// Properties for defining a `CfnTagAssociation`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack Stack
//   var accountId string
//
//
//   tagKey := "aws"
//   tagValues := []*string{
//   	"dev",
//   }
//
//   database := awscdkgluealpha.NewDatabase(this, jsii.String("Database"))
//
//   table := awscdkgluealpha.NewS3Table(this, jsii.String("Table"), &S3TableProps{
//   	Database: Database,
//   	Columns: []Column{
//   		&Column{
//   			Name: jsii.String("col1"),
//   			Type: awscdkgluealpha.Schema_STRING(),
//   		},
//   		&Column{
//   			Name: jsii.String("col2"),
//   			Type: awscdkgluealpha.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: awscdkgluealpha.DataFormat_CSV(),
//   })
//
//   synthesizer := stack.Synthesizer.(DefaultStackSynthesizer)
//   awscdk.NewCfnDataLakeSettings(this, jsii.String("DataLakeSettings"), &CfnDataLakeSettingsProps{
//   	Admins: []interface{}{
//   		&DataLakePrincipalProperty{
//   			DataLakePrincipalIdentifier: stack.FormatArn(&ArnComponents{
//   				Service: jsii.String("iam"),
//   				Resource: jsii.String("role"),
//   				Region: jsii.String(""),
//   				Account: accountId,
//   				ResourceName: jsii.String("Admin"),
//   			}),
//   		},
//   		&DataLakePrincipalProperty{
//   			// The CDK cloudformation execution role.
//   			DataLakePrincipalIdentifier: synthesizer.cloudFormationExecutionRoleArn.replace(jsii.String("${AWS::Partition}"), jsii.String("aws")),
//   		},
//   	},
//   })
//
//   tag := awscdk.NewCfnTag(this, jsii.String("Tag"), &CfnTagProps{
//   	CatalogId: accountId,
//   	TagKey: jsii.String(TagKey),
//   	TagValues: TagValues,
//   })
//
//   lfTagPairProperty := &LFTagPairProperty{
//   	CatalogId: accountId,
//   	TagKey: jsii.String(TagKey),
//   	TagValues: TagValues,
//   }
//
//   tagAssociation := awscdk.NewCfnTagAssociation(this, jsii.String("TagAssociation"), &CfnTagAssociationProps{
//   	LfTags: []interface{}{
//   		lfTagPairProperty,
//   	},
//   	Resource: &ResourceProperty{
//   		TableWithColumns: &TableWithColumnsResourceProperty{
//   			DatabaseName: database.databaseName,
//   			ColumnNames: []*string{
//   				jsii.String("col1"),
//   				jsii.String("col2"),
//   			},
//   			CatalogId: accountId,
//   			Name: table.tableName,
//   		},
//   	},
//   })
//
//   tagAssociation.Node.AddDependency(tag)
//   tagAssociation.Node.AddDependency(table)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-tagassociation.html
//
type CfnTagAssociationProps struct {
	// A structure containing an LF-tag key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-tagassociation.html#cfn-lakeformation-tagassociation-lftags
	//
	LfTags interface{} `field:"required" json:"lfTags" yaml:"lfTags"`
	// UTF-8 string (valid values: `DATABASE | TABLE` ).
	//
	// The resource for which the LF-tag policy applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-tagassociation.html#cfn-lakeformation-tagassociation-resource
	//
	Resource interface{} `field:"required" json:"resource" yaml:"resource"`
}

