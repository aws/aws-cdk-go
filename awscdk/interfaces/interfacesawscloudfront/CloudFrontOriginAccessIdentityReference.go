package interfacesawscloudfront


// A reference to a CloudFrontOriginAccessIdentity resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudFrontOriginAccessIdentityReference := &CloudFrontOriginAccessIdentityReference{
//   	CloudFrontOriginAccessIdentityId: jsii.String("cloudFrontOriginAccessIdentityId"),
//   }
//
type CloudFrontOriginAccessIdentityReference struct {
	// The Id of the CloudFrontOriginAccessIdentity resource.
	CloudFrontOriginAccessIdentityId *string `field:"required" json:"cloudFrontOriginAccessIdentityId" yaml:"cloudFrontOriginAccessIdentityId"`
}

