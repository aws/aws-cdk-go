package mixinsawsrefactorspaces


// The configuration for the default route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   defaultRouteInputProperty := &DefaultRouteInputProperty{
//   	ActivationState: jsii.String("activationState"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-defaultrouteinput.html
//
type CfnRoutePropsMixin_DefaultRouteInputProperty struct {
	// If set to `ACTIVE` , traffic is forwarded to this route’s service after the route is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-refactorspaces-route-defaultrouteinput.html#cfn-refactorspaces-route-defaultrouteinput-activationstate
	//
	ActivationState *string `field:"optional" json:"activationState" yaml:"activationState"`
}

