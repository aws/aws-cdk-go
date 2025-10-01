package awspcaconnectorscep


// A reference to a Challenge resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   challengeReference := &ChallengeReference{
//   	ChallengeArn: jsii.String("challengeArn"),
//   }
//
type ChallengeReference struct {
	// The ChallengeArn of the Challenge resource.
	ChallengeArn *string `field:"required" json:"challengeArn" yaml:"challengeArn"`
}

