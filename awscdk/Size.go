package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the amount of digital storage.
//
// The amount can be specified either as a literal value (e.g: `10`) which
// cannot be negative, or as an unresolved number token.
//
// When the amount is passed as a token, unit conversion is not possible.
//
// Example:
//   var bucket bucket
//   // Provide a Lambda function that will transform records before delivery, with custom
//   // buffering and retry configuration
//   lambdaFunction := lambda.NewFunction(this, jsii.String("Processor"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("process-records"))),
//   })
//   lambdaProcessor := firehose.NewLambdaFunctionProcessor(lambdaFunction, &DataProcessorProps{
//   	BufferInterval: awscdk.Duration_Minutes(jsii.Number(5)),
//   	BufferSize: awscdk.Size_Mebibytes(jsii.Number(5)),
//   	Retries: jsii.Number(5),
//   })
//   s3Destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
//   	Processor: lambdaProcessor,
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destinations: []iDestination{
//   		s3Destination,
//   	},
//   })
//
type Size interface {
	// Checks if size is a token or a resolvable object.
	IsUnresolved() *bool
	// Return this storage as a total number of bytes.
	//
	// Returns: the quantity expressed in bytes.
	ToBytes(opts *SizeConversionOptions) *float64
	// Return this storage as a total number of gibibytes.
	//
	// Returns: the quantity of bytes expressed in gibibytes.
	ToGibibytes(opts *SizeConversionOptions) *float64
	// Return this storage as a total number of kibibytes.
	//
	// Returns: the quantity of bytes expressed in kibibytes.
	ToKibibytes(opts *SizeConversionOptions) *float64
	// Return this storage as a total number of mebibytes.
	//
	// Returns: the quantity of bytes expressed in mebibytes.
	ToMebibytes(opts *SizeConversionOptions) *float64
	// Return this storage as a total number of pebibytes.
	//
	// Returns: the quantity of bytes expressed in pebibytes.
	ToPebibytes(opts *SizeConversionOptions) *float64
	// Return this storage as a total number of tebibytes.
	//
	// Returns: the quantity of bytes expressed in tebibytes.
	ToTebibytes(opts *SizeConversionOptions) *float64
}

// The jsii proxy struct for Size
type jsiiProxy_Size struct {
	_ byte // padding
}

// Create a Storage representing an amount bytes.
//
// Returns: a new `Size` instance.
func Size_Bytes(amount *float64) Size {
	_init_.Initialize()

	if err := validateSize_BytesParameters(amount); err != nil {
		panic(err)
	}
	var returns Size

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Size",
		"bytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount gibibytes.
//
// 1 GiB = 1024 MiB.
//
// Returns: a new `Size` instance.
func Size_Gibibytes(amount *float64) Size {
	_init_.Initialize()

	if err := validateSize_GibibytesParameters(amount); err != nil {
		panic(err)
	}
	var returns Size

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Size",
		"gibibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount kibibytes.
//
// 1 KiB = 1024 bytes.
//
// Returns: a new `Size` instance.
func Size_Kibibytes(amount *float64) Size {
	_init_.Initialize()

	if err := validateSize_KibibytesParameters(amount); err != nil {
		panic(err)
	}
	var returns Size

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Size",
		"kibibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount mebibytes.
//
// 1 MiB = 1024 KiB.
//
// Returns: a new `Size` instance.
func Size_Mebibytes(amount *float64) Size {
	_init_.Initialize()

	if err := validateSize_MebibytesParameters(amount); err != nil {
		panic(err)
	}
	var returns Size

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Size",
		"mebibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount pebibytes.
//
// 1 PiB = 1024 TiB.
//
// Returns: a new `Size` instance.
func Size_Pebibytes(amount *float64) Size {
	_init_.Initialize()

	if err := validateSize_PebibytesParameters(amount); err != nil {
		panic(err)
	}
	var returns Size

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Size",
		"pebibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount tebibytes.
//
// 1 TiB = 1024 GiB.
//
// Returns: a new `Size` instance.
func Size_Tebibytes(amount *float64) Size {
	_init_.Initialize()

	if err := validateSize_TebibytesParameters(amount); err != nil {
		panic(err)
	}
	var returns Size

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Size",
		"tebibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) IsUnresolved() *bool {
	var returns *bool

	_jsii_.Invoke(
		s,
		"isUnresolved",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) ToBytes(opts *SizeConversionOptions) *float64 {
	if err := s.validateToBytesParameters(opts); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"toBytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) ToGibibytes(opts *SizeConversionOptions) *float64 {
	if err := s.validateToGibibytesParameters(opts); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"toGibibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) ToKibibytes(opts *SizeConversionOptions) *float64 {
	if err := s.validateToKibibytesParameters(opts); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"toKibibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) ToMebibytes(opts *SizeConversionOptions) *float64 {
	if err := s.validateToMebibytesParameters(opts); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"toMebibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) ToPebibytes(opts *SizeConversionOptions) *float64 {
	if err := s.validateToPebibytesParameters(opts); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"toPebibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Size) ToTebibytes(opts *SizeConversionOptions) *float64 {
	if err := s.validateToTebibytesParameters(opts); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"toTebibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

