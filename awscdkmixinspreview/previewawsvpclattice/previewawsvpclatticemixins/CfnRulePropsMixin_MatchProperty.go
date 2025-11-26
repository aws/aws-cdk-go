package previewawsvpclatticemixins


// Describes a rule match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   matchProperty := &MatchProperty{
//   	HttpMatch: &HttpMatchProperty{
//   		HeaderMatches: []interface{}{
//   			&HeaderMatchProperty{
//   				CaseSensitive: jsii.Boolean(false),
//   				Match: &HeaderMatchTypeProperty{
//   					Contains: jsii.String("contains"),
//   					Exact: jsii.String("exact"),
//   					Prefix: jsii.String("prefix"),
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Method: jsii.String("method"),
//   		PathMatch: &PathMatchProperty{
//   			CaseSensitive: jsii.Boolean(false),
//   			Match: &PathMatchTypeProperty{
//   				Exact: jsii.String("exact"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-match.html
//
type CfnRulePropsMixin_MatchProperty struct {
	// The HTTP criteria that a rule must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-match.html#cfn-vpclattice-rule-match-httpmatch
	//
	HttpMatch interface{} `field:"optional" json:"httpMatch" yaml:"httpMatch"`
}

