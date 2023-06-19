package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Controls the countries in which content is distributed.
//
// Example:
//   // Adding restrictions to a Cloudfront Web Distribution.
//   var sourceBucket bucket
//
//   cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyDistribution"), &CloudFrontWebDistributionProps{
//   	OriginConfigs: []sourceConfiguration{
//   		&sourceConfiguration{
//   			S3OriginSource: &S3OriginConfig{
//   				S3BucketSource: sourceBucket,
//   			},
//   			Behaviors: []behavior{
//   				&behavior{
//   					IsDefaultBehavior: jsii.Boolean(true),
//   				},
//   			},
//   		},
//   	},
//   	GeoRestriction: cloudfront.GeoRestriction_Allowlist(jsii.String("US"), jsii.String("GB")),
//   })
//
// Experimental.
type GeoRestriction interface {
	// Two-letter, uppercase country code for a country that you want to allow/deny.
	//
	// Include one element for each country.
	// See ISO 3166-1-alpha-2 code on the *International Organization for Standardization* website.
	// Experimental.
	Locations() *[]*string
	// Specifies the restriction type to impose.
	// Experimental.
	RestrictionType() *string
}

// The jsii proxy struct for GeoRestriction
type jsiiProxy_GeoRestriction struct {
	_ byte // padding
}

func (j *jsiiProxy_GeoRestriction) Locations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeoRestriction) RestrictionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionType",
		&returns,
	)
	return returns
}


// Allow specific countries which you want CloudFront to distribute your content.
// Experimental.
func GeoRestriction_Allowlist(locations ...*string) GeoRestriction {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range locations {
		args = append(args, a)
	}

	var returns GeoRestriction

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.GeoRestriction",
		"allowlist",
		args,
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `denylist`.
func GeoRestriction_Blacklist(locations ...*string) GeoRestriction {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range locations {
		args = append(args, a)
	}

	var returns GeoRestriction

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.GeoRestriction",
		"blacklist",
		args,
		&returns,
	)

	return returns
}

// Deny specific countries which you don't want CloudFront to distribute your content.
// Experimental.
func GeoRestriction_Denylist(locations ...*string) GeoRestriction {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range locations {
		args = append(args, a)
	}

	var returns GeoRestriction

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.GeoRestriction",
		"denylist",
		args,
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `allowlist`.
func GeoRestriction_Whitelist(locations ...*string) GeoRestriction {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range locations {
		args = append(args, a)
	}

	var returns GeoRestriction

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.GeoRestriction",
		"whitelist",
		args,
		&returns,
	)

	return returns
}

