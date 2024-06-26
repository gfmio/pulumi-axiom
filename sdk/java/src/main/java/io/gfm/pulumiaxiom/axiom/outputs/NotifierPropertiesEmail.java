// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.gfm.pulumiaxiom.axiom.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class NotifierPropertiesEmail {
    /**
     * @return The emails to be notified
     * 
     */
    private List<String> emails;

    private NotifierPropertiesEmail() {}
    /**
     * @return The emails to be notified
     * 
     */
    public List<String> emails() {
        return this.emails;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(NotifierPropertiesEmail defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> emails;
        public Builder() {}
        public Builder(NotifierPropertiesEmail defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.emails = defaults.emails;
        }

        @CustomType.Setter
        public Builder emails(List<String> emails) {
            if (emails == null) {
              throw new MissingRequiredPropertyException("NotifierPropertiesEmail", "emails");
            }
            this.emails = emails;
            return this;
        }
        public Builder emails(String... emails) {
            return emails(List.of(emails));
        }
        public NotifierPropertiesEmail build() {
            final var _resultValue = new NotifierPropertiesEmail();
            _resultValue.emails = emails;
            return _resultValue;
        }
    }
}
