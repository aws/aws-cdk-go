package interfacesawscomprehend


// A reference to a DocumentClassifierEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentClassifierEndpointReference := &DocumentClassifierEndpointReference{
//   	DocumentClassifierEndpointArn: jsii.String("documentClassifierEndpointArn"),
//   }
//
type DocumentClassifierEndpointReference struct {
	// The Arn of the DocumentClassifierEndpoint resource.
	DocumentClassifierEndpointArn *string `field:"required" json:"documentClassifierEndpointArn" yaml:"documentClassifierEndpointArn"`
}

