package awsathena


// Properties for defining a `CfnNamedQuery`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNamedQueryProps := &CfnNamedQueryProps{
//   	Database: jsii.String("database"),
//   	QueryString: jsii.String("queryString"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	WorkGroup: jsii.String("workGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html
//
type CfnNamedQueryProps struct {
	// The database to which the query belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-database
	//
	Database *string `field:"required" json:"database" yaml:"database"`
	// The SQL statements that make up the query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-querystring
	//
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// The query description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The query name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the workgroup that contains the named query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-workgroup
	//
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

