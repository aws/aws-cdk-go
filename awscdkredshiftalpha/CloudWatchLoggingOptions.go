package awscdkredshiftalpha


// Options for CloudWatch logging.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkredshiftalpha"
//
//   var vpc IVpc
//
//
//   cluster := awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
//   	MasterUser: &Login{
//   		MasterUsername: jsii.String("admin"),
//   	},
//   	Vpc: Vpc,
//   	Logging: awscdkredshiftalpha.ClusterLogging_CloudWatch(&CloudWatchLoggingOptions{
//   		LogExports: []LogExport{
//   			awscdkredshiftalpha.LogExport_CONNECTION_LOG,
//   			awscdkredshiftalpha.LogExport_USER_LOG,
//   		},
//   	}),
//   })
//
// Experimental.
type CloudWatchLoggingOptions struct {
	// The types of logs to export.
	// Default: - [LogExport.CONNECTION_LOG, LogExport.USER_LOG, LogExport.USER_ACTIVITY_LOG]
	//
	// Experimental.
	LogExports *[]LogExport `field:"optional" json:"logExports" yaml:"logExports"`
}

