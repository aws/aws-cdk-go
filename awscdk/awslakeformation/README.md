# AWS::LakeFormation Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

```go
import lakeformation "github.com/aws/aws-cdk-go/awscdk"
```

<!--BEGIN CFNONLY DISCLAIMER-->

There are no official hand-written ([L2](https://docs.aws.amazon.com/cdk/latest/guide/constructs.html#constructs_lib)) constructs for this service yet. Here are some suggestions on how to proceed:

* Search [Construct Hub for LakeFormation construct libraries](https://constructs.dev/search?q=lakeformation)
* Use the automatically generated [L1](https://docs.aws.amazon.com/cdk/latest/guide/constructs.html#constructs_l1_using) constructs, in the same way you would use [the CloudFormation AWS::LakeFormation resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/AWS_LakeFormation.html) directly.

<!--BEGIN CFNONLY DISCLAIMER-->

There are no hand-written ([L2](https://docs.aws.amazon.com/cdk/latest/guide/constructs.html#constructs_lib)) constructs for this service yet.
However, you can still use the automatically generated [L1](https://docs.aws.amazon.com/cdk/latest/guide/constructs.html#constructs_l1_using) constructs, and use this service exactly as you would using CloudFormation directly.

For more information on the resources and properties available for this service, see the [CloudFormation documentation for AWS::LakeFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/AWS_LakeFormation.html).

(Read the [CDK Contributing Guide](https://github.com/aws/aws-cdk/blob/main/CONTRIBUTING.md) and submit an RFC if you are interested in contributing to this construct library.)

<!--END CFNONLY DISCLAIMER-->

## Example

Here is an example of creating a glue table and putting lakeformation tags on it. Note: this example uses deprecated constructs and overly permissive IAM roles. This example is meant to give a general idea of using the L1s; it is not production level.

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdkgluealpha"
import "github.com/aws/aws-cdk-go/awscdk"

var stack stack
var accountId string


tagKey := "aws"
tagValues := []*string{
	"dev",
}

database := awscdkgluealpha.NewDatabase(this, jsii.String("Database"))

table := awscdkgluealpha.NewS3Table(this, jsii.String("Table"), &S3TableProps{
	Database: Database,
	Columns: []column{
		&column{
			Name: jsii.String("col1"),
			Type: awscdkgluealpha.Schema_STRING(),
		},
		&column{
			Name: jsii.String("col2"),
			Type: awscdkgluealpha.Schema_STRING(),
		},
	},
	DataFormat: awscdkgluealpha.DataFormat_CSV(),
})

synthesizer := stack.Synthesizer.(defaultStackSynthesizer)
awscdk.NewCfnDataLakeSettings(this, jsii.String("DataLakeSettings"), &CfnDataLakeSettingsProps{
	Admins: []interface{}{
		&DataLakePrincipalProperty{
			DataLakePrincipalIdentifier: stack.FormatArn(&ArnComponents{
				Service: jsii.String("iam"),
				Resource: jsii.String("role"),
				Region: jsii.String(""),
				Account: accountId,
				ResourceName: jsii.String("Admin"),
			}),
		},
		&DataLakePrincipalProperty{
			// The CDK cloudformation execution role.
			DataLakePrincipalIdentifier: synthesizer.cloudFormationExecutionRoleArn.replace(jsii.String("${AWS::Partition}"), jsii.String("aws")),
		},
	},
})

tag := awscdk.NewCfnTag(this, jsii.String("Tag"), &CfnTagProps{
	CatalogId: accountId,
	TagKey: jsii.String(TagKey),
	TagValues: TagValues,
})

lfTagPairProperty := &LFTagPairProperty{
	CatalogId: accountId,
	TagKey: jsii.String(TagKey),
	TagValues: TagValues,
}

tagAssociation := awscdk.NewCfnTagAssociation(this, jsii.String("TagAssociation"), &CfnTagAssociationProps{
	LfTags: []interface{}{
		lfTagPairProperty,
	},
	Resource: &ResourceProperty{
		TableWithColumns: &TableWithColumnsResourceProperty{
			DatabaseName: database.DatabaseName,
			ColumnNames: []*string{
				jsii.String("col1"),
				jsii.String("col2"),
			},
			CatalogId: accountId,
			Name: table.TableName,
		},
	},
})

tagAssociation.Node.AddDependency(tag)
tagAssociation.Node.AddDependency(table)
```

Additionally, you may need to use the lakeformation console to give permissions, particularly to give the cdk-exec-role tagging permissions.
