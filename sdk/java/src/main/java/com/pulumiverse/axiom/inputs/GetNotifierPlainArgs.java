// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.gfm.pulumiaxiom.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetNotifierPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNotifierPlainArgs Empty = new GetNotifierPlainArgs();

    /**
     * Notifier identifier
     * 
     */
    @Import(name="id", required=true)
    private String id;

    /**
     * @return Notifier identifier
     * 
     */
    public String id() {
        return this.id;
    }

    private GetNotifierPlainArgs() {}

    private GetNotifierPlainArgs(GetNotifierPlainArgs $) {
        this.id = $.id;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNotifierPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNotifierPlainArgs $;

        public Builder() {
            $ = new GetNotifierPlainArgs();
        }

        public Builder(GetNotifierPlainArgs defaults) {
            $ = new GetNotifierPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id Notifier identifier
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            $.id = id;
            return this;
        }

        public GetNotifierPlainArgs build() {
            if ($.id == null) {
                throw new MissingRequiredPropertyException("GetNotifierPlainArgs", "id");
            }
            return $;
        }
    }

}
