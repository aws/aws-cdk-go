package interfacesawsroute53globalresolver


// A reference to a AccessToken resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessTokenReference := &AccessTokenReference{
//   	AccessTokenArn: jsii.String("accessTokenArn"),
//   	AccessTokenId: jsii.String("accessTokenId"),
//   }
//
type AccessTokenReference struct {
	// The ARN of the AccessToken resource.
	AccessTokenArn *string `field:"required" json:"accessTokenArn" yaml:"accessTokenArn"`
	// The AccessTokenId of the AccessToken resource.
	AccessTokenId *string `field:"required" json:"accessTokenId" yaml:"accessTokenId"`
}

