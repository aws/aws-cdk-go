package awsapigateway


// A reference to a BasePathMappingV2 resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   basePathMappingV2Reference := &BasePathMappingV2Reference{
//   	BasePathMappingArn: jsii.String("basePathMappingArn"),
//   }
//
type BasePathMappingV2Reference struct {
	// The BasePathMappingArn of the BasePathMappingV2 resource.
	BasePathMappingArn *string `field:"required" json:"basePathMappingArn" yaml:"basePathMappingArn"`
}

