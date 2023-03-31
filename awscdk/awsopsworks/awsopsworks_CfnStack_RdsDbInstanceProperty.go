package awsopsworks


// Describes an Amazon RDS instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rdsDbInstanceProperty := &rdsDbInstanceProperty{
//   	dbPassword: jsii.String("dbPassword"),
//   	dbUser: jsii.String("dbUser"),
//   	rdsDbInstanceArn: jsii.String("rdsDbInstanceArn"),
//   }
//
type CfnStack_RdsDbInstanceProperty struct {
	// AWS OpsWorks Stacks returns `*****FILTERED*****` instead of the actual value.
	DbPassword *string `field:"required" json:"dbPassword" yaml:"dbPassword"`
	// The master user name.
	DbUser *string `field:"required" json:"dbUser" yaml:"dbUser"`
	// The instance's ARN.
	RdsDbInstanceArn *string `field:"required" json:"rdsDbInstanceArn" yaml:"rdsDbInstanceArn"`
}

