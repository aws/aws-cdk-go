package awsappsync


// Properties for an AppSync Event API.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var handler Function
//
//
//   iamProvider := &AppSyncAuthProvider{
//   	AuthorizationType: appsync.AppSyncAuthorizationType_IAM,
//   }
//
//   apiKeyProvider := &AppSyncAuthProvider{
//   	AuthorizationType: appsync.AppSyncAuthorizationType_API_KEY,
//   }
//
//   /* API with IAM and API Key providers.
//    * Connection, default publish and default subscribe
//    * can be done with either IAM and API Key.
//    */
//   api := appsync.NewEventApi(this, jsii.String("api"), &EventApiProps{
//   	ApiName: jsii.String("api"),
//   	AuthorizationConfig: &EventApiAuthConfig{
//   		// set auth providers
//   		AuthProviders: []AppSyncAuthProvider{
//   			iamProvider,
//   			apiKeyProvider,
//   		},
//   	},
//   })
//
//   api.AddChannelNamespace(jsii.String("default"))
//
type EventApiProps struct {
	// the name of the Event API.
	ApiName *string `field:"required" json:"apiName" yaml:"apiName"`
	// Optional authorization configuration.
	// Default: - API Key authorization.
	//
	AuthorizationConfig *EventApiAuthConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The domain name configuration for the Event API.
	//
	// The Route 53 hosted zone and CName DNS record must be configured in addition to this setting to
	// enable custom domain URL.
	// Default: - no domain name.
	//
	DomainName *AppSyncDomainOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Logging configuration for this api.
	// Default: - None.
	//
	LogConfig *AppSyncLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// The owner contact information for an API resource.
	//
	// This field accepts any string input with a length of 0 - 256 characters.
	// Default: - No owner contact.
	//
	OwnerContact *string `field:"optional" json:"ownerContact" yaml:"ownerContact"`
}

