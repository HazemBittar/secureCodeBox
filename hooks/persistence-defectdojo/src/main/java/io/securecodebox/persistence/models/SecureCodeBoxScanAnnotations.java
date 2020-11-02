package io.securecodebox.persistence.models;

import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
public enum SecureCodeBoxScanAnnotations {
  ENGAGEMENT_NAME("defectdojo.securecodebox.io/engagement-name"),
  ENGAGEMENT_VERSION("defectdojo.securecodebox.io/engagement-version"),
  ENGAGEMENT_TAGS("defectdojo.securecodebox.io/engagement-tags"),
  PRODUCT_NAME("defectdojo.securecodebox.io/product-name");

  @Getter
  private final String label;


  @Override
  public String toString() {
    return this.label;
  }
}
