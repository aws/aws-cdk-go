package awslambda


// List of signing profiles that can sign a code package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allowedPublishersProperty := &AllowedPublishersProperty{
//   	SigningProfileVersionArns: []*string{
//   		jsii.String("signingProfileVersionArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-codesigningconfig-allowedpublishers.html
//
type CfnCodeSigningConfig_AllowedPublishersProperty struct {
	// The Amazon Resource Name (ARN) for each of the signing profiles.
	//
	// A signing profile defines a trusted user who can sign a code package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-codesigningconfig-allowedpublishers.html#cfn-lambda-codesigningconfig-allowedpublishers-signingprofileversionarns
	//
	SigningProfileVersionArns *[]*string `field:"required" json:"signingProfileVersionArns" yaml:"signingProfileVersionArns"`
}

