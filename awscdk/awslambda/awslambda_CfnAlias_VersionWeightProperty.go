package awslambda


// The [traffic-shifting](https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html) configuration of a Lambda function alias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   versionWeightProperty := &versionWeightProperty{
//   	functionVersion: jsii.String("functionVersion"),
//   	functionWeight: jsii.Number(123),
//   }
//
type CfnAlias_VersionWeightProperty struct {
	// The qualifier of the second version.
	FunctionVersion *string `field:"required" json:"functionVersion" yaml:"functionVersion"`
	// The percentage of traffic that the alias routes to the second version.
	FunctionWeight *float64 `field:"required" json:"functionWeight" yaml:"functionWeight"`
}

