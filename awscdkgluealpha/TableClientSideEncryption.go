package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Client-side encryption for an `S3Table`'s data.
//
// Independent of the bucket's server-side encryption and of who owns the bucket:
// the data is encrypted by the client before it is written to S3. When set, the
// `grant*` methods also grant the relevant KMS permissions on the key.
//
// Example:
//   var myDatabase Database
//
//   // KMS key is created automatically
//   // KMS key is created automatically
//   glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
//   	ClientSideEncryption: glue.TableClientSideEncryption_Kms(),
//   	// ...
//   	Database: myDatabase,
//   	Columns: []Column{
//   		&Column{
//   			Name: jsii.String("col1"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   })
//
//   // with an explicit KMS key
//   // with an explicit KMS key
//   glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
//   	ClientSideEncryption: glue.TableClientSideEncryption_*Kms(kms.NewKey(this, jsii.String("MyKey"))),
//   	// ...
//   	Database: myDatabase,
//   	Columns: []Column{
//   		&Column{
//   			Name: jsii.String("col1"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   })
//
// Experimental.
type TableClientSideEncryption interface {
}

// The jsii proxy struct for TableClientSideEncryption
type jsiiProxy_TableClientSideEncryption struct {
	_ byte // padding
}

// Client-side encryption (CSE-KMS) with an AWS KMS key managed by the account owner.
// Experimental.
func TableClientSideEncryption_Kms(key awskms.IKey) TableClientSideEncryption {
	_init_.Initialize()

	var returns TableClientSideEncryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.TableClientSideEncryption",
		"kms",
		[]interface{}{key},
		&returns,
	)

	return returns
}

