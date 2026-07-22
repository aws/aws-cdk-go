package awscertificatemanager


// Properties for CfnAcmeDomainValidationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAcmeDomainValidationMixinProps := &CfnAcmeDomainValidationMixinProps{
//   	AcmeEndpointArn: jsii.String("acmeEndpointArn"),
//   	DomainName: jsii.String("domainName"),
//   	PrevalidationOptions: &PrevalidationOptionsProperty{
//   		DnsPrevalidation: &DnsPrevalidationOptionsProperty{
//   			DomainScope: &DomainScopeProperty{
//   				ExactDomain: jsii.String("exactDomain"),
//   				Subdomains: jsii.String("subdomains"),
//   				Wildcards: jsii.String("wildcards"),
//   			},
//   			HostedZoneId: jsii.String("hostedZoneId"),
//   		},
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmedomainvalidation.html
//
type CfnAcmeDomainValidationMixinProps struct {
	// The ARN of the ACME endpoint this domain validation is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmedomainvalidation.html#cfn-certificatemanager-acmedomainvalidation-acmeendpointarn
	//
	AcmeEndpointArn *string `field:"optional" json:"acmeEndpointArn" yaml:"acmeEndpointArn"`
	// The domain name to validate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmedomainvalidation.html#cfn-certificatemanager-acmedomainvalidation-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Prevalidation method configuration.
	//
	// Currently only DNS-based prevalidation is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmedomainvalidation.html#cfn-certificatemanager-acmedomainvalidation-prevalidationoptions
	//
	PrevalidationOptions interface{} `field:"optional" json:"prevalidationOptions" yaml:"prevalidationOptions"`
	// Tags associated with the domain validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-acmedomainvalidation.html#cfn-certificatemanager-acmedomainvalidation-tags
	//
	Tags *[]*CfnAcmeDomainValidationPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

