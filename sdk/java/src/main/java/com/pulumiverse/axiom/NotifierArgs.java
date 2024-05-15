// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.gfm.pulumiaxiom;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import io.gfm.pulumiaxiom.inputs.NotifierPropertiesArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NotifierArgs extends com.pulumi.resources.ResourceArgs {

    public static final NotifierArgs Empty = new NotifierArgs();

    /**
     * Notifier name
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Notifier name
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The properties of the notifier
     * 
     */
    @Import(name="properties", required=true)
    private Output<NotifierPropertiesArgs> properties;

    /**
     * @return The properties of the notifier
     * 
     */
    public Output<NotifierPropertiesArgs> properties() {
        return this.properties;
    }

    private NotifierArgs() {}

    private NotifierArgs(NotifierArgs $) {
        this.name = $.name;
        this.properties = $.properties;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NotifierArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NotifierArgs $;

        public Builder() {
            $ = new NotifierArgs();
        }

        public Builder(NotifierArgs defaults) {
            $ = new NotifierArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Notifier name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Notifier name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param properties The properties of the notifier
         * 
         * @return builder
         * 
         */
        public Builder properties(Output<NotifierPropertiesArgs> properties) {
            $.properties = properties;
            return this;
        }

        /**
         * @param properties The properties of the notifier
         * 
         * @return builder
         * 
         */
        public Builder properties(NotifierPropertiesArgs properties) {
            return properties(Output.of(properties));
        }

        public NotifierArgs build() {
            if ($.properties == null) {
                throw new MissingRequiredPropertyException("NotifierArgs", "properties");
            }
            return $;
        }
    }

}
