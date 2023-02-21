//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::Route53Resolver
package awscdkroute53resolveralpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FirewallDomains) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateFirewallDomains_FromAssetParameters(assetPath *string) error {
	if assetPath == nil {
		return fmt.Errorf("parameter assetPath is required, but nil was provided")
	}

	return nil
}

func validateFirewallDomains_FromListParameters(list *[]*string) error {
	if list == nil {
		return fmt.Errorf("parameter list is required, but nil was provided")
	}

	return nil
}

func validateFirewallDomains_FromS3Parameters(bucket awss3.IBucket, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateFirewallDomains_FromS3UrlParameters(url *string) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}

