package previewawsathenamixins


// Properties for CfnNamedQueryPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNamedQueryMixinProps := &CfnNamedQueryMixinProps{
//   	Database: jsii.String("database"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	QueryString: jsii.String("queryString"),
//   	WorkGroup: jsii.String("workGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html
//
type CfnNamedQueryMixinProps struct {
	// The database to which the query belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-database
	//
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The query description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The query name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The SQL statements that make up the query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-querystring
	//
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// The name of the workgroup that contains the named query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-workgroup
	//
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

