package mixinsawssso


// The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes.
//
// We recomend that you use  AccessControlAttributes property instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instanceAccessControlAttributeConfigurationProperty := &InstanceAccessControlAttributeConfigurationProperty{
//   	AccessControlAttributes: []interface{}{
//   		&AccessControlAttributeProperty{
//   			Key: jsii.String("key"),
//   			Value: &AccessControlAttributeValueProperty{
//   				Source: []*string{
//   					jsii.String("source"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-instanceaccesscontrolattributeconfiguration.html
//
type CfnInstanceAccessControlAttributeConfigurationPropsMixin_InstanceAccessControlAttributeConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-instanceaccesscontrolattributeconfiguration-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-instanceaccesscontrolattributeconfiguration-accesscontrolattributes
	//
	AccessControlAttributes interface{} `field:"optional" json:"accessControlAttributes" yaml:"accessControlAttributes"`
}

