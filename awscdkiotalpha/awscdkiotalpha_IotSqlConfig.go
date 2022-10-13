// The CDK Construct Library for AWS::IoT
package awscdkiotalpha


// The type returned from the `bind()` method in {@link IotSql}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iot_alpha "github.com/aws/aws-cdk-go/awscdkiotalpha"
//
//   iotSqlConfig := &iotSqlConfig{
//   	awsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   	sql: jsii.String("sql"),
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

