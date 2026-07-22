package awscertificatemanager


// Properties for CfnAcmeExternalAccountBindingPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAcmeExternalAccountBindingMixinProps := &CfnAcmeExternalAccountBindingMixinProps{
//   	AcmeEndpointArn: jsii.String("acmeEndpointArn"),
//   	Expiration: &ExpirationProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeexternalaccountbinding.html
//
type CfnAcmeExternalAccountBindingMixinProps struct {
	// The ARN of the ACME endpoint this binding is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeexternalaccountbinding.html#cfn-certificatemanager-acmeexternalaccountbinding-acmeendpointarn
	//
	AcmeEndpointArn *string `field:"optional" json:"acmeEndpointArn" yaml:"acmeEndpointArn"`
	// The expiration configuration for the external account binding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeexternalaccountbinding.html#cfn-certificatemanager-acmeexternalaccountbinding-expiration
	//
	Expiration interface{} `field:"optional" json:"expiration" yaml:"expiration"`
	// The IAM role ARN for cross-account access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeexternalaccountbinding.html#cfn-certificatemanager-acmeexternalaccountbinding-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags associated with the external account binding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeexternalaccountbinding.html#cfn-certificatemanager-acmeexternalaccountbinding-tags
	//
	Tags *[]*CfnAcmeExternalAccountBindingPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

