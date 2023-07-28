// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package esxi

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/edmondshtogu/pulumi-esxi-native/sdk/v3/go/esxi/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "esxi-native:index:PortGroup":
		r = &PortGroup{}
	case "esxi-native:index:ResourcePool":
		r = &ResourcePool{}
	case "esxi-native:index:VirtualDisk":
		r = &VirtualDisk{}
	case "esxi-native:index:VirtualMachine":
		r = &VirtualMachine{}
	case "esxi-native:index:VirtualSwitch":
		r = &VirtualSwitch{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:esxi-native" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"esxi-native",
		"index",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"esxi-native",
		&pkg{version},
	)
}
