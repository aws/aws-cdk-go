package awssso


// A structure that describes how IAM Identity Center represents the application provider in the portal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   displayDataProperty := &DisplayDataProperty{
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	IconUrl: jsii.String("iconUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-applicationprovider-displaydata.html
//
type CfnApplicationProviderPropsMixin_DisplayDataProperty struct {
	// The description of the application provider that appears in the portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-applicationprovider-displaydata.html#cfn-sso-applicationprovider-displaydata-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the application provider that appears in the portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-applicationprovider-displaydata.html#cfn-sso-applicationprovider-displaydata-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// A URL that points to an icon that represents the application provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-applicationprovider-displaydata.html#cfn-sso-applicationprovider-displaydata-iconurl
	//
	IconUrl *string `field:"optional" json:"iconUrl" yaml:"iconUrl"`
}

