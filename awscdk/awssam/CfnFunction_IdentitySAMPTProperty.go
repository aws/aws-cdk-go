package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identitySAMPTProperty := &IdentitySAMPTProperty{
//   	IdentityName: jsii.String("identityName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-identitysampt.html
//
type CfnFunction_IdentitySAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-identitysampt.html#cfn-serverless-function-identitysampt-identityname
	//
	IdentityName *string `field:"required" json:"identityName" yaml:"identityName"`
}

