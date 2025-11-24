package mixinsawsappsync


// Properties for CfnDomainNameApiAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDomainNameApiAssociationMixinProps := &CfnDomainNameApiAssociationMixinProps{
//   	ApiId: jsii.String("apiId"),
//   	DomainName: jsii.String("domainName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainnameapiassociation.html
//
type CfnDomainNameApiAssociationMixinProps struct {
	// The API ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainnameapiassociation.html#cfn-appsync-domainnameapiassociation-apiid
	//
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// The domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainnameapiassociation.html#cfn-appsync-domainnameapiassociation-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
}

