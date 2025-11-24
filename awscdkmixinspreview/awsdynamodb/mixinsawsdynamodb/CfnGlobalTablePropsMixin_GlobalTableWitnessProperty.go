package mixinsawsdynamodb


// The witness Region for the MRSC global table.
//
// A MRSC global table can be configured with either three replicas, or with two replicas and one witness.
//
// The witness must be in a different Region than the replicas and within the same Region set:
//
// - US Region set: US East (N. Virginia), US East (Ohio), US West (Oregon)
// - EU Region set: Europe (Ireland), Europe (London), Europe (Paris), Europe (Frankfurt)
// - AP Region set: Asia Pacific (Tokyo), Asia Pacific (Seoul), Asia Pacific (Osaka).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   globalTableWitnessProperty := &GlobalTableWitnessProperty{
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globaltablewitness.html
//
type CfnGlobalTablePropsMixin_GlobalTableWitnessProperty struct {
	// The name of the AWS Region that serves as a witness for the MRSC global table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globaltablewitness.html#cfn-dynamodb-globaltable-globaltablewitness-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

