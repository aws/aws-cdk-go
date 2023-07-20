package awssso


// The value used for mapping a specified attribute to an identity source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessControlAttributeValueProperty := &AccessControlAttributeValueProperty{
//   	Source: []*string{
//   		jsii.String("source"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributevalue.html
//
type CfnInstanceAccessControlAttributeConfiguration_AccessControlAttributeValueProperty struct {
	// The identity source to use when mapping a specified attribute to IAM Identity Center .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributevalue.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributevalue-source
	//
	Source *[]*string `field:"required" json:"source" yaml:"source"`
}

