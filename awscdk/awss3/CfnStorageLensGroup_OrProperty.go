package awss3


// This resource contains the `Or` logical operator, which allows multiple filter conditions to be joined for more complex comparisons of Storage Lens group data.
//
// Objects can match any of the listed filter conditions that are joined by the `Or` logical operator. Only one of each filter condition is allowed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   orProperty := &OrProperty{
//   	MatchAnyPrefix: []*string{
//   		jsii.String("matchAnyPrefix"),
//   	},
//   	MatchAnySuffix: []*string{
//   		jsii.String("matchAnySuffix"),
//   	},
//   	MatchAnyTag: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MatchObjectAge: &MatchObjectAgeProperty{
//   		DaysGreaterThan: jsii.Number(123),
//   		DaysLessThan: jsii.Number(123),
//   	},
//   	MatchObjectSize: &MatchObjectSizeProperty{
//   		BytesGreaterThan: jsii.Number(123),
//   		BytesLessThan: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-or.html
//
type CfnStorageLensGroup_OrProperty struct {
	// This property contains a list of prefixes.
	//
	// At least one prefix must be specified. Up to 10 prefixes are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-or.html#cfn-s3-storagelensgroup-or-matchanyprefix
	//
	MatchAnyPrefix *[]*string `field:"optional" json:"matchAnyPrefix" yaml:"matchAnyPrefix"`
	// This property contains the list of suffixes.
	//
	// At least one suffix must be specified. Up to 10 suffixes are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-or.html#cfn-s3-storagelensgroup-or-matchanysuffix
	//
	MatchAnySuffix *[]*string `field:"optional" json:"matchAnySuffix" yaml:"matchAnySuffix"`
	// This property contains the list of S3 object tags.
	//
	// At least one object tag must be specified. Up to 10 object tags are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-or.html#cfn-s3-storagelensgroup-or-matchanytag
	//
	MatchAnyTag interface{} `field:"optional" json:"matchAnyTag" yaml:"matchAnyTag"`
	// This property filters objects that match the specified object age range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-or.html#cfn-s3-storagelensgroup-or-matchobjectage
	//
	MatchObjectAge interface{} `field:"optional" json:"matchObjectAge" yaml:"matchObjectAge"`
	// This property contains the `BytesGreaterThan` and `BytesLessThan` values to define the object size range (minimum and maximum number of Bytes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-or.html#cfn-s3-storagelensgroup-or-matchobjectsize
	//
	MatchObjectSize interface{} `field:"optional" json:"matchObjectSize" yaml:"matchObjectSize"`
}

