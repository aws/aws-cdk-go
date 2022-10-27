package awsapigateway


// Example:
//   var api restApi
//
//
//   domainName := apigateway.domainName.fromDomainNameAttributes(this, jsii.String("DomainName"), &domainNameAttributes{
//   	domainName: jsii.String("domainName"),
//   	domainNameAliasHostedZoneId: jsii.String("domainNameAliasHostedZoneId"),
//   	domainNameAliasTarget: jsii.String("domainNameAliasTarget"),
//   })
//
//   apigateway.NewBasePathMapping(this, jsii.String("BasePathMapping"), &basePathMappingProps{
//   	domainName: domainName,
//   	restApi: api,
//   })
//
type BasePathMappingProps struct {
	// Whether to attach the base path mapping to a stage.
	//
	// Use this property to create a base path mapping without attaching it to the Rest API default stage.
	// This property is ignored if `stage` is provided.
	AttachToStage *bool `field:"optional" json:"attachToStage" yaml:"attachToStage"`
	// The base path name that callers of the API must provide in the URL after the domain name (e.g. `example.com/base-path`). If you specify this property, it can't be an empty string.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The Deployment stage of API [disable-awslint:ref-via-interface].
	Stage Stage `field:"optional" json:"stage" yaml:"stage"`
	// The DomainName to associate with this base path mapping.
	DomainName IDomainName `field:"required" json:"domainName" yaml:"domainName"`
	// The RestApi resource to target.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
}

