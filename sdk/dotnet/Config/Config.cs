// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Snowflake
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly Pulumi.Config __config = new Pulumi.Config("snowflake");

        private static readonly __Value<string?> _account = new __Value<string?>(() => __config.Get("account"));
        public static string? Account
        {
            get => _account.Get();
            set => _account.Set(value);
        }

        private static readonly __Value<bool?> _browserAuth = new __Value<bool?>(() => __config.GetBoolean("browserAuth"));
        public static bool? BrowserAuth
        {
            get => _browserAuth.Get();
            set => _browserAuth.Set(value);
        }

        private static readonly __Value<string?> _oauthAccessToken = new __Value<string?>(() => __config.Get("oauthAccessToken"));
        public static string? OauthAccessToken
        {
            get => _oauthAccessToken.Get();
            set => _oauthAccessToken.Set(value);
        }

        private static readonly __Value<string?> _oauthClientId = new __Value<string?>(() => __config.Get("oauthClientId"));
        public static string? OauthClientId
        {
            get => _oauthClientId.Get();
            set => _oauthClientId.Set(value);
        }

        private static readonly __Value<string?> _oauthClientSecret = new __Value<string?>(() => __config.Get("oauthClientSecret"));
        public static string? OauthClientSecret
        {
            get => _oauthClientSecret.Get();
            set => _oauthClientSecret.Set(value);
        }

        private static readonly __Value<string?> _oauthEndpoint = new __Value<string?>(() => __config.Get("oauthEndpoint"));
        public static string? OauthEndpoint
        {
            get => _oauthEndpoint.Get();
            set => _oauthEndpoint.Set(value);
        }

        private static readonly __Value<string?> _oauthRedirectUrl = new __Value<string?>(() => __config.Get("oauthRedirectUrl"));
        public static string? OauthRedirectUrl
        {
            get => _oauthRedirectUrl.Get();
            set => _oauthRedirectUrl.Set(value);
        }

        private static readonly __Value<string?> _oauthRefreshToken = new __Value<string?>(() => __config.Get("oauthRefreshToken"));
        public static string? OauthRefreshToken
        {
            get => _oauthRefreshToken.Get();
            set => _oauthRefreshToken.Set(value);
        }

        private static readonly __Value<string?> _password = new __Value<string?>(() => __config.Get("password"));
        public static string? Password
        {
            get => _password.Get();
            set => _password.Set(value);
        }

        private static readonly __Value<string?> _privateKey = new __Value<string?>(() => __config.Get("privateKey"));
        public static string? PrivateKey
        {
            get => _privateKey.Get();
            set => _privateKey.Set(value);
        }

        private static readonly __Value<string?> _privateKeyPassphrase = new __Value<string?>(() => __config.Get("privateKeyPassphrase"));
        /// <summary>
        /// Supports the encryption ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, and
        /// des-ede3-cbc
        /// </summary>
        public static string? PrivateKeyPassphrase
        {
            get => _privateKeyPassphrase.Get();
            set => _privateKeyPassphrase.Set(value);
        }

        private static readonly __Value<string?> _privateKeyPath = new __Value<string?>(() => __config.Get("privateKeyPath"));
        public static string? PrivateKeyPath
        {
            get => _privateKeyPath.Get();
            set => _privateKeyPath.Set(value);
        }

        private static readonly __Value<string?> _region = new __Value<string?>(() => __config.Get("region"));
        public static string? Region
        {
            get => _region.Get();
            set => _region.Set(value);
        }

        private static readonly __Value<string?> _role = new __Value<string?>(() => __config.Get("role"));
        public static string? Role
        {
            get => _role.Get();
            set => _role.Set(value);
        }

        private static readonly __Value<string?> _username = new __Value<string?>(() => __config.Get("username"));
        public static string? Username
        {
            get => _username.Get();
            set => _username.Set(value);
        }

    }
}
