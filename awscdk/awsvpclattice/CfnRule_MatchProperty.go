package awsvpclattice


// Describes a rule match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchProperty := &MatchProperty{
//   	HttpMatch: &HttpMatchProperty{
//   		HeaderMatches: []interface{}{
//   			&HeaderMatchProperty{
//   				Match: &HeaderMatchTypeProperty{
//   					Contains: jsii.String("contains"),
//   					Exact: jsii.String("exact"),
//   					Prefix: jsii.String("prefix"),
//   				},
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				CaseSensitive: jsii.Boolean(false),
//   			},
//   		},
//   		Method: jsii.String("method"),
//   		PathMatch: &PathMatchProperty{
//   			Match: &PathMatchTypeProperty{
//   				Exact: jsii.String("exact"),
//   				Prefix: jsii.String("prefix"),
//   			},
//
//   			// the properties below are optional
//   			CaseSensitive: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-match.html
//
type CfnRule_MatchProperty struct {
	// The HTTP criteria that a rule must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-match.html#cfn-vpclattice-rule-match-httpmatch
	//
	HttpMatch interface{} `field:"required" json:"httpMatch" yaml:"httpMatch"`
}

