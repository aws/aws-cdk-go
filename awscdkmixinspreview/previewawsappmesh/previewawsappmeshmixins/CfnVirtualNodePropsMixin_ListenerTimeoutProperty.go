package previewawsappmeshmixins


// An object that represents timeouts for different protocols.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   listenerTimeoutProperty := &ListenerTimeoutProperty{
//   	Grpc: &GrpcTimeoutProperty{
//   		Idle: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		PerRequest: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   	},
//   	Http: &HttpTimeoutProperty{
//   		Idle: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		PerRequest: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   	},
//   	Http2: &HttpTimeoutProperty{
//   		Idle: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		PerRequest: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   	},
//   	Tcp: &TcpTimeoutProperty{
//   		Idle: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertimeout.html
//
type CfnVirtualNodePropsMixin_ListenerTimeoutProperty struct {
	// An object that represents types of timeouts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertimeout.html#cfn-appmesh-virtualnode-listenertimeout-grpc
	//
	Grpc interface{} `field:"optional" json:"grpc" yaml:"grpc"`
	// An object that represents types of timeouts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertimeout.html#cfn-appmesh-virtualnode-listenertimeout-http
	//
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// An object that represents types of timeouts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertimeout.html#cfn-appmesh-virtualnode-listenertimeout-http2
	//
	Http2 interface{} `field:"optional" json:"http2" yaml:"http2"`
	// An object that represents types of timeouts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-listenertimeout.html#cfn-appmesh-virtualnode-listenertimeout-tcp
	//
	Tcp interface{} `field:"optional" json:"tcp" yaml:"tcp"`
}

