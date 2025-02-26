package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ioTRuleEventProperty := &IoTRuleEventProperty{
//   	Sql: jsii.String("sql"),
//
//   	// the properties below are optional
//   	AwsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-iotruleevent.html
//
type CfnFunction_IoTRuleEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-iotruleevent.html#cfn-serverless-function-iotruleevent-sql
	//
	Sql *string `field:"required" json:"sql" yaml:"sql"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-iotruleevent.html#cfn-serverless-function-iotruleevent-awsiotsqlversion
	//
	AwsIotSqlVersion *string `field:"optional" json:"awsIotSqlVersion" yaml:"awsIotSqlVersion"`
}

