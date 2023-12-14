package awsapigatewayv2


// Options for DomainMapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var domainName domainName
//
//   domainMappingOptions := &DomainMappingOptions{
//   	DomainName: domainName,
//
//   	// the properties below are optional
//   	MappingKey: jsii.String("mappingKey"),
//   }
//
type DomainMappingOptions struct {
	// The domain name for the mapping.
	DomainName IDomainName `field:"required" json:"domainName" yaml:"domainName"`
	// The API mapping key.
	//
	// Leave it undefined for the root path mapping.
	// Default: - empty key for the root path mapping.
	//
	MappingKey *string `field:"optional" json:"mappingKey" yaml:"mappingKey"`
}

