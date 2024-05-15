// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.gfm.pulumiaxiom.axiom.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetNotifierPropertiesOpsgenie {
    /**
     * @return The opsgenie API key
     * 
     */
    private String apiKey;
    /**
     * @return The opsgenie is EU
     * 
     */
    private Boolean isEu;

    private GetNotifierPropertiesOpsgenie() {}
    /**
     * @return The opsgenie API key
     * 
     */
    public String apiKey() {
        return this.apiKey;
    }
    /**
     * @return The opsgenie is EU
     * 
     */
    public Boolean isEu() {
        return this.isEu;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNotifierPropertiesOpsgenie defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String apiKey;
        private Boolean isEu;
        public Builder() {}
        public Builder(GetNotifierPropertiesOpsgenie defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiKey = defaults.apiKey;
    	      this.isEu = defaults.isEu;
        }

        @CustomType.Setter
        public Builder apiKey(String apiKey) {
            if (apiKey == null) {
              throw new MissingRequiredPropertyException("GetNotifierPropertiesOpsgenie", "apiKey");
            }
            this.apiKey = apiKey;
            return this;
        }
        @CustomType.Setter
        public Builder isEu(Boolean isEu) {
            if (isEu == null) {
              throw new MissingRequiredPropertyException("GetNotifierPropertiesOpsgenie", "isEu");
            }
            this.isEu = isEu;
            return this;
        }
        public GetNotifierPropertiesOpsgenie build() {
            final var _resultValue = new GetNotifierPropertiesOpsgenie();
            _resultValue.apiKey = apiKey;
            _resultValue.isEu = isEu;
            return _resultValue;
        }
    }
}