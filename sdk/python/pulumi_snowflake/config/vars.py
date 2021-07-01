# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'account',
    'browser_auth',
    'oauth_access_token',
    'oauth_client_id',
    'oauth_client_secret',
    'oauth_endpoint',
    'oauth_redirect_url',
    'oauth_refresh_token',
    'password',
    'private_key',
    'private_key_path',
    'region',
    'role',
    'username',
]

__config__ = pulumi.Config('snowflake')

account = __config__.get('account')

browser_auth = __config__.get('browserAuth')

oauth_access_token = __config__.get('oauthAccessToken')

oauth_client_id = __config__.get('oauthClientId')

oauth_client_secret = __config__.get('oauthClientSecret')

oauth_endpoint = __config__.get('oauthEndpoint')

oauth_redirect_url = __config__.get('oauthRedirectUrl')

oauth_refresh_token = __config__.get('oauthRefreshToken')

password = __config__.get('password')

private_key = __config__.get('privateKey')

private_key_path = __config__.get('privateKeyPath')

region = __config__.get('region')

role = __config__.get('role')

username = __config__.get('username')

