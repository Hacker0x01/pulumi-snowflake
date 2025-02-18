// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func GetAccount(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:account")
}
func GetBrowserAuth(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "snowflake:browserAuth")
}
func GetOauthAccessToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:oauthAccessToken")
}
func GetOauthClientId(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:oauthClientId")
}
func GetOauthClientSecret(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:oauthClientSecret")
}
func GetOauthEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:oauthEndpoint")
}
func GetOauthRedirectUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:oauthRedirectUrl")
}
func GetOauthRefreshToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:oauthRefreshToken")
}
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:password")
}
func GetPrivateKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:privateKey")
}

// Supports the encryption ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, and
// des-ede3-cbc
func GetPrivateKeyPassphrase(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:privateKeyPassphrase")
}
func GetPrivateKeyPath(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:privateKeyPath")
}
func GetRegion(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:region")
}
func GetRole(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:role")
}
func GetUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "snowflake:username")
}
