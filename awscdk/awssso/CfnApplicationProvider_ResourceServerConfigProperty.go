package awssso


// A structure that describes the configuration of a resource server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceServerConfigProperty := &ResourceServerConfigProperty{
//   	Scopes: map[string]interface{}{
//   		"scopesKey": &ResourceServerScopeDetailsProperty{
//   			"detailedTitle": jsii.String("detailedTitle"),
//   			"longDescription": jsii.String("longDescription"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-applicationprovider-resourceserverconfig.html
//
type CfnApplicationProvider_ResourceServerConfigProperty struct {
	// A map of IAM Identity Center access scopes associated with this resource server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-applicationprovider-resourceserverconfig.html#cfn-sso-applicationprovider-resourceserverconfig-scopes
	//
	Scopes interface{} `field:"optional" json:"scopes" yaml:"scopes"`
}

