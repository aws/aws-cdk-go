package awsathena


// Properties for defining a `CfnNamedQuery`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNamedQueryProps := &cfnNamedQueryProps{
//   	database: jsii.String("database"),
//   	queryString: jsii.String("queryString"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	workGroup: jsii.String("workGroup"),
//   }
//
type CfnNamedQueryProps struct {
	// The database to which the query belongs.
	Database *string `field:"required" json:"database" yaml:"database"`
	// The SQL statements that make up the query.
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// The query description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The query name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the workgroup that contains the named query.
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

