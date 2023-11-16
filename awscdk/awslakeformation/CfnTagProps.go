package awslakeformation


// Properties for defining a `CfnTag`.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
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
//   	Columns: []column{
//   		&column{
//   			Name: jsii.String("col1"),
//   			Type: awscdkgluealpha.Schema_STRING(),
//   		},
//   		&column{
//   			Name: jsii.String("col2"),
//   			Type: awscdkgluealpha.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: awscdkgluealpha.DataFormat_CSV(),
//   })
//
//   synthesizer := stack.Synthesizer.(defaultStackSynthesizer)
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
//   			DatabaseName: database.DatabaseName,
//   			ColumnNames: []*string{
//   				jsii.String("col1"),
//   				jsii.String("col2"),
//   			},
//   			CatalogId: accountId,
//   			Name: table.TableName,
//   		},
//   	},
//   })
//
//   tagAssociation.Node.AddDependency(tag)
//   tagAssociation.Node.AddDependency(table)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-tag.html
//
type CfnTagProps struct {
	// UTF-8 string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.html) .
	//
	// The key-name for the LF-tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-tag.html#cfn-lakeformation-tag-tagkey
	//
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// An array of UTF-8 strings, not less than 1 or more than 50 strings.
	//
	// A list of possible values of the corresponding `TagKey` of an LF-tag key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-tag.html#cfn-lakeformation-tag-tagvalues
	//
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
	// Catalog id string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.html) .
	//
	// The identifier for the Data Catalog . By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your AWS Lake Formation environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-tag.html#cfn-lakeformation-tag-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
}

