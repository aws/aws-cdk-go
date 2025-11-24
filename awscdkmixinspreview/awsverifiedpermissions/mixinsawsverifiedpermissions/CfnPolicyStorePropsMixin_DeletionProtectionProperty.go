package mixinsawsverifiedpermissions


// Specifies whether the policy store can be deleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deletionProtectionProperty := &DeletionProtectionProperty{
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-deletionprotection.html
//
type CfnPolicyStorePropsMixin_DeletionProtectionProperty struct {
	// Specifies whether the policy store can be deleted. If enabled, the policy store can't be deleted.
	//
	// The default state is `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-deletionprotection.html#cfn-verifiedpermissions-policystore-deletionprotection-mode
	//
	// Default: - "DISABLED".
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

