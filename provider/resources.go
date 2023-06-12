// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unifi

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	provider "github.com/paultyng/terraform-provider-unifi/shim"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumiverse/pulumi-unifi/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "unifi"
	// modules:
	mainMod     = "index" // the unifi module
	firewallMod = "Firewall"
	portMod     = "Port"
	settingMod  = "Setting"
	iamMod      = "IAM"
)

var namespaceMap = map[string]string{
	mainPkg: "Unifi",
}

// unifiMember manufactures a type token for the IBM package and the given module, file name, and type.
func unifiMember(moduleTitle string, fn string, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(moduleTitle)
	namespaceMap[moduleName] = moduleTitle
	if fn != "" {
		moduleName += "/" + fn
	}
	return tokens.ModuleMember(mainPkg + ":" + moduleName + ":" + mem)
}

// unifiType manufactures a type token for the IBM package and the given module, file name, and type.
func unifiType(mod string, fn string, typ string) tokens.Type {
	return tokens.Type(unifiMember(mod, fn, typ))
}

// unifiTypeDefaultFile manufactures a standard resource token given a module and resource name.
func unifiTypeDefaultFile(mod string, typ string) tokens.Type {
	fn := string(unicode.ToLower(rune(typ[0]))) + typ[1:]
	return unifiType(mod, fn, typ)
}

// unifiDataSource manufactures a standard resource token given a module and resource name.
func unifiDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return unifiMember(mod, fn, res)
}

// unifiResource manufactures a standard resource token given a module and resource name.
func unifiResource(mod string, res string) tokens.Type {
	return unifiTypeDefaultFile(mod, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(provider.NewProvider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "unifi",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Unifi",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Pulumiverse",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// https://www.pulumi.com/docs/guides/pulumi-packages/how-to-author/#support-for-github-releases
		PluginDownloadURL: "github://api.github.com/pulumiverse",
		Description:       "A Pulumi package for creating and managing Unifi network resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "unifi", "category/network"},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/pulumiverse",
		Repository: "https://github.com/pulumiverse/pulumi-unifi",

		// The GitHub Org for the provider - defaults to `terraform-providers`
		GitHubOrg: "paultyng",
		Config: map[string]*tfbridge.SchemaInfo{
			"username": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UNIFI_USERNAME"},
				},
			},
			"password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UNIFI_PASSWORD"},
				},
			},
			"api_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UNIFI_API"},
				},
			},
			"site": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UNIFI_SITE"},
				},
			},
			"allow_insecure": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"UNIFI_INSECURE"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"unifi_account":        {Tok: unifiResource(mainMod, "Account")},
			"unifi_device":         {Tok: unifiResource(mainMod, "Device")},
			"unifi_dynamic_dns":    {Tok: unifiResource(mainMod, "DynamicDNS")},
			"unifi_firewall_group": {Tok: unifiResource(firewallMod, "Group")},
			"unifi_firewall_rule":  {Tok: unifiResource(firewallMod, "Rule")},
			"unifi_network":        {Tok: unifiResource(mainMod, "Network")},
			"unifi_port_forward":   {Tok: unifiResource(portMod, "Forward")},
			"unifi_port_profile":   {Tok: unifiResource(portMod, "Profile")},
			"unifi_radius_profile": {Tok: unifiResource(mainMod, "RadiusProfile")},
			"unifi_setting_mgmt":   {Tok: unifiResource(settingMod, "Mgmt")},
			"unifi_setting_usg":    {Tok: unifiResource(settingMod, "USG")},
			"unifi_setting_radius": {Tok: unifiResource(settingMod, "Radius")},
			"unifi_site":           {Tok: unifiResource(mainMod, "Site")},
			"unifi_static_route":   {Tok: unifiResource(mainMod, "StaticRoute")},
			"unifi_user":           {Tok: unifiResource(iamMod, "User")},
			"unifi_user_group":     {Tok: unifiResource(iamMod, "Group")},
			"unifi_wlan":           {Tok: unifiResource(mainMod, "Wlan")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"unifi_account":        {Tok: unifiDataSource(mainMod, "getAccount")},
			"unifi_ap_group":       {Tok: unifiDataSource(mainMod, "getApGroup")},
			"unifi_network":        {Tok: unifiDataSource(mainMod, "getNetwork")},
			"unifi_port_profile":   {Tok: unifiDataSource(portMod, "getProfile")},
			"unifi_radius_profile": {Tok: unifiDataSource(mainMod, "getRadiusProfile")},
			"unifi_user":           {Tok: unifiDataSource(iamMod, "getUser")},
			"unifi_user_group":     {Tok: unifiDataSource(iamMod, "getGroup")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@maienm/pulumi-unifi",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PackageName: "pulumiverse_unifi",
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			RootNamespace: "Pulumiverse",
			Namespaces:    namespaceMap,
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.pulumiverse",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
