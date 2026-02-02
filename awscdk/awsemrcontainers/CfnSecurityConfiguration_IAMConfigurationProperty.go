package awsemrcontainers


// IAM configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iAMConfigurationProperty := map[string]*string{
//   	"systemRole": jsii.String("systemRole"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-iamconfiguration.html
//
type CfnSecurityConfiguration_IAMConfigurationProperty struct {
	// The system role ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-iamconfiguration.html#cfn-emrcontainers-securityconfiguration-iamconfiguration-systemrole
	//
	SystemRole *string `field:"optional" json:"systemRole" yaml:"systemRole"`
}

