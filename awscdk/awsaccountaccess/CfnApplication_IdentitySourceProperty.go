package awsaccountaccess


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identitySourceProperty := &IdentitySourceProperty{
//   	IdentityCenter: &IdentityCenterProperty{
//   		InstanceArn: jsii.String("instanceArn"),
//
//   		// the properties below are optional
//   		ApplicationArn: jsii.String("applicationArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-application-identitysource.html
//
type CfnApplication_IdentitySourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-application-identitysource.html#cfn-accountaccess-application-identitysource-identitycenter
	//
	IdentityCenter interface{} `field:"required" json:"identityCenter" yaml:"identityCenter"`
}

