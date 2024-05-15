// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.gfm.pulumiaxiom.axiom.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class NotifierPropertiesEmailArgs extends com.pulumi.resources.ResourceArgs {

    public static final NotifierPropertiesEmailArgs Empty = new NotifierPropertiesEmailArgs();

    /**
     * The emails to be notified
     * 
     */
    @Import(name="emails", required=true)
    private Output<List<String>> emails;

    /**
     * @return The emails to be notified
     * 
     */
    public Output<List<String>> emails() {
        return this.emails;
    }

    private NotifierPropertiesEmailArgs() {}

    private NotifierPropertiesEmailArgs(NotifierPropertiesEmailArgs $) {
        this.emails = $.emails;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NotifierPropertiesEmailArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NotifierPropertiesEmailArgs $;

        public Builder() {
            $ = new NotifierPropertiesEmailArgs();
        }

        public Builder(NotifierPropertiesEmailArgs defaults) {
            $ = new NotifierPropertiesEmailArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param emails The emails to be notified
         * 
         * @return builder
         * 
         */
        public Builder emails(Output<List<String>> emails) {
            $.emails = emails;
            return this;
        }

        /**
         * @param emails The emails to be notified
         * 
         * @return builder
         * 
         */
        public Builder emails(List<String> emails) {
            return emails(Output.of(emails));
        }

        /**
         * @param emails The emails to be notified
         * 
         * @return builder
         * 
         */
        public Builder emails(String... emails) {
            return emails(List.of(emails));
        }

        public NotifierPropertiesEmailArgs build() {
            if ($.emails == null) {
                throw new MissingRequiredPropertyException("NotifierPropertiesEmailArgs", "emails");
            }
            return $;
        }
    }

}