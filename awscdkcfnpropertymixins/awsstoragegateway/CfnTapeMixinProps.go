package awsstoragegateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTapePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTapeMixinProps := &CfnTapeMixinProps{
//   	GatewayArn: jsii.String("gatewayArn"),
//   	KmsEncrypted: jsii.Boolean(false),
//   	KmsKey: jsii.String("kmsKey"),
//   	PoolId: jsii.String("poolId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TapeBarcode: jsii.String("tapeBarcode"),
//   	TapeSizeInBytes: jsii.Number(123),
//   	Worm: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tape.html
//
type CfnTapeMixinProps struct {
	// The Amazon Resource Name (ARN) of the Tape Gateway that hosts the virtual tape.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tape.html#cfn-storagegateway-tape-gatewayarn
	//
	GatewayArn *string `field:"optional" json:"gatewayArn" yaml:"gatewayArn"`
	// Set to true to use Amazon S3 server-side encryption with your own KMS key, or false to use a key managed by Amazon S3.
	//
	// Optional.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tape.html#cfn-storagegateway-tape-kmsencrypted
	//
	KmsEncrypted interface{} `field:"optional" json:"kmsEncrypted" yaml:"kmsEncrypted"`
	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for Amazon S3 server-side encryption.
	//
	// This value must be set if KMSEncrypted is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tape.html#cfn-storagegateway-tape-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The ID of the pool that you want to add your tape to for archiving.
	//
	// Tapes in this pool are archived in the S3 storage class that is associated with the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tape.html#cfn-storagegateway-tape-poolid
	//
	PoolId *string `field:"optional" json:"poolId" yaml:"poolId"`
	// A list of up to 50 tags to assign to the virtual tape.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tape.html#cfn-storagegateway-tape-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The barcode that you want to assign to the virtual tape.
	//
	// Barcodes cannot be reused, even after a tape is deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tape.html#cfn-storagegateway-tape-tapebarcode
	//
	TapeBarcode *string `field:"optional" json:"tapeBarcode" yaml:"tapeBarcode"`
	// The size, in bytes, of the virtual tape that you want to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tape.html#cfn-storagegateway-tape-tapesizeinbytes
	//
	TapeSizeInBytes *float64 `field:"optional" json:"tapeSizeInBytes" yaml:"tapeSizeInBytes"`
	// Set to true to create a write-once-read-many (WORM) virtual tape.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tape.html#cfn-storagegateway-tape-worm
	//
	Worm interface{} `field:"optional" json:"worm" yaml:"worm"`
}

