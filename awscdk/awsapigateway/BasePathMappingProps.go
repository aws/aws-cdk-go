package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var domainName domainName
//   var restApi restApi
//   var stage stage
//
//   basePathMappingProps := &BasePathMappingProps{
//   	DomainName: domainName,
//   	RestApi: restApi,
//
//   	// the properties below are optional
//   	BasePath: jsii.String("basePath"),
//   	Stage: stage,
//   }
//
// Experimental.
type BasePathMappingProps struct {
	// The base path name that callers of the API must provide in the URL after the domain name (e.g. `example.com/base-path`). If you specify this property, it can't be an empty string.
	// Experimental.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The Deployment stage of API [disable-awslint:ref-via-interface].
	// Experimental.
	Stage Stage `field:"optional" json:"stage" yaml:"stage"`
	// The DomainName to associate with this base path mapping.
	// Experimental.
	DomainName IDomainName `field:"required" json:"domainName" yaml:"domainName"`
	// The RestApi resource to target.
	// Experimental.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
}

