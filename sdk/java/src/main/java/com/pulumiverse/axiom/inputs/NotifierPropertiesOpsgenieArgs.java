// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.gfm.pulumiaxiom.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class NotifierPropertiesOpsgenieArgs extends com.pulumi.resources.ResourceArgs {

    public static final NotifierPropertiesOpsgenieArgs Empty = new NotifierPropertiesOpsgenieArgs();

    /**
     * The opsgenie API key
     * 
     */
    @Import(name="apiKey", required=true)
    private Output<String> apiKey;

    /**
     * @return The opsgenie API key
     * 
     */
    public Output<String> apiKey() {
        return this.apiKey;
    }

    /**
     * The opsgenie is EU
     * 
     */
    @Import(name="isEu", required=true)
    private Output<Boolean> isEu;

    /**
     * @return The opsgenie is EU
     * 
     */
    public Output<Boolean> isEu() {
        return this.isEu;
    }

    private NotifierPropertiesOpsgenieArgs() {}

    private NotifierPropertiesOpsgenieArgs(NotifierPropertiesOpsgenieArgs $) {
        this.apiKey = $.apiKey;
        this.isEu = $.isEu;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NotifierPropertiesOpsgenieArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NotifierPropertiesOpsgenieArgs $;

        public Builder() {
            $ = new NotifierPropertiesOpsgenieArgs();
        }

        public Builder(NotifierPropertiesOpsgenieArgs defaults) {
            $ = new NotifierPropertiesOpsgenieArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiKey The opsgenie API key
         * 
         * @return builder
         * 
         */
        public Builder apiKey(Output<String> apiKey) {
            $.apiKey = apiKey;
            return this;
        }

        /**
         * @param apiKey The opsgenie API key
         * 
         * @return builder
         * 
         */
        public Builder apiKey(String apiKey) {
            return apiKey(Output.of(apiKey));
        }

        /**
         * @param isEu The opsgenie is EU
         * 
         * @return builder
         * 
         */
        public Builder isEu(Output<Boolean> isEu) {
            $.isEu = isEu;
            return this;
        }

        /**
         * @param isEu The opsgenie is EU
         * 
         * @return builder
         * 
         */
        public Builder isEu(Boolean isEu) {
            return isEu(Output.of(isEu));
        }

        public NotifierPropertiesOpsgenieArgs build() {
            if ($.apiKey == null) {
                throw new MissingRequiredPropertyException("NotifierPropertiesOpsgenieArgs", "apiKey");
            }
            if ($.isEu == null) {
                throw new MissingRequiredPropertyException("NotifierPropertiesOpsgenieArgs", "isEu");
            }
            return $;
        }
    }

}
