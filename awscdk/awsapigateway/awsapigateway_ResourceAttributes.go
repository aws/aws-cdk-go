package awsapigateway


// Attributes that can be specified when importing a Resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var restApi restApi
//
//   resourceAttributes := &resourceAttributes{
//   	path: jsii.String("path"),
//   	resourceId: jsii.String("resourceId"),
//   	restApi: restApi,
//   }
//
type ResourceAttributes struct {
	// The full path of this resource.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The ID of the resource.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The rest API that this resource is part of.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
}

