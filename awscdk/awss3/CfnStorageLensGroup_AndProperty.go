package awss3


// The Storage Lens group will include objects that match all of the specified filter values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   andProperty := &AndProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-and.html
//
type CfnStorageLensGroup_AndProperty struct {
	// Filter to match any of the specified prefixes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-and.html#cfn-s3-storagelensgroup-and-matchanyprefix
	//
	MatchAnyPrefix *[]*string `field:"optional" json:"matchAnyPrefix" yaml:"matchAnyPrefix"`
	// Filter to match any of the specified suffixes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-and.html#cfn-s3-storagelensgroup-and-matchanysuffix
	//
	MatchAnySuffix *[]*string `field:"optional" json:"matchAnySuffix" yaml:"matchAnySuffix"`
	// Filter to match any of the specified object tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-and.html#cfn-s3-storagelensgroup-and-matchanytag
	//
	MatchAnyTag interface{} `field:"optional" json:"matchAnyTag" yaml:"matchAnyTag"`
	// Filter to match all of the specified values for the minimum and maximum object age.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-and.html#cfn-s3-storagelensgroup-and-matchobjectage
	//
	MatchObjectAge interface{} `field:"optional" json:"matchObjectAge" yaml:"matchObjectAge"`
	// Filter to match all of the specified values for the minimum and maximum object size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-and.html#cfn-s3-storagelensgroup-and-matchobjectsize
	//
	MatchObjectSize interface{} `field:"optional" json:"matchObjectSize" yaml:"matchObjectSize"`
}

