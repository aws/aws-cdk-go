package awslambda


// Code signing configuration [policies](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html#config-codesigning-policies) specify the validation failure action for signature mismatch or expiry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeSigningPoliciesProperty := &codeSigningPoliciesProperty{
//   	untrustedArtifactOnDeployment: jsii.String("untrustedArtifactOnDeployment"),
//   }
//
type CfnCodeSigningConfig_CodeSigningPoliciesProperty struct {
	// Code signing configuration policy for deployment validation failure.
	//
	// If you set the policy to `Enforce` , Lambda blocks the deployment request if signature validation checks fail. If you set the policy to `Warn` , Lambda allows the deployment and creates a CloudWatch log.
	//
	// Default value: `Warn`.
	UntrustedArtifactOnDeployment *string `field:"required" json:"untrustedArtifactOnDeployment" yaml:"untrustedArtifactOnDeployment"`
}

