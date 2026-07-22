package awscertificatemanager


// Properties for defining a `CfnAcmeExternalAccountBinding`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAcmeExternalAccountBindingProps := &CfnAcmeExternalAccountBindingProps{
//   	AcmeEndpointArn: jsii.String("acmeEndpointArn"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Expiration: &ExpirationProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
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
type CfnAcmeExternalAccountBindingProps struct {
	// The ARN of the ACME endpoint this binding is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeexternalaccountbinding.html#cfn-certificatemanager-acmeexternalaccountbinding-acmeendpointarn
	//
	AcmeEndpointArn *string `field:"required" json:"acmeEndpointArn" yaml:"acmeEndpointArn"`
	// The IAM role ARN for cross-account access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeexternalaccountbinding.html#cfn-certificatemanager-acmeexternalaccountbinding-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The expiration configuration for the external account binding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeexternalaccountbinding.html#cfn-certificatemanager-acmeexternalaccountbinding-expiration
	//
	Expiration interface{} `field:"optional" json:"expiration" yaml:"expiration"`
	// Tags associated with the external account binding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeexternalaccountbinding.html#cfn-certificatemanager-acmeexternalaccountbinding-tags
	//
	Tags *[]*CfnAcmeExternalAccountBinding_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

