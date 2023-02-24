package awsinspectorv2


// Contains information on the details of a package filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   packageFilterProperty := &PackageFilterProperty{
//   	Architecture: &StringFilterProperty{
//   		Comparison: jsii.String("comparison"),
//   		Value: jsii.String("value"),
//   	},
//   	Epoch: &NumberFilterProperty{
//   		LowerInclusive: jsii.Number(123),
//   		UpperInclusive: jsii.Number(123),
//   	},
//   	Name: &StringFilterProperty{
//   		Comparison: jsii.String("comparison"),
//   		Value: jsii.String("value"),
//   	},
//   	Release: &StringFilterProperty{
//   		Comparison: jsii.String("comparison"),
//   		Value: jsii.String("value"),
//   	},
//   	SourceLayerHash: &StringFilterProperty{
//   		Comparison: jsii.String("comparison"),
//   		Value: jsii.String("value"),
//   	},
//   	Version: &StringFilterProperty{
//   		Comparison: jsii.String("comparison"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
type CfnFilter_PackageFilterProperty struct {
	// An object that contains details on the package architecture type to filter on.
	Architecture interface{} `field:"optional" json:"architecture" yaml:"architecture"`
	// An object that contains details on the package epoch to filter on.
	Epoch interface{} `field:"optional" json:"epoch" yaml:"epoch"`
	// An object that contains details on the name of the package to filter on.
	Name interface{} `field:"optional" json:"name" yaml:"name"`
	// An object that contains details on the package release to filter on.
	Release interface{} `field:"optional" json:"release" yaml:"release"`
	// An object that contains details on the source layer hash to filter on.
	SourceLayerHash interface{} `field:"optional" json:"sourceLayerHash" yaml:"sourceLayerHash"`
	// The package version to filter on.
	Version interface{} `field:"optional" json:"version" yaml:"version"`
}

