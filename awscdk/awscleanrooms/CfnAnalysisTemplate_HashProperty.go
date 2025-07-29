package awscleanrooms


// Hash.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hashProperty := &HashProperty{
//   	Sha256: jsii.String("sha256"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-hash.html
//
type CfnAnalysisTemplate_HashProperty struct {
	// The SHA-256 hash value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-hash.html#cfn-cleanrooms-analysistemplate-hash-sha256
	//
	Sha256 *string `field:"optional" json:"sha256" yaml:"sha256"`
}

