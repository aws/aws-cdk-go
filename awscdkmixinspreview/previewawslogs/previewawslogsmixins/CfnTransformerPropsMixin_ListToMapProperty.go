package previewawslogsmixins


// This processor takes a list of objects that contain key fields, and converts them into a map of target keys.
//
// For more information about this processor including examples, see [listToMap](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-listToMap) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   listToMapProperty := &ListToMapProperty{
//   	Flatten: jsii.Boolean(false),
//   	FlattenedElement: jsii.String("flattenedElement"),
//   	Key: jsii.String("key"),
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//   	ValueKey: jsii.String("valueKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-listtomap.html
//
type CfnTransformerPropsMixin_ListToMapProperty struct {
	// A Boolean value to indicate whether the list will be flattened into single items.
	//
	// Specify `true` to flatten the list. The default is `false`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-listtomap.html#cfn-logs-transformer-listtomap-flatten
	//
	Flatten interface{} `field:"optional" json:"flatten" yaml:"flatten"`
	// If you set `flatten` to `true` , use `flattenedElement` to specify which element, `first` or `last` , to keep.
	//
	// You must specify this parameter if `flatten` is `true`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-listtomap.html#cfn-logs-transformer-listtomap-flattenedelement
	//
	FlattenedElement *string `field:"optional" json:"flattenedElement" yaml:"flattenedElement"`
	// The key of the field to be extracted as keys in the generated map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-listtomap.html#cfn-logs-transformer-listtomap-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The key in the log event that has a list of objects that will be converted to a map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-listtomap.html#cfn-logs-transformer-listtomap-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The key of the field that will hold the generated map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-listtomap.html#cfn-logs-transformer-listtomap-target
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
	// If this is specified, the values that you specify in this parameter will be extracted from the `source` objects and put into the values of the generated map.
	//
	// Otherwise, original objects in the source list will be put into the values of the generated map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-listtomap.html#cfn-logs-transformer-listtomap-valuekey
	//
	ValueKey *string `field:"optional" json:"valueKey" yaml:"valueKey"`
}

