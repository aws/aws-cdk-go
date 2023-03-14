package awslambda


// List of signing profiles that can sign a code package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allowedPublishersProperty := &allowedPublishersProperty{
//   	signingProfileVersionArns: []*string{
//   		jsii.String("signingProfileVersionArns"),
//   	},
//   }
//
type CfnCodeSigningConfig_AllowedPublishersProperty struct {
	// The Amazon Resource Name (ARN) for each of the signing profiles.
	//
	// A signing profile defines a trusted user who can sign a code package.
	SigningProfileVersionArns *[]*string `field:"required" json:"signingProfileVersionArns" yaml:"signingProfileVersionArns"`
}

