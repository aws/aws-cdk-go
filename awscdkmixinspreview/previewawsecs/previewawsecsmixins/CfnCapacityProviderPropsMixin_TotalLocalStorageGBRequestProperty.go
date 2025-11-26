package previewawsecsmixins


// The minimum and maximum total local storage in gigabytes (GB) for instance types with local storage.
//
// This is useful for workloads that require local storage for temporary data or caching.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   totalLocalStorageGBRequestProperty := &TotalLocalStorageGBRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-totallocalstoragegbrequest.html
//
type CfnCapacityProviderPropsMixin_TotalLocalStorageGBRequestProperty struct {
	// The maximum total local storage in GB.
	//
	// Instance types with more local storage are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-totallocalstoragegbrequest.html#cfn-ecs-capacityprovider-totallocalstoragegbrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum total local storage in GB.
	//
	// Instance types with less local storage are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-totallocalstoragegbrequest.html#cfn-ecs-capacityprovider-totallocalstoragegbrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

