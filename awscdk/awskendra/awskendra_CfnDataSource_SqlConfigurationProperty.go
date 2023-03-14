package awskendra


// Provides information that configures Amazon Kendra to use a SQL database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqlConfigurationProperty := &sqlConfigurationProperty{
//   	queryIdentifiersEnclosingOption: jsii.String("queryIdentifiersEnclosingOption"),
//   }
//
type CfnDataSource_SqlConfigurationProperty struct {
	// Determines whether Amazon Kendra encloses SQL identifiers for tables and column names in double quotes (") when making a database query.
	//
	// You can set the value to `DOUBLE_QUOTES` or `NONE` .
	//
	// By default, Amazon Kendra passes SQL identifiers the way that they are entered into the data source configuration. It does not change the case of identifiers or enclose them in quotes.
	//
	// PostgreSQL internally converts uppercase characters to lower case characters in identifiers unless they are quoted. Choosing this option encloses identifiers in quotes so that PostgreSQL does not convert the character's case.
	//
	// For MySQL databases, you must enable the ansi_quotes option when you set this field to `DOUBLE_QUOTES` .
	QueryIdentifiersEnclosingOption *string `field:"optional" json:"queryIdentifiersEnclosingOption" yaml:"queryIdentifiersEnclosingOption"`
}

