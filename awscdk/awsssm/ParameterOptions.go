package awsssm


// Properties needed to create a new SSM Parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterOptions := &ParameterOptions{
//   	AllowedPattern: jsii.String("allowedPattern"),
//   	Description: jsii.String("description"),
//   	ParameterName: jsii.String("parameterName"),
//   	SimpleName: jsii.Boolean(false),
//   	Tier: awscdk.Aws_ssm.ParameterTier_ADVANCED,
//   }
//
type ParameterOptions struct {
	// A regular expression used to validate the parameter value.
	//
	// For example, for String types with values restricted to
	// numbers, you can specify the following: ``^\d+$``.
	// Default: no validation is performed.
	//
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// Information about the parameter that you want to add to the system.
	// Default: none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter.
	// Default: - a name will be generated by CloudFormation.
	//
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// Indicates whether the parameter name is a simple name.
	//
	// A parameter name
	// without any "/" is considered a simple name. If the parameter name includes
	// "/", setting simpleName to true might cause unintended issues such
	// as duplicate "/" in the resulting ARN.
	//
	// This is required only if `parameterName` is a token, which means we
	// are unable to detect if the name is simple or "path-like" for the purpose
	// of rendering SSM parameter ARNs.
	//
	// If `parameterName` is not specified, `simpleName` must be `true` (or
	// undefined) since the name generated by AWS CloudFormation is always a
	// simple name.
	// Default: - auto-detect based on `parameterName`.
	//
	SimpleName *bool `field:"optional" json:"simpleName" yaml:"simpleName"`
	// The tier of the string parameter.
	// Default: - undefined.
	//
	Tier ParameterTier `field:"optional" json:"tier" yaml:"tier"`
}

