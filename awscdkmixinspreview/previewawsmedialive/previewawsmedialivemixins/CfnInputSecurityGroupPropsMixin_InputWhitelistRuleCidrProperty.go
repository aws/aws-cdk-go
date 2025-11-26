package previewawsmedialivemixins


// An IPv4 CIDR range to include in this input security group.
//
// The parent of this entity is InputSecurityGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputWhitelistRuleCidrProperty := &InputWhitelistRuleCidrProperty{
//   	Cidr: jsii.String("cidr"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-inputsecuritygroup-inputwhitelistrulecidr.html
//
type CfnInputSecurityGroupPropsMixin_InputWhitelistRuleCidrProperty struct {
	// An IPv4 CIDR range to include in this input security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-inputsecuritygroup-inputwhitelistrulecidr.html#cfn-medialive-inputsecuritygroup-inputwhitelistrulecidr-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

