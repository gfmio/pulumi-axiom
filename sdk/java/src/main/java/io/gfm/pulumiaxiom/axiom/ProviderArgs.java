// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.gfm.pulumiaxiom.axiom;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * The Axiom API token.
     * 
     */
    @Import(name="apiToken", required=true)
    private Output<String> apiToken;

    /**
     * @return The Axiom API token.
     * 
     */
    public Output<String> apiToken() {
        return this.apiToken;
    }

    /**
     * The base url of the axiom api.
     * 
     */
    @Import(name="baseUrl")
    private @Nullable Output<String> baseUrl;

    /**
     * @return The base url of the axiom api.
     * 
     */
    public Optional<Output<String>> baseUrl() {
        return Optional.ofNullable(this.baseUrl);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.apiToken = $.apiToken;
        this.baseUrl = $.baseUrl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiToken The Axiom API token.
         * 
         * @return builder
         * 
         */
        public Builder apiToken(Output<String> apiToken) {
            $.apiToken = apiToken;
            return this;
        }

        /**
         * @param apiToken The Axiom API token.
         * 
         * @return builder
         * 
         */
        public Builder apiToken(String apiToken) {
            return apiToken(Output.of(apiToken));
        }

        /**
         * @param baseUrl The base url of the axiom api.
         * 
         * @return builder
         * 
         */
        public Builder baseUrl(@Nullable Output<String> baseUrl) {
            $.baseUrl = baseUrl;
            return this;
        }

        /**
         * @param baseUrl The base url of the axiom api.
         * 
         * @return builder
         * 
         */
        public Builder baseUrl(String baseUrl) {
            return baseUrl(Output.of(baseUrl));
        }

        public ProviderArgs build() {
            if ($.apiToken == null) {
                throw new MissingRequiredPropertyException("ProviderArgs", "apiToken");
            }
            return $;
        }
    }

}