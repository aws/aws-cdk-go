package awslambda


// The [traffic-shifting](https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html) configuration of a Lambda function alias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   versionWeightProperty := &VersionWeightProperty{
//   	FunctionVersion: jsii.String("functionVersion"),
//   	FunctionWeight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-versionweight.html
//
type CfnAlias_VersionWeightProperty struct {
	// The qualifier of the second version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-versionweight.html#cfn-lambda-alias-versionweight-functionversion
	//
	FunctionVersion *string `field:"required" json:"functionVersion" yaml:"functionVersion"`
	// The percentage of traffic that the alias routes to the second version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-versionweight.html#cfn-lambda-alias-versionweight-functionweight
	//
	FunctionWeight *float64 `field:"required" json:"functionWeight" yaml:"functionWeight"`
}

