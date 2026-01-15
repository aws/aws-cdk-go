package previewawsssomixins


// Properties for CfnInstanceAccessControlAttributeConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceAccessControlAttributeConfigurationMixinProps := &CfnInstanceAccessControlAttributeConfigurationMixinProps{
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
//   	InstanceAccessControlAttributeConfiguration: &InstanceAccessControlAttributeConfigurationProperty{
//   		AccessControlAttributes: []interface{}{
//   			&AccessControlAttributeProperty{
//   				Key: jsii.String("key"),
//   				Value: &AccessControlAttributeValueProperty{
//   					Source: []*string{
//   						jsii.String("source"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	InstanceArn: jsii.String("instanceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html
//
type CfnInstanceAccessControlAttributeConfigurationMixinProps struct {
	// Lists the attributes that are configured for ABAC in the specified  instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributes
	//
	AccessControlAttributes interface{} `field:"optional" json:"accessControlAttributes" yaml:"accessControlAttributes"`
	// The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes.
	//
	// We recomend that you use  AccessControlAttributes property instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-instanceaccesscontrolattributeconfiguration
	//
	// Deprecated: this property has been deprecated.
	InstanceAccessControlAttributeConfiguration interface{} `field:"optional" json:"instanceAccessControlAttributeConfiguration" yaml:"instanceAccessControlAttributeConfiguration"`
	// The ARN of the  instance under which the operation will be executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
}

