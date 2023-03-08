package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ioTRuleEventProperty := &ioTRuleEventProperty{
//   	sql: jsii.String("sql"),
//
//   	// the properties below are optional
//   	awsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   }
//
type CfnFunction_IoTRuleEventProperty struct {
	// `CfnFunction.IoTRuleEventProperty.Sql`.
	Sql *string `field:"required" json:"sql" yaml:"sql"`
	// `CfnFunction.IoTRuleEventProperty.AwsIotSqlVersion`.
	AwsIotSqlVersion *string `field:"optional" json:"awsIotSqlVersion" yaml:"awsIotSqlVersion"`
}

