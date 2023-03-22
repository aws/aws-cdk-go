package awsapplicationinsights


// The `AWS::ApplicationInsights::Application Log` property type specifies a log to monitor for the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logProperty := &logProperty{
//   	logType: jsii.String("logType"),
//
//   	// the properties below are optional
//   	encoding: jsii.String("encoding"),
//   	logGroupName: jsii.String("logGroupName"),
//   	logPath: jsii.String("logPath"),
//   	patternSet: jsii.String("patternSet"),
//   }
//
type CfnApplication_LogProperty struct {
	// The log type decides the log patterns against which Application Insights analyzes the log.
	//
	// The log type is selected from the following: `SQL_SERVER` , `MYSQL` , `MYSQL_SLOW_QUERY` , `POSTGRESQL` , `ORACLE_ALERT` , `ORACLE_LISTENER` , `IIS` , `APPLICATION` , `WINDOWS_EVENTS` , `WINDOWS_EVENTS_ACTIVE_DIRECTORY` , `WINDOWS_EVENTS_DNS` , `WINDOWS_EVENTS_IIS` , `WINDOWS_EVENTS_SHAREPOINT` , `SQL_SERVER_ALWAYSON_AVAILABILITY_GROUP` , `SQL_SERVER_FAILOVER_CLUSTER_INSTANCE` , `STEP_FUNCTION` , `API_GATEWAY_ACCESS` , `API_GATEWAY_EXECUTION` , `SAP_HANA_LOGS` , `SAP_HANA_TRACE` , `SAP_HANA_HIGH_AVAILABILITY` , and `DEFAULT` .
	LogType *string `field:"required" json:"logType" yaml:"logType"`
	// The type of encoding of the logs to be monitored.
	//
	// The specified encoding should be included in the list of CloudWatch agent supported encodings. If not provided, CloudWatch Application Insights uses the default encoding type for the log type:
	//
	// - `APPLICATION/DEFAULT` : utf-8 encoding
	// - `SQL_SERVER` : utf-16 encoding
	// - `IIS` : ascii encoding.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The CloudWatch log group name to be associated with the monitored log.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The path of the logs to be monitored.
	//
	// The log path must be an absolute Windows or Linux system file path. For more information, see [CloudWatch Agent Configuration File: Logs Section](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Agent-Configuration-File-Details.html#CloudWatch-Agent-Configuration-File-Logssection) .
	LogPath *string `field:"optional" json:"logPath" yaml:"logPath"`
	// The log pattern set.
	PatternSet *string `field:"optional" json:"patternSet" yaml:"patternSet"`
}

