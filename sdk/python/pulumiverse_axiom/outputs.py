# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'NotifierProperties',
    'NotifierPropertiesDiscord',
    'NotifierPropertiesDiscordWebhook',
    'NotifierPropertiesEmail',
    'NotifierPropertiesOpsgenie',
    'NotifierPropertiesPagerduty',
    'NotifierPropertiesSlack',
    'NotifierPropertiesWebhook',
    'GetNotifierPropertiesResult',
    'GetNotifierPropertiesDiscordResult',
    'GetNotifierPropertiesDiscordWebhookResult',
    'GetNotifierPropertiesEmailResult',
    'GetNotifierPropertiesOpsgenieResult',
    'GetNotifierPropertiesPagerdutyResult',
    'GetNotifierPropertiesSlackResult',
    'GetNotifierPropertiesWebhookResult',
]

@pulumi.output_type
class NotifierProperties(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "discordWebhook":
            suggest = "discord_webhook"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotifierProperties. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotifierProperties.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotifierProperties.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 discord: Optional['outputs.NotifierPropertiesDiscord'] = None,
                 discord_webhook: Optional['outputs.NotifierPropertiesDiscordWebhook'] = None,
                 email: Optional['outputs.NotifierPropertiesEmail'] = None,
                 opsgenie: Optional['outputs.NotifierPropertiesOpsgenie'] = None,
                 pagerduty: Optional['outputs.NotifierPropertiesPagerduty'] = None,
                 slack: Optional['outputs.NotifierPropertiesSlack'] = None,
                 webhook: Optional['outputs.NotifierPropertiesWebhook'] = None):
        if discord is not None:
            pulumi.set(__self__, "discord", discord)
        if discord_webhook is not None:
            pulumi.set(__self__, "discord_webhook", discord_webhook)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if opsgenie is not None:
            pulumi.set(__self__, "opsgenie", opsgenie)
        if pagerduty is not None:
            pulumi.set(__self__, "pagerduty", pagerduty)
        if slack is not None:
            pulumi.set(__self__, "slack", slack)
        if webhook is not None:
            pulumi.set(__self__, "webhook", webhook)

    @property
    @pulumi.getter
    def discord(self) -> Optional['outputs.NotifierPropertiesDiscord']:
        return pulumi.get(self, "discord")

    @property
    @pulumi.getter(name="discordWebhook")
    def discord_webhook(self) -> Optional['outputs.NotifierPropertiesDiscordWebhook']:
        return pulumi.get(self, "discord_webhook")

    @property
    @pulumi.getter
    def email(self) -> Optional['outputs.NotifierPropertiesEmail']:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def opsgenie(self) -> Optional['outputs.NotifierPropertiesOpsgenie']:
        return pulumi.get(self, "opsgenie")

    @property
    @pulumi.getter
    def pagerduty(self) -> Optional['outputs.NotifierPropertiesPagerduty']:
        return pulumi.get(self, "pagerduty")

    @property
    @pulumi.getter
    def slack(self) -> Optional['outputs.NotifierPropertiesSlack']:
        return pulumi.get(self, "slack")

    @property
    @pulumi.getter
    def webhook(self) -> Optional['outputs.NotifierPropertiesWebhook']:
        return pulumi.get(self, "webhook")


@pulumi.output_type
class NotifierPropertiesDiscord(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "discordChannel":
            suggest = "discord_channel"
        elif key == "discordToken":
            suggest = "discord_token"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotifierPropertiesDiscord. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotifierPropertiesDiscord.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotifierPropertiesDiscord.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 discord_channel: str,
                 discord_token: str):
        """
        :param str discord_channel: The discord channel
        :param str discord_token: The discord token
        """
        pulumi.set(__self__, "discord_channel", discord_channel)
        pulumi.set(__self__, "discord_token", discord_token)

    @property
    @pulumi.getter(name="discordChannel")
    def discord_channel(self) -> str:
        """
        The discord channel
        """
        return pulumi.get(self, "discord_channel")

    @property
    @pulumi.getter(name="discordToken")
    def discord_token(self) -> str:
        """
        The discord token
        """
        return pulumi.get(self, "discord_token")


@pulumi.output_type
class NotifierPropertiesDiscordWebhook(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "discordWebhookUrl":
            suggest = "discord_webhook_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotifierPropertiesDiscordWebhook. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotifierPropertiesDiscordWebhook.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotifierPropertiesDiscordWebhook.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 discord_webhook_url: str):
        """
        :param str discord_webhook_url: The discord webhook URL
        """
        pulumi.set(__self__, "discord_webhook_url", discord_webhook_url)

    @property
    @pulumi.getter(name="discordWebhookUrl")
    def discord_webhook_url(self) -> str:
        """
        The discord webhook URL
        """
        return pulumi.get(self, "discord_webhook_url")


@pulumi.output_type
class NotifierPropertiesEmail(dict):
    def __init__(__self__, *,
                 emails: Sequence[str]):
        """
        :param Sequence[str] emails: The emails to be notified
        """
        pulumi.set(__self__, "emails", emails)

    @property
    @pulumi.getter
    def emails(self) -> Sequence[str]:
        """
        The emails to be notified
        """
        return pulumi.get(self, "emails")


@pulumi.output_type
class NotifierPropertiesOpsgenie(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "apiKey":
            suggest = "api_key"
        elif key == "isEu":
            suggest = "is_eu"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotifierPropertiesOpsgenie. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotifierPropertiesOpsgenie.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotifierPropertiesOpsgenie.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 api_key: str,
                 is_eu: bool):
        """
        :param str api_key: The opsgenie API key
        :param bool is_eu: The opsgenie is EU
        """
        pulumi.set(__self__, "api_key", api_key)
        pulumi.set(__self__, "is_eu", is_eu)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> str:
        """
        The opsgenie API key
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="isEu")
    def is_eu(self) -> bool:
        """
        The opsgenie is EU
        """
        return pulumi.get(self, "is_eu")


@pulumi.output_type
class NotifierPropertiesPagerduty(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "routingKey":
            suggest = "routing_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotifierPropertiesPagerduty. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotifierPropertiesPagerduty.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotifierPropertiesPagerduty.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 routing_key: str,
                 token: str):
        """
        :param str routing_key: The pagerduty routing key
        :param str token: The pager duty token
        """
        pulumi.set(__self__, "routing_key", routing_key)
        pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="routingKey")
    def routing_key(self) -> str:
        """
        The pagerduty routing key
        """
        return pulumi.get(self, "routing_key")

    @property
    @pulumi.getter
    def token(self) -> str:
        """
        The pager duty token
        """
        return pulumi.get(self, "token")


@pulumi.output_type
class NotifierPropertiesSlack(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "slackUrl":
            suggest = "slack_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotifierPropertiesSlack. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotifierPropertiesSlack.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotifierPropertiesSlack.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 slack_url: str):
        """
        :param str slack_url: The slack URL
        """
        pulumi.set(__self__, "slack_url", slack_url)

    @property
    @pulumi.getter(name="slackUrl")
    def slack_url(self) -> str:
        """
        The slack URL
        """
        return pulumi.get(self, "slack_url")


@pulumi.output_type
class NotifierPropertiesWebhook(dict):
    def __init__(__self__, *,
                 url: str):
        """
        :param str url: The webhook URL
        """
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The webhook URL
        """
        return pulumi.get(self, "url")


@pulumi.output_type
class GetNotifierPropertiesResult(dict):
    def __init__(__self__, *,
                 discord: 'outputs.GetNotifierPropertiesDiscordResult',
                 discord_webhook: 'outputs.GetNotifierPropertiesDiscordWebhookResult',
                 email: 'outputs.GetNotifierPropertiesEmailResult',
                 opsgenie: 'outputs.GetNotifierPropertiesOpsgenieResult',
                 pagerduty: 'outputs.GetNotifierPropertiesPagerdutyResult',
                 slack: 'outputs.GetNotifierPropertiesSlackResult',
                 webhook: 'outputs.GetNotifierPropertiesWebhookResult'):
        pulumi.set(__self__, "discord", discord)
        pulumi.set(__self__, "discord_webhook", discord_webhook)
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "opsgenie", opsgenie)
        pulumi.set(__self__, "pagerduty", pagerduty)
        pulumi.set(__self__, "slack", slack)
        pulumi.set(__self__, "webhook", webhook)

    @property
    @pulumi.getter
    def discord(self) -> 'outputs.GetNotifierPropertiesDiscordResult':
        return pulumi.get(self, "discord")

    @property
    @pulumi.getter(name="discordWebhook")
    def discord_webhook(self) -> 'outputs.GetNotifierPropertiesDiscordWebhookResult':
        return pulumi.get(self, "discord_webhook")

    @property
    @pulumi.getter
    def email(self) -> 'outputs.GetNotifierPropertiesEmailResult':
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def opsgenie(self) -> 'outputs.GetNotifierPropertiesOpsgenieResult':
        return pulumi.get(self, "opsgenie")

    @property
    @pulumi.getter
    def pagerduty(self) -> 'outputs.GetNotifierPropertiesPagerdutyResult':
        return pulumi.get(self, "pagerduty")

    @property
    @pulumi.getter
    def slack(self) -> 'outputs.GetNotifierPropertiesSlackResult':
        return pulumi.get(self, "slack")

    @property
    @pulumi.getter
    def webhook(self) -> 'outputs.GetNotifierPropertiesWebhookResult':
        return pulumi.get(self, "webhook")


@pulumi.output_type
class GetNotifierPropertiesDiscordResult(dict):
    def __init__(__self__, *,
                 discord_channel: str,
                 discord_token: str):
        """
        :param str discord_channel: The discord channel
        :param str discord_token: The discord token
        """
        pulumi.set(__self__, "discord_channel", discord_channel)
        pulumi.set(__self__, "discord_token", discord_token)

    @property
    @pulumi.getter(name="discordChannel")
    def discord_channel(self) -> str:
        """
        The discord channel
        """
        return pulumi.get(self, "discord_channel")

    @property
    @pulumi.getter(name="discordToken")
    def discord_token(self) -> str:
        """
        The discord token
        """
        return pulumi.get(self, "discord_token")


@pulumi.output_type
class GetNotifierPropertiesDiscordWebhookResult(dict):
    def __init__(__self__, *,
                 discord_webhook_url: str):
        """
        :param str discord_webhook_url: The discord webhook URL
        """
        pulumi.set(__self__, "discord_webhook_url", discord_webhook_url)

    @property
    @pulumi.getter(name="discordWebhookUrl")
    def discord_webhook_url(self) -> str:
        """
        The discord webhook URL
        """
        return pulumi.get(self, "discord_webhook_url")


@pulumi.output_type
class GetNotifierPropertiesEmailResult(dict):
    def __init__(__self__, *,
                 emails: Sequence[str]):
        """
        :param Sequence[str] emails: The emails to be notified
        """
        pulumi.set(__self__, "emails", emails)

    @property
    @pulumi.getter
    def emails(self) -> Sequence[str]:
        """
        The emails to be notified
        """
        return pulumi.get(self, "emails")


@pulumi.output_type
class GetNotifierPropertiesOpsgenieResult(dict):
    def __init__(__self__, *,
                 api_key: str,
                 is_eu: bool):
        """
        :param str api_key: The opsgenie API key
        :param bool is_eu: The opsgenie is EU
        """
        pulumi.set(__self__, "api_key", api_key)
        pulumi.set(__self__, "is_eu", is_eu)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> str:
        """
        The opsgenie API key
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="isEu")
    def is_eu(self) -> bool:
        """
        The opsgenie is EU
        """
        return pulumi.get(self, "is_eu")


@pulumi.output_type
class GetNotifierPropertiesPagerdutyResult(dict):
    def __init__(__self__, *,
                 routing_key: str,
                 token: str):
        """
        :param str routing_key: The pagerduty routing key
        :param str token: The pager duty token
        """
        pulumi.set(__self__, "routing_key", routing_key)
        pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="routingKey")
    def routing_key(self) -> str:
        """
        The pagerduty routing key
        """
        return pulumi.get(self, "routing_key")

    @property
    @pulumi.getter
    def token(self) -> str:
        """
        The pager duty token
        """
        return pulumi.get(self, "token")


@pulumi.output_type
class GetNotifierPropertiesSlackResult(dict):
    def __init__(__self__, *,
                 slack_url: str):
        """
        :param str slack_url: The slack URL
        """
        pulumi.set(__self__, "slack_url", slack_url)

    @property
    @pulumi.getter(name="slackUrl")
    def slack_url(self) -> str:
        """
        The slack URL
        """
        return pulumi.get(self, "slack_url")


@pulumi.output_type
class GetNotifierPropertiesWebhookResult(dict):
    def __init__(__self__, *,
                 url: str):
        """
        :param str url: The webhook URL
        """
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The webhook URL
        """
        return pulumi.get(self, "url")


