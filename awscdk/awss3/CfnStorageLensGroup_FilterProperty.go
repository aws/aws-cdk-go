package awss3


// This resource sets the criteria for the Storage Lens group data that is displayed.
//
// For multiple filter conditions, the `AND` or `OR` logical operator is used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &FilterProperty{
//   	And: &AndProperty{
//   		MatchAnyPrefix: []*string{
//   			jsii.String("matchAnyPrefix"),
//   		},
//   		MatchAnySuffix: []*string{
//   			jsii.String("matchAnySuffix"),
//   		},
//   		MatchAnyTag: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MatchObjectAge: &MatchObjectAgeProperty{
//   			DaysGreaterThan: jsii.Number(123),
//   			DaysLessThan: jsii.Number(123),
//   		},
//   		MatchObjectSize: &MatchObjectSizeProperty{
//   			BytesGreaterThan: jsii.Number(123),
//   			BytesLessThan: jsii.Number(123),
//   		},
//   	},
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
//   	Or: &OrProperty{
//   		MatchAnyPrefix: []*string{
//   			jsii.String("matchAnyPrefix"),
//   		},
//   		MatchAnySuffix: []*string{
//   			jsii.String("matchAnySuffix"),
//   		},
//   		MatchAnyTag: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MatchObjectAge: &MatchObjectAgeProperty{
//   			DaysGreaterThan: jsii.Number(123),
//   			DaysLessThan: jsii.Number(123),
//   		},
//   		MatchObjectSize: &MatchObjectSizeProperty{
//   			BytesGreaterThan: jsii.Number(123),
//   			BytesLessThan: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-filter.html
//
type CfnStorageLensGroup_FilterProperty struct {
	// This property contains the `And` logical operator, which allows multiple filter conditions to be joined for more complex comparisons of Storage Lens group data.
	//
	// Objects must match all of the listed filter conditions that are joined by the `And` logical operator. Only one of each filter condition is allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-filter.html#cfn-s3-storagelensgroup-filter-and
	//
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// This property contains a list of prefixes.
	//
	// At least one prefix must be specified. Up to 10 prefixes are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-filter.html#cfn-s3-storagelensgroup-filter-matchanyprefix
	//
	MatchAnyPrefix *[]*string `field:"optional" json:"matchAnyPrefix" yaml:"matchAnyPrefix"`
	// This property contains a list of suffixes.
	//
	// At least one suffix must be specified. Up to 10 suffixes are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-filter.html#cfn-s3-storagelensgroup-filter-matchanysuffix
	//
	MatchAnySuffix *[]*string `field:"optional" json:"matchAnySuffix" yaml:"matchAnySuffix"`
	// This property contains the list of S3 object tags.
	//
	// At least one object tag must be specified. Up to 10 object tags are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-filter.html#cfn-s3-storagelensgroup-filter-matchanytag
	//
	MatchAnyTag interface{} `field:"optional" json:"matchAnyTag" yaml:"matchAnyTag"`
	// This property contains `DaysGreaterThan` and `DaysLessThan` to define the object age range (minimum and maximum number of days).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-filter.html#cfn-s3-storagelensgroup-filter-matchobjectage
	//
	MatchObjectAge interface{} `field:"optional" json:"matchObjectAge" yaml:"matchObjectAge"`
	// This property contains `BytesGreaterThan` and `BytesLessThan` to define the object size range (minimum and maximum number of Bytes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-filter.html#cfn-s3-storagelensgroup-filter-matchobjectsize
	//
	MatchObjectSize interface{} `field:"optional" json:"matchObjectSize" yaml:"matchObjectSize"`
	// This property contains the `Or` logical operator, which allows multiple filter conditions to be joined.
	//
	// Objects can match any of the listed filter conditions, which are joined by the `Or` logical operator. Only one of each filter condition is allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-filter.html#cfn-s3-storagelensgroup-filter-or
	//
	Or interface{} `field:"optional" json:"or" yaml:"or"`
}

