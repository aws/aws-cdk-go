package awsfsx


// The configuration that Amazon FSx for Windows File Server uses to audit and log user accesses of files, folders, and file shares on the Amazon FSx for Windows File Server file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auditLogConfigurationProperty := &auditLogConfigurationProperty{
//   	fileAccessAuditLogLevel: jsii.String("fileAccessAuditLogLevel"),
//   	fileShareAccessAuditLogLevel: jsii.String("fileShareAccessAuditLogLevel"),
//
//   	// the properties below are optional
//   	auditLogDestination: jsii.String("auditLogDestination"),
//   }
//
type CfnFileSystem_AuditLogConfigurationProperty struct {
	// Sets which attempt type is logged by Amazon FSx for file and folder accesses.
	//
	// - `SUCCESS_ONLY` - only successful attempts to access files or folders are logged.
	// - `FAILURE_ONLY` - only failed attempts to access files or folders are logged.
	// - `SUCCESS_AND_FAILURE` - both successful attempts and failed attempts to access files or folders are logged.
	// - `DISABLED` - access auditing of files and folders is turned off.
	FileAccessAuditLogLevel *string `field:"required" json:"fileAccessAuditLogLevel" yaml:"fileAccessAuditLogLevel"`
	// Sets which attempt type is logged by Amazon FSx for file share accesses.
	//
	// - `SUCCESS_ONLY` - only successful attempts to access file shares are logged.
	// - `FAILURE_ONLY` - only failed attempts to access file shares are logged.
	// - `SUCCESS_AND_FAILURE` - both successful attempts and failed attempts to access file shares are logged.
	// - `DISABLED` - access auditing of file shares is turned off.
	FileShareAccessAuditLogLevel *string `field:"required" json:"fileShareAccessAuditLogLevel" yaml:"fileShareAccessAuditLogLevel"`
	// The Amazon Resource Name (ARN) for the destination of the audit logs.
	//
	// The destination can be any Amazon CloudWatch Logs log group ARN or Amazon Kinesis Data Firehose delivery stream ARN.
	//
	// The name of the Amazon CloudWatch Logs log group must begin with the `/aws/fsx` prefix. The name of the Amazon Kinesis Data Firehouse delivery stream must begin with the `aws-fsx` prefix.
	//
	// The destination ARN (either CloudWatch Logs log group or Kinesis Data Firehose delivery stream) must be in the same AWS partition, AWS Region , and AWS account as your Amazon FSx file system.
	AuditLogDestination *string `field:"optional" json:"auditLogDestination" yaml:"auditLogDestination"`
}

