package awsec2


// Properties for CfnSqlHaStandbyDetectedInstancePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSqlHaStandbyDetectedInstanceMixinProps := &CfnSqlHaStandbyDetectedInstanceMixinProps{
//   	InstanceId: jsii.String("instanceId"),
//   	SqlServerCredentials: jsii.String("sqlServerCredentials"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-sqlhastandbydetectedinstance.html
//
type CfnSqlHaStandbyDetectedInstanceMixinProps struct {
	// The ID of the EC2 instance to enable for SQL Server high availability standby detection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-sqlhastandbydetectedinstance.html#cfn-ec2-sqlhastandbydetectedinstance-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The ARN of the AWS Secrets Manager secret containing SQL Server access credentials to the EC2 instance.
	//
	// If not specified, AWS Systems Manager agent will use default local user credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-sqlhastandbydetectedinstance.html#cfn-ec2-sqlhastandbydetectedinstance-sqlservercredentials
	//
	SqlServerCredentials *string `field:"optional" json:"sqlServerCredentials" yaml:"sqlServerCredentials"`
}

