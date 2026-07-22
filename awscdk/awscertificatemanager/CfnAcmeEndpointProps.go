package awscertificatemanager


// Properties for defining a `CfnAcmeEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAcmeEndpointProps := &CfnAcmeEndpointProps{
//   	AuthorizationBehavior: jsii.String("authorizationBehavior"),
//   	CertificateAuthority: &CertificateAuthorityProperty{
//   		PublicCertificateAuthority: &PublicCertificateAuthorityProperty{
//   			AllowedKeyAlgorithms: []*string{
//   				jsii.String("allowedKeyAlgorithms"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	CertificateTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Contact: jsii.String("contact"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeendpoint.html
//
type CfnAcmeEndpointProps struct {
	// The authorization behavior for the ACME endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeendpoint.html#cfn-certificatemanager-acmeendpoint-authorizationbehavior
	//
	AuthorizationBehavior *string `field:"required" json:"authorizationBehavior" yaml:"authorizationBehavior"`
	// The certificate authority configuration for the ACME endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeendpoint.html#cfn-certificatemanager-acmeendpoint-certificateauthority
	//
	CertificateAuthority interface{} `field:"required" json:"certificateAuthority" yaml:"certificateAuthority"`
	// Tags applied to certificates issued via this endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeendpoint.html#cfn-certificatemanager-acmeendpoint-certificatetags
	//
	CertificateTags interface{} `field:"optional" json:"certificateTags" yaml:"certificateTags"`
	// Whether contact information is required for the ACME endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeendpoint.html#cfn-certificatemanager-acmeendpoint-contact
	//
	Contact *string `field:"optional" json:"contact" yaml:"contact"`
	// Tags associated with the ACME endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmeendpoint.html#cfn-certificatemanager-acmeendpoint-tags
	//
	Tags *[]*CfnAcmeEndpoint_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

