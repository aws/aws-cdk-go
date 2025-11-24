package mixinsawsappmesh


// An object representing the TCP route to match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tcpRouteMatchProperty := &TcpRouteMatchProperty{
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-tcproutematch.html
//
type CfnRoutePropsMixin_TcpRouteMatchProperty struct {
	// The port number to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-tcproutematch.html#cfn-appmesh-route-tcproutematch-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

