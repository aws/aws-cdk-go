package awsmediapackage


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingestEndpointProperty := &ingestEndpointProperty{
//   	id: jsii.String("id"),
//   	password: jsii.String("password"),
//   	url: jsii.String("url"),
//   	username: jsii.String("username"),
//   }
//
type CfnChannel_IngestEndpointProperty struct {
	// `CfnChannel.IngestEndpointProperty.Id`.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// `CfnChannel.IngestEndpointProperty.Password`.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// `CfnChannel.IngestEndpointProperty.Url`.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// `CfnChannel.IngestEndpointProperty.Username`.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

