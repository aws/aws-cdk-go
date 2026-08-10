package awssso


// Details for an IAM Identity Center access scope associated with a resource server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   resourceServerScopeDetailsProperty := &ResourceServerScopeDetailsProperty{
//   	DetailedTitle: jsii.String("detailedTitle"),
//   	LongDescription: jsii.String("longDescription"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-applicationprovider-resourceserverscopedetails.html
//
type CfnApplicationProviderPropsMixin_ResourceServerScopeDetailsProperty struct {
	// The title of an access scope for a resource server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-applicationprovider-resourceserverscopedetails.html#cfn-sso-applicationprovider-resourceserverscopedetails-detailedtitle
	//
	DetailedTitle *string `field:"optional" json:"detailedTitle" yaml:"detailedTitle"`
	// The description of an access scope for a resource server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-applicationprovider-resourceserverscopedetails.html#cfn-sso-applicationprovider-resourceserverscopedetails-longdescription
	//
	LongDescription *string `field:"optional" json:"longDescription" yaml:"longDescription"`
}

