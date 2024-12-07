package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDomainNameAccessAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainNameAccessAssociationProps := &CfnDomainNameAccessAssociationProps{
//   	AccessAssociationSource: jsii.String("accessAssociationSource"),
//   	AccessAssociationSourceType: jsii.String("accessAssociationSourceType"),
//   	DomainNameArn: jsii.String("domainNameArn"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnameaccessassociation.html
//
type CfnDomainNameAccessAssociationProps struct {
	// The source of the domain name access association resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnameaccessassociation.html#cfn-apigateway-domainnameaccessassociation-accessassociationsource
	//
	AccessAssociationSource *string `field:"required" json:"accessAssociationSource" yaml:"accessAssociationSource"`
	// The source type of the domain name access association resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnameaccessassociation.html#cfn-apigateway-domainnameaccessassociation-accessassociationsourcetype
	//
	AccessAssociationSourceType *string `field:"required" json:"accessAssociationSourceType" yaml:"accessAssociationSourceType"`
	// The amazon resource name (ARN) of the domain name resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnameaccessassociation.html#cfn-apigateway-domainnameaccessassociation-domainnamearn
	//
	DomainNameArn *string `field:"required" json:"domainNameArn" yaml:"domainNameArn"`
	// An array of arbitrary tags (key-value pairs) to associate with the domainname access association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnameaccessassociation.html#cfn-apigateway-domainnameaccessassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

