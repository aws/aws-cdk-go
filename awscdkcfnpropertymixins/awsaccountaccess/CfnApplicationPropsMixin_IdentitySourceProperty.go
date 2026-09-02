package awsaccountaccess


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   identitySourceProperty := &IdentitySourceProperty{
//   	IdentityCenter: &IdentityCenterProperty{
//   		ApplicationArn: jsii.String("applicationArn"),
//   		InstanceArn: jsii.String("instanceArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-application-identitysource.html
//
type CfnApplicationPropsMixin_IdentitySourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-application-identitysource.html#cfn-accountaccess-application-identitysource-identitycenter
	//
	IdentityCenter interface{} `field:"optional" json:"identityCenter" yaml:"identityCenter"`
}

