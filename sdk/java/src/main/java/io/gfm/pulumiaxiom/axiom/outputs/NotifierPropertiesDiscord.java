// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.gfm.pulumiaxiom.axiom.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class NotifierPropertiesDiscord {
    /**
     * @return The discord channel
     * 
     */
    private String discordChannel;
    /**
     * @return The discord token
     * 
     */
    private String discordToken;

    private NotifierPropertiesDiscord() {}
    /**
     * @return The discord channel
     * 
     */
    public String discordChannel() {
        return this.discordChannel;
    }
    /**
     * @return The discord token
     * 
     */
    public String discordToken() {
        return this.discordToken;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(NotifierPropertiesDiscord defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String discordChannel;
        private String discordToken;
        public Builder() {}
        public Builder(NotifierPropertiesDiscord defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.discordChannel = defaults.discordChannel;
    	      this.discordToken = defaults.discordToken;
        }

        @CustomType.Setter
        public Builder discordChannel(String discordChannel) {
            if (discordChannel == null) {
              throw new MissingRequiredPropertyException("NotifierPropertiesDiscord", "discordChannel");
            }
            this.discordChannel = discordChannel;
            return this;
        }
        @CustomType.Setter
        public Builder discordToken(String discordToken) {
            if (discordToken == null) {
              throw new MissingRequiredPropertyException("NotifierPropertiesDiscord", "discordToken");
            }
            this.discordToken = discordToken;
            return this;
        }
        public NotifierPropertiesDiscord build() {
            final var _resultValue = new NotifierPropertiesDiscord();
            _resultValue.discordChannel = discordChannel;
            _resultValue.discordToken = discordToken;
            return _resultValue;
        }
    }
}
