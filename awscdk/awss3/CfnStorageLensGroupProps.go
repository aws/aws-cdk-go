package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnStorageLensGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStorageLensGroupProps := &CfnStorageLensGroupProps{
//   	Filter: &FilterProperty{
//   		And: &AndProperty{
//   			MatchAnyPrefix: []*string{
//   				jsii.String("matchAnyPrefix"),
//   			},
//   			MatchAnySuffix: []*string{
//   				jsii.String("matchAnySuffix"),
//   			},
//   			MatchAnyTag: []interface{}{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MatchObjectAge: &MatchObjectAgeProperty{
//   				DaysGreaterThan: jsii.Number(123),
//   				DaysLessThan: jsii.Number(123),
//   			},
//   			MatchObjectSize: &MatchObjectSizeProperty{
//   				BytesGreaterThan: jsii.Number(123),
//   				BytesLessThan: jsii.Number(123),
//   			},
//   		},
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
//   		Or: &OrProperty{
//   			MatchAnyPrefix: []*string{
//   				jsii.String("matchAnyPrefix"),
//   			},
//   			MatchAnySuffix: []*string{
//   				jsii.String("matchAnySuffix"),
//   			},
//   			MatchAnyTag: []interface{}{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MatchObjectAge: &MatchObjectAgeProperty{
//   				DaysGreaterThan: jsii.Number(123),
//   				DaysLessThan: jsii.Number(123),
//   			},
//   			MatchObjectSize: &MatchObjectSizeProperty{
//   				BytesGreaterThan: jsii.Number(123),
//   				BytesLessThan: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelensgroup.html
//
type CfnStorageLensGroupProps struct {
	// This property contains the criteria for the Storage Lens group data that is displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelensgroup.html#cfn-s3-storagelensgroup-filter
	//
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
	// This property contains the Storage Lens group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelensgroup.html#cfn-s3-storagelensgroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// This property contains the AWS resource tags that you're adding to your Storage Lens group.
	//
	// This parameter is optional.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelensgroup.html#cfn-s3-storagelensgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

