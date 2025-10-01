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
//   	FilePath: &StringFilterProperty{
//   		Comparison: jsii.String("comparison"),
//   		Value: jsii.String("value"),
//   	},
//   	Name: &StringFilterProperty{
//   		Comparison: jsii.String("comparison"),
//   		Value: jsii.String("value"),
//   	},
//   	Release: &StringFilterProperty{
//   		Comparison: jsii.String("comparison"),
//   		Value: jsii.String("value"),
//   	},
//   	SourceLambdaLayerArn: &StringFilterProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-packagefilter.html
//
type CfnFilter_PackageFilterProperty struct {
	// An object that contains details on the package architecture type to filter on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-packagefilter.html#cfn-inspectorv2-filter-packagefilter-architecture
	//
	Architecture interface{} `field:"optional" json:"architecture" yaml:"architecture"`
	// An object that contains details on the package epoch to filter on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-packagefilter.html#cfn-inspectorv2-filter-packagefilter-epoch
	//
	Epoch interface{} `field:"optional" json:"epoch" yaml:"epoch"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-packagefilter.html#cfn-inspectorv2-filter-packagefilter-filepath
	//
	FilePath interface{} `field:"optional" json:"filePath" yaml:"filePath"`
	// An object that contains details on the name of the package to filter on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-packagefilter.html#cfn-inspectorv2-filter-packagefilter-name
	//
	Name interface{} `field:"optional" json:"name" yaml:"name"`
	// An object that contains details on the package release to filter on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-packagefilter.html#cfn-inspectorv2-filter-packagefilter-release
	//
	Release interface{} `field:"optional" json:"release" yaml:"release"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-packagefilter.html#cfn-inspectorv2-filter-packagefilter-sourcelambdalayerarn
	//
	SourceLambdaLayerArn interface{} `field:"optional" json:"sourceLambdaLayerArn" yaml:"sourceLambdaLayerArn"`
	// An object that contains details on the source layer hash to filter on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-packagefilter.html#cfn-inspectorv2-filter-packagefilter-sourcelayerhash
	//
	SourceLayerHash interface{} `field:"optional" json:"sourceLayerHash" yaml:"sourceLayerHash"`
	// The package version to filter on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-packagefilter.html#cfn-inspectorv2-filter-packagefilter-version
	//
	Version interface{} `field:"optional" json:"version" yaml:"version"`
}

