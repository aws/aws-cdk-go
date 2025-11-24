package mixinsawslambda


// Code signing configuration [policies](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html#config-codesigning-policies) specify the validation failure action for signature mismatch or expiry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeSigningPoliciesProperty := &CodeSigningPoliciesProperty{
//   	UntrustedArtifactOnDeployment: jsii.String("untrustedArtifactOnDeployment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-codesigningconfig-codesigningpolicies.html
//
type CfnCodeSigningConfigPropsMixin_CodeSigningPoliciesProperty struct {
	// Code signing configuration policy for deployment validation failure.
	//
	// If you set the policy to `Enforce` , Lambda blocks the deployment request if signature validation checks fail. If you set the policy to `Warn` , Lambda allows the deployment and issues a new Amazon CloudWatch metric ( `SignatureValidationErrors` ) and also stores the warning in the CloudTrail log.
	//
	// Default value: `Warn`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-codesigningconfig-codesigningpolicies.html#cfn-lambda-codesigningconfig-codesigningpolicies-untrustedartifactondeployment
	//
	// Default: - "Warn".
	//
	UntrustedArtifactOnDeployment *string `field:"optional" json:"untrustedArtifactOnDeployment" yaml:"untrustedArtifactOnDeployment"`
}

