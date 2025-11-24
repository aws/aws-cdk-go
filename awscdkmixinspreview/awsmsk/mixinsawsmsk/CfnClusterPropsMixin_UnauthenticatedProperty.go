package mixinsawsmsk


// Details for allowing no client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   unauthenticatedProperty := &UnauthenticatedProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-unauthenticated.html
//
type CfnClusterPropsMixin_UnauthenticatedProperty struct {
	// Unauthenticated is enabled or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-unauthenticated.html#cfn-msk-cluster-unauthenticated-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

