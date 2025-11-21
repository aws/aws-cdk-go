package awsappsync


// Properties for defining a `CfnDomainNameApiAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainNameApiAssociationProps := &CfnDomainNameApiAssociationProps{
//   	ApiId: jsii.String("apiId"),
//   	DomainName: jsii.String("domainName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainnameapiassociation.html
//
type CfnDomainNameApiAssociationProps struct {
	// The API ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainnameapiassociation.html#cfn-appsync-domainnameapiassociation-apiid
	//
	ApiId interface{} `field:"required" json:"apiId" yaml:"apiId"`
	// The domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainnameapiassociation.html#cfn-appsync-domainnameapiassociation-domainname
	//
	DomainName interface{} `field:"required" json:"domainName" yaml:"domainName"`
}

