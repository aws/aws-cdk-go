package awsiot


// The type returned from the `bind()` method in {@link IotSql}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotSqlConfig := &IotSqlConfig{
//   	AwsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   	Sql: jsii.String("sql"),
//   }
//
// Experimental.
type IotSqlConfig struct {
	// The version of the SQL rules engine to use when evaluating the rule.
	// Experimental.
	AwsIotSqlVersion *string `field:"required" json:"awsIotSqlVersion" yaml:"awsIotSqlVersion"`
	// The SQL statement used to query the topic.
	// Experimental.
	Sql *string `field:"required" json:"sql" yaml:"sql"`
}

