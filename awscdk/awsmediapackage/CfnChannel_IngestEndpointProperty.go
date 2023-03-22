package awsmediapackage


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingestEndpointProperty := &IngestEndpointProperty{
//   	Id: jsii.String("id"),
//   	Password: jsii.String("password"),
//   	Url: jsii.String("url"),
//   	Username: jsii.String("username"),
//   }
//
type CfnChannel_IngestEndpointProperty struct {
	// `CfnChannel.IngestEndpointProperty.Id`.
	Id *string `field:"required" json:"id" yaml:"id"`
	// `CfnChannel.IngestEndpointProperty.Password`.
	Password *string `field:"required" json:"password" yaml:"password"`
	// `CfnChannel.IngestEndpointProperty.Url`.
	Url *string `field:"required" json:"url" yaml:"url"`
	// `CfnChannel.IngestEndpointProperty.Username`.
	Username *string `field:"required" json:"username" yaml:"username"`
}

