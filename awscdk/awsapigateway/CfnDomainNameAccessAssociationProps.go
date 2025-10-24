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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnameaccessassociation.html
//
type CfnDomainNameAccessAssociationProps struct {
	// The identifier of the domain name access association source.
	//
	// For a `VPCE` , the value is the VPC endpoint ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnameaccessassociation.html#cfn-apigateway-domainnameaccessassociation-accessassociationsource
	//
	AccessAssociationSource *string `field:"required" json:"accessAssociationSource" yaml:"accessAssociationSource"`
	// The type of the domain name access association source.
	//
	// Only `VPCE` is currently supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnameaccessassociation.html#cfn-apigateway-domainnameaccessassociation-accessassociationsourcetype
	//
	AccessAssociationSourceType *string `field:"required" json:"accessAssociationSourceType" yaml:"accessAssociationSourceType"`
	// The ARN of the domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnameaccessassociation.html#cfn-apigateway-domainnameaccessassociation-domainnamearn
	//
	DomainNameArn *string `field:"required" json:"domainNameArn" yaml:"domainNameArn"`
	// The collection of tags.
	//
	// Each tag element is associated with a given resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnameaccessassociation.html#cfn-apigateway-domainnameaccessassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

