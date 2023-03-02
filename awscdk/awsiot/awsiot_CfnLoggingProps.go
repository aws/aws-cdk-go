package awsiot


// Properties for defining a `CfnLogging`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoggingProps := &cfnLoggingProps{
//   	accountId: jsii.String("accountId"),
//   	defaultLogLevel: jsii.String("defaultLogLevel"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnLoggingProps struct {
	// The account ID.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The default log level.
	//
	// Valid Values: `DEBUG | INFO | ERROR | WARN | DISABLED`.
	DefaultLogLevel *string `field:"required" json:"defaultLogLevel" yaml:"defaultLogLevel"`
	// The role ARN used for the log.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

