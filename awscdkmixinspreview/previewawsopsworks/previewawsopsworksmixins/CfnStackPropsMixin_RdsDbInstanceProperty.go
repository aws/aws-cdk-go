package previewawsopsworksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rdsDbInstanceProperty := &RdsDbInstanceProperty{
//   	DbPassword: jsii.String("dbPassword"),
//   	DbUser: jsii.String("dbUser"),
//   	RdsDbInstanceArn: jsii.String("rdsDbInstanceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html
//
type CfnStackPropsMixin_RdsDbInstanceProperty struct {
	// OpsWorks Stacks returns `*****FILTERED*****` instead of the actual value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-dbpassword
	//
	DbPassword *string `field:"optional" json:"dbPassword" yaml:"dbPassword"`
	// The master user name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-dbuser
	//
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// The instance's ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-rdsdbinstancearn
	//
	RdsDbInstanceArn *string `field:"optional" json:"rdsDbInstanceArn" yaml:"rdsDbInstanceArn"`
}

