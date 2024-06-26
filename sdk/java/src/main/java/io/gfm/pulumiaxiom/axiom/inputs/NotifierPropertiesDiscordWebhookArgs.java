// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.gfm.pulumiaxiom.axiom.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class NotifierPropertiesDiscordWebhookArgs extends com.pulumi.resources.ResourceArgs {

    public static final NotifierPropertiesDiscordWebhookArgs Empty = new NotifierPropertiesDiscordWebhookArgs();

    /**
     * The discord webhook URL
     * 
     */
    @Import(name="discordWebhookUrl", required=true)
    private Output<String> discordWebhookUrl;

    /**
     * @return The discord webhook URL
     * 
     */
    public Output<String> discordWebhookUrl() {
        return this.discordWebhookUrl;
    }

    private NotifierPropertiesDiscordWebhookArgs() {}

    private NotifierPropertiesDiscordWebhookArgs(NotifierPropertiesDiscordWebhookArgs $) {
        this.discordWebhookUrl = $.discordWebhookUrl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NotifierPropertiesDiscordWebhookArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NotifierPropertiesDiscordWebhookArgs $;

        public Builder() {
            $ = new NotifierPropertiesDiscordWebhookArgs();
        }

        public Builder(NotifierPropertiesDiscordWebhookArgs defaults) {
            $ = new NotifierPropertiesDiscordWebhookArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param discordWebhookUrl The discord webhook URL
         * 
         * @return builder
         * 
         */
        public Builder discordWebhookUrl(Output<String> discordWebhookUrl) {
            $.discordWebhookUrl = discordWebhookUrl;
            return this;
        }

        /**
         * @param discordWebhookUrl The discord webhook URL
         * 
         * @return builder
         * 
         */
        public Builder discordWebhookUrl(String discordWebhookUrl) {
            return discordWebhookUrl(Output.of(discordWebhookUrl));
        }

        public NotifierPropertiesDiscordWebhookArgs build() {
            if ($.discordWebhookUrl == null) {
                throw new MissingRequiredPropertyException("NotifierPropertiesDiscordWebhookArgs", "discordWebhookUrl");
            }
            return $;
        }
    }

}
